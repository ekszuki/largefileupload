package core_test

import (
	"context"
	"ekszuki/uploader/portDomain/app/contracts"
	"ekszuki/uploader/portDomain/app/core"

	"ekszuki/uploader/portDomain/app/models"
	mockport "ekszuki/uploader/portDomain/mocks/port"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPortApplication_SaveOrUpdate(t *testing.T) {
	type args struct {
		port *models.Port
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		mockedErr     error
		expectedError error
	}{
		{
			name: "saved successfully",
			args: args{
				port: &models.Port{
					Key:         "some key",
					Name:        "some name",
					City:        "some city",
					Country:     "some country",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{},
					Province:    "some province",
					Timezone:    "some timezone",
					Unlocs:      []string{},
					Code:        "some code",
				},
			},
			wantErr: false,
		},
		{
			name: "port parameter nil",
			args: args{
				port: nil,
			},
			wantErr:       true,
			mockedErr:     fmt.Errorf("parameter port could not be nil"),
			expectedError: fmt.Errorf("parameter port could not be nil"),
		},
		{
			name: "generic error on database",
			args: args{
				port: &models.Port{
					Key:         "some key",
					Name:        "some name",
					City:        "some city",
					Country:     "some country",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{},
					Province:    "some province",
					Timezone:    "some timezone",
					Unlocs:      []string{},
					Code:        "some code",
				},
			},
			wantErr:       true,
			mockedErr:     fmt.Errorf("error on save or update port on database"),
			expectedError: fmt.Errorf("error on save or update port on database"),
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock Repositorio
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPort := mockport.NewMockPortRepository(mockCtrl)
			mockPort.EXPECT().
				SaveOrUpdate(ctx, tt.args.port).
				MinTimes(0).
				Return(tt.mockedErr)

			a := core.NewPortApplication(mockPort)
			err := a.SaveOrUpdate(ctx, tt.args.port)

			if tt.wantErr {
				assert.Error(t, err, "error couldn't be nil")
				assert.EqualError(t, err, tt.expectedError.Error(), "Divergent erros")
			} else {
				assert.NoError(t, err, "error should be nil")
			}
		})
	}
}

func TestNewPortApplication(t *testing.T) {
	type args struct {
		portRepo contracts.PortRepository
	}
	tests := []struct {
		name     string
		args     args
		expected *core.PortApplication
	}{
		{
			name: "check Port Application Type",
			args: args{
				portRepo: nil,
			},
			expected: &core.PortApplication{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.NewPortApplication(tt.args.portRepo)
			assert.NotNil(t, got, "obj couldn't be nil")
			assert.IsType(t, tt.expected, got, "invalid type received")
		})
	}
}
