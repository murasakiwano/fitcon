package fitconner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordIsBeingHashed(t *testing.T) {
	password := "test-password1"

	fc, err := New("C123456", "Joca", password, "BFIGHT", "", "", "", "", "", 1)
	if err != nil {
		t.Fatalf("Error creating FitConner: %s", err)
	}

	if !assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(fc.HashedPassword), []byte("test-password1"))) {
		t.Fatal("Password is not being hashed")
	}
}
