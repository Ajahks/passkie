package hash

import (
	"reflect"
	"testing"
)

func TestUrlHash(t *testing.T) {
    url := "testurl.com"
    salt := []byte("TestSalt")

    hashedUrl := HashUrl(url, salt)

    if string(hashedUrl) == url {
        t.Fatalf(
            "Hashed url matches url!  This should not be the case! original url: %v, hashedUrl: %v\n",
            url,
            string(hashedUrl),
        )
    }
}

func TestUrlHashWithDifferentSaltShouldHashToDifferentValue(t *testing.T) {
    url := "testurl.com"
    salt1 := []byte("TestSalt")
    salt2 := []byte("TestSalt2")

    hashedUrl1 := HashUrl(url, salt1)
    hashedUrl2 := HashUrl(url, salt2)

    if reflect.DeepEqual(hashedUrl1, hashedUrl2) {
        t.Fatalf(
            "Hashed urls with different salts should not match! hashedUrl1: %v, hashedUrl2: %v\n",
            string(hashedUrl1),
            string(hashedUrl2),
        )
    }
}

func TestUrlHashWithSameSaltShouldResultInTheSameHashValue(t *testing.T) {
    url := "testurl.com"
    salt1 := []byte("TestSalt")
    salt2 := []byte("TestSalt")

    hashedUrl1 := HashUrl(url, salt1)
    hashedUrl2 := HashUrl(url, salt2)

    if !reflect.DeepEqual(hashedUrl1, hashedUrl2) {
        t.Fatalf(
            "Hashed urls with same salts should match! hashedUrl1: %v, hashedUrl2: %v\n",
            string(hashedUrl1),
            string(hashedUrl2),
        )
    }
}
