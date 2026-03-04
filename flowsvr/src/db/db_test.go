package db

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Errorf("InitDB failed: %v", err)
	}
}
