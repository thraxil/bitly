package bitly

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Connection struct {
	SSLHost     string
	AccessToken string
}

func NewConnection(accessToken string) *Connection {
	return &Connection{"api-ssl.bit.ly", accessToken}
}

func (c Connection) Shorten(uri string) (string, error) {
	params := url.Values{}
	params.Set("access_token", c.AccessToken)
	params.Set("longUrl", uri)
	params.Set("format", "txt")

	res, err := http.Get(c.shortenURL() + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(b), err
}

func (c Connection) shortenURL() string {
	return "https://" + c.SSLHost + "/v3/shorten"
}
