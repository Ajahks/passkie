package urldb 

import (
	"os"
	"testing"

    "github.com/Ajahks/passkie/storage/localStorage"
)

func TestAddActiveUrlCreatesANewDbFileIfNonExistent(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url := "https://testurl.com/"

    AddActiveUrl(url)

    _, err := os.ReadFile(getFilePath())
    if err != nil {
        t.Fatalf("AddActiveUrl did not create a new file: %s", getFilePath())        
    }
}

func TestAddActiveUrlAndIsUserActiveReturnsTrueForNewlyActiveUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url := "https://testurl.com/"
    
    AddActiveUrl(url)
    result := IsUrlActive(url)

    if result == false {
        t.Fatal("Test user hash is not active in the DB")
    }
}

func TestIsUrlActiveReturnsFalseForUserThatWasNeverAdded(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url := "url.com"

    result := IsUrlActive(url)

    if result == true {
        t.Fatal("Test url should not be active in the DB")
    }
}

func TestAddActiveUrlForMultipleUserShowsThatBothUsersAreActive(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url1 := "url1.com"
    url2 := "url2.com"

    AddActiveUrl(url1)
    AddActiveUrl(url2)
    result1 := IsUrlActive(url1)
    result2 := IsUrlActive(url2)

    if result1 == false {
        t.Fatal("url1 was not active in the DB")
    }
    if result2 == false {
        t.Fatal("url2 was not active in the DB")
    }
}

func TestRemoveActiveUrlForNonExistentUserDoesNotFail(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    RemoveActiveUrl("url.com ")

    localstorage.CleanDB()
}

func TestAddActiveUrlAndRemoveActiveUserRemovesUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url := "url.com"

    AddActiveUrl(url)
    RemoveActiveUrl(url)
    result := IsUrlActive(url)

    if result == true {
        t.Fatal("Removed url is still active in the DB")
    }
}

func TestAddActiveUrlsAndRemoveOneUserDoesNotRemoveTheOther(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    url1 := "url1.com"
    url2 := "url2.com"

    AddActiveUrl(url1)
    AddActiveUrl(url2)
    RemoveActiveUrl(url1)
    result1 := IsUrlActive(url1)
    result2 := IsUrlActive(url2)

    if result1 == true {
        t.Fatal("Removed url is still active in the DB")
    }
    if result2 == false {
        t.Fatal("Non removed url is no longer active in DB")
    }
}

