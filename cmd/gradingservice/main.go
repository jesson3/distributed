package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/jesson3/distributed/grades"
	"github.com/jesson3/distributed/log"
	"github.com/jesson3/distributed/registry"
	"github.com/jesson3/distributed/service"
)

func main() {
	host, port := "localhost", "9878"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registeration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
