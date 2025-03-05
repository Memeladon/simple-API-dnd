package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Engine - структура для работы с MongoDB
type Engine struct {
	Client *mongo.Client
}

// NewMongoEngine создает новый экземпляр MongoEngine
func NewMongoEngine(uri string) (*Engine, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return &Engine{Client: client}, nil
}

// GetCollection возвращает коллекцию по имени
func (e *Engine) GetCollection(dbName, collectionName string) *mongo.Collection {
	return e.Client.Database(dbName).Collection(collectionName)
}

// Close закрывает подключение к MongoDB
func (e *Engine) Close(ctx context.Context) error {
	return e.Client.Disconnect(ctx)
}
