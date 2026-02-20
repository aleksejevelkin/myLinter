package checkers

import (
	"testing"
)

func TestCheckSpecialChars(t *testing.T) {
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
			input:   "something went wrong",
			wantErr: false,
		},
		{
			name:    "одиночный восклицательный знак",
			input:   "error!",
			wantErr: false,
		},
		{
			name:    "двойной восклицательный знак",
			input:   "error!!",
			wantErr: true,
		},
		{
			name:    "тройной восклицательный знак",
			input:   "error!!!",
			wantErr: true,
		},
		{
			name:    "тройной вопросительный знак",
			input:   "what???",
			wantErr: true,
		},
		{
			name:    "тройная точка (многоточие)",
			input:   "loading...",
			wantErr: true,
		},
		{
			name:    "символ @",
			input:   "user@domain.com",
			wantErr: true,
		},
		{
			name:    "символ #",
			input:   "error #42",
			wantErr: true,
		},
		{
			name:    "символ $",
			input:   "cost is $100",
			wantErr: true,
		},
		{
			name:    "символ %",
			input:   "100% complete",
			wantErr: true,
		},
		{
			name:    "символ ^",
			input:   "power^2",
			wantErr: true,
		},
		{
			name:    "символ &",
			input:   "cats & dogs",
			wantErr: true,
		},
		{
			name:    "символ *",
			input:   "wildcard *",
			wantErr: true,
		},
		{
			name:    "символ ~",
			input:   "path ~/home",
			wantErr: true,
		},
		{
			name:    "эмодзи — смайл",
			input:   "great job 😀",
			wantErr: true,
		},
		{
			name:    "эмодзи — ракета",
			input:   "launching 🚀",
			wantErr: true,
		},
		{
			name:    "две точки — не многоточие",
			input:   "version 1.2.3",
			wantErr: false,
		},
		{
			name:    "дефис допустим",
			input:   "well-known error",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckSpecialChars(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckSpecialChars(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
