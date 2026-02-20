package checkers

import (
	"testing"
)

func TestCheckNoSensitiveData(t *testing.T) {
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
			name:    "обычное сообщение",
			input:   "connection timeout",
			wantErr: false,
		},
		{
			name:    "password с двоеточием",
			input:   "password: 12345",
			wantErr: true,
		},
		{
			name:    "password со знаком равно",
			input:   "password=12345",
			wantErr: true,
		},
		{
			name:    "password с пробелом",
			input:   "password 12345",
			wantErr: true,
		},
		{
			name:    "token с двоеточием",
			input:   "token: abc123",
			wantErr: true,
		},
		{
			name:    "api_key со знаком равно",
			input:   "api_key=secret",
			wantErr: true,
		},
		{
			name:    "secret с двоеточием",
			input:   "secret: value",
			wantErr: true,
		},
		{
			name:    "jwt с пробелом",
			input:   "jwt eyJhbGciOiJIUzI1NiJ9",
			wantErr: true,
		},
		{
			name:    "bearer с пробелом",
			input:   "bearer token123",
			wantErr: true,
		},
		{
			name:    "keyword без разделителя — не триггер",
			input:   "passwordcheck failed",
			wantErr: false,
		},
		{
			name:    "keyword в верхнем регистре",
			input:   "PASSWORD: 12345",
			wantErr: true,
		},
		{
			name:    "authorization с двоеточием",
			input:   "authorization: Bearer token",
			wantErr: true,
		},
		{
			name:    "cvv с двоеточием",
			input:   "cvv: 123",
			wantErr: true,
		},
		{
			name:    "credit_card со знаком равно",
			input:   "credit_card=1234567890",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckNoSensitiveData(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckNoSensitiveData(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
