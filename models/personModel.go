package models

import "time"

type Person struct {
	PersonID  string    `json:"person_id" bson:"person_id"`
	Name      string    `json:"name" bson:"name" validate:"required"`
	Age       int       `json:"age" bson:"age" validate:"required"`
	Hobbies   []string  `json:"hobbies" bson:"hobbies"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
