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
		{
			name: "Parse OpenBSD operating system",
			args: args{
				b: []byte(`
name: OpenBSD
version: "7.1"
architecture: amd64
release: "7.1"
downloads:
    - https://mirrors.ocf.berkeley.edu/pub/OpenBSD/7.1/amd64/install71.iso
documentation: https://www.openbsd.org/faq/index.html
hardware:
    memory: 2048
    storage: 8192
    processors:
        count: 1
        cores: 1
boot:
    http:
        commands: |
            a<enter><wait>
            <wait10>
            {{ .ServerUrl }}/install.conf<enter><wait>
            <wait5>
            i<enter><wait>
        files:
            install.conf.pkrtpl.hcl: |-
                System hostname = openbsd
                Password for root account = ${ ssh_password }

                Change the default console to com0? = yes
                Which speed should com0 use? (or 'done') = 19200

                Setup a user? = ${ packer_username }
                Password for user ${ packer_username }? = ${ packer_password }

                What timezone are you in? = ${ timezone }

                Which disk is the root disk? = sd0
                Use DUIDs rather than device names in fstab? = yes
                Use (W)hole disk or (E)dit the MBR? = W
                Use (A)uto layout, (E)dit auto layout, or create (C)ustom layout? = a
                Which disk do you wish to initialize? = done

                Do you expect to run the X Window System = no

                Start sshd(8) by default? = yes
                Allow root ssh login = yes

                # Location of sets = http
                # HTTP Server = cdn.openbsd.org
                Set name(s) = -game*

                Start ntpd(8) by default? = no

                Directory does not contain SHA256.sig. Continue without verification = yes
        delay: 30s
        variables:
            - name: packer_password
              type: string
            - name: packer_username
              default: phaka
              type: string
            - name: ssh_password
              type: string
            - name: timezone
              default: America/Los_Angeles
              type: string
hypervisors:
    vsphere:
        guest_os_type: otherGuest64
        disk_controller_type: lsilogic
        network_adapter_type: e1000
        firmware: bios

`),
			},
			want: &OperatingSystem{
				Name:             "OpenBSD",
				Id:               "openbsd",
				Version:          "7.1",
				Architecture:     "amd64",
				Release:          "7.1",
				DownloadURLs:     []string{"https://mirrors.ocf.berkeley.edu/pub/OpenBSD/7.1/amd64/install71.iso"},
				Flavor:           "",
				DocumentationURL: "https://www.openbsd.org/faq/index.html",
				RecommendedHardware: Hardware{
					Memory:  2048,
					Storage: 8192,
					Processors: Processors{
						Count: 1,
						Cores: 1,
					},
				},
				Hypervisors: Hypervisors{
					HypervisorKindVMware: &Hypervisor{
						"disk_controller_type": "lsilogic",
						"firmware":             "bios",
						"guest_os_type":        "otherGuest64",
						"network_adapter_type": "e1000",
					},
				},
				BootMethods: BootMethods{
					BootMethodKindHttp: &BootMethod{
						Username: "root",
						Name:     "HTTP",
						Commands: `a<enter><wait>
<wait10>
{{ .ServerUrl }}/install.conf<enter><wait>
<wait5>
i<enter><wait>
`,
						Files: map[string]string{
							"install.conf.pkrtpl.hcl": `System hostname = openbsd
Password for root account = ${ ssh_password }

Change the default console to com0? = yes
Which speed should com0 use? (or 'done') = 19200

Setup a user? = ${ packer_username }
Password for user ${ packer_username }? = ${ packer_password }

What timezone are you in? = ${ timezone }

Which disk is the root disk? = sd0
Use DUIDs rather than device names in fstab? = yes
Use (W)hole disk or (E)dit the MBR? = W
Use (A)uto layout, (E)dit auto layout, or create (C)ustom layout? = a
Which disk do you wish to initialize? = done

Do you expect to run the X Window System = no

Start sshd(8) by default? = yes
Allow root ssh login = yes

# Location of sets = http
# HTTP Server = cdn.openbsd.org
Set name(s) = -game*

Start ntpd(8) by default? = no

Directory does not contain SHA256.sig. Continue without verification = yes`,
						},
						Delay: "30s",
						Variables: []*Variable{
							{
								Name:         "packer_password",
								Type:         "string",
								DefaultValue: "",
							},
							{
								Name:         "packer_username",
								Type:         "string",
								DefaultValue: "phaka",
							},
							{
								Name:         "ssh_password",
								Type:         "string",
								DefaultValue: "",
							},
							{
								Name:         "timezone",
								Type:         "string",
								DefaultValue: "America/Los_Angeles",
							},
						},
					},
				},
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
					t.Errorf(
						"Parse() got = \n%s\n%v\n%s\n, want \n%s\n%v\n%s\n%s",
						s,
						got,
						s,
						s,
						tt.want,
						s,
						string(tt.args.b),
					)
				}
			},
		)
	}
}
