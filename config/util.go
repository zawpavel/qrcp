package config

import (
	"github.com/zawpavel/qrcp/application"
	"github.com/zawpavel/qrcp/util"
)

func chooseInterface(flags application.Flags) (string, error) {
	iface, err := util.GetInterface(flags.ListAllInterfaces)
	if err != nil {
		return "", err
	}
	return iface, nil
}
