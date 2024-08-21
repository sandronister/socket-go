package ruptela

import (
	"fmt"

	"github.com/sandronister/socket-go/pkg/devices/service/fmi"
)

func (d *Device) ParseOneByte() {
	var (
		id           uint16
		value1b      byte
		oneByteCount byte
	)

	d.Read1B(&oneByteCount)

	for range oneByteCount {
		d.Read2Bu(&id)
		d.Read1B(&value1b)

		fmiField := fmi.GetFmi(id)
		thisLog := fmt.Sprintf(" \tFMI:[%04d] [%020d] \t[%s]", id, value1b, fmiField)

		switch id {
		case 5:
			d.Parameter.IgnitionKey = int32(value1b)

		case 27:
			d.Parameter.GsmStatus = int32(value1b)

		case 35:
			d.Parameter.ClutchPedalStatus = int32(value1b)

		case 36:
			d.Parameter.ServiceBrakePedalStatus = int32(value1b)

		case 37:
			d.Parameter.CruiseControlStatus = int32(value1b)

		case 49:
			d.Parameter.AccelerometerX = int32(value1b)

		case 50:
			d.Parameter.AccelerometerY = int32(value1b)

		case 51:
			d.Parameter.AccelerometerZ = int32(value1b)

		case 56:
			d.Parameter.Lls1Temp = int32(value1b)

		case 57:
			d.Parameter.Lls2Temp = int32(value1b)

		case 58:
			d.Parameter.Lls3Temp = int32(value1b)

		case 61:
			d.Parameter.Lls6Temp = int32(value1b)

		case 76:
			d.Parameter.Lls5Temp = int32(value1b)

		case 115:
			d.Parameter.CoolantTemperature = int32(value1b) - 40

		case 176:
			d.Parameter.GpsSpeed = int32(value1b)

		case 193:
			d.Parameter.Input_1 = int32(value1b)

		case 194:
			d.Parameter.Input_2 = int32(value1b)

		case 195:
			d.Parameter.Input_3 = int32(value1b)

		case 196:
			d.Parameter.Input_4 = int32(value1b)

		case 201:
			d.Parameter.Adblue = float64(value1b) * 0.4

		case 206:
			d.Parameter.AccelerationPedalPosition = float64(value1b) * 0.4

		case 207:
			d.Parameter.CanFuelLevel = float64(value1b) * 0.4
			d.Parameter.Lls6Value = int32(float64(value1b) * 0.4)

		case 212:
			d.Parameter.Lls4Temp = int32(value1b)

		case 219:
			d.Parameter.AxleIndex = int32(value1b)

		case 362:
			d.Parameter.ParkingBrakeStatus = int32(value1b)

		case 365:
			d.Parameter.DoorStatus = int32(value1b)

		case 367:
			d.Parameter.Gear = float64(value1b) - 125

		case 418:
			d.Parameter.GpsStatus = int32(value1b)

		case 482:
			d.Parameter.ServiceBrakePedalPosition = float64(value1b) * 0.4

		default:
			thisLog += " \t\t <= not handled by go-parser"
		}
	}
}
