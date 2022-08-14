package platform

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *OperatingSystem
		wantErr bool
	}{
		{
			name: "Parse minimal operating system",
			args: args{
				b: []byte(`
name: "MyOS"
`),
			},
			want: &OperatingSystem{
				Name:                "MyOS",
				Id:                  "myos",
				Version:             defaultOperatingSystem.Version,
				Architecture:        defaultOperatingSystem.Architecture,
				Release:             defaultOperatingSystem.Release,
				DownloadURLs:        defaultOperatingSystem.DownloadURLs,
				Flavor:              defaultOperatingSystem.Flavor,
				DocumentationURL:    defaultOperatingSystem.DocumentationURL,
				RecommendedHardware: defaultHardware,
				Hypervisors:         defaultHypervisors,
				BootMethods:         defaultBootMethods,
			},
		},
		{
			name: "Parse operating system with empty buffer",
			args: args{
				b: []byte(``),
			},
			wantErr: false,
		},
		{
			name: "Parse operating system with no name",
			args: args{
				b: []byte(`
name: ""
`),
			},
			wantErr: true,
		},
		{
			name: "Parse operating system with padded name, version",
			args: args{
				b: []byte(`
name: "MyOS    "
id: "myos   "
version: "  1.0   "
release: "  2.0   "
architecture: "  amd64   "
`),
			},
			want: &OperatingSystem{
				Name:                "MyOS",
				Id:                  "myos",
				Version:             "1.0",
				Architecture:        "amd64",
				Release:             "2.0",
				DownloadURLs:        defaultOperatingSystem.DownloadURLs,
				Flavor:              defaultOperatingSystem.Flavor,
				DocumentationURL:    defaultOperatingSystem.DocumentationURL,
				RecommendedHardware: defaultHardware,
				Hypervisors:         defaultHypervisors,
				BootMethods:         defaultBootMethods,
			},

		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := Parse(tt.args.b)
				if (err != nil) != tt.wantErr {
					t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					s := strings.Repeat("-", 80)
					t.Errorf("Parse() got = \n%s\n%v\n%s\n, want \n%s\n%v\n%s", s, got, s, s, tt.want, s)
				}
			},
		)
	}
}
