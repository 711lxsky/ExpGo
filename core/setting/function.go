package setting

var (
	Sin     = "sin"
	Cos     = "cos"
	Tan     = "tan"
	Cot     = "cot"
	ArcSin  = "asin"
	ArcCos  = "acos"
	ArcTan  = "atan"
	SinH    = "sinh"
	CosH    = "cosh"
	TanH    = "tanh"
	Abs     = "abs"
	Log     = "log"
	Log10   = "log10"
	Log2    = "log2"
	Log1p   = "log1p"
	Ceil    = "ceil"
	Floor   = "floor"
	Sqrt    = "sqrt"
	Cbrt    = "cbrt"
	Pow     = "pow"
	Exp     = "exp"
	Expm1   = "expm1"
	Signum  = "signum"
	Csc     = "csc"
	Sec     = "sec"
	CscH    = "csch"
	SecH    = "sech"
	CotH    = "coth"
	Deg2Rad = "deg2rad"
	Rad2Deg = "rad2deg"
)

var AllowedFunctions = []string{
	Sin, Cos, Tan, Cot, ArcSin, ArcCos, ArcTan, SinH, CosH, TanH,
	Abs,
	Log, Log10, Log2, Log1p,
	Ceil, Floor,
	Sqrt, Cbrt,
	Pow, Exp, Expm1,
	Signum,
	Csc, Sec, CscH, SecH, CotH,
	Deg2Rad, Rad2Deg,
}
