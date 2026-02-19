package ulid

import "github.com/oklog/ulid/v2"

func Make() string {
	return ulid.Make().String()
}
