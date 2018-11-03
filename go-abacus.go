package goabacus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultSendTimeout = time.Second * 30
	baseURL            = "https://api.abacus.co"
	invoicesURL        = "invoices"
)

// Abacus The main struct of this package
type Abacus struct {
	ClientSecret string
	Timeout      time.Duration
}

// NewClient will create a Abacus client with default values
func NewClient(clientSecret string) *Abacus {
	return &Abacus{
		ClientSecret: clientSecret,
		Timeout:      defaultSendTimeout,
	}
}

// GetInvoices will return the invoices of the Client Secret
func (v *Abacus) GetInvoices(page int, limit int, lastUpdate time.Time) (*Invoices, error) {
	client := &http.Client{}
	client.CheckRedirect = checkRedirectFunc

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to build Abacus invoices: %v", err)
	}

	u.Path = invoicesURL
	urlStr := fmt.Sprintf("%v", u)

	r, err := http.NewRequest("GET", urlStr, nil)

	r.Header = http.Header(make(map[string][]string))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", v.ClientSecret))

	data := url.Values{}
	data.Add("limit", strconv.Itoa(limit))
	data.Add("page", strconv.Itoa(page))
	data.Add("lastUpdated", lastUpdate.Format("2006-01-02T15:04:05.00Z"))
	r.URL.RawQuery = data.Encode()

	fmt.Println(r.URL.RawQuery)

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to call Abacus invoices: %v", err)
	}

	if res.StatusCode == 200 {
		rawResBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read Abacus invoices: %v", err)
		}
		//test
		//fmt.Println("rawResBody", string(rawResBody))
		var resp Invoices
		err = json.Unmarshal(rawResBody, &resp)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal Abacus invoices: %v", err)
		}
		return &resp, nil

	}
	return nil, fmt.Errorf("Failed to get Abacus invoices: %s", res.Status)
}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	if req.Header.Get("Authorization") == "" {
		req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	}
	return nil
}
