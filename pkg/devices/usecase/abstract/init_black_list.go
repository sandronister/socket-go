package abstract

func (a *AbstractUseCase) InitBlackListDevice() error {
	result, err := a.DeviceRepo.GetBlackList()

	if err != nil {
		return err
	}

	a.BlackListDevice = make(map[string]bool)

	for _, v := range result {
		a.BlackListDevice[v.Imei] = v.Blocked
	}

	return nil
}
