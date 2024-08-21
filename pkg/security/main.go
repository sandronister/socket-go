package security

import (
	"github.com/getrak/crc16"
)

const (
	identityHash = "2162249098780e7b30b48f0afebd7c82"
)

func MatchSecurityWord(tag string) bool {
	return tag == identityHash
}

func MatchCRC(byteArray []byte, len, crc uint16) bool {
	calcCrc := Crc(byteArray, len)
	return calcCrc == crc
}

func Crc(byteArray []byte, len uint16) uint16 {
	bufWithoutCRCAndPackedLen := byteArray[2 : len-2]
	table := crc16.MakeTable(crc16.CRC16_KERMIT)
	return crc16.Checksum(bufWithoutCRCAndPackedLen, table)
}
