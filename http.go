// Author: Turing Zhu
// Date: 6/11/21 2:47 PM
// File: http.go

package shamrock

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"github.com/andybalholm/brotli"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	client = http.Client{
		Transport: &http.Transport{
			// DisableKeepAlives:  true,
			// DisableCompression: true,
		},
	}
)

type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           http.Header
	Body             []byte
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Trailer          http.Header
}

// send http request
func Request(method, url string, headers map[string]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	contentEncoding := res.Header.Get("Content-Encoding")
	data, err := parseHttpRespBody(res.Body, contentEncoding)
	if err != nil {
		return nil, err
	}

	// Content-Type: application/javascript; charset=GB18030
	if strings.Contains(res.Header.Get("Content-Type"), "GB18030") {
		decoder := simplifiedchinese.GB18030.NewDecoder()
		tmpData, err := decoder.Bytes(data)
		if err != nil {
			return nil, err
		} else {
			data = tmpData
		}
	}

	resp := &Response{
		Status:           res.Status,
		StatusCode:       res.StatusCode,
		Proto:            res.Proto,
		ProtoMajor:       res.ProtoMajor,
		ProtoMinor:       res.ProtoMinor,
		Header:           res.Header,
		Body:             data,
		ContentLength:    res.ContentLength,
		TransferEncoding: res.TransferEncoding,
		Close:            res.Close,
		Uncompressed:     res.Uncompressed,
		Trailer:          res.Trailer,
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("status=%d", res.StatusCode)
	}
	return resp, err
}

func parseHttpRespBody(input io.ReadCloser, contentEncoding string) (data []byte, err error) {
	data = make([]byte, 0)
	switch contentEncoding {
	case "gzip":
		gzipReader, err := gzip.NewReader(input)
		if err != nil {
			return nil, err
		}
		data, err = ioutil.ReadAll(gzipReader)
	case "br":
		brReader := brotli.NewReader(input)
		if brReader == nil {
			return nil, fmt.Errorf("create br reader failed")
		}
		data, err = ioutil.ReadAll(brReader)
	case "compress":

	case "deflate":
		deflateReader := flate.NewReader(input)
		if deflateReader == nil {
			return nil, fmt.Errorf("create deflate reader failed")
		}
		data, err = ioutil.ReadAll(deflateReader)
	default:
		data, err = ioutil.ReadAll(input)
		if err != nil {
			return nil, err
		}

	}
	return data, nil
}
