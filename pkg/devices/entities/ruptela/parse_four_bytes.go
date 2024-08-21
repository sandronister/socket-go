package ruptela

import (
	"fmt"

	"github.com/sandronister/socket-go/pkg/devices/service/fmi"
)

func (d *Device) Parse4Bytes() {

	var id uint16
	var value4b int32
	var fourByteCount byte

	d.Read1B(&fourByteCount)

	for range fourByteCount {

		d.Read2Bu(&id)
		d.Read4B(&value4b)

		fmiField := fmi.GetFmi(id)
		thisLog := fmt.Sprintf(" \tFMI:[%04d] [%020d] \t[%s]", id, value4b, fmiField)

		switch id {
		case 65:
			d.Parameter.MileageVirt = float64(value4b) * 0.001

		case 77:
			d.Parameter.Mileage = float64(value4b) * 0.001

		case 92:
			d.Parameter.TotalFuelConsumptionHighResolution = float64(value4b) * 0.001

		case 114:
			d.Parameter.TotalMileage = float64(value4b) * 0.005

		case 203:
			d.Parameter.TotalTimeOfEngineOperation = float64(value4b) * 0.05

		case 208:
			d.Parameter.TotalFuelConsumption = float64(value4b) * 0.5

		case 484:
			d.Parameter.CanTotalFuelGaseous = float64(value4b) * 0.5

		default:
			thisLog += " \t\t\t <= not handled by go-parser"
		}

	}

}
