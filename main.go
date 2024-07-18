package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

type MessagePayload struct {
	ID int `json:"id"`
}

var (
	connectionString string
	queueName        string
	maxProducer      int = 5
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

	messages := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

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

	/* // Create a message to send
	sessionID := "session-id-123"
	message := &azservicebus.Message{
		Body:      []byte("Hello, Service Bus!"),
		SessionID: &sessionID,
	}

	// Send the message
	err = sender.SendMessage(context.TODO(), message, nil)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	} */
	sendConcurrentMessage(sender, messages)

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

func sendConcurrentMessage(sender *azservicebus.Sender, messages []int) {
	var wg sync.WaitGroup

	queueMessage := make(chan int, len(messages))

	for i := 0; i < maxProducer; i++ {
		wg.Add(1)

		sessionID := fmt.Sprintf("SID-%d", i)

		go func(i int, sessionID string) {
			defer wg.Done()

			for msg := range queueMessage {
				payload := MessagePayload{
					ID: msg,
				}

				body, err1 := json.Marshal(payload)
				if err1 != nil {
					log.Fatalf("Failed to marshal JSON: %v", err1)
				}
				// body = []byte(strconv.Itoa(msg))

				message := &azservicebus.Message{
					Body:      body,
					SessionID: &sessionID,
				}

				fmt.Printf("Producer-%d is sending => SessionId: %s, Message: %d\n", i, sessionID, msg)

				err := sender.SendMessage(context.TODO(), message, nil)
				if err != nil {
					log.Fatalf("Failed to send message: %v", err)
				}
				// time.Sleep(1 * time.Second)

				fmt.Printf("Producer-%d is send success => SessionId: %s, Message: %d\n", i, sessionID, msg)
			}
		}(i, sessionID)
	}

	for _, msg := range messages {
		queueMessage <- msg
	}
	close(queueMessage)

	wg.Wait()
}
