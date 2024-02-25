package konsole

type Formatter interface {
	FormatMessage(string) string
	FormatAlert(string) string
}

type NopFormatter struct{}

func (NopFormatter) FormatMessage(string) string { return "Console messages not implemented" }

func (NopFormatter) FormatAlert(string) string { return "Console alerts not implemented" }
