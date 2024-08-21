package ruptela

import (
	"fmt"

	"github.com/sandronister/socket-go/pkg/devices/service/fmi"
)

func (d *Device) Parse2Bytes() {

	var id uint16
	var value2b int16
	var twoBytesNum byte

	d.Read1B(&twoBytesNum)

	for range twoBytesNum {

		d.Read2Bu(&id)
		d.Read2B(&value2b)

		fmiField := fmi.GetFmi(id)
		thisLog := fmt.Sprintf(" \tFMI:[%04d] [%020d] \t[%s]", id, value2b, fmiField)

		switch id {
		case 29:
			d.Parameter.Voltage = float64(value2b)

		case 30:
			d.Parameter.BateryLife = int32(value2b)

		case 41:
			d.Parameter.AxleLoad_1 = float64(value2b) * 0.5

		case 52:
			d.Parameter.AxleLoad_2 = float64(value2b) * 0.5

		case 53:
			d.Parameter.AxleLoad_3 = float64(value2b) * 0.5

		case 54:
			d.Parameter.AxleLoad_4 = float64(value2b) * 0.5

		case 55:
			d.Parameter.AxleLoad_5 = float64(value2b) * 0.5

		case 66:
			d.Parameter.Lls1Value = int32(value2b)

		case 67:
			d.Parameter.Lls5Value = int32(value2b)

		case 68:
			d.Parameter.Lls2Value = int32(value2b)

		case 69:
			d.Parameter.Lls3Value = int32(value2b)

		case 90:
			d.Parameter.EventOfInstantaneousFuelEconomy = int32(float64(value2b) / 512)

		case 197:
			d.Parameter.Rpm = float64(value2b) * 0.125

		case 204:
			d.Parameter.MileageUntilTheNextMaintenance = int32((float64(value2b) * 5) - 16063)

		case 211:
			d.Parameter.Lls4Value = int32(value2b)

		case 213:
			d.Parameter.CanSpeed = int32(float64(value2b) / 256)

		case 419:
			d.Parameter.Height = int32(value2b)

		default:
			thisLog += " \t\t <= not handled by go-parser"
		}

	}

}
