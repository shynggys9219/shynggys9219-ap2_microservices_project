package mongo

import (
	"fmt"
)

func (m Config) genConnectURL() string {
	var url string
	if m.Username == "" || m.Password == "" {
		url = fmt.Sprintf("mongodb://%s/?tls=%t&retryWrites=false", m.URI, m.TLSEnable)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s@%s/?tls=%t&retryWrites=false", m.Username, m.Password, m.URI, m.TLSEnable)
	}

	return url
}
