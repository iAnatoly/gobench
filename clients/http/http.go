package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/iAnatoly/gobench/executor"
	"github.com/iAnatoly/gobench/executor/metrics"
)

type HttpClient struct {
	prefix string
	client *http.Client
}

type HttpClientConfig struct {
	dnsResolveOverride string
	verifyTLS          bool
}

func NewHttpClient(ctx context.Context, prefix string, config ...HttpClientConfig) (HttpClient, error) {
	group := metrics.Group{
		Name: "HTTP (" + prefix + ")",
		Graphs: []metrics.Graph{
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".http_ok",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_fail",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_other_fail",
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".latency",
						Type:  metrics.Histogram,
					},
				},
			},
		},
	}
	groups := []metrics.Group{
		group,
	}

	tr := &http.Transport{
		MaxIdleConnsPerHost: 300,
	}

	if len(config) > 0 {
		if config[0].dnsResolveOverride != "" {
			dialer := &net.Dialer{
				Timeout:   10 * time.Second, // TODO: make configurable in the config
				KeepAlive: 10 * time.Second,
				DualStack: true,
			}
			tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
				addr = config[0].dnsResolveOverride
				return dialer.DialContext(ctx, network, addr)
			}
		}
		tr.TLSClientConfig.InsecureSkipVerify = !config[0].verifyTLS
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10, // TODO: use this or the dialer
	}

	httpClient := HttpClient{
		prefix: prefix,
		client: client,
	}

	if err := executor.Setup(groups); err != nil {
		return httpClient, err
	}

	return httpClient, nil
}

func (h *HttpClient) do(method, url string, body []byte, headers map[string]string) (buf []byte, err error) {
	begin := time.Now()
	otherFail := h.prefix + ".http_other_fail"
	fail := h.prefix + ".http_fail"
	success := h.prefix + ".http_ok"
	latency := h.prefix + ".latency"
	var res *http.Response

	defer func() {
		diff := time.Since(begin)
		executor.Notify(latency, diff.Microseconds())
		if err != nil {
			executor.Notify(otherFail, 1)
			return
		}
		if res.StatusCode >= 300 || res.StatusCode < 200 {
			executor.Notify(fail, 1)
			err = fmt.Errorf("request failed with status code %d", res.StatusCode)
			return
		}
		executor.Notify(success, 1)
	}()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// add headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err = h.client.Do(req)
	if err != nil {
		return
	}

	// io.Copy(ioutil.Discard, res.Body)

	defer res.Body.Close()

	buf, err = ioutil.ReadAll(res.Body)

	return
}

// Get makes http get request and record the metrics
func (h *HttpClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return h.do("GET", url, nil, headers)
}

// Post makes http post request and record the metrics
func (h *HttpClient) Post(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return h.do("POST", url, body, headers)
}
