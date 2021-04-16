package idxn_test

import (
	"fmt"
	"testing"

	idx "github.com/Mobilpadde/idxn"
)

func TestAlpha(t *testing.T) {
	var tests = []struct {
		idx  int
		len  int
		want string
	}{
		{
			idx:  0,
			want: "a",
		},
		{
			idx:  8,
			want: "i",
		},
		{
			idx:  100,
			want: "wabc",
		},
		{
			idx:  250,
			want: "qabcdefghi",
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf(`%d_%s`, tt.idx, tt.want)

		t.Run(name, func(t *testing.T) {
			ans := idx.Alpha(tt.idx)

			if ans != tt.want {
				t.Errorf(`got "%s", want "%s"`, ans, tt.want)
			}
		})
	}
}

func TestAlphanumeric(t *testing.T) {
	var tests = []struct {
		idx  int
		len  int
		want string
	}{
		{
			idx:  0,
			want: "a",
		},
		{
			idx:  8,
			want: "i",
		},
		{
			idx:  100,
			want: "3ab",
		},
		{
			idx:  250,
			want: "9abcdef",
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf(`%d_%s`, tt.idx, tt.want)

		t.Run(name, func(t *testing.T) {
			ans := idx.Alphanumeric(tt.idx)

			if ans != tt.want {
				t.Errorf(`got "%s", want "%s"`, ans, tt.want)
			}
		})
	}
}
