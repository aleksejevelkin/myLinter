package checkers

import (
	"testing"
)

func TestCheckEnglishOnly(t *testing.T) {
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
			name:    "только английские символы",
			input:   "hello world",
			wantErr: false,
		},
		{
			name:    "английские символы с цифрами",
			input:   "error code 42",
			wantErr: false,
		},
		{
			name:    "английские символы с пунктуацией",
			input:   "something went wrong!",
			wantErr: false,
		},
		{
			name:    "символы новой строки и табуляция",
			input:   "line1\nline2\ttab",
			wantErr: false,
		},
		{
			name:    "русские символы",
			input:   "привет мир",
			wantErr: true,
		},
		{
			name:    "смешанный текст: английский и русский",
			input:   "hello мир",
			wantErr: true,
		},
		{
			name:    "эмодзи",
			input:   "hello 😀",
			wantErr: true,
		},
		{
			name:    "китайские иероглифы",
			input:   "你好",
			wantErr: true,
		},
		{
			name:    "все printable ASCII символы",
			input:   "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckEnglishOnly(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckEnglishOnly(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
