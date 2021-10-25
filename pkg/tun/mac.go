// +build darwin

package tun

import "github.com/sgostarter/i/logger"

func DeviceSetup(localCIDR, deviceName string) (TunDevice, error) {
	return nil, nil
}

func ClientExtInit(isTargetVPN bool, logger logger.Wrapper) {

}
