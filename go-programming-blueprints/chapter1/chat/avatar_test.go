package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/markbates/goth"
)

func TestAuthAvatar(t *testing.T) {
	// LOOK THERE IS NO INITIALIZATION HERE
	var authAvatar AuthAvatar

	// don't know what to do with that ???
	// testUser := &gothic_test.TestUser{}
	// testUser.On("AvatarURL").Return("", ErrNoAvatarURL)

	testGothUser := &goth.User{}
	// testGothUser.On("AvatarURL").Return("", ErrNoAvatarURL)

	testChatUser := &chatUser{User: testGothUser}

	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrorAvatarURL when no value present")
	}

	testURL := "http://url-to-gravatar/"

	testGothUser.AvatarURL = testURL

	url, err = authAvatar.GetAvatarURL(testChatUser)

	log.Println(testChatUser.GetAvatarURL(), url)

	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}

	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar

	testChatUser := &chatUser{uniqueID: "0bc83cb571cd1c50ba6f3e8a78ef1346"}

	url, err := gravatarAvatar.GetAvatarURL(testChatUser)

	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}

	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}

	// client := new(client)
	// client.UserData = map[string]interface{}{"Email": "MyEmailAddress@example.com", "UserID": "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	// url, err := gravatarAvatar.GetAvatarURL(client)
	// if err != nil {
	// 	t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	// }
	// if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
	// 	t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	// }
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()
	var fileSystemAvatar FileSystemAvatar

	testChatUser := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}

	// client := new(client)
	// client.UserData = map[string]interface{}{"UserID": "abc"}
	// url, err := fileSystemAvatar.GetAvatarURL(client)
	// if err != nil {
	// 	t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	// }
	// if url != "/avatars/abc.jpg" {
	// 	t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	// }
}
