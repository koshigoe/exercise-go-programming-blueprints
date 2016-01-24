package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL は Avatar インスタンスがアバターの URL を返す事ができない場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat: アバターの URL を取得できません。")
// Avatar はユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURL は指定されたクライアントのアバターの URL を返します。
	// 問題が発生した場合にはエラーを返します。特に、URL を取得できなかった場合に ErrNoAvatarURL を返します。
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct {}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct {}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if match, _ := filepath.Match(useridStr + "*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
