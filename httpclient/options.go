package httpclient

import (
	"net/http"
)

type HeaderMap map[string]string

func WithHeader(header HeaderMap) ReqOption {
	return func(r *http.Request) {
		r.Header = r.Header.Clone()
		for k, v := range header {
			r.Header.Set(k, v)
		}
	}
}

func WithStream() ReqOption {
	return func(r *http.Request) {
		r.Header.Set("Accept", "text/event-stream")
	}
}

func WithTokenHeaderOption(token string) ReqOption {
	m := map[string]string{"Authorization": "Bearer " + token}
	return WithHeader(m)
}

// func WithTimeout(timeout time.Duration) ClientOption {
// 	return func(c *HTTPCli) {
// 		c.client.Timeout = timeout
// 	}
// }
