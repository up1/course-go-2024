package hello_test

import (
	"hello"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess_HelloWorld(t *testing.T) {
	res := hello.SayHi()
	if res != "Hello, World!" {
		t.Errorf("Expected: Hello, World! Got: %s", res)
	}
	assert.Equal(t, "Hello, World!", res, "they should be equal")
}
