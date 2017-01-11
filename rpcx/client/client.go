package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/frankyoceanwing/rpc_selection/common"
	"github.com/montanaflynn/stats"
	"github.com/smallnest/rpcx"
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

	totalTime := time.Now().UnixNano()
	times := make([]int64, 0, *total)
	for i := 0; i < n; i++ {
		go func(i int) {
			s := &rpcx.DirectClientSelector{Network: "tcp", Address: *host}
			client := rpcx.NewClient(s)

			args := &common.Args{A: 7, B: 8}
			var reply int
			// 预热
			for j := 0; j < 10; j++ {
				client.Call("Arith.Multiply", args, &reply)
			}

			for k := 0; k < m; k++ {
				singleTime := time.Now().UnixNano()
				err = client.Call("Arith.Multiply", args, &reply)
				singleTime = time.Now().UnixNano() - singleTime
				times = append(times, singleTime)
				if err != nil {
					log.Fatal("arith error:", err)
				}
				//fmt.Printf("[%d-%d/%d]arith: %d*%d=%d %dns\n", (i + 1), (k + 1), m, args.A, args.B, reply, singleTime)
				wg.Done()
			}

			err = client.Close()
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
