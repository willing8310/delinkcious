package social_graph_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	om "github.com/willing8310/delinkcious/pkg/object_model"
)

func NewClient(baseURL string) (om.SocialGraphManager, error) {
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	followEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/follow"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse,
	).Endpoint()

	return EndpointSet{
		FollowEndpoint: followEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
