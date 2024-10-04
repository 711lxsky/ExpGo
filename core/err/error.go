package err

var (
	CanNotParseToken           = "Can not parse char '%c' at [%d]"
	CheckTheName               = "Check the name '%c' at [%d]"
	MissSeparatorOrParenthesis = "Miss separator or parenthesis at [%d]"
	UnknownTokenType           = "Unknown token type at [%d]"
	ExpressionCanNotBeEmpty    = "Expression can not be empty"
	VariableCanNotHasSameName  = "Variable can not has same name with function"
	VariableMissValue          = "The variable '%s' miss value"
	HaveNotEnoughArguments     = "Have not enough arguments for function '%s'"
	TooManyOperators           = "Too many operators"
	TooManyArguments           = "Too many arguments"
	LackSufficientOperands     = "Lack sufficient operands for '%s'"
	InvalidItemsInStack        = "Invalid items in stack"
)
