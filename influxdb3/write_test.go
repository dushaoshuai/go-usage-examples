package influxdb3

import (
	"testing"
)

// INFLUX_HOST=http://localhost:8181 INFLUX_TOKEN=dummy INFLUX_DATABASE=testdb go test -test.run ^TestWrite$
func TestWrite(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// influxdb3 query --database testdb "SELECT * FROM home"
