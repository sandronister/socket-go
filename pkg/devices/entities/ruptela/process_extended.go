package ruptela

import (
	"time"

	"dario.cat/mergo"
	"github.com/google/uuid"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities"
)

func (d *Device) appendOrMerge() {

	if recordsLen := len(d.Parameters); recordsLen > 0 {

		lastRecord := (d.Parameters)[recordsLen-1]

		if lastRecord.Time == d.Parameter.Time {

			if err := mergo.Merge(&lastRecord, d.Parameter); err != nil {
				return
			}

			(d.Parameters)[recordsLen-1] = lastRecord
			return
		}
	}

	d.Parameters = append(d.Parameters, d.Parameter)
}

func (d *Device) readBufferExtended() {
	var (
		sat          byte
		lon          int32
		lat          int32
		hdop         byte
		alt          uint16
		notUsed      byte
		speed        uint16
		course       uint16
		gpstime      uint32
		eventIDExt   uint16
		timestampExt byte
	)
	for range d.Header.RecordsCount {
		d.Parameter = &entities.TParameters{}

		d.Parameter.Id = uuid.NewString()

		d.Read4Bu(&gpstime)
		d.Read1B(&timestampExt)

		d.Read1B(&notUsed)
		d.Read1B(&notUsed)

		d.Read4B(&lon)
		d.Read4B(&lat)
		d.Read2Bu(&alt)
		d.Read2Bu(&course)
		d.Read1B(&sat)
		d.Read2Bu(&speed)
		d.Read1B(&hdop)

		d.Read2Bu(&eventIDExt)

		d.Parameter.Gpstime = int32(gpstime)
		d.Parameter.GpsSpeed = int32(speed)
		d.Parameter.Height = int32(alt) / 10
		d.Parameter.SatelliteQuatity = int32(sat)
		d.Parameter.GpsCourse = int32(course) / 100
		d.Parameter.Latitude = float64(lat) / 10000000
		d.Parameter.Longitude = float64(lon) / 10000000

		d.Parameter.Imei = d.Header.Imei
		d.Parameter.DeviceId = d.Header.Imei
		d.Parameter.Protocol = string(MANUFACTURER_RUPTELA)

		d.ParseOneByte()
		d.Parse2Bytes()
		d.Parse4Bytes()
		d.Parse8Bytes()

		d.TransformAdditionalValues(hdop)

		d.Parameter.TotalMileageFilled = d.Parameter.TotalMileage

		tm := time.Unix(int64(gpstime), 0)
		d.Parameter.Time = tm.Format("2006-01-02 15:04:05")

		d.Parameter.IsValid = d.IsValidRecord(tm)

		d.appendOrMerge()

	}

}

func (d *Device) ProcessExtended() *customerrors.Error {
	var (
		recordsLeft  byte
		recordsCount byte
	)

	d.Read1B(&recordsLeft)
	d.Read1B(&recordsCount)

	d.Header.RecordsLeft = recordsLeft
	d.Header.RecordsCount = recordsCount
	d.readBufferExtended()

	crc := uint16(0)
	d.Read2Bu(&crc)

	if !d.checkCRC(crc) {
		d.Ack = extSuccessNack
		return &customerrors.ErrInvalidCRC
	}

	if !d.IsValidSize() {
		d.Ack = extSuccessNack
		return &customerrors.ErrInvalidSize
	}

	d.Success = true
	d.Ack = extSuccessAck
	d.Success = len(d.Parameters) > 0

	return nil
}
