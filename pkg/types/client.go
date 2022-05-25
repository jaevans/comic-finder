package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("target", "https://www.comicvine.com/api")
}

type ComicVineClient struct {
	apiKey  string
	baseUrl string
}

func NewComicVineClient(apiKey string) *ComicVineClient {
	cvApi := &ComicVineClient{
		apiKey:  apiKey,
		baseUrl: viper.GetString("target"),
	}
	return cvApi
}

func (a *ComicVineClient) genericGetPaged(resource ResourceType, opts GetOptions) (Response, bool, error) {
	fields := strings.Join(opts.Fields, ",")
	filters := strings.Join(opts.Filters, ",")
	var getType string
	resp := Response{}
	switch resource {
	case ResourceTypeVolume:
		getType = "volumes"
	case ResourceTypeIssue:
		getType = "issues"
	default:
		return resp, false, fmt.Errorf("unknown type for genericGet")
	}
	getUrl := fmt.Sprintf("%s/%s?format=json&api_key=%s&field_list=%s&limit=%d&offset=%d&sort=%s&filter=%s", a.baseUrl, getType, a.apiKey, fields, opts.Limit, opts.Offset, opts.Sort, filters)
	response, err := http.Get(getUrl)
	if err != nil {
		return resp, false, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return resp, false, err
	}
	err = json.Unmarshal(responseData, &resp)
	// var more bool
	more := resp.Offset+resp.PageResults < resp.TotalResults
	return resp, more, err
}

func (a *ComicVineClient) genericGet(resource ResourceType, opts GetOptions) (Response, error) {
	fields := strings.Join(opts.Fields, ",")
	filters := strings.Join(opts.Filters, ",")
	var getType string
	resp := Response{}
	switch resource {
	case ResourceTypeVolume:
		getType = "volumes"
	case ResourceTypeIssue:
		getType = "issues"
	default:
		return resp, fmt.Errorf("unknown type for genericGet")
	}
	getUrl := fmt.Sprintf("%s/%s?format=json&api_key=%s&field_list=%s&limit=%d&offset=%d&sort=%s&filter=%s", a.baseUrl, getType, a.apiKey, fields, opts.Limit, opts.Offset, opts.Sort, filters)
	response, err := http.Get(getUrl)
	if err != nil {
		return resp, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(responseData, &resp)
	return resp, err
}
