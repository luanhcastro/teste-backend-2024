package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ms-go/app/models"
	"ms-go/db"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "rails-to-go",
		GroupID: "go-consumer-group",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var product models.Product
		if err := json.Unmarshal(msg.Value, &product); err != nil {
			log.Fatal(err)
		}

		if err := saveToMongoDB(product); err != nil {
			log.Printf("Erro ao salvar o produto no MongoDB: %v", err)
		} else {
			fmt.Printf("Produto salvo no MongoDB: %+v\n", product)
		}
	}
}

func saveToMongoDB(product models.Product) error {
	collection := db.Connection()

	filter := bson.M{"id": product.ID}

	_, err := collection.ReplaceOne(
		context.TODO(),
		filter,
		product,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return fmt.Errorf("falha ao realizar upsert no MongoDB: %v", err)
	}

	return nil
}
