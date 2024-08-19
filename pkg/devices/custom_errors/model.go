package customerrors

type ErrorMessage string

type Error struct {
	Code    int
	Message ErrorMessage
}

var (
	ErrSilentClose         = Error{Code: 0, Message: "silent close"}
	ErrInvalidImei         = Error{Code: 1, Message: "invalid imei"}
	ErrInvalidCommand      = Error{Code: 2, Message: "invalid command"}
	ErrInvalidHeader       = Error{Code: 3, Message: "invalid header"}
	ErrInvalidSize         = Error{Code: 4, Message: "invalid size"}
	ErrInvalidCRC          = Error{Code: 5, Message: "invalid crc"}
	ErrBlockedImei         = Error{Code: 6, Message: "blocked imei"}
	ErrInvalidSecurityWord = Error{Code: 7, Message: "invalid security word"}
)
