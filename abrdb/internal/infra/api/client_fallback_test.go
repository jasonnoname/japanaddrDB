package api

import "testing"

func TestPublicS3FallbackURL(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
		ok   bool
	}{
		{
			name: "ABR CDN URL",
			in:   "https://data.address-br.digital.go.jp/mt_city/pref/mt_city_pref47.csv.zip",
			want: "https://gov-csv-export-public.s3.ap-northeast-1.amazonaws.com/mt_city/pref/mt_city_pref47.csv.zip",
			ok:   true,
		},
		{name: "unrelated host", in: "https://example.com/file.zip", ok: false},
		{name: "insecure URL", in: "http://data.address-br.digital.go.jp/file.zip", ok: false},
		{name: "lookalike host", in: "https://data.address-br.digital.go.jp.example.com/file.zip", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := publicS3FallbackURL(tt.in)
			if ok != tt.ok || got != tt.want {
				t.Fatalf("publicS3FallbackURL(%q) = (%q, %v), want (%q, %v)", tt.in, got, ok, tt.want, tt.ok)
			}
		})
	}
}
