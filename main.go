package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

var (
	connectionString string
	queueName        string
)

func main() {
	utils.LoadEnv()

	// employees, _ := employeeService.GetEmployeeByServerId([]string{"1", "2"})
	// utils.PrintExistJsonIndent(employees, true)

	// utils.ExportModel("hms_api.comp_employee")

	runProducer()
}

func runProducer() {
	connectionString = os.Getenv("CONNECTION_STRING")
	queueName = os.Getenv("QUEUE_NAME")

	// สร้าง Service Bus Client
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a sender for the queue
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("Failed to create Service Bus Sender: %v", err)
	}

	// Create a message to send
	sessionID := "session-id-123"
	message := &azservicebus.Message{
		Body:      []byte("Hello, Service Bus!"),
		SessionID: &sessionID,
	}

	// Send the message
	err = sender.SendMessage(context.TODO(), message, nil)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message send successfully")

	// Close the sender
	err = sender.Close(context.TODO())
	if err != nil {
		log.Fatalf("Failed to close sender: %v", err)
	}

	// Close the client
	err = client.Close(context.TODO())
	if err != nil {
		log.Fatalf("Failed to close client: %v", err)
	}
}
