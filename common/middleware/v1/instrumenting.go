package v1

import (
	"fmt"
	"time"
)

type Instrumenting struct {
	Client DataDogClient
}

type DataDogClient interface {
	Incr(name string, tags []string, rate float64) error
	Timing(name string, value time.Duration, tags []string, rate float64) error
}

//Latency measure request latency
func (mw Instrumenting) Latency(begin time.Time, url, method string) {
	if mw.Client == nil {
		return
	}
	// nolint
	mw.Client.Timing("http.request.duration", time.Since(begin), []string{
		fmt.Sprintf("url: %s", url),
		fmt.Sprintf("method: %s", method),
	}, 1)
}

//Counter increment request and error counter
func (mw Instrumenting) Counter(err error, url, method string) {
	if mw.Client == nil {
		return
	}

	mw.Client.Incr("http.request.total", []string{
		fmt.Sprintf("url: %s", url),
		fmt.Sprintf("method: %s", method),
	}, 1)

	if err != nil {
		//nolint
		mw.Client.Incr("http.request.error", []string{
			fmt.Sprintf("url: %s", url),
			fmt.Sprintf("method: %s", method),
		}, 1)
	}
}
