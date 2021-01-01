package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cookie represents a type of cookie
type Cookie struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Size        string             `json:"size" bson:"size,omitempty"`
	Ingredients []string           `json:"ingredients" bson:"ingredients,omitempty"`
	Calories    string             `json:"calories" bson:"calories,omitempty"`
	Location    string             `json:"location" bson:"location,omitempty"`
	Vegetarian  bool               `json:"vegetarian" bson:"vegetarian,omitempty"`
}
