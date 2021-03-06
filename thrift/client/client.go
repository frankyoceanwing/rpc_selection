package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/frankyoceanwing/rpc_selection/thrift/gen-go/tf"
	"github.com/montanaflynn/stats"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")
var host = flag.String("s", common.ServerAddress, "listened ip and port")

func main() {
	flag.Parse()

	n := *concurrency
	m := *total / n

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	wg.Add(n * m)

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	totalTime := time.Now().UnixNano()
	times := make([]int64, 0, *total)
	for i := 0; i < n; i++ {
		go func(i int) {
			var transport thrift.TTransport
			var err error

			transport, err = thrift.NewTSocket(*host)
			if err != nil {
				log.Fatal("creating socket error:", err)
			}

			transport = transportFactory.GetTransport(transport)
			if err = transport.Open(); err != nil {
				log.Fatal("opening socket error:", err)
			}

			client := tf.NewArithClientFactory(transport, protocolFactory)

			args := &tf.Args_{A: 7, B: 8}
			// 预热
			for j := 0; j < 10; j++ {
				client.Multiply(args)
			}

			for k := 0; k < m; k++ {
				singleTime := time.Now().UnixNano()
				_, err = client.Multiply(args)
				singleTime = time.Now().UnixNano() - singleTime
				times = append(times, singleTime)
				if err != nil {
					log.Fatal("arith error:", err)
				}
				//fmt.Printf("[%d-%d/%d]arith: %d*%d=%d %dns\n", (i + 1), (k + 1), m, args.A, args.B, reply, singleTime)
				wg.Done()
			}

			err = transport.Close()
			if err != nil {
				log.Fatal("close error:", err)
			}
		}(i)
	}

	wg.Wait()

	totalTime = (time.Now().UnixNano() - totalTime) / 1000000

	l := len(times)
	times2 := make([]float64, l, l)
	for i := 0; i < l; i++ {
		times2[i] = float64(times[i])
	}

	mean, _ := stats.Mean(times2)
	median, _ := stats.Median(times2)
	max, _ := stats.Max(times2)
	min, _ := stats.Min(times2)
	p99, _ := stats.Percentile(times2, 99.9)

	fmt.Printf("concurrency         : %d\n", n)
	fmt.Printf("requests per client : %d\n", m)
	fmt.Printf("took                : %d ms\n", totalTime)
	fmt.Printf("throughput  (TPS)   : %d\n", int64(n*m)*1000/totalTime)
	fmt.Printf("mean: %.f ns, median: %.f ns, max: %.f ns, min: %.f ns, p99: %.f ns\n", mean, median, max, min, p99)
	fmt.Printf("mean: %d ms, median: %d ms, max: %d ms, min: %d ms, p99: %d ms\n", int64(mean/1000000), int64(median/1000000), int64(max/1000000), int64(min/1000000), int64(p99/1000000))
}
