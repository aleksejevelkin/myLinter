package checkers

import (
	"testing"
)

func TestCheckLowercase(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "пустая строка",
			input:   "",
			wantErr: false,
		},
		{
			name:    "строка только из пробелов",
			input:   "   ",
			wantErr: false,
		},
		{
			name:    "начинается с строчной буквы",
			input:   "hello world",
			wantErr: false,
		},
		{
			name:    "начинается с цифры",
			input:   "42 errors found",
			wantErr: false,
		},
		{
			name:    "начинается с заглавной буквы",
			input:   "Hello world",
			wantErr: true,
		},
		{
			name:    "все символы в верхнем регистре",
			input:   "ERROR OCCURRED",
			wantErr: true,
		},
		{
			name:    "начинается с пробела, потом заглавная",
			input:   "  Hello world",
			wantErr: true,
		},
		{
			name:    "начинается с пробела, потом строчная",
			input:   "  hello world",
			wantErr: false,
		},
		{
			name:    "начинается со спецсимвола",
			input:   "!important message",
			wantErr: false,
		},
		{
			name:    "одна строчная буква",
			input:   "a",
			wantErr: false,
		},
		{
			name:    "одна заглавная буква",
			input:   "A",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckLowercase(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLowercase(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
