package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase *mongo.Database = (MongodbInitialization().Database(os.Getenv("MONGO_DATABASE")))

func MongodbInitialization() *mongo.Client {
	return Connect()
}

func Connect() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // context.Background() create empty context
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	println("Connected to MongoDB!")
	return client
}

func Disconnect(client *mongo.Client, ctx context.Context) {
	client.Disconnect(ctx)
}
