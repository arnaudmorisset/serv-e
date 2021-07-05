package pkg

import "errors"

/// Returned if the template layout file has not been found
var ErrCannotFindHTMLLayoutFile = errors.New("unable to find an HTML layout file")
