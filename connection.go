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
	apiURL := c.shortenURL()
	params := url.Values{}

	params.Set("access_token", c.AccessToken)
	params.Set("longUrl", uri)
	params.Set("format", "txt")

	requestURL := apiURL + "?" + params.Encode()
	res, err := http.Get(requestURL)
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
