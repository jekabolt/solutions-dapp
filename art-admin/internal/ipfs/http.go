package ipfs

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

func initHTTPClient(timeout time.Duration) *fasthttp.Client {
	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, timeout)
		},
	}
}

func (ipfs *IPFS) makeURL(path string) string {
	ipfs.BaseURL.Path = path
	return ipfs.BaseURL.String()
}

func (ipfs *IPFS) post(path string, reqBody []byte, data interface{}) error {

	req := fasthttp.AcquireRequest()
	req.SetBody(reqBody)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetRequestURI(ipfs.makeURL(path))
	req.Header.Set("X-API-Key", ipfs.c.APIKey)

	res := fasthttp.AcquireResponse()
	if err := ipfs.cli.Do(req, res); err != nil {
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
