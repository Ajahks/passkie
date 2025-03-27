package urldb

import (
	"os"
	"testing"

	"github.com/Ajahks/passkie/storage/localStorage"
)

func TestAddActiveUrlForUserCreatesANewDbFileIfNonExistent(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "https://testurl.com/"
	username := "testUser"

	AddActiveUrlForUser(url, username)

	_, err := os.ReadFile(getFilePath(username))
	if err != nil {
		t.Fatalf("AddActiveUrlForUser did not create a new file: %s", getFilePath(username))
	}
}

func TestAddActiveUrlForUserAndIsUserActiveReturnsTrueForNewlyActiveUser(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "https://testurl.com/"
	username := "testUser"

	AddActiveUrlForUser(url, username)
	result := IsUrlActiveForUser(url, username)

	if result == false {
		t.Fatal("Test user hash is not active in the DB")
	}
}

func TestIsUrlActiveForUserReturnsFalseForUserThatWasNeverAdded(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "url.com"
	username := "testUser"

	result := IsUrlActiveForUser(url, username)

	if result == true {
		t.Fatal("Test url should not be active in the DB")
	}
}

func TestAddActiveUrlForUserForMultipleUrlsShowsThatBothUrlsAreActive(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url1 := "url1.com"
	url2 := "url2.com"
	username := "username"

	AddActiveUrlForUser(url1, username)
	AddActiveUrlForUser(url2, username)
	result1 := IsUrlActiveForUser(url1, username)
	result2 := IsUrlActiveForUser(url2, username)

	if result1 == false {
		t.Fatal("url1 was not active in the DB")
	}
	if result2 == false {
		t.Fatal("url2 was not active in the DB")
	}
}

func TestAddActiveUrlForUserForMultipleUrlsShowsThatBothUrlsCanBeListed(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url1 := "url1.com"
	url2 := "url2.com"
	username := "username"

	AddActiveUrlForUser(url1, username)
	AddActiveUrlForUser(url2, username)
	urlList, err := ListUrlsForUser(username)

	if err != nil {
		t.Fatalf("Error returned from listing user's urls: %v", err)
	}
	if len(urlList) != 2 {
		t.Fatalf("Expecting 2 elements in url list.  Actual list: %v", urlList)
	}
	if !contains(urlList, url1) {
		t.Fatalf("urlList does not contain url1! %v", url1)
	}
	if !contains(urlList, url2) {
		t.Fatalf("urlList does not contain url2! %v", url2)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestRemoveActiveUrlForUserForNonExistentUserDoesNotFail(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	RemoveActiveUrlForUser("url.com ", "username")

	localstorage.CleanDB()
}

func TestAddActiveUrlForUserAndRemoveActiveUrl(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "url.com"
	username := "username"

	AddActiveUrlForUser(url, username)
	RemoveActiveUrlForUser(url, username)
	result := IsUrlActiveForUser(url, username)

	if result == true {
		t.Fatal("Removed url is still active in the DB")
	}
}

func TestAddActiveUrlForUsersAndRemoveOneUrlDoesNotRemoveTheOther(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url1 := "url1.com"
	url2 := "url2.com"
	username := "testUser"

	AddActiveUrlForUser(url1, username)
	AddActiveUrlForUser(url2, username)
	RemoveActiveUrlForUser(url1, username)
	result1 := IsUrlActiveForUser(url1, username)
	result2 := IsUrlActiveForUser(url2, username)

	if result1 == true {
		t.Fatal("Removed url is still active in the DB")
	}
	if result2 == false {
		t.Fatal("Non removed url is no longer active in DB")
	}
}

func TestAddActiveUrlForOneUserDoesNotAddItForTheOtherUser(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "url.com"
	username1 := "testUser1"
	username2 := "testUser2"

	AddActiveUrlForUser(url, username1)
	user1UrlResult := IsUrlActiveForUser(url, username1)
	user2Urls, _ := ListUrlsForUser(username2)

	if len(user2Urls) > 0 {
		t.Fatalf("User2's url list should have been empty but was %v", user2Urls)
	}
	if user1UrlResult == false {
		t.Fatal("User1's list does not contain the url it added")
	}
}

func TestUrlAddedToBothUsersButRemovedFromOneShouldOnlyRemoveFromOne(t *testing.T) {
	localstorage.SetTestDb()
	defer localstorage.CleanDB()
	url := "url.com"
	username1 := "testUser1"
	username2 := "testUser2"

	AddActiveUrlForUser(url, username1)
	AddActiveUrlForUser(url, username2)
	RemoveActiveUrlForUser(url, username1)

	user1Result := IsUrlActiveForUser(url, username1)
	user2Result := IsUrlActiveForUser(url, username2)
	if user1Result == true {
		t.Fatal("Url should have been deleted from user 1, but it is still active")
	}
	if user2Result == false {
		t.Fatal("Url should not have been deleted from user 2, but it is no longer active")
	}
}
