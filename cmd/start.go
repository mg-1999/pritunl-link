package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/pritunl/pritunl-link/config"
	"github.com/pritunl/pritunl-link/constants"
	"github.com/pritunl/pritunl-link/sync"
	"time"
)

func Start() (err error) {
	logrus.WithFields(logrus.Fields{
		"version":     constants.Version,
		"public_addr": config.Config.PublicAddress,
	}).Info("cmd.start: Starting link")

	sync.Init()

	for {
		time.Sleep(1 * time.Second)
	}

	return
}
