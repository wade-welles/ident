package main

import (
	"errors"
)

var (
	ErrMissingSeq  = errors.New("ident: dynamodb: missing seq after increment")
	ErrRanOutOfIDs = errors.New("ident: ran out of identifiers")
)
