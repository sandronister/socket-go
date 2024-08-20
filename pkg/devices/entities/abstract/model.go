package abstract

import (
	"bytes"
)

type Device struct {
	OriginalBuff []byte
	BytesReaded  uint32
	Buffer       *bytes.Buffer
}
