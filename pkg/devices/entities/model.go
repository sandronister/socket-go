package entities

type TSensor struct {
	ID    int
	Value *int64
}

type TParameters struct {
	Imei                                 string  `parquet:"name=imei, type=UTF8, encoding=PLAIN_DICTIONARY"`
	DeviceId                             string  `parquet:"name=deviceId, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Protocol                             string  `parquet:"name=protocol, type=UTF8, encoding=PLAIN_DICTIONARY"`
	SatelliteQuatity                     int32   `parquet:"name=satellite_quatity, type=INT32"`
	Latitude                             float64 `parquet:"name=latitude, type=DOUBLE"`
	Longitude                            float64 `parquet:"name=longitude, type=DOUBLE"`
	GpsCourse                            int32   `parquet:"name=gps_course, type=INT32"`
	Height                               int32   `parquet:"name=height, type=INT32"`
	GpsSpeed                             int32   `parquet:"name=gps_speed, type=INT32"`
	IgnitionKey                          int32   `parquet:"name=ignition_key, type=INT32"`
	BateryLife                           int32   `parquet:"name=batery_life, type=INT32"`
	GsmStatus                            int32   `parquet:"name=gsm_status, type=INT32"`
	GpsStatus                            int32   `parquet:"name=gps_status, type=INT32"`
	Voltage                              float64 `parquet:"name=voltage, type=DOUBLE"`
	Mileage                              float64 `parquet:"name=mileage, type=DOUBLE"`
	MileageVirt                          float64 `parquet:"name=mileage_virt, type=DOUBLE"`
	Acceleration                         float64 `parquet:"name=acceleration, type=DOUBLE"`
	Input_0                              int32   `parquet:"name=input_0, type=INT32"`
	Input_1                              int32   `parquet:"name=input_1, type=INT32"`
	Input_2                              int32   `parquet:"name=input_2, type=INT32"`
	Input_3                              int32   `parquet:"name=input_3, type=INT32"`
	Input_4                              int32   `parquet:"name=input_4, type=INT32"`
	Rpm                                  float64 `parquet:"name=rpm, type=DOUBLE"`
	Lls1Value                            int32   `parquet:"name=lls1_value, type=INT32"`
	Lls1Status                           int32   `parquet:"name=lls1_status, type=INT32"`
	Lls1Temp                             int32   `parquet:"name=lls1_temp, type=INT32"`
	Lls2Value                            int32   `parquet:"name=lls2_value, type=INT32"`
	Lls2Status                           int32   `parquet:"name=lls2_status, type=INT32"`
	Lls2Temp                             int32   `parquet:"name=lls2_temp, type=INT32"`
	Lls3Value                            int32   `parquet:"name=lls3_value, type=INT32"`
	Lls3Status                           int32   `parquet:"name=lls3_status, type=INT32"`
	Lls3Temp                             int32   `parquet:"name=lls3_temp, type=INT32"`
	Lls4Value                            int32   `parquet:"name=lls4_value, type=INT32"`
	Lls4Status                           int32   `parquet:"name=lls4_status, type=INT32"`
	Lls4Temp                             int32   `parquet:"name=lls4_temp, type=INT32"`
	Lls5Value                            int32   `parquet:"name=lls5_value, type=INT32"`
	Lls5Status                           int32   `parquet:"name=lls5_status, type=INT32"`
	Lls5Temp                             int32   `parquet:"name=lls5_temp, type=INT32"`
	Lls6Value                            int32   `parquet:"name=lls6_value, type=INT32"`
	Lls6Status                           int32   `parquet:"name=lls6_status, type=INT32"`
	Lls6Temp                             int32   `parquet:"name=lls6_temp, type=INT32"`
	CanFuelLevel                         float64 `parquet:"name=can_fuel_level, type=DOUBLE"`
	AccelerometerX                       int32   `parquet:"name=accelerometer_x, type=INT32"`
	AccelerometerY                       int32   `parquet:"name=accelerometer_y, type=INT32"`
	AccelerometerZ                       int32   `parquet:"name=accelerometer_z, type=INT32"`
	CanTotalFuelGaseous                  float64 `parquet:"name=can_total_fuel_gaseous, type=DOUBLE"`
	ParkingBrakeStatus                   int32   `parquet:"name=parking_brake_status, type=INT32"`
	AccelerationPedalPosition            float64 `parquet:"name=acceleration_pedal_position, type=DOUBLE"`
	EngineOilPressure                    int32   `parquet:"name=engine_oil_pressure, type=INT32"`
	CoolantTemperature                   int32   `parquet:"name=coolant_temperature, type=INT32"`
	FuelTemperature                      int32   `parquet:"name=fuel_temperature, type=INT32"`
	EngineOilTemperature                 int32   `parquet:"name=engine_oil_temperature, type=INT32"`
	DailyFuelConsumption                 int32   `parquet:"name=daily_fuel_consumption, type=INT32"`
	EventOfInstantaneousFuelEconomy      int32   `parquet:"name=event_of_instantaneous_fuel_economy, type=INT32"`
	DailyMileage                         float64 `parquet:"name=daily_mileage, type=DOUBLE"`
	TotalMileage                         float64 `parquet:"name=total_mileage, type=DOUBLE"`
	TotalMileageFilled                   float64 `parquet:"name=total_mileage_filled, type=DOUBLE"`
	TotalTimeOfEngineOperation           float64 `parquet:"name=total_time_of_engine_operation, type=DOUBLE"`
	TotalFuelConsumption                 float64 `parquet:"name=total_fuel_consumption, type=DOUBLE"`
	TotalFuelConsumptionHighResolution   float64 `parquet:"name=total_fuel_consumption_high_resolution, type=DOUBLE"`
	ServiceBrakePedalPosition            float64 `parquet:"name=service_brake_pedal_position, type=DOUBLE"`
	ClutchPedalPosition                  float64 `parquet:"name=clutch_pedal_position, type=DOUBLE"`
	CruiseControlStatus                  int32   `parquet:"name=cruise_control_status, type=INT32"`
	AxlePressure                         int32   `parquet:"name=axle_pressure, type=INT32"`
	AxleLoad_1                           float64 `parquet:"name=axle_load_1, type=DOUBLE"`
	AxleLoad_2                           float64 `parquet:"name=axle_load_2, type=DOUBLE"`
	AxleLoad_3                           float64 `parquet:"name=axle_load_3, type=DOUBLE"`
	AxleLoad_4                           float64 `parquet:"name=axle_load_4, type=DOUBLE"`
	AxleLoad_5                           float64 `parquet:"name=axle_load_5, type=DOUBLE"`
	Adblue                               float64 `parquet:"name=adblue, type=DOUBLE"`
	Gear                                 float64 `parquet:"name=gear, type=DOUBLE"`
	ServiceBrakePedalStatus              int32   `parquet:"name=service_brake_pedal_status, type=INT32"`
	ClutchPedalStatus                    int32   `parquet:"name=clutch_pedal_status, type=INT32"`
	MileageUntilTheNextMaintenance       int32   `parquet:"name=mileage_until_the_next_maintenance, type=INT32"`
	EngineRunTimeUntilTheNextMaintenance int32   `parquet:"name=engine_run_time_until_the_next_maintenance, type=INT32"`
	AxleIndex                            int32   `parquet:"name=axle_index, type=INT32"`
	CanSpeed                             int32   `parquet:"name=can_speed, type=INT32"`
	DoorStatus                           int32   `parquet:"name=door_status, type=INT32"`
	SeatbeltStatus                       int32   `parquet:"name=seatbelt_status, type=INT32"`
	DeviceOpen                           int32   `parquet:"name=device_open, type=INT32"`
	OutputStatus                         int32   `parquet:"name=output_status, type=INT32"`
	Gpstime                              int32   `parquet:"name=gpstime, type=INT32"`
	IsValid                              bool    `parquet:"name=IsValid, type=BOOLEAN"`
	Id                                   string  `parquet:"name=id, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Time                                 string  `parquet:"name=time, type=UTF8, encoding=PLAIN_DICTIONARY"`
}

type TDtcParameters struct {
	Id      string
	Gpstime int32
	Source  uint8
	Lat     float64
	Lon     float64
	Status  uint8
	Text    string
	Time    string
}
