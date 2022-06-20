package ipfs

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/valyala/fasthttp"
)

type Config struct {
	APIKey  string `env:"MORALIS_API_KEY"`
	Timeout string `env:"MORALIS_TIMEOUT" envDefault:"10s"`
	BaseURL string `env:"MORALIS_BASE_URL" envDefault:"https://deep-index.moralis.io/api/v2/"`
}

type Moralis struct {
	cli     *fasthttp.Client
	c       *Config
	BaseURL *url.URL
	desc    *descriptions.Store
}

func (c *Config) Init(desc *descriptions.Store) (*Moralis, error) {
	tOut, err := time.ParseDuration(c.Timeout)
	if err != nil && c.Timeout != "" {
		return nil, fmt.Errorf("InitMoralis:time.ParseDuration [%s]", err.Error())
	}
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("initSUTClient:url.Parse %s", err)
	}
	hc := initHTTPClient(tOut)

	return &Moralis{
		c:       c,
		cli:     hc,
		BaseURL: baseURL,
		desc:    desc,
	}, nil
}

func initHTTPClient(timeout time.Duration) *fasthttp.Client {
	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, timeout)
		},
	}
}

func (m *Moralis) makeURL(path string) string {
	m.BaseURL.Path = path
	return m.BaseURL.String()
}

func (m *Moralis) post(path string, reqBody []byte, data interface{}) error {

	req := fasthttp.AcquireRequest()
	req.SetBody(reqBody)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetRequestURI(m.makeURL(path))
	req.Header.Set("X-API-Key", m.c.APIKey)

	res := fasthttp.AcquireResponse()
	if err := m.cli.Do(req, res); err != nil {
		return fmt.Errorf("request:m.cli.Do %s", err.Error())
	}
	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("request:bad status code %d %s", res.StatusCode(), res.Body())
	}
	fasthttp.ReleaseRequest(req)

	defer fasthttp.ReleaseResponse(res)

	err := json.Unmarshal(res.Body(), &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal [%v]", err.Error())
	}
	return err
}
