package ruptela

import (
	"strconv"

	"github.com/sandronister/socket-go/pkg/utils"
)

func (r *Device) ProcessHeader() TCommandProtocol {

	var (
		imei       uint64
		command    byte
		packageLen uint16
	)

	r.Read2Bu(&packageLen)
	r.Read8Bu(&imei)
	r.Read1B(&command)

	r.Header.Command = command
	r.Header.PackageLen = uint32(packageLen)
	r.Header.Imei = utils.PadLeft(strconv.FormatUint(imei, 10), "0", 15)

	retCmd := TCommandProtocol(command)

	if err := r.IsValidHeader(); err != nil {
		retCmd = COMMAND_UNKNOWN
		r.Success = false

	}

	return retCmd
}
