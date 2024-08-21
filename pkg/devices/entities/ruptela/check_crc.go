package ruptela

import "github.com/sandronister/socket-go/pkg/security"

func (d *Device) checkCRC(crc uint16) bool {
	return security.MatchCRC(d.OriginalBuff, uint16(d.BytesReaded), crc)
}
