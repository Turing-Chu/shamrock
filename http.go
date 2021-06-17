// Author: Turing Zhu
// Date: 6/11/21 2:47 PM
// File: http.go

package shamrock

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
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
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status=%d", res.StatusCode)
	}

	contentEncoding := res.Header.Get("Content-Encoding")
	data := make([]byte, 0)
	switch contentEncoding {
	case "gzip":
		gzipReader, err := gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		data, err = ioutil.ReadAll(gzipReader)
	case "br":
	case "compress":
	case "deflate":

	default:
		data, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

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
	return &Response{
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
	}, nil
}
