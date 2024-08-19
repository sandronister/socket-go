package devices

type ProtocolBehave uint8

const (
	BEHAVE_UNKNOWN              ProtocolBehave = 0x00
	BEHAVE_CLOSE_CONN           ProtocolBehave = 0x00
	BEHAVE_REPLY_AND_KEEP_ALIVE ProtocolBehave = 0x01
	BEHAVE_REPLY_AND_CLOSE_CONN ProtocolBehave = 0x02
)
