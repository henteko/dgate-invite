package main

import (
	"bytes"
	"crypto/tls"
	"github.com/mreiferson/go-httpclient"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func httpGet(uri string, params map[string]string) (*http.Response, error) {
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	defer transport.Close()

	client := &http.Client{Transport: transport}

	reqParams := paramsToReqestParams(params)
	return client.Get(uri + reqParams)
}

func httpPost(uri string, params map[string]string) (*http.Response, error) {
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	defer transport.Close()

	client := &http.Client{Transport: transport}

	values := url.Values{}
	for key, val := range params {
		values.Set(key, val)
	}
	return client.PostForm(uri, values)
}

func httpFilesPost(uri string, params map[string]string, filePaths map[string]string) (*http.Response, error) {
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	defer transport.Close()

	client := &http.Client{Transport: transport}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(key, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)
	return client.Do(req)
}

func httpDelete(uri string, params map[string]string) (*http.Response, error) {
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	defer transport.Close()

	client := &http.Client{Transport: transport}

	reqParams := paramsToReqestParams(params)
	req, _ := http.NewRequest("DELETE", uri+reqParams, nil)

	return client.Do(req)
}

func paramsToReqestParams(params map[string]string) string {
	var buffer bytes.Buffer
	for key, val := range params {
		buffer.WriteString("&" + key + "=" + val)
	}
	reqParams := buffer.String()

	re, _ := regexp.Compile("^&")
	return re.ReplaceAllString(reqParams, "?")
}
