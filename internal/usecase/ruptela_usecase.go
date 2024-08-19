package usecase

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/internal/entities"
	"github.com/sandronister/socket-go/internal/repositories"
)

type RuptelaUsecase struct {
	DeviceRepo      repositories.IDeviceRepository
	blackListDevice map[string]bool
}

func NewRuptelaUsecase(deviceRepo repositories.IDeviceRepository) *RuptelaUsecase {
	r := &RuptelaUsecase{
		DeviceRepo: deviceRepo,
	}

	r.initBlackListDevice()
	return r
}

func (u *RuptelaUsecase) initBlackListDevice() error {
	result, err := u.DeviceRepo.GetBlackList()

	if err != nil {
		return err
	}

	u.blackListDevice = make(map[string]bool)

	for _, v := range result {
		u.blackListDevice[v.Imei] = v.Blocked
	}

	return nil
}

func (u *RuptelaUsecase) getMD5(buff []byte) string {
	hasher := md5.New()
	hasher.Write(buff)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *RuptelaUsecase) Handle(buff []byte) *dto.DeviceResponse {
	md5Str := u.getMD5(buff)
	u.DeviceRepo.SaveMessage(md5Str)
	r := entities.Ruptela{
		OriginalBuff: buff,
	}

	return &dto.DeviceResponse{
		Error:          "",
		Ack:            r.Ack,
		ProtocolBehave: 0,
		Success:        r.Success,
		Imei:           r.Header.Imei,
	}
}
