package main

import (
	"context"
	"errors"
	apporder "lesson-1/internal/app/order"
	"lesson-1/internal/domain/order"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"lesson-1/fixture"
	appcard "lesson-1/internal/app/card"
	appitem "lesson-1/internal/app/item"
	domaincard "lesson-1/internal/domain/card"
	domainitem "lesson-1/internal/domain/item"
	repositories "lesson-1/internal/repository"
)

func main() {

	// Service
	itemRepository := repositories.NewItemRepository(fixture.Items)
	itemService := domainitem.NewItemService(itemRepository)

	cardRepository := repositories.NewCardRepository()
	cardService := domaincard.NewCardService(cardRepository)

	orderRepository := repositories.NewOrderRepository()
	orderService := order.NewOrderService(orderRepository)

	// Application
	itemServer := appitem.NewItemServerImplementation(itemService)
	cardServer := appcard.NewCardServerImplementation(cardService, itemService)
	orderServer := apporder.NewOrderImplementation(orderService, cardService)

	// Mux
	mux := http.NewServeMux()

	appitem.RegisterRoutes(mux, itemServer)
	appcard.RegisterRoutes(mux, cardServer)
	apporder.RegisterRoutes(mux, orderServer)

	// Server
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		log.Printf("server listening at %s", server.Addr)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	_ = server.Shutdown(ctx)
}
