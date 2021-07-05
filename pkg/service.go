package pkg

import "errors"

/// Returned if the template layout file has not been found
var ErrCannotFindHTMLLayoutFile = errors.New("unable to find an HTML layout file")

/// Returned if the body is unreadable
var ErrCannotReadBody = errors.New("cannot read body input")
