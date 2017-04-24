package httpClient

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL *url.URL
	client  *http.Client
}

func NewClient(host string) *Client {
	baseURL, err := url.Parse("http://" + host + ":32400")
	if err != nil {
		log.Fatal(err)
	}

	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &Client{
		BaseURL: baseURL,
		client:  &http.Client{Transport: transportConfig},
	}

	return client
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	relURL, err := url.Parse(path)
	u := c.BaseURL.ResolveReference(relURL)
	req, err := http.NewRequest(method, u.String(), body)
	return req, err
}

func (c *Client) NewGet(path string, body io.Reader) (*http.Request, error) {
	return c.newRequest("GET", path, body)
}

func (c *Client) NewPost(path string, body io.Reader) (*http.Request, error) {
	return c.newRequest("POST", path, body)
}

func (c *Client) NewDelete(path string, body io.Reader) (*http.Request, error) {
	return c.newRequest("DELETE", path, body)
}

func (c *Client) NewPut(path string, body io.Reader) (*http.Request, error) {
	return c.newRequest("PUT", path, body)
}

func (c *Client) Invoke(request *http.Request) (*http.Response, error) {
	return c.client.Do(request)
}
