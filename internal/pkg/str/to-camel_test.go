package str

import "testing"

func TestToCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should success",
			args: args{
				s: "user_repository",
			},
			want: "UserRepository",
		},
		{
			name: "should success",
			args: args{
				s: "user repository",
			},
			want: "UserRepository",
		},
		{
			name: "should success",
			args: args{
				s: "user-repository",
			},
			want: "UserRepository",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPascal(tt.args.s); got != tt.want {
				t.Errorf("ToPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}
