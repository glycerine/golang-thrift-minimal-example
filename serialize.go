package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/glycerine/golang-thrift-minimal-example/vendor/tutorial"
)

func main() {
	binary()
	compact()
}

func binary() {
	fmt.Printf("\n ==== demo Thrift Binary serialization ====\n")
	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t)

	tser := &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}

	str := "hi there"
	a := &tutorial.Work{
		Num1:    12,
		Num2:    24,
		Comment: &str,
	}

	by, err := tser.Write(a)
	panicOn(err)
	fmt.Printf("by = '%v', length %v\n", string(by), len(by))

	t2 := thrift.NewTMemoryBufferLen(1024)
	p2 := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t2)

	deser := &thrift.TDeserializer{
		Transport: t2,
		Protocol:  p2,
	}

	b := tutorial.NewWork()
	deser.Transport.Close() // resets underlying bytes.Buffer
	err = deser.Read(b, by)
	panicOn(err)
	fmt.Printf("b = '%#v'\n", b)
}

func compact() {
	fmt.Printf("\n ==== demo Thrift Compact Binary serialization ====\n")

	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTCompactProtocolFactory().GetProtocol(t)

	tser := &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}

	str := "hi there"
	a := &tutorial.Work{
		Num1:    12,
		Num2:    24,
		Comment: &str,
	}

	by, err := tser.Write(a)
	panicOn(err)
	fmt.Printf("by = '%v', length %v\n", string(by), len(by))

	t2 := thrift.NewTMemoryBufferLen(1024)
	p2 := thrift.NewTCompactProtocolFactory().GetProtocol(t2)

	deser := &thrift.TDeserializer{
		Transport: t2,
		Protocol:  p2,
	}

	b := tutorial.NewWork()
	deser.Transport.Close() // resets underlying bytes.Buffer
	err = deser.Read(b, by)
	panicOn(err)
	fmt.Printf("b = '%#v'\n", b)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
