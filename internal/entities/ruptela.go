package entities

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"unsafe"
)

type Header struct {
	Imei string
}

type KeyCountRuptela struct {
	Data  string
	Count int
}

type Ruptela struct {
	OriginalBuff []byte
	Md5Str       string
	Header       Header
	Buffer       *bytes.Buffer
	Ack          []byte
	Success      bool
	BytesReaded  uint32
}

func (r *Ruptela) SetMD5() {
	hasher := md5.New()
	hasher.Write(r.OriginalBuff)
	r.Md5Str = hex.EncodeToString(hasher.Sum(nil))
}

func (r *Ruptela) Read1B(output *uint8) {
	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read2Bu(output *uint16) {
	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read2B(output *int16) {
	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read4Bu(output *uint32) {

	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read4B(output *int32) {

	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read8Bu(output *uint64) {

	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(output)
	r.BytesReaded += uint32(sz)
}

func (r *Ruptela) Read8B(output *int64) {
	binary.Read(r.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(output)
	r.BytesReaded += uint32(sz)
}
