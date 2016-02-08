package cachet

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	urlparser "net/url"
)

func makeRequest(requestType string, url string, reqBody []byte) (*http.Response, []byte, error) {
	req, err := http.NewRequest(requestType, Config.APIUrl + url, bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Cachet-Token", Config.APIToken)

	client := &http.Client{}
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment}
	if Config.InsecureAPI == true {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	if Config.Proxy != "" {
		proxyUrl, err := urlparser.Parse(Config.Proxy)
		if err != nil {
			return nil, []byte{}, err
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	client.Transport = transport

	res, err := client.Do(req)
	if err != nil {
		return nil, []byte{}, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return res, body, nil
}
