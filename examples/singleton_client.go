package examples

import (
	"github.com/federicoleon/go-httpclient/gohttp"
	"time"
)

var (
	httpClient  = getHttpClient()
	httpClient2 = getHttpClient()
	c3          = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
