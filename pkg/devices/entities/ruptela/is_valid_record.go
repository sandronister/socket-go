package ruptela

import "time"

const minSATELITE_QUALITY = 3

func (d *Device) isValidCoordinates(lat_float float64, lon_float float64) bool {
	if lon_float == 0 || lon_float < -180 || lon_float > 180 || lat_float == 0 || lat_float < -90 || lat_float > 90 {
		return false
	}
	return true
}

func (d *Device) isValidDate(date time.Time) bool {

	now := time.Now()
	nowMinus90Days := now.Add(-(time.Hour * 24) * 90)
	nowPlus24Hours := now.Add(time.Hour * 24)

	return date.After(nowMinus90Days) && date.Before(nowPlus24Hours)
}

func (d *Device) IsValidRecord(tm time.Time) bool {

	if (d.Parameter.SatelliteQuatity > minSATELITE_QUALITY) && d.Parameter.Latitude != 0 && d.Parameter.Longitude != 0 {
		return d.isValidDate(tm) && d.isValidCoordinates(d.Parameter.Latitude, d.Parameter.Longitude)
	}

	return false
}
