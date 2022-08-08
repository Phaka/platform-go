package platform

import (
	"reflect"
	"testing"
)

func TestOperatingSystemDescriptor_GetArchitecture(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "x86_64",
			fields: fields{
				Architecture: "x86_64",
			},
			want: "x86_64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetArchitecture(); got != tt.want {
				t.Errorf("GetArchitecture() = %v, want %v", got, tt.want)
			}
		})
	}
}

var foobar = "foobar"

func TestOperatingSystemDescriptor_GetDocumentationURL(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Nil",
			fields: fields{
				DocumentationURL: nil,
			},
			want: "",
		},
		{
			name: "Custom",
			fields: fields{
				DocumentationURL: &foobar,
			},
			want: "foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetDocumentationURL(); got != tt.want {
				t.Errorf("GetDocumentationURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetDownloadURLs(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Nil",
			fields: fields{
				DownloadURLs: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetDownloadURLs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDownloadURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetFlavor(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Nil",
			fields: fields{
				Flavor: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetFlavor(); got != tt.want {
				t.Errorf("GetFlavor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetName(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Custom",
			fields: fields{
				Name: "Dummy",
			},
			want: "Dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetRecommendedHardware(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   Hardware
	}{
		{
			name: "Nil",
			fields: fields{
				RecommendedHardware: nil,
			},
			want: &HardwareDescriptor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetRecommendedHardware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecommendedHardware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetRelease(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Nil",
			fields: fields{
				Release: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetRelease(); got != tt.want {
				t.Errorf("GetRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_GetVersion(t *testing.T) {
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Nil",
			fields: fields{
				Version: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if got := o.GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatingSystemDescriptor_Validate(t *testing.T) {
	version := "1.0.0"
	release := "1"
	flavor := "test"
	documentation := "https://example.com"
	type fields struct {
		Name                string
		Version             *string
		Architecture        string
		Release             *string
		DownloadURLs        []string
		Flavor              *string
		DocumentationURL    *string
		RecommendedHardware *HardwareDescriptor
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid",
			fields: fields{
				Name:                "test",
				Version:             &version,
				Architecture:        "amd64",
				Release:             &release,
				DownloadURLs:        []string{"https://example.com/test"},
				Flavor:              &flavor,
				DocumentationURL:    &documentation,
				RecommendedHardware: &HardwareDescriptor{},
			},
			wantErr: false,
		},
		{
			name: "Invalid Name",
			fields: fields{
				Name:                "",
				Version:             &version,
				Architecture:        "amd64",
				Release:             &release,
				DownloadURLs:        []string{"https://example.com/test"},
				Flavor:              &flavor,
				DocumentationURL:    &documentation,
				RecommendedHardware: &HardwareDescriptor{},
			},
			wantErr: true,
		},
		{
			name: "Invalid Architecture",
			fields: fields{
				Name:                "test",
				Version:             &version,
				Architecture:        "",
				Release:             &release,
				DownloadURLs:        []string{"https://example.com/test"},
				Flavor:              &flavor,
				DocumentationURL:    &documentation,
				RecommendedHardware: &HardwareDescriptor{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatingSystemDescriptor{
				Name:                tt.fields.Name,
				Version:             tt.fields.Version,
				Architecture:        tt.fields.Architecture,
				Release:             tt.fields.Release,
				DownloadURLs:        tt.fields.DownloadURLs,
				Flavor:              tt.fields.Flavor,
				DocumentationURL:    tt.fields.DocumentationURL,
				RecommendedHardware: tt.fields.RecommendedHardware,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
