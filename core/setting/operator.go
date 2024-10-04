package setting

var (
	UnaryOrPowerPriority  = 1000
	AddOrSubPriority      = 10
	MulOrDivOrModPriority = 100
)

var AllowedOperators = []rune{
	'+', '-', '*', '/', '%', '^', '!',
	'#', '§', '$', '&', ';', ':', '~',
	'<', '>', '|', '=', '÷', '√', '∛',
	'⌈', '⌊',
}
