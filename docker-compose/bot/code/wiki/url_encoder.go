package wiki

import (
	"net/url"
)

//urlEncoded encodes a string to be used in a query part of a URL
func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}