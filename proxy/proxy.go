package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"rhea/errors"
	responsestore "rhea/response-store"
)

type Proxy struct {
	TargetRawUrl string
	Store        *responsestore.ResponseStore
}

func (p *Proxy) Handler(c *gin.Context) {
	req := c.Request
	if data, err := p.Store.GetData(req.URL.Path); err == nil {
		for k, v := range data.Headers {
			c.Header(k, v)
		}
		_, _ = c.Writer.Write(data.Content)
		return
	}
	proxy, err := url.Parse(p.TargetRawUrl)
	if err != nil {
		c.JSON(500, errors.ParseUrlError{Url: p.TargetRawUrl})
		return
	}
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	urlValue, _ := json.Marshal(req.URL)
	fmt.Println(string(urlValue))

	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(req)

	if err != nil {
		c.JSON(500, errors.RoundtripError{})
		return
	}

	headers := map[string]string{}
	for k, vv := range resp.Header {
		for _, v := range vv {
			headers[k] = v
			c.Header(k, v)
		}
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	p.Store.Add(req.URL.Path, &responsestore.ResponseData{
		Headers: headers,
		Content: bodyBytes,
	})
	_, _ = c.Writer.Write(bodyBytes)
	return
}
