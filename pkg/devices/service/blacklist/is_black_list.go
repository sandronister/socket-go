package blacklist

func (s *blacklistservice) IsBlackListed(imei string) bool {
	if s.blackListDevice == nil {
		return false
	}

	if _, ok := s.blackListDevice[imei]; ok {
		return true
	}
	return false
}
