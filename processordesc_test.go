package platform

import "testing"

func TestProcessorsDescriptor_GetCores(t *testing.T) {
	type fields struct {
		Count *int
		Cores *int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Nil",
			fields: fields{
				Count: nil,
				Cores: nil,
			},
			want: 1,
		},
		{
			name: "Custom",
			fields: fields{
				Count: nil,
				Cores: &two,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProcessorsDescriptor{
				Count: tt.fields.Count,
				Cores: tt.fields.Cores,
			}
			if got := p.GetCores(); got != tt.want {
				t.Errorf("GetCores() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessorsDescriptor_GetCount(t *testing.T) {
	type fields struct {
		Count *int
		Cores *int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Nil",
			fields: fields{
				Count: nil,
				Cores: nil,
			},
			want: 1,
		},
		{
			name: "Custom",
			fields: fields{
				Count: &two,
				Cores: nil,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProcessorsDescriptor{
				Count: tt.fields.Count,
				Cores: tt.fields.Cores,
			}
			if got := p.GetCount(); got != tt.want {
				t.Errorf("GetCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessorsDescriptor_Validate(t *testing.T) {
	type fields struct {
		Count *int
		Cores *int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Nil",
			fields: fields{
				Count: nil,
				Cores: nil,
			},
			wantErr: false,
		},
		{
			name: "Custom Cores",
			fields: fields{
				Count: nil,
				Cores: &two,
			},
			wantErr: false,
		},
		{
			name: "Custom Count",
			fields: fields{
				Count: &two,
				Cores: nil,
			},
			wantErr: false,
		},
		{
			name: "Negative Cores",
			fields: fields{
				Count: nil,
				Cores: &negative,
			},
			wantErr: true,
		},
		{
			name: "Negative Count",
			fields: fields{
				Count: &negative,
				Cores: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProcessorsDescriptor{
				Count: tt.fields.Count,
				Cores: tt.fields.Cores,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
