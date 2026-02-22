package analyzer

// Config задаёт, какие правила включены для проверки лог-сообщений.

type Config struct {
	Lowercase         *bool    `json:"lowercase" mapstructure:"lowercase"`
	EnglishOnly       *bool    `json:"englishOnly" mapstructure:"englishOnly"`
	SpecialChars      *bool    `json:"specialChars" mapstructure:"specialChars"`
	Sensitive         *bool    `json:"sensitive" mapstructure:"sensitive"`
	SensitiveKeywords []string `json:"sensitiveKeywords" mapstructure:"sensitiveKeywords"`
}

func defaultEnabled(v *bool) bool {
	if v == nil {
		return true
	}
	return *v
}

func (c Config) IsLowercaseEnabled() bool    { return defaultEnabled(c.Lowercase) }
func (c Config) IsEnglishOnlyEnabled() bool  { return defaultEnabled(c.EnglishOnly) }
func (c Config) IsSpecialCharsEnabled() bool { return defaultEnabled(c.SpecialChars) }
func (c Config) IsSensitiveEnabled() bool    { return defaultEnabled(c.Sensitive) }
