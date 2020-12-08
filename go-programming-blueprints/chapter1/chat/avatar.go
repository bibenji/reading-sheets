package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL is the error returned when Avatar instance is unable to provide an avatar URL
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing user profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars to try each strategy
type TryAvatars []Avatar

// GetAvatarURL for TryAvatars
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar struct of AuthAvatar
type AuthAvatar struct{}

// UserAuthAvatar var of AuthAvatar for user
var UserAuthAvatar AuthAvatar

// GetAvatarURL do the thing for user
// func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
// 	if url, ok := c.UserData["AvatarURL"]; ok {
// 		if urlStr, ok := url.(string); ok {
// 			return urlStr, nil
// 		}
// 	}
// 	return "", ErrNoAvatarURL
// }

// GetAvatarURL do the thing for user
// refacto with line of sight, the happy at the end and not too much if else imbricated
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.GetAvatarURL()

	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}

	return url, nil

	// if !ok {
	// 	return "", ErrNoAvatarURL
	// }

	// urlStr, ok := url.(string)
	// if !ok {
	// 	return "", ErrNoAvatarURL
	// }

	// return urlStr, nil
}

// GravatarAvatar struc for gravatar
type GravatarAvatar struct{}

// UserGravatar gravatar for user
var UserGravatar GravatarAvatar

// GetAvatarURL for gravatar do the thing for the user
// func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
// 	if email, ok := c.UserData["Email"]; ok {
// 		if emailStr, ok := email.(string); ok {
// 			m := md5.New()
// 			io.WriteString(m, strings.ToLower(emailStr))
// 			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
// 		}
// 	}

// 	return "", ErrNoAvatarURL
// }

// GetAvatarURL for gravatar do the thing for the user (with refacto)
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	uniqueID := u.UniqueID()

	if len(uniqueID) == 0 {
		return "", ErrNoAvatarURL
	}

	return fmt.Sprintf("//www.gravatar.com/avatar/%s", uniqueID), nil

	// userID, ok := c.UserData["UserID"]
	// if !ok {
	// 	return "", ErrNoAvatarURL
	// }

	// userIDStr, ok := userID.(string)
	// if !ok {
	// 	return "", ErrNoAvatarURL
	// }

	// return fmt.Sprintf("//www.gravatar.com/avatar/%s", userIDStr), nil
}

// FileSystemAvatar a file system avatar
type FileSystemAvatar struct{}

// UserFileSystemAvatar a user file system avatar
var UserFileSystemAvatar FileSystemAvatar

// GetAvatarURL getAvatarURL for FileSystemAvatar
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}

	return "", ErrNoAvatarURL

	// if userid, ok := c.UserData["UserID"]; ok {
	// 	if useridStr, ok := userid.(string); ok {
	// 		files, err := ioutil.ReadDir("avatars")
	// 		if err != nil {
	// 			return "", ErrNoAvatarURL
	// 		}
	// 		for _, file := range files {
	// 			if file.IsDir() {
	// 				continue
	// 			}
	// 			if match, _ := path.Match(useridStr+"*", file.Name()); match {
	// 				return "/avatars/" + file.Name(), nil
	// 			}
	// 		}
	// 	}
	// }
	// return "", ErrNoAvatarURL
}
