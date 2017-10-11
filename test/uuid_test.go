package test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUID(t *testing.T) {
	uuidStr := uuid.NewV4().String()
	if len(uuidStr) == 0 {
		t.Fatalf("length is 0")
	}
}
