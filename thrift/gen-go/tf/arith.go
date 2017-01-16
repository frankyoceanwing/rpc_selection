// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tf

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - A
//  - B
type Args_ struct {
  A int32 `thrift:"A,1" db:"A" json:"A"`
  B int32 `thrift:"B,2" db:"B" json:"B"`
}

func NewArgs_() *Args_ {
  return &Args_{}
}


func (p *Args_) GetA() int32 {
  return p.A
}

func (p *Args_) GetB() int32 {
  return p.B
}
func (p *Args_) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Args_)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.A = v
}
  return nil
}

func (p *Args_)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.B = v
}
  return nil
}

func (p *Args_) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Args_) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("A", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:A: ", p), err) }
  if err := oprot.WriteI32(int32(p.A)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.A (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:A: ", p), err) }
  return err
}

func (p *Args_) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("B", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:B: ", p), err) }
  if err := oprot.WriteI32(int32(p.B)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.B (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:B: ", p), err) }
  return err
}

func (p *Args_) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Args_(%+v)", *p)
}

// Attributes:
//  - R
type Reply struct {
  R int32 `thrift:"R,1" db:"R" json:"R"`
}

func NewReply() *Reply {
  return &Reply{}
}


func (p *Reply) GetR() int32 {
  return p.R
}
func (p *Reply) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Reply)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.R = v
}
  return nil
}

func (p *Reply) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Reply"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Reply) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("R", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:R: ", p), err) }
  if err := oprot.WriteI32(int32(p.R)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.R (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:R: ", p), err) }
  return err
}

func (p *Reply) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Reply(%+v)", *p)
}

// Attributes:
//  - Quo
//  - Rem
type Quotient struct {
  Quo int32 `thrift:"Quo,1" db:"Quo" json:"Quo"`
  Rem int32 `thrift:"Rem,2" db:"Rem" json:"Rem"`
}

func NewQuotient() *Quotient {
  return &Quotient{}
}


func (p *Quotient) GetQuo() int32 {
  return p.Quo
}

func (p *Quotient) GetRem() int32 {
  return p.Rem
}
func (p *Quotient) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Quotient)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Quo = v
}
  return nil
}

func (p *Quotient)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Rem = v
}
  return nil
}

