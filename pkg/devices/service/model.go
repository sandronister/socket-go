package service

type IBlackListService interface {
	IsBlackListed(imei string) bool
}
