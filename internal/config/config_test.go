package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		wantC   Config
		wantErr bool
	}{
		{
			name: "load",
			wantC: Config{
				ConnString: "mongodb+srv://localhost",
				Database:   "program",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Load() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
