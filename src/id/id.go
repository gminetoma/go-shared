package id

import "github.com/gminetoma/go-shared/src/ulid"

func Make() string {
	return ulid.Make()
}
