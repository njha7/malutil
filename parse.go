package malutil

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

const (
	usersPage       = "http://myanimelist.net/users.php"
	userProfilePath = "/profile/"
)

// GetUsers returns a list of MAL users from /users.php
// It is assumed that usersPage is valid, UTF-8 encoded HTML
func GetUsers(usersPage io.ReadCloser) []string {
	defer usersPage.Close()
	usersList := make([]string, 20)
	parser := html.NewTokenizer(usersPage)
	for {
		tt := parser.Next()
		if tt == html.ErrorToken {
			break
		}
		token := parser.Token()
		if len(token.Attr) > 0 {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, userProfilePath) {
						usersList = append(usersList, attr.Val)
					}
				}
			}
		}
	}
	return usersList
}
