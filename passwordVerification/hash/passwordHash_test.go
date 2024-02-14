package hash

import (
	"reflect"
	"testing"
)

func TestPasswordHash(t *testing.T) {
    password := "testPassword"
    salt := []byte("TestSalt")

    hashedPassword := HashPassword(password, salt)

    if string(hashedPassword) == password {
        t.Fatalf(
            "Hashed password matches original passowrd!  This should not be the case! original password: %v, hashedPassword: %v\n",
            password,
            string(hashedPassword),
        )
    }
}

func TestPasswordHashWithDifferentSaltShouldHashToDifferentValue(t *testing.T) {
    password := "testPassword"
    salt1 := []byte("TestSalt")
    salt2 := []byte("TestSalt2")

    hashedPassword1 := HashPassword(password, salt1)
    hashedPassword2 := HashPassword(password, salt2)

    if reflect.DeepEqual(hashedPassword1, hashedPassword2) {
        t.Fatalf(
            "Hashed passwords with different salts should not match! hashedPassword1: %v, hashedUrl2: %v\n",
            string(hashedPassword1),
            string(hashedPassword2),
        )
    }
}

func TestPasswordHashWithSameSaltShouldResultInTheSameHashValue(t *testing.T) {
    password := "testPassword"
    salt1 := []byte("TestSalt")
    salt2 := []byte("TestSalt")

    hashedPassword1 := HashPassword(password, salt1)
    hashedPassword2 := HashPassword(password, salt2)

    if !reflect.DeepEqual(hashedPassword1, hashedPassword2) {
        t.Fatalf(
            "Hashed passwords with same salts should match! hashedPassword1: %v, hashedUrl2: %v\n",
            string(hashedPassword1),
            string(hashedPassword2),
        )
    }
}

