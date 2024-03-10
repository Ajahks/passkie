package activeuserdb

import (
	"os"
	"testing"

    "github.com/Ajahks/passkie/storage/localStorage"
)

func TestAddActiveUserCreatesANewDbFileIfNonExistent(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash := "userhash"

    AddActiveUser(userhash)

    _, err := os.ReadFile(getFilePath())
    if err != nil {
        t.Fatalf("AddActiveUser did not create a new file: %s", getFilePath())        
    }
}

func TestAddActiveUserAndIsUserActiveReturnsTrueForNewlyActiveUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash := "userhash"
    
    AddActiveUser(userhash)
    result := IsUserHashActive(userhash)

    if result == false {
        t.Fatal("Test user hash is not active in the DB")
    }
}

func TestIsUserActiveReturnsFalseForUserThatWasNeverAdded(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash := "userhash"

    result := IsUserHashActive(userhash)

    if result == true {
        t.Fatal("Test userhash should not be active in the DB")
    }
}

func TestAddActiveUserForMultipleUserShowsThatBothUsersAreActive(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash1 := "userhash1"
    userhash2 := "userhash2"

    AddActiveUser(userhash1)
    AddActiveUser(userhash2)
    result1 := IsUserHashActive(userhash1)
    result2 := IsUserHashActive(userhash2)

    if result1 == false {
        t.Fatal("userhash1 was not active in the DB")
    }
    if result2 == false {
        t.Fatal("userhash2 was not active in the DB")
    }
}

func TestRemoveActiveUserForNonExistentUserDoesNotFail(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    RemoveActiveUser("testuserhash")

    localstorage.CleanDB()
}

func TestAddActiveUserAndRemoveActiveUserRemovesUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash := "userhash"

    AddActiveUser(userhash)
    RemoveActiveUser(userhash)
    result := IsUserHashActive(userhash)

    if result == true {
        t.Fatal("Removed user is still active in the DB")
    }
}

func TestAddActiveUsersAndRemoveOneUserDoesNotRemoveTheOther(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    userhash1 := "userhash1"
    userhash2 := "userhash2"

    AddActiveUser(userhash1)
    AddActiveUser(userhash2)
    RemoveActiveUser(userhash1)
    result1 := IsUserHashActive(userhash1)
    result2 := IsUserHashActive(userhash2)

    if result1 == true {
        t.Fatal("Removed user is still active in the DB")
    }
    if result2 == false {
        t.Fatal("Non removed user is no longer active in DB")
    }
}

