package blacklist

func (s *blacklistservice) InitBlackListDevice() error {
	result, err := s.repository.GetBlackList()

	if err != nil {
		return err
	}

	s.blackListDevice = make(map[string]bool)

	for _, v := range result {
		s.blackListDevice[v.Imei] = v.Blocked
	}

	return nil
}
