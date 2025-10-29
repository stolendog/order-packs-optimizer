package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/stolendog/order-packs-optimizer/internal/api"
	"github.com/stolendog/order-packs-optimizer/internal/app"
	"github.com/stolendog/order-packs-optimizer/internal/calculator"
	"github.com/stolendog/order-packs-optimizer/internal/domain"
	"github.com/stolendog/order-packs-optimizer/internal/infra"
)

func main() {
	repo := infra.NewMemoryRepository()
	calculator := calculator.NewPackDPCalculator()
	app := app.NewApp(calculator, repo)

	defaultPacks := GetDefaultPacks()
	// load default packs into storage
	err := app.StorePackList(defaultPacks)
	if err != nil {
		panic(err)
	}

	// TODO fix graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := api.Start(app); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
}

func GetDefaultPacks() []domain.Pack {
	var packs []domain.Pack
	for _, size := range []int{250, 500, 1000, 2000, 100} {
		pack, err := domain.NewPack(size)
		if err != nil {
			panic(err)
		}
		packs = append(packs, pack)
	}
	return packs
}
