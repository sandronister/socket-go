package ruptela

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/google/uuid"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities"
)

func (d *Device) readBuffer(recordsCount int, buff *bytes.Buffer) {
	for range recordsCount {
		var (
			source  byte
			lat     float64
			lon     float64
			status  byte
			gpstime uint32
			text4b  uint32
			text1b  byte
		)

		d.DtcParameters = &entities.TDtcParameters{}
		d.DtcParameters.Id = uuid.New().String()

		binary.Read(buff, binary.BigEndian, &source)
		binary.Read(buff, binary.BigEndian, &gpstime)
		binary.Read(buff, binary.BigEndian, &lon)
		binary.Read(buff, binary.BigEndian, &lat)
		binary.Read(buff, binary.BigEndian, &status)

		binary.Read(buff, binary.BigEndian, &text4b)
		binary.Read(buff, binary.BigEndian, &text1b)

		d.DtcParameters.Gpstime = int32(gpstime)
		d.DtcParameters.Lat = lat / 10000000
		d.DtcParameters.Lon = lon / 10000000
		d.DtcParameters.Status = status
		d.DtcParameters.Source = source
		d.DtcParameters.Text = fmt.Sprintf("%d %d", text4b, text1b)

		tm := time.Unix(int64(d.Parameter.Gpstime), 0)
		d.DtcParameters.Time = tm.Format("2006-01-02 15:04:05")

	}

	d.Success = true
	d.Ack = defaultDtcSuccessAck

}

func (d *Device) ProcessDTC() *customerrors.Error {

	var recordsCount byte

	binary.Read(d.Buffer, binary.BigEndian, &recordsCount)

	d.readBuffer(int(recordsCount), d.Buffer)

	return nil
}
