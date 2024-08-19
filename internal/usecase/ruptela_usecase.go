package usecase

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/internal/entities"
	"github.com/sandronister/socket-go/internal/repositories"
)

type RuptelaUsecase struct {
	deviceRepo      repositories.IDeviceRepository
	broker          types.IBroker
	blackListDevice map[string]bool
}

func NewRuptelaUsecase(deviceRepo repositories.IDeviceRepository, broker types.IBroker) *RuptelaUsecase {
	r := &RuptelaUsecase{
		deviceRepo: deviceRepo,
		broker:     broker,
	}

	r.initBlackListDevice()
	return r
}

func (u *RuptelaUsecase) initBlackListDevice() error {
	result, err := u.deviceRepo.GetBlackList()

	if err != nil {
		return err
	}

	u.blackListDevice = make(map[string]bool)

	for _, v := range result {
		u.blackListDevice[v.Imei] = v.Blocked
	}

	return nil
}

func (u *RuptelaUsecase) Handle(buff []byte) *dto.DeviceResponse {
	r := entities.Ruptela{
		OriginalBuff: buff,
	}
	r.SetMD5()
	u.deviceRepo.SaveMessage(r)
	return &dto.DeviceResponse{
		Error:          "",
		Ack:            r.Ack,
		ProtocolBehave: 0,
		Success:        r.Success,
		Imei:           r.Header.Imei,
	}
}
