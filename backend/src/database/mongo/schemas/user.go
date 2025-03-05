package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

// User представляет модель пользователя
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}
