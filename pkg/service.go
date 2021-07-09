package pkg

import "errors"

/// Returned if the template layout file has not been found
var ErrCannotFindHTMLLayoutFile = errors.New("unable to find an HTML layout file")

/// Returned if the body is unreadable
var ErrCannotReadBody = errors.New("cannot read body input")

/// Returned if an internal template can't be parsed by the go HTML templates engine
var ErrCannotParseTemplate = errors.New("cannot parse internal template")

/// Returned when executing an internal HTML template failed
var ErrTemplateExecution = errors.New("error when executing internal template")
