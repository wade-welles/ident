package dynamodb

import (
	"errors"
)

var (
	ErrMissingSeq  = errors.New("dynamodb: missing seq after increment")
	ErrRanOutOfIDs = errors.New("dynamodb: ran out of identifiers")
)
