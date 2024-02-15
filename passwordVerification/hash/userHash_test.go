package hash

import (
    "reflect"
    "testing"
)

func TestHashUsername(t *testing.T) {
    username := "testUsername"
    password := "testPassword" 

    hashedUsername := HashUsername(username, password)

    if string(hashedUsername) == username {
        t.Fatalf(
            "Hashed username matches original username!  This should not be the case! original username: %v, hashedUsername: %v\n",
            username,
            string(hashedUsername),
        )
    }

    if string(hashedUsername) == password {
        t.Fatalf(
            "Hashed username matches password salt!  This should not be the case! password salt: %v, hashedUsername: %v\n",
            password,
            string(hashedUsername),
        )
    }
}

func TestHashUsernameWithDifferentPasswordsShouldResultInDifferentHashValue(t *testing.T) {
    username := "testUsername"
    password1 := "testPassword1"
    password2 := "testPassword2"

    hashedUsername1 := HashUsername(username, password1)
    hashedUsername2 := HashUsername(username, password2)

    if reflect.DeepEqual(hashedUsername1, hashedUsername2) {
        t.Fatalf(
            "Hashed usernames with different passwords should not match! hashedUsername1: %v, hashedUsername2: %v\n",
            string(hashedUsername1),
            string(hashedUsername2),
        )
    }
}

func TestHashUsernameWithDifferentUsernamesButSameMasterPasswordShouldNotMatch(t *testing.T) {
    username1 := "testUsername1"
    username2 := "testUsername2"
    password := "testPassword" 

    hashedUsername1 := HashUsername(username1, password)
    hashedUsername2 := HashUsername(username2, password)

    if reflect.DeepEqual(hashedUsername1, hashedUsername2) {
        t.Fatalf(
            "Hashed different usernames with same passwords should not match! hashedUsername1: %v, hashedUsername2: %v\n",
            string(hashedUsername1),
            string(hashedUsername2),
        )
    }
}

func TestUsernameHashWithSamePasswordShouldResultInTheSameHashValue(t *testing.T) {
    username := "testUsername"
    password1 := "testPassword"
    password2 := "testPassword"

    hashedUsername1 := HashUsername(username, password1)
    hashedUsername2 := HashUsername(username, password2)

    if !reflect.DeepEqual(hashedUsername1, hashedUsername2) {
        t.Fatalf(
            "Hashed usernames with same passwords should match! hashedUsername1: %v, hashedUsername2: %v\n",
            string(hashedUsername1),
            string(hashedUsername2),
        )
    }
}

