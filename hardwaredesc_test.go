package platform

import (
	"reflect"
	"testing"
)

var zero = 0

var _4096 = 4096
var _10240 = 1024 * 10
var _20480 = 2048 * 10

func TestHardwareDescriptor_GetMemory(t *testing.T) {
	type fields struct {
		Memory     *int
		Storage    *int
		Processors *ProcessorsDescriptor
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Nil",
			fields: fields{
				Memory: nil,
			},
			want: 2048,
		},
		{
			name: "Zero",
			fields: fields{
				Memory: &zero,
			},
			want: 2048,
		},
		{
			name: "Custom",
			fields: fields{
				Memory: &_4096,
			},
			want: 4096,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HardwareDescriptor{
				Memory:     tt.fields.Memory,
				Storage:    tt.fields.Storage,
				Processors: tt.fields.Processors,
			}
			if got := h.GetMemory(); got != tt.want {
				t.Errorf("GetMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHardwareDescriptor_GetProcessors(t *testing.T) {
	type fields struct {
		Memory     *int
		Storage    *int
		Processors *ProcessorsDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   Processors
	}{
		{
			name: "Nil",
			fields: fields{
				Processors: nil,
			},
			want: defaultProcessors,
		},
		{
			name: "Custom",
			fields: fields{
				Processors: &ProcessorsDescriptor{
					Count: &two,
					Cores: &two,
				},
			},
			want: &ProcessorsDescriptor{
				Count: &two,
				Cores: &two,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HardwareDescriptor{
				Memory:     tt.fields.Memory,
				Storage:    tt.fields.Storage,
				Processors: tt.fields.Processors,
			}
			if got := h.GetProcessors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProcessors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHardwareDescriptor_GetStorage(t *testing.T) {
	type fields struct {
		Memory     *int
		Storage    *int
		Processors *ProcessorsDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Nil",
			fields: fields{
				Storage: nil,
			},
			want: 10240,
		},
		{
			name: "Zero",
			fields: fields{
				Storage: &zero,
			},
			want: 10240,
		},
		{
			name: "Custom",
			fields: fields{
				Storage: &_20480,
			},
			want: 20480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HardwareDescriptor{
				Memory:     tt.fields.Memory,
				Storage:    tt.fields.Storage,
				Processors: tt.fields.Processors,
			}
			if got := h.GetStorage(); got != tt.want {
				t.Errorf("GetStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

var negative = -1

func TestHardwareDescriptor_Validate(t *testing.T) {
	type fields struct {
		Memory     *int
		Storage    *int
		Processors *ProcessorsDescriptor
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Nil",
			fields: fields{
				Memory:     nil,
				Storage:    nil,
				Processors: nil,
			},
			wantErr: false,
		},
		{
			name: "Zero",
			fields: fields{
				Memory:     &zero,
				Storage:    &zero,
				Processors: nil,
			},
			wantErr: false,
		},
		{
			name: "Negative_Memory",
			fields: fields{
				Memory:     &negative,
				Storage:    &zero,
				Processors: nil,
			},
			wantErr: true,
		},
		{
			name: "Negative_Storage",
			fields: fields{
				Memory:     &zero,
				Storage:    &negative,
				Processors: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HardwareDescriptor{
				Memory:     tt.fields.Memory,
				Storage:    tt.fields.Storage,
				Processors: tt.fields.Processors,
			}
			if err := h.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
