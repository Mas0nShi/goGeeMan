package MHttp

import (
	"io/ioutil"
	"net/http"
)

type MHttp struct {}
type response struct {
	body  []byte
	httpCode int
	headers  http.Header
	cookies  []*http.Cookie
}
func (res response) GetHttpCode() int {
	return res.httpCode
}
func (res response) GetResponseBody() []byte {
	return res.body
}
func (res response) GetResponseText() string {
	return Bytes2str(res.body)
}
func (res response) GetResponseHeader(key string) []string {
	return res.headers[key]
}
func (res response) GetResponseHeaders() map[string][]string {
	return res.headers
}
func (res response) GetCookie(key string) string {
	for _, cookie := range res.cookies {
		if cookie.Name == key {
			return cookie.String()
		}
	}
	return ""
}
func (res response) GetCookies() string {
	cookies := ""
	for _, cookie := range res.cookies {
		cookies += cookie.String()
	}
	return cookies
}

func (h *MHttp) Get(url string, headers *map[string]string, cookies *map[string]string) response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// set headers
	if headers!=nil && len(*headers) > 0 {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	// set cookies
	if cookies!=nil && len(*cookies) > 0 {
		for key, value := range *cookies {
			req.AddCookie(&http.Cookie{Name: key,Value: value, HttpOnly: true})
		}
	}

	// send http requests
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// read response
	resp, err := ioutil.ReadAll(res.Body)
	return response{resp, res.StatusCode, res.Header, res.Cookies()}
}

