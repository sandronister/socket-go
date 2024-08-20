package ruptela

import "github.com/sandronister/socket-go/pkg/devices/entities/abstract"

var extSuccesAck = []uint8{0x00, 0x02, 0x64, 0x01, 0x13, 0xbc} //{0x00, 0x02, 0x64, 0x01, 0x13, 0xbc}
var extSuccesNack = []uint8{0x00, 0x02, 0x64, 0x00, 0x2, 0x35}

type KeyCountRuptela struct {
	Data  string
	Count int
}

type Header struct {
	Imei         string
	Command      byte
	RecordsLeft  byte
	RecordsCount byte
	PackageLen   uint32
}

type Device struct {
	abstract.Device
	Header  Header
	Ack     []byte
	Success bool
}

type TCommandProtocol uint8

const (
	COMMAND_UNKNOWN                            TCommandProtocol = 0x00
	COMMAND_RECORDS                            TCommandProtocol = 0x01
	COMMAND_DEVICE_CONFIG_DATA                 TCommandProtocol = 0x02
	COMMAND_DEVICE_VERSION_INFO                TCommandProtocol = 0x03
	COMMAND_DEVICE_FIRMWARE_UPDATE             TCommandProtocol = 0x04
	COMMAND_SMART_CARD_DATA                    TCommandProtocol = 0x05
	COMMAND_START_CARD_DATA_SIZETMSP           TCommandProtocol = 0x06
	COMMAND_SMS_VIA_GPRS                       TCommandProtocol = 0x07
	COMMAND_DIAGNOSTIC_TROUBLE_CODES           TCommandProtocol = 0x09
	COMMAND_TACHOGRAPH_COMMUNICATION           TCommandProtocol = 0x0A
	COMMAND_INFORMATION_PACKET_TACHOGRAPH_DATA TCommandProtocol = 0x0B
	COMMAND_TRANSPARENT_CHANNEL_DATA           TCommandProtocol = 0x0E
	COMMAND_INFORMATION_PACKET                 TCommandProtocol = 0x0F
	COMMAND_HEARTBEAT                          TCommandProtocol = 0x10
	COMMAND_SET_IO_VALUE                       TCommandProtocol = 0x11
	COMMAND_DYNAMIC_IDENTFIC_PACKET            TCommandProtocol = 0x12
	COMMAND_2ND_GEN_SMART_CARD_DATA_SIZETMSP   TCommandProtocol = 0x13
	COMMAND_GARMIN_DEVICE_REQUEST_STATUS       TCommandProtocol = 0x1E
	COMMAND_GARMIN_DEVICE_DATA                 TCommandProtocol = 0x1F
	COMMAND_WEIGHTING_SYSTEM_DATA              TCommandProtocol = 0x20
	COMMAND_FLS_COMMUNICATION_CHANNEL          TCommandProtocol = 0x21
	COMMAND_SD_CARD_LOGGING_FUNCTIONALITY      TCommandProtocol = 0x22
	COMMAND_ACCIDENT_RECONSTRUCTION            TCommandProtocol = 0x23
	COMMAND_FILES                              TCommandProtocol = 0x25
	COMMAND_BEACON_DATA                        TCommandProtocol = 0x26
	COMMAND_EXTENDED_PROTOCOL_RECORDS          TCommandProtocol = 0x44
	//Proprietary commands
	COMMAND_GOBRAX_PROPRIETARY  TCommandProtocol = 0xFF
	COMMAND_GOBRAX_BLOCKED_IMEI TCommandProtocol = COMMAND_GOBRAX_PROPRIETARY - 1
)

var allCommands = []TCommandProtocol{
	COMMAND_RECORDS,
	COMMAND_DEVICE_CONFIG_DATA,
	COMMAND_DEVICE_VERSION_INFO,
	COMMAND_DEVICE_FIRMWARE_UPDATE,
	COMMAND_SMART_CARD_DATA,
	COMMAND_START_CARD_DATA_SIZETMSP,
	COMMAND_SMS_VIA_GPRS,
	COMMAND_DIAGNOSTIC_TROUBLE_CODES,
	COMMAND_TACHOGRAPH_COMMUNICATION,
	COMMAND_INFORMATION_PACKET_TACHOGRAPH_DATA,
	COMMAND_TRANSPARENT_CHANNEL_DATA,
	COMMAND_INFORMATION_PACKET,
	COMMAND_HEARTBEAT,
	COMMAND_SET_IO_VALUE,
	COMMAND_DYNAMIC_IDENTFIC_PACKET,
	COMMAND_2ND_GEN_SMART_CARD_DATA_SIZETMSP,
	COMMAND_GARMIN_DEVICE_REQUEST_STATUS,
	COMMAND_GARMIN_DEVICE_DATA,
	COMMAND_WEIGHTING_SYSTEM_DATA,
	COMMAND_FLS_COMMUNICATION_CHANNEL,
	COMMAND_SD_CARD_LOGGING_FUNCTIONALITY,
	COMMAND_ACCIDENT_RECONSTRUCTION,
	COMMAND_FILES,
	COMMAND_BEACON_DATA,
	COMMAND_EXTENDED_PROTOCOL_RECORDS,
}

type TDynamicProtocol uint16

var (
	DYNAMIC_FIRM_VERSION TDynamicProtocol = 0x02
	DYNAMIC_CFG_TAG      TDynamicProtocol = 0x0B
)
