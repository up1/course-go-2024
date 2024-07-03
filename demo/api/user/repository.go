package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetById(id string)
}

type User struct {
	Name string `json:"name", bson:"name"`
}

type MyRepo struct {
	client *mongo.Client
}

func (m *MyRepo) GetById(id string) {
	// Call repository
	u := User{
		Name: "Demo",
	}
	res, err := m.client.Database("test").Collection("user").InsertOne(context.Background(), u)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.InsertedID)

	// Find by id
	res2 := m.client.Database("test").Collection("user").FindOne(context.Background(), bson.M{"_id": res.InsertedID})
	if res2.Err() != nil {
		panic(res2.Err())
	}
	res2.Decode(&u)
	fmt.Printf("User: %+v\n", u)
}
