package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	base   string
	jar    http.CookieJar
	client *http.Client
}

func NewClient(base string) *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		base: base,
		jar:  jar,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Jar:     jar,
		},
	}
}

func (c *Client) Warmup(path string) error {
	req, _ := http.NewRequest("GET", c.base+path, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("warmup status %d", resp.StatusCode)
	}
	return nil
}

func (c *Client) cookie(name string) (string, bool) {
	u, _ := url.Parse(c.base)
	for _, ck := range c.jar.Cookies(u) {
		if ck.Name == name {
			return ck.Value, true
		}
	}
	return "", false
}

func (c *Client) PostJSON(path, refererPath string, payload any) ([]byte, error) {
	b, _ := json.Marshal(payload)
	
	token, ok := c.cookie("_EXPOSURE_TOKEN_")
	if !ok || token == "" {
		return nil, fmt.Errorf("missing _EXPOSURE_TOKEN_ after warmup")
	}
	
	req, _ := http.NewRequest("POST", c.base+path, bytes.NewReader(b))
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Exposure-Token", token)
	req.Header.Set("Origin", c.base)
	req.Header.Set("Referer", c.base+refererPath)
	
	resp, err := c.client.Do(req)
	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	out, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(out))
	}
	return out, nil
}

func main() {
	
	apiPath := "/youth-basketball-events"
	referer := "/youth-basketball-events"
	
	c := NewClient("https://basketball.exposureevents.com")
	err := c.Warmup(apiPath)
	if err != nil {
		panic(err)
	}
	
	reqBody := &SearchReq{
		SearchToken:     "",
		StartDateString: "8/10/2025", // MM/DD/YYYY
		Page:            1,
		SportType:       "1",
	}
	fmt.Println("Outside loop")
	out, err := c.PostJSON(apiPath, referer, reqBody)
	if err != nil {
		panic(err)
	}
	
	var getResult GetResult
	if err := json.Unmarshal(out, &getResult); err != nil {
		panic(err)
	}
	
	var finalResults []Event
	
	for i := 1; i <= getResult.Total; i++ {
		reqBody := &SearchReq{
			Gender:          "-1",
			InviteType:      "0",
			Page:            i,
			SportType:       "1",
			StartDateString: "1/1/2025",
			EndDateString:   "12/31/2025",
		}
		fmt.Printf("Fetching page %d\n", i)
		resp, err := c.PostJSON(apiPath, referer, reqBody)
		if err != nil {
			panic(err)
		}
		var results Results
		if err := json.Unmarshal(resp, &results); err != nil {
			panic(err)
		}
		
		for _, event := range results.Results {
			finalResults = append(finalResults, event)
		}
	}
	
	for _, event := range finalResults {
		fmt.Printf("Event ID: %d, Name: %s\n", event.ID, event.Name)
	}
}
