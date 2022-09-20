package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"runtime"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL) // localhost:4222
	if err != nil {
		panic(err)
	}
	count := 0
	conn.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Printf("%v %v\n", string(msg.Data), count)
		count++
		conn.Publish(msg.Reply, []byte("Message delivered"))
	})
	runtime.Goexit()
	defer conn.Close()
}
