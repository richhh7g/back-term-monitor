package localization

type Language byte

const (
	PT_BR Language = iota
	EN_US
)

func (l *Language) String() string {
	return []string{"pt_br", "en_us"}[*l]
}
