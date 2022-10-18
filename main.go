package main

import (
	"github.com/high-performance-payment-gateway/order-service/order"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("devops/.env")
	balanceModule := balance.NewModule()
	balanceModule.Start()
}
