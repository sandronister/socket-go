package ruptela

import customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"

func (d *Device) checks(paramsRecStr map[TDynamicProtocol]string, crc uint16) *customerrors.Error {
	if !d.CheckWord(paramsRecStr) {
		d.Ack = dyn_DispNotAllowed
		return &customerrors.ErrBlockedImei
	}

	if !d.checkCRC(crc) {
		d.Ack = dyn_DispNotAllowed
		return &customerrors.ErrInvalidCRC
	}

	if !d.IsValidSize() {
		d.Ack = dyn_DispNotAllowed
		return &customerrors.ErrInvalidSize
	}

	return nil
}

func (d *Device) ProcessDynamic() *customerrors.Error {

	var (
		version    uint8
		paramId    uint8
		paramLen   uint8
		paramCount uint8
		crc        uint16
	)

	d.Read1B(&version)
	d.Read1B(&paramCount)

	paramsRecStr := map[TDynamicProtocol]string{}

	for range int(paramCount) {
		d.Read1B(&paramId)
		d.Read1B(&paramLen)

		params := make([]uint8, paramLen)

		for j := range int(paramLen) {
			d.Read1B(&params[j])
		}

		cmd := TDynamicProtocol(paramId)
		paramsRecStr[cmd] = string(params)
	}

	return d.checks(paramsRecStr, crc)

}
