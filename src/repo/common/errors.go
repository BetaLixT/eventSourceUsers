package common

import (
	"fmt"
)

type RepoError struct {
  code int
}

func (err *RepoError) Error() string {
  return fmt.Sprintf("RepoError - %d", err.code)
}

var _ error = (*RepoError)(nil)

const (
  ERROR_STREAM_CONFLICT = 1000
)
