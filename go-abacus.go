package goabacus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

	data := url.Values{}
	data.Add("limit", string(limit))
	data.Add("page", string(page))
	data.Add("lastUpdated", lastUpdate.Format(time.RFC3339))

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = invoicesURL
	urlStr := fmt.Sprintf("%v", u)

	r, err := http.NewRequest("GET", urlStr, nil)

	r.Header = http.Header(make(map[string][]string))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", v.ClientSecret)

	r.URL.RawQuery = data.Encode()

	fmt.Println("URL", r.URL.String())

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("rawResBody", string(rawResBody))
	fmt.Println("urlStr", urlStr)

	return nil, nil

}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	if req.Header.Get("Authorization") == "" {
		req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	}
	return nil
}
