// Sina weibo SDK in golang.
// Implemented based on the OAuth2.0 protocal.
package weibo

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
)

const OAUTH_SERVER = "https://api.weibo.com"

// Get your ApiKey, ApiSecret from http://api.weibo.com.
// Don't forget to set a RedirectUri.
type Client struct {
	ApiKey      string
	ApiSecret   string
	RedirectUri string

	accessToken string
}

// Get the authentication url you should redirect your user to.
func (c *Client) GetAuthUrl() (string, error) {

	baseUrlString := fmt.Sprintf("%s/oauth2/authorize", OAUTH_SERVER)
	baseUrl, _ := url.Parse(baseUrlString)
	parmas := url.Values{}
	parmas.Add("client_id", c.ApiKey)
	parmas.Add("response_type", "code")
	parmas.Add("redirect_uri", c.RedirectUri)

	baseUrl.RawQuery = parmas.Encode()
	return baseUrl.String(), nil

}

// Once the user logined, Sina will redirect him to RedirectUri with a code.
func (c *Client) GetAccessToken(code string) (string, error) {

	baseUrlString := fmt.Sprintf("%s/oauth2/access_token", OAUTH_SERVER)

	baseUrl, _ := url.Parse(baseUrlString)

	parmas := url.Values{}
	parmas.Add("grant_type", "authorization_code")
	parmas.Add("redirect_uri", c.RedirectUri)
	parmas.Add("code", code)
	parmas.Add("client_id", c.ApiKey)
	parmas.Add("client_secret", c.ApiSecret)

	baseUrl.RawQuery = parmas.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", baseUrl.String(), nil)

	res, err := client.Do(req)

	if err != nil {
		return res.Status, err
	}

	defer res.Body.Close()

	json, _ := simplejson.NewFromReader(res.Body)

	c.accessToken, err = json.Get("access_token").String()

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}

// Do GET request
func (c *Client) Get(endpoint string, params url.Values) (*simplejson.Json, error) {
	baseUrlString := fmt.Sprintf("%s/2/%s.json", OAUTH_SERVER, endpoint)

	baseUrl, _ := url.Parse(baseUrlString)

	params.Add("access_token", c.accessToken)

	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())

	defer res.Body.Close()

	if err != nil {
		return nil, error(err)
	}

	json, _ := simplejson.NewFromReader(res.Body)

	return json, nil
}

// Do POST request
func (c *Client) Post(endpoint string, params url.Values) (*simplejson.Json, error) {
	baseUrlString := fmt.Sprintf("%s/2/%s.json", OAUTH_SERVER, endpoint)

	baseUrl, _ := url.Parse(baseUrlString)

	params.Add("access_token", c.accessToken)

	res, err := http.PostForm(baseUrl.String(), params)

	defer res.Body.Close()

	if err != nil {
		return nil, error(err)
	}

	json, _ := simplejson.NewFromReader(res.Body)

	return json, nil
}
