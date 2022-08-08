package platform

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadOperatingSystem(t *testing.T) {
	version := "7.1"
	release := "7"

	documentation := "https://www.openbsd.org/faq/index.html"
	memory := 2048
	storage := 8192
	processorCount := 1
	processorCoreCount := 1

	path, _ := os.Getwd()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    OperatingSystem
		wantErr bool
	}{
		{
			name: "Load operating system",
			args: args{
				path: filepath.Join(path, "data", "openbsd-7.1-amd64.yml"),
			},
			want: &OperatingSystemDescriptor{
				Name:         "OpenBSD",
				Version:      &version,
				Architecture: "amd64",
				Release:      &release,
				DownloadURLs: []string{
					"https://mirrors.ocf.berkeley.edu/pub/OpenBSD/7.1/amd64/install71.iso",
				},
				Flavor:           nil,
				DocumentationURL: &documentation,
				RecommendedHardware: &HardwareDescriptor{
					Memory:  &memory,
					Storage: &storage,
					Processors: &ProcessorsDescriptor{
						Count: &processorCount,
						Cores: &processorCoreCount,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadOperatingSystem(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadOperatingSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadOperatingSystem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
