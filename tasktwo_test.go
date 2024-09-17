package golangtryout

import "testing"

func TestBatchScan(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestBatchScan",
			args: args{
				urls: []string{
					"https://www.google.com",
					"https://www.facebook.com",
					"https://www.twitter.com",
					"https://www.instagram.com",
					"https://www.linkedin.com",
					"https://www.youtube.com",
					"https://www.pinterest.com",
					"https://www.tumblr.com",
					"https://www.reddit.com",
					"https://www.snapchat.com",
					"https://www.whatsapp.com",
					"https://www.viber.com",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BatchScan(tt.args.urls)
		})
	}
}
