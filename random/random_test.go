package random

import "testing"

func Test_randFromRunes(t *testing.T) {
	type args struct {
		runes  []rune
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{runes: []rune{'a', 'b', 'c'}, lenght: 10},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randFromRunes(tt.args.runes, tt.args.lenght); got != tt.want {
				t.Errorf("randFromRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromRunes(t *testing.T) {
	type args struct {
		runes  []rune
		lenght int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromRunes(tt.args.runes, tt.args.lenght)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromRunes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.lenght); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphaNumeric(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphaNumeric(tt.args.lenght); got != tt.want {
				t.Errorf("AlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Number(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Numeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumeric(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Numeric(tt.args.lenght); got != tt.want {
				t.Errorf("Numeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlpha(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alpha(tt.args.lenght); got != tt.want {
				t.Errorf("Alpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphaLower(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphaLower(tt.args.lenght); got != tt.want {
				t.Errorf("AlphaLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphaUpper(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphaUpper(tt.args.lenght); got != tt.want {
				t.Errorf("AlphaUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMixin(t *testing.T) {
	type args struct {
		lenght int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mixin(tt.args.lenght); got != tt.want {
				t.Errorf("Mixin() = %v, want %v", got, tt.want)
			}
		})
	}
}
