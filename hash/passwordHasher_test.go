package hash

import (
    "reflect"
    "testing"
)

func TestPasswordHashIsNotTheSameAsPassword(t *testing.T) {
    password := "testPassword"

    hashedPassword := HashPassword(password) 

    if string(hashedPassword) == password {
        t.Fatalf(
            "Hashed password matches password!  This should not be the case! original password: %v, hashedPassword: %v\n",
            password,
            string(hashedPassword),
        )
    }
}

func TestPasswordHashesOfDifferentPasswordsShouldBeDifferent(t *testing.T) {
    password1 := "testPassword1"
    password2 := "testPassword2"

    hashedPassword1 := HashPassword(password1)
    hashedPassword2 := HashPassword(password2)

    if reflect.DeepEqual(hashedPassword1, hashedPassword2) {
        t.Fatalf(
            "Hashes of two passwords should not be the same!  Password1: %v, Hash1: %v, Password2: %v, Hash2: %v\n",
            password1,
            string(hashedPassword1),
            password2,
            string(hashedPassword2),
        )
    }
}

func TestPasswordsHashesOfSamePasswordShouldBeTheSame(t *testing.T) {
    password1 := "testPassword1"
    password2 := "testPassword1"

    hashedPassword1 := HashPassword(password1)
    hashedPassword2 := HashPassword(password2)

    if !reflect.DeepEqual(hashedPassword1, hashedPassword2) {
        t.Fatalf(
            "Hashes of two same passwords should be the same!  Password1: %v, Hash1: %v, Password2: %v, Hash2: %v\n",
            password1,
            string(hashedPassword1),
            password2,
            string(hashedPassword2),
        )
    }
}
