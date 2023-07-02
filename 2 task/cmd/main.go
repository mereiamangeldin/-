package main

import (
	"context"
	"fmt"
	"github.com/mereiamangeldin/OnePlusTask2/repository"
	"github.com/mereiamangeldin/OnePlusTask2/service"
	"github.com/mereiamangeldin/OnePlusTask2/transport"
	"github.com/mereiamangeldin/OnePlusTask2/transport/handler"
	"log"
)

const url = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down: %s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repo, err := repository.New(url)
	if err != nil {
		return err
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		return err
	}
	h := handler.NewManager(svc)
	srv := transport.NewServer(h)
	return srv.Run(ctx)
}
