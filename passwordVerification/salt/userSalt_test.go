package salt

import (
	"bytes"
	"os"
	"testing"
)

func TestGetSaltForUserHashGeneratesReturnsASalt(t *testing.T) {
    testUserHash := []byte("testUserHash")
    expectedSaltLen := 32

    salt := GetSaltForUserHash(testUserHash)

    if len(salt) != expectedSaltLen {
        t.Fatalf("Generated salt len is not correct.  Salt len found: %v", len(salt))
    }
    if bytes.Equal(salt, make([]byte, expectedSaltLen)) {
        t.Fatalf("Generated salt is just an empty byte buffer")
    }

    cleanDb()
}

func TestPutSaltForUserThenGetSaltForUserRetrievesPutSalt(t *testing.T) {
    testUserHash := []byte("testUserHash")
    testSalt := []byte("testSalt")

    PutSaltForUserHash(testUserHash, testSalt)
    retrievedSalt := GetSaltForUserHash(testUserHash)

    if !bytes.Equal(testSalt, retrievedSalt) {
        t.Fatalf("Got salt that is different from put salt: Retrieved Salt: %s, Expected: %s", string(retrievedSalt), string(testSalt))
    }

    cleanDb()
}

func TestGetSaltForUserHashTwiceReturnsTheSameSaltGeneratedTheFirstTime(t *testing.T) {
    testUserHash := []byte("testUserHash")

    salt1 := GetSaltForUserHash(testUserHash)
    salt2 := GetSaltForUserHash(testUserHash)

    if !bytes.Equal(salt1, salt2) {
        t.Fatalf("GetSaltForUserHash does not return the same salt when called twice! Salt 1: %s, Salt2: %s", salt1, salt2)
    }

    cleanDb()
}

func TestGetSaltForUserGeneratesNewHashAfterDelete(t *testing.T) {
    testUserHash := []byte("testUserHash")

    salt1 := GetSaltForUserHash(testUserHash)
    RemoveSaltForUserHash(testUserHash)
    salt2 := GetSaltForUserHash(testUserHash)

    if bytes.Equal(salt1, salt2) {
        t.Fatalf("GetSaltForUserHash does not generate a new hash after delete! Salt 1: %s, Salt2: %s", salt1, salt2)
    }

    cleanDb()
}

func cleanDb() {
    os.RemoveAll("localDb")
}
