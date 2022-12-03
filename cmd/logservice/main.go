package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jesson3/distributed/log"
	"github.com/jesson3/distributed/registry"
	"github.com/jesson3/distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "9876"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registeration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
