package main

import (
	"fmt"
	"log"

	"github.com/qasaur/gremgo"
)

func main() {
	errs := make(chan error)
	go func(chan error) {
		err := <-errs
		log.Fatal("Lost connection to the database: " + err.Error())
	}(errs) // Example of connection error handling logic

	dialer := gremgo.NewDialer("ws://127.0.0.1:8182/gremlin") // Returns a WebSocket dialer to connect to Gremlin Server
	g, err := gremgo.Dial(dialer, errs)                       // Returns a gremgo client to interact with
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := g.Execute( // Sends a query to Gremlin Server with bindings
		"g.V().count()",
		map[string]string{},
		map[string]string{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
