//go:build unit

package echo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	echotransport "github.com/kikihakiem/gkit/transport/echo"
	"github.com/labstack/echo/v4"
)

func TestSetResponseHeader(t *testing.T) {
	const (
		key = "X-Foo"
		val = "12345"
	)

	req := httptest.NewRequest(http.MethodPost, "/dummy", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	echotransport.SetResponseHeader(key, val)(context.Background(), c)
	if want, have := val, c.Response().Header().Get(key); want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}

func TestSetContentType(t *testing.T) {
	const contentType = echo.MIMEApplicationJSON
	req := httptest.NewRequest(http.MethodPost, "/dummy", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	echotransport.SetContentType(contentType)(context.Background(), c)
	if want, have := contentType, c.Response().Header().Get("Content-Type"); want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}

func TestSetRequestHeader(t *testing.T) {
	const contentType = echo.MIMEApplicationJSON
	req := httptest.NewRequest(http.MethodPost, "/dummy", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	echotransport.SetRequestHeader(echo.HeaderAccept, echo.MIMEApplicationJSON)(context.Background(), c)
	if want, have := contentType, c.Request().Header.Get("Accept"); want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}
