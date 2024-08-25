package ruptela

import (
	"bytes"
	"encoding/gob"
)

func (r *Device) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
