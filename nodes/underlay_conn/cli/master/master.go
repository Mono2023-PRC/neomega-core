package main

import (
	"fmt"
	"time"

	"github.com/OmineDev/neomega-core/nodes"
	"github.com/OmineDev/neomega-core/nodes/defines"
	"github.com/OmineDev/neomega-core/nodes/underlay_conn"
)

func Master() {
	fmt.Println("server start")
	server, err := underlay_conn.NewServerFromBasicNet("tcp://0.0.0.0:7333")
	if err != nil {
		panic(err)
	}
	master := nodes.NewMasterNode(server)
	master.ExposeAPI("echo_master").InstantAPI(func(args defines.Values) (result defines.Values, err error) {
		return defines.FromString("echo_master echo").Extend(args), nil
	})
	master.ExposeAPI("echo").InstantAPI(func(args defines.Values) (result defines.Values, err error) {
		fmt.Println("master %v recv:", args.ToStrings())
		return defines.FromString("master echo").Extend(args), nil
	})
	<-master.WaitClosed()
	fmt.Println("master closed")
}

func Slave(id string) {
	fmt.Println("client start")
	client, err := underlay_conn.NewClientFromBasicNet("tcp://127.0.0.1:7333", time.Second)
	if err != nil {
		panic(err)
	}
	slave, err := nodes.NewSlaveNode(client)
	if err != nil {
		panic(err)
	}
	go func() {
		ret, err := slave.CallWithResponse("echo_master", defines.FromStrings("hello", "world", fmt.Sprintf("%v", id))).BlockGetResult()
		fmt.Printf("slave %v call echo_master get response: %v %v\n", id, ret.ToStrings(), err)
	}()
	slave.ExposeAPI(fmt.Sprintf("slave-echo-%v", id)).InstantAPI(func(args defines.Values) (result defines.Values, err error) {
		fmt.Println(fmt.Sprintf("slave-echo %v recv:", id), args.ToStrings())
		return defines.FromString(fmt.Sprintf("slave-echo-%v echo", id)).Extend(args), nil
	})
	slave.ExposeAPI("echo").InstantAPI(func(args defines.Values) (result defines.Values, err error) {
		fmt.Println(fmt.Sprintf("slave %v recv:", id), args.ToStrings())
		return defines.FromString(fmt.Sprintf("slave %v echo", id)).Extend(args), nil
	})
	ret, err := slave.CallWithResponse(fmt.Sprintf("slave-echo-%v", id), defines.FromStrings("hello", "world", fmt.Sprintf("%v", id))).BlockGetResult()
	fmt.Printf("slave %v call slave-echo-%v get response: %v %v\n", id, id, ret.ToStrings(), err)
	ret, err = slave.CallWithResponse("echo", defines.FromStrings("hello", "world", fmt.Sprintf("%v", id))).BlockGetResult()
	fmt.Printf("slave %v call echo get response: %v %v\n", id, ret.ToStrings(), err)
	<-slave.WaitClosed()
	fmt.Printf("slave %v closed\n", id)
}

func main() {
	go Master()
	time.Sleep(time.Second)
	go Slave("1")
	go Slave("2")
	c := make(chan struct{})
	<-c
}
