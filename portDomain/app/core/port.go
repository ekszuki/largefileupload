package core

import (
	"context"
	"ekszuki/uploader/portDomain/app/contracts"
	"ekszuki/uploader/portDomain/app/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

type PortApplication struct {
	portRepository contracts.PortRepository
}

func NewPortApplication(portRepo contracts.PortRepository) *PortApplication {
	return &PortApplication{
		portRepository: portRepo,
	}
}

func (a *PortApplication) SaveOrUpdate(ctx context.Context, port *models.Port) error {
	logCtx := logrus.WithFields(
		logrus.Fields{"package": "core", "function": "SaveOrUpdate"},
	)

	if port == nil {
		logCtx.Warn("nil is invalid port parameter")
		return fmt.Errorf("parameter port could not be nil")
	}

	err := a.portRepository.SaveOrUpdate(ctx, port)
	if err != nil {
		logCtx.Errorf("could not save or update port on database: %v", err)
		return fmt.Errorf("error on save or update port on database")
	}

	return nil
}
