package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/EduardoPPCaldas/fullcycle/internal/order/infra/database"
	"github.com/EduardoPPCaldas/fullcycle/internal/order/usecase"
	"github.com/EduardoPPCaldas/fullcycle/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()	
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println(outputDTO)
		time.Sleep(500 * time.Millisecond)
	}
}