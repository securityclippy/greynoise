package client

import (
	"net/http"
	"time"
	"github.com/securityclippy/greynoise/conf"
	"io"
	"github.com/sirupsen/logrus"
	"path"
	"net/url"
	"fmt"
	"io/ioutil"
)

var (
	logger = logrus.WithField("client", "client")
)


type Client struct {
	NetClient	http.Client
	Header  	http.Header
	APIKey		string
	Endpoint	string
}

func NewClient(config conf.Config) (*Client) {
	c := http.Client{
		Timeout: time.Second * 5,

	}

	h := http.Header{}
	h.Set("Content-Type", "application/json")
	h.Set("user-agent", "greynoise-goclient")
	h.Set("key", config.APIKey)

	client := Client{
		NetClient: c,
		Endpoint: config.Endpoint,
		APIKey: config.APIKey,
		Header: h,
	}
	return &client
}



func (c Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header = c.Header
	return req, nil
}


func (c Client) buildURL(params map[string]string, paths ...string) string {
	reqURL := c.Endpoint
	urlPath := path.Join(paths...)
	reqURL = reqURL + urlPath
	if params != nil {
		q := url.Values{}
		for k, v := range params {
			q.Add(k, string(v))
		}
		logger.Infof("encoded values: %s", q.Encode())
		reqURL = fmt.Sprintf("%s?%s", reqURL, q.Encode())

	}
	logger.Infof("reqURL: %s", reqURL)
	return reqURL
}

func (c Client) doRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.NetClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c Client) ParseResponse(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
