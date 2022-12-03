package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jesson3/distributed/grades"
	"github.com/jesson3/distributed/registry"
	"github.com/jesson3/distributed/service"
)

func main() {
	host, port := "localhost", "9878"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registeration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
