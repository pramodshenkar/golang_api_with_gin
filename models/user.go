package models

import (
	"github.com/pramodshenkar/golang_todo_api/forms"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
}

// UserModel defines the model structure
type UserModel struct{}

// Signup handles registering a user
func (u *UserModel) Signup(data forms.SignupUserCommand) error {
	// Connect to the user collection
	collection := dbConnect.Use(databaseName, "user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": data.Password,
		// This will come later when adding verification
		"is_verified": false,
	})

	return err
}
