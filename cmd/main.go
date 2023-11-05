package main

import (
	"trendyol/internal/domain"
	"trendyol/internal/domain/promotion"
	"trendyol/internal/handler"
	"trendyol/internal/service"
	"trendyol/pkg"
)

func main() {

	handler := handler.NewCart(domain.NewCart(), service.NewPromotionService(
		promotion.NewPromotionAppliers(),
	))

	commandExecutor := pkg.NewCommandExecutor(handler)
	commandProcessor := pkg.NewFileCommandProcessor(pkg.CommandsInputFileName, pkg.ResponsesInputFileName)
	requests, err := commandProcessor.ReadFile()
	if err != nil {
		panic(err)
	}

	responses := commandExecutor.Run(requests)
	for _, response := range responses {
		commandProcessor.AppendResponse(response)
	}

	if err := commandProcessor.WriteFile(); err != nil {
		panic(err)
	}
}

// :)