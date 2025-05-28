package parser

import "testing"

func Test_defaultPackageName(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				dir: "../../",
			},
			want: "base",
		},
		{
			name: "test",
			args: args{
				dir: "www/xxx/111",
			},
			want: "base",
		},
		{
			name: "test",
			args: args{
				dir: "www/xxx/111AAB",
			},
			want: "aab",
		},
		{
			name: "test",
			args: args{
				dir: "www/xxx/111AA--B-xxx",
			},
			want: "aabxxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultPackageName(tt.args.dir); got != tt.want {
				t.Errorf("defaultPackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
