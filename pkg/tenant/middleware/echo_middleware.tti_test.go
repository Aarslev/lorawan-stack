// Copyright © 2019 The Things Industries B.V.

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestEchoMiddleware(t *testing.T) {
	testCases := []struct {
		desc string
		req  func() *http.Request
	}{
		{
			desc: "Forwarded Host",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.Header.Set("X-Forwarded-Host", "foo-bar.nz.cluster.ttn")
				return req
			},
		},
		{
			desc: "Host",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.Host = "foo-bar.identity.ttn"
				return req
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assertions.New(t).So(fromRequest(tc.req()).TenantID, should.Equal, "foo-bar")
		})
	}
}
