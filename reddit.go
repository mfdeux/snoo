// Package snoo provides Reddit API wrapper utilities.
package snoo

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func newHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func newHTTPClientWithProxy(timeout int, proxyURL string) (*http.Client, error) {
	parsedProxyURL, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(parsedProxyURL),
	}
	netClient := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: transport,
	}
	return netClient, nil
}

const (
	baseAuthURL string = "https://oauth.reddit.com"
	baseURL     string = "http://reddit.com"
	// UpVote is an upvote for a submission or comment
	UpVote int = 1
	// DownVote is an upvote for a submission or comment
	DownVote int = -1
	// NoVote is no vote for a submission or comment
	NoVote int = 0
	// HTTPTimeout in seconds
	HTTPTimeout int = 5
)

// Client is the client for interacting with the Reddit API.
type Client struct {
	http      *http.Client
	userAgent string
}

func NewPublicClient(userAgent string) *Client {
	httpClient := newHTTPClient(HTTPTimeout)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}
}

func NewPublicClientWithProxy(userAgent string, proxyURL string) (*Client, error) {
	httpClient, err := newHTTPClientWithProxy(HTTPTimeout, proxyURL)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}, err
}

func NewPreAuthorizedClient(accessKey, refreshKey, userAgent string) *Client {
	httpClient := newHTTPClient(HTTPTimeout)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}
}

func NewPreAuthorizedClientWithProxy(accessKey, refreshKey, userAgent, proxyURL string) (*Client, error) {
	httpClient, err := newHTTPClientWithProxy(HTTPTimeout, proxyURL)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}, err
}

func (c *Client) commentOnThing(thingID string, text string) error {
	data := url.Values{}
	data.Set("thing_id", thingID)
	data.Set("text", text)
	data.Set("api_type", "json")
	url := fmt.Sprintf("%s/api/comment", baseAuthURL)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) deleteThing(thingID string) error {
	data := url.Values{}
	data.Set("id", thingID)
	url := fmt.Sprintf("%s/api/del", baseAuthURL)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) editThingText(thingID string, text string) error {
	data := url.Values{}
	data.Set("thing_id", thingID)
	data.Set("text", text)
	data.Set("api_type", "json")
	url := fmt.Sprintf("%s/api/editusertext", baseAuthURL)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) voteOnThing(thingID string, direction int) error {
	data := url.Values{}
	data.Set("thing_id", thingID)
	data.Set("dir", strconv.Itoa(direction))
	data.Set("api_type", "json")
	url := fmt.Sprintf("%s/api/vote", baseAuthURL)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) saveThing(thingID string, category string) error {
	data := url.Values{}
	data.Set("thing_id", thingID)
	data.Set("category", category)
	data.Set("api_type", "json")
	url := fmt.Sprintf("%s/api/save", baseAuthURL)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}
