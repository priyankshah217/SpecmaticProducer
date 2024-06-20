package controllers

import "testing"

func Test_isNonStringParam(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test isNonStringParam with string",
			args: args{param: "test"},
			want: false,
		},
		{
			name: "Test isNonStringParam with int",
			args: args{param: "1"},
			want: true,
		},
		{
			name: "Test isNonStringParam with bool",
			args: args{param: "true"},
			want: true,
		},
		{
			name: "Test isNonStringParam with float",
			args: args{param: "1.1"},
			want: true,
		},
		{
			name: "Test isNonStringParam with empty string",
			args: args{param: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNonStringParam(tt.args.param); got != tt.want {
				t.Errorf("isNonStringParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
