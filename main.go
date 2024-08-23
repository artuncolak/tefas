package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TefasClient struct {
	baseUrl *url.URL
	client  *http.Client
}

const (
	baseUrl    = "https://www.tefas.gov.tr"
	referer    = "http://www.tefas.gov.tr/TarihselVeriler.aspx"
	dateFormat = "02-01-2006"
)

/*
New creates a new instance of TefasClient.
*/
func New() *TefasClient {
	parsedURL, _ := url.Parse(baseUrl)

	return &TefasClient{
		baseUrl: parsedURL,
		client:  &http.Client{},
	}
}

func (c *TefasClient) request(method, path string, urlParams url.Values) (*http.Response, error) {
	relPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	finalUrl := c.baseUrl.ResolveReference(relPath)

	bodyReader := strings.NewReader(urlParams.Encode())
	req, err := http.NewRequest(method, finalUrl.String(), bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Origin", baseUrl)
	req.Header.Add("Referer", referer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	return c.client.Do(req)
}

func (c *TefasClient) unmarshalResponseBody(body io.Reader, v interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		return err
	}

	return nil
}

/*
GetFundInfo retrieves fund information based on the provided FundInfoRequest.

It sends a POST request to the "/api/DB/BindHistoryInfo" endpoint with the specified parameters.

The function returns a slice of Fund objects and an error if any occurred during the request or response handling.
*/
func (c *TefasClient) GetFundInfo(fundInfoRequest FundInfoRequest) ([]Fund, error) {
	params := url.Values{}
	params.Add("fontip", string(fundInfoRequest.Type))
	params.Add("fonkod", strings.ToUpper(fundInfoRequest.Code))
	params.Add("bastarih", fundInfoRequest.StartDate.Format(dateFormat))
	params.Add("bittarih", fundInfoRequest.EndDate.Format(dateFormat))

	resp, err := c.request("POST", "/api/DB/BindHistoryInfo", params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fundInfoResponse FundInfoResponse
	err = c.unmarshalResponseBody(resp.Body, &fundInfoResponse)

	if err != nil {
		return nil, err
	}

	return fundInfoResponse.Data, nil
}
