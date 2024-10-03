package setting

var (
	UnaryPlusPriority      = 1000
	AdditionPriority       = 10
	UnaryMinusPriority     = 1000
	SubtractionPriority    = 10
	MultiplicationPriority = 100
	DivisionPriority       = 100
	PowerPriority          = 1000
	ModuloPriority         = 100
)

var AllowedOperators = []rune{
	'+', '-', '*', '/', '%', '^', '!',
	'#', '§', '$', '&', ';', ':', '~',
	'<', '>', '|', '=', '÷', '√', '∛',
	'⌈', '⌊',
}
