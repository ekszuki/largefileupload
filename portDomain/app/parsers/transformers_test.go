package parsers

import (
	"ekszuki/uploader/portDomain/app/models"
	protoport "ekszuki/uploader/portDomain/protos/port"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromUpLoadPortRequestToDomain(t *testing.T) {
	type args struct {
		req *protoport.UpLoadPortRequest
	}
	tests := []struct {
		name     string
		args     args
		expected *models.Port
	}{
		{
			name: "check conversion of load port request to domain",
			args: args{
				req: &protoport.UpLoadPortRequest{
					Key:         "some key",
					Name:        "some name",
					City:        "some city",
					Country:     "some country",
					Alias:       []string{"alias 1", "alias 2"},
					Regions:     []string{"region 1"},
					Coordinates: []float64{10.0001, 20.0002},
					Province:    "some province",
					Timezone:    "some timezone",
					Unlocs:      []string{"some unlocs"},
					Code:        "123456",
				},
			},
			expected: &models.Port{
				Key:         "some key",
				Name:        "some name",
				City:        "some city",
				Country:     "some country",
				Alias:       []string{"alias 1", "alias 2"},
				Regions:     []string{"region 1"},
				Coordinates: []float64{10.0001, 20.0002},
				Province:    "some province",
				Timezone:    "some timezone",
				Unlocs:      []string{"some unlocs"},
				Code:        "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromUpLoadPortRequestToDomain(tt.args.req)
			assert.NotNil(t, got, "the object should not be nil")
			assert.EqualValuesf(t, tt.expected, got, "divergent values on the object")
		})
	}
}
