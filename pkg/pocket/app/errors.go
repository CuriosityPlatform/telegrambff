package app

import "errors"

var (
	ErrURLFailedToProceed       = errors.New("failed to proceed url")
	ErrFailedToParseForMetadata = errors.New("failed to parse for metadata")
)
