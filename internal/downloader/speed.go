package download

import (
	"net/http"
	"time"
)

type RateLimitedTransport struct {
	http.RoundTripper
	SpeedLimit int64
}

func (r *RateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	res, err := r.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	contentLength := res.ContentLength
	if contentLength <= 0 {
		return res, nil
	}

	elapsed := time.Since(start)
	if int64(elapsed) < contentLength/r.SpeedLimit {
		time.Sleep(time.Duration(contentLength/r.SpeedLimit - int64(elapsed)))
	}

	return res, nil
}
