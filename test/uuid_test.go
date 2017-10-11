package test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUID(t *testing.T) {
	uuidV4 := uuid.NewV4()
	if len(uuidV4.String()) == 0 {
		t.Fatalf("length is 0")
	}

	uuidV1 := uuid.NewV1()
	if len(uuidV1.String()) == 0 {
		t.Fatalf("length is 0")
	}
}
