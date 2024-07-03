package hello

import "fmt"

type User struct {
	Id   string
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("Id=%s, Name=%s", u.Id, u.Name)
}

// Composition
type UserV2 struct {
	Age int
	User
}

func NewUserV2(id, name string, age int) UserV2 {
	return UserV2{
		Age:  age,
		User: User{id, name},
	}
}

// Builder/Creational pattern
func NewUser(id, name string) User {
	return User{
		Id:   id,
		Name: name,
	}
}

// Receiver function
func (u User) GetById() string {
	u.Id = "xxxxxx"
	return "get by id " + u.Id
}
