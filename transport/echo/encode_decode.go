package echo

import (
	"context"

	"github.com/labstack/echo/v4"
)

// EncodeResponseFunc encodes the passed response object to the HTTP response
// writer. It's designed to be used in HTTP servers, for server-side
// endpoints. One straightforward EncodeResponseFunc could be something that
// JSON encodes the object directly to the response body.
type EncodeResponseFunc[Res any] func(context.Context, echo.Context, Res) error
