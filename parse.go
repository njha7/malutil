package malutil

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	userProfilePath = "/profile/"
	malusers        = "http://myanimelist.net/users.php"
)

// GetUsers returns a slice of MAL usernames from /users.php
func GetUsers() ([]string, error) {
	resp, err := http.Get(malusers)
	if err != nil {
		return nil, err
	}
	return GetUsersFromPage(resp.Body), nil
}

// GetUsersFromPage returns a slice of MAL usernames from /users.php
// It is assumed that usersPage is valid, UTF-8 encoded HTML
func GetUsersFromPage(usersPage io.ReadCloser) []string {
	defer usersPage.Close()
	// MAL users page has a 5 x 4 table of users
	usersList := make([]string, 0, 20)
	usersSet := make(map[string]bool, 20)
	parser := html.NewTokenizer(usersPage)
	for {
		tt := parser.Next()
		// EOF is represented as an Error
		if tt == html.ErrorToken {
			break
		}
		token := parser.Token()
		if len(token.Attr) > 0 {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, userProfilePath) {
						user := strings.Split(attr.Val, "/profile/")[1]
						_, ok := usersSet[user]
						if !ok {
							usersList = append(usersList, user)
							usersSet[user] = true
						}
					}
				}
			}
		}
	}
	return usersList
}
