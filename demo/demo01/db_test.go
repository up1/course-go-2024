package hello_test

import (
	"hello"
	"testing"
)

type MyDBMock struct {
}

func (db *MyDBMock) Get() (string, error) {
	return "Called mock db", nil
}

func TestCalledRealDB(t *testing.T) {
	db := &hello.MyDB{}
	demo := hello.Demo{Db: db}
	if result, _ := demo.Db.Get(); result != "Called real db" {
		t.Errorf("Expected Called real db, got %s", result)
	}
}

func TestCalledStubDB(t *testing.T) {
	db := &MyDBMock{}
	demo := hello.Demo{Db: db}
	if result, _ := demo.Db.Get(); result != "Called real db" {
		t.Errorf("Expected Called real db, got %s", result)
	}
}
