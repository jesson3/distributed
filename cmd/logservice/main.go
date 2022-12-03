package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jesson3/distributed/log"
	"github.com/jesson3/distributed/service"
)

func main() {
	log.Run(".distributed.log")
	host, port := "localhost", "9876"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
