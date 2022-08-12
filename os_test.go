package platform

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadOperatingSystem(t *testing.T) {
	version := "7.1"
	release := "7.1"

	documentation := "https://www.openbsd.org/faq/index.html"
	memory := 2048
	storage := 8192
	processorCount := 1
	processorCoreCount := 1

	path, _ := os.Getwd()

	type args struct {
		path string
	}
	otherGuest64 := "otherGuest64"
	lsilogic := "lsilogic"
	vmxnet3 := "vmxnet3"
	bios := "bios"
	this := "this"
	that := "that"
	zat := "zat"
	lines1 := "line 1\nline 2"
	lines2 := "line 3\nline 4"
	lines3 := "line 5\nline 6"
	var tests = []struct {
		name    string
		args    args
		want    OperatingSystem
		wantErr bool
	}{
		{
			name: "Load operating system",
			args: args{
				path: filepath.Join(path, "os.yml"),
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
				BootMethods: &BootMethodsDescriptor{
					Http: &BootMethodDescriptor{
						Commands: &this,
						Files: map[string]string{
							"file1.conf": lines1,
						},
					},
					Cdrom: &BootMethodDescriptor{
						Commands: &that,
						Files: map[string]string{
							"file2.conf": lines2,
						},
					},
					Floppy: &BootMethodDescriptor{
						Commands: &zat,
						Files: map[string]string{
							"file3.conf": lines3,
						},
					},
				},
				Hypervisors: &HypervisorsDescriptor{
					VSphere: &VSphereHypervisorDescriptor{
						GuestOSType:        &otherGuest64,
						DiskControllerType: &lsilogic,
						NetworkAdapterType: &vmxnet3,
						Firmware:           &bios,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := LoadOperatingSystem(tt.args.path)
				if (err != nil) != tt.wantErr {
					t.Errorf("LoadOperatingSystem() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LoadOperatingSystem() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
