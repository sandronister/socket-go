package ruptela

import (
	"fmt"

	"github.com/sandronister/socket-go/pkg/devices/service/fmi"
)

func (d *Device) Parse8Bytes() {

	var id uint16
	var value8b int64
	var eightBytesCount byte

	d.Read1B(&eightBytesCount)

	for i := 0; i < int(eightBytesCount); i++ {

		d.Read2Bu(&id)
		d.Read8B(&value8b)

		fmiField := fmi.GetFmi(id)
		thisLog := fmt.Sprintf(" \tFMI:[%04d] [%020d] \t[%s]", id, value8b, fmiField)

		//TODO
		switch id {
		// case 65:
		// 	parameter.MileageVirt = float64(value4b) * 0.001

		// case 77:
		// 	parameter.Mileage = float64(value4b) * 0.001

		// case 92:
		// 	parameter.TotalFuelConsumptionHighResolution = float64(value4b) * 0.001

		// case 114:
		// 	parameter.TotalMileage = float64(value4b) * 0.005

		// case 203:
		// 	parameter.TotalTimeOfEngineOperation = float64(value4b) * 0.05

		// case 208:
		// 	parameter.TotalFuelConsumption = float64(value4b) * 0.5

		default:
			thisLog += " \t\t <= not handled by go-parser"
		}

		// fmt.Println(thisLog)
	}
}
