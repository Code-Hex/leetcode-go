package solution

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				S: "ab#c",
				T: "ad#c",
			},
			want: true,
		},
		{
			args: args{
				S: "ab##",
				T: "c#d#",
			},
			want: true,
		},
		{
			args: args{
				S: "a##c",
				T: "#a#c",
			},
			want: true,
		},
		{
			args: args{
				S: "a#c",
				T: "b",
			},
			want: false,
		},
		{
			args: args{
				S: "a#######",
				T: "",
			},
			want: true,
		},
		{
			args: args{
				S: "abcdefghijklmnopqrstuvwxyz",
				T: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			args: args{
				S: "abcde#####",
				T: "",
			},
			want: true,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
