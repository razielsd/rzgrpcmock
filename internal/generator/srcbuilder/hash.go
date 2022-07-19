package srcbuilder

import (
	"crypto/md5"
	"fmt"
)

func MakeHash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data)) //nolint:gosec
}
