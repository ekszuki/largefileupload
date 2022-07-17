package contracts

import (
	"context"
	"ekszuki/uploader/portDomain/app/models"
)

// Repository interface to store data on database
type PortRepository interface {
	SaveOrUpdate(ctx context.Context, port *models.Port) error
}
