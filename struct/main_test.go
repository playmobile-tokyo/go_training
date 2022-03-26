package main

import "testing"

func Test_display(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test ok",
			args: args{msg: "string"},
			want: "[string]",
		},
		{
			name: "test ng",
			args: args{msg: "string"},
			want: "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := display(tt.args.msg); got != tt.want {
				t.Errorf("display() = %v, want %v", got, tt.want)
			}
		})
	}
}
