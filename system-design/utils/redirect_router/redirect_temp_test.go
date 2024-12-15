package redirect_router

import "testing"

func TestRedirectList_GetFullLink(t *testing.T) {
	type args struct {
		key string
	}
	var tests = []struct {
		name    string
		rl      RedirectList
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ExistingShorLink",
			rl: RedirectList{
				Redirect{
					ShortLink: "abc",
					FullLink:  "https://example.com/v1/page1",
				},
				Redirect{
					ShortLink: "xyz",
					FullLink:  "https://example.com/v1/page2",
				},
			},
			args:    args{"abc"},
			want:    "https://example.com/v1/page1",
			wantErr: false,
		},

		{
			name: "NonExistingShorLink",
			rl: RedirectList{
				Redirect{
					ShortLink: "abc",
					FullLink:  "https://example.com/v1/page1",
				},
				Redirect{
					ShortLink: "xyz",
					FullLink:  "https://example.com/v1/page2",
				},
			},
			args:    args{"def"},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rl.GetFullLink(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFullLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
