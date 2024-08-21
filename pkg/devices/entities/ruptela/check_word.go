package ruptela

import "github.com/sandronister/socket-go/pkg/security"

func (r *Device) CheckWord(inMap map[TDynamicProtocol]string) bool {
	tag, ok := inMap[DYNAMIC_CFG_TAG]

	if !ok {
		return false
	}

	return security.MatchSecurityWord(tag)
}
