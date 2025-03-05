package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"simple-API-dnd/src/database/mongo/schemas"
)

// UserRepository - репозиторий для работы с пользователями
type UserRepository struct {
	Collection *mongo.Collection
}

// NewUserRepository создает новый экземпляр UserRepository
func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: collection}
}

// InsertUser вставляет нового пользователя в коллекцию
func (r *UserRepository) InsertUser(ctx context.Context, user schemas.User) (*mongo.InsertOneResult, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Failed to insert user:", err)
		return nil, err
	}
	return result, nil
}

// FindUserByUsername ищет пользователя по имени
func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*schemas.User, error) {
	var user schemas.User
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("User not found:", username)
			return nil, nil
		}
		log.Println("Failed to find user:", err)
		return nil, err
	}
	return &user, nil
}

// DeleteUserByUsername удаляет пользователя по имени
func (r *UserRepository) DeleteUserByUsername(ctx context.Context, username string) (*mongo.DeleteResult, error) {
	result, err := r.Collection.DeleteOne(ctx, bson.M{"username": username})
	if err != nil {
		log.Println("Failed to delete user:", err)
		return nil, err
	}
	return result, nil
}