func (p *Quotient) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Quotient"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Quotient) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("Quo", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Quo: ", p), err) }
  if err := oprot.WriteI32(int32(p.Quo)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.Quo (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Quo: ", p), err) }
  return err
}

func (p *Quotient) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("Rem", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Rem: ", p), err) }
  if err := oprot.WriteI32(int32(p.Rem)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.Rem (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Rem: ", p), err) }
  return err
}

func (p *Quotient) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Quotient(%+v)", *p)
}

type Arith interface {
  // Parameters:
  //  - Arg
  Multiply(arg *Args_) (r *Reply, err error)
  // Parameters:
  //  - Arg
  Divide(arg *Args_) (r *Quotient, err error)
}

type ArithClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewArithClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ArithClient {
  return &ArithClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewArithClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ArithClient {
  return &ArithClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Arg
func (p *ArithClient) Multiply(arg *Args_) (r *Reply, err error) {
  if err = p.sendMultiply(arg); err != nil { return }
  return p.recvMultiply()
}

func (p *ArithClient) sendMultiply(arg *Args_)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("Multiply", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := ArithMultiplyArgs{
  Arg : arg,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *ArithClient) recvMultiply() (value *Reply, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "Multiply" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Multiply failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Multiply failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Multiply failed: invalid message type")
    return
  }
  result := ArithMultiplyResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Arg
func (p *ArithClient) Divide(arg *Args_) (r *Quotient, err error) {
  if err = p.sendDivide(arg); err != nil { return }
  return p.recvDivide()
}

func (p *ArithClient) sendDivide(arg *Args_)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("Divide", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := ArithDivideArgs{
  Arg : arg,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *ArithClient) recvDivide() (value *Quotient, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "Divide" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Divide failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Divide failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Divide failed: invalid message type")
    return
  }
  result := ArithDivideResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type ArithProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Arith
}

func (p *ArithProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ArithProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ArithProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewArithProcessor(handler Arith) *ArithProcessor {

  self4 := &ArithProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["Multiply"] = &arithProcessorMultiply{handler:handler}
  self4.processorMap["Divide"] = &arithProcessorDivide{handler:handler}
return self4
}

func (p *ArithProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x5

}

type arithProcessorMultiply struct {
  handler Arith
}

func (p *arithProcessorMultiply) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ArithMultiplyArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Multiply", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ArithMultiplyResult{}
var retval *Reply
  var err2 error
  if retval, err2 = p.handler.Multiply(args.Arg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Multiply: " + err2.Error())
    oprot.WriteMessageBegin("Multiply", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Multiply", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type arithProcessorDivide struct {
  handler Arith
}

func (p *arithProcessorDivide) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ArithDivideArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Divide", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ArithDivideResult{}
var retval *Quotient
  var err2 error
  if retval, err2 = p.handler.Divide(args.Arg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Divide: " + err2.Error())
    oprot.WriteMessageBegin("Divide", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Divide", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Arg
type ArithMultiplyArgs struct {
  Arg *Args_ `thrift:"arg,1" db:"arg" json:"arg"`
}

func NewArithMultiplyArgs() *ArithMultiplyArgs {
  return &ArithMultiplyArgs{}
}

var ArithMultiplyArgs_Arg_DEFAULT *Args_
func (p *ArithMultiplyArgs) GetArg() *Args_ {
  if !p.IsSetArg() {
    return ArithMultiplyArgs_Arg_DEFAULT
  }
return p.Arg
}
func (p *ArithMultiplyArgs) IsSetArg() bool {
  return p.Arg != nil
}

func (p *ArithMultiplyArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ArithMultiplyArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Arg = &Args_{}
  if err := p.Arg.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Arg), err)
  }
  return nil
}

func (p *ArithMultiplyArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Multiply_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ArithMultiplyArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("arg", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:arg: ", p), err) }
  if err := p.Arg.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Arg), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:arg: ", p), err) }
  return err
}

func (p *ArithMultiplyArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ArithMultiplyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ArithMultiplyResult struct {
  Success *Reply `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewArithMultiplyResult() *ArithMultiplyResult {
  return &ArithMultiplyResult{}
}

var ArithMultiplyResult_Success_DEFAULT *Reply
func (p *ArithMultiplyResult) GetSuccess() *Reply {
  if !p.IsSetSuccess() {
    return ArithMultiplyResult_Success_DEFAULT
  }
return p.Success
}
func (p *ArithMultiplyResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ArithMultiplyResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ArithMultiplyResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Reply{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ArithMultiplyResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Multiply_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ArithMultiplyResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ArithMultiplyResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ArithMultiplyResult(%+v)", *p)
}

// Attributes:
//  - Arg
type ArithDivideArgs struct {
  Arg *Args_ `thrift:"arg,1" db:"arg" json:"arg"`
}

func NewArithDivideArgs() *ArithDivideArgs {
  return &ArithDivideArgs{}
}

var ArithDivideArgs_Arg_DEFAULT *Args_
func (p *ArithDivideArgs) GetArg() *Args_ {
  if !p.IsSetArg() {
    return ArithDivideArgs_Arg_DEFAULT
  }
return p.Arg
}
func (p *ArithDivideArgs) IsSetArg() bool {
  return p.Arg != nil
}

func (p *ArithDivideArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ArithDivideArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Arg = &Args_{}
  if err := p.Arg.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Arg), err)
  }
  return nil
}

func (p *ArithDivideArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Divide_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ArithDivideArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("arg", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:arg: ", p), err) }
  if err := p.Arg.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Arg), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:arg: ", p), err) }
  return err
}

func (p *ArithDivideArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ArithDivideArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ArithDivideResult struct {
  Success *Quotient `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewArithDivideResult() *ArithDivideResult {
  return &ArithDivideResult{}
}

var ArithDivideResult_Success_DEFAULT *Quotient
func (p *ArithDivideResult) GetSuccess() *Quotient {
  if !p.IsSetSuccess() {
    return ArithDivideResult_Success_DEFAULT
  }
return p.Success
}
func (p *ArithDivideResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ArithDivideResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ArithDivideResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Quotient{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ArithDivideResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Divide_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ArithDivideResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ArithDivideResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ArithDivideResult(%+v)", *p)
}


