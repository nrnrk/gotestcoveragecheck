package main

import "testing"

func Test_feeling(t *testing.T) {
	type args struct {
		typ int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `type_0`,
			args: args{
				typ: 0,
			},
			want: `so-so`,
		},
		{
			name: `type_1`,
			args: args{
				typ: 1,
			},
			want: `fine`,
		},
		{
			name: `type_2`,
			args: args{
				typ: 2,
			},
			want: `great`,
		},
		{
			name: `type_3`,
			args: args{
				typ: 3,
			},
			want: `super`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := feeling(tt.args.typ); got != tt.want {
				t.Errorf("feeling() = %v, want %v", got, tt.want)
			}
		})
	}
}
