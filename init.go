package decide

// init handle setting up a default lexer and parser.
func init() {
	lexer = &Lexer{}

	lexer.Specials = map[string]bool{
		"=": true,
		"<": true,
		">": true,
		"!": true,
		"~": true,
		"&": true,
		"|": true,
	}

	lexer.Delims = map[string]bool{
		" ": true,
	}

	lexer.Scopes = map[string]bool{
		"(": true,
		")": true,
	}

	parser = &Parser{}

	parser.Ops = map[string]Operator{
		"=":  Eq,
		"!=": Neq,
		">":  Gt,
		"<":  Lt,
		">=": Gteq,
		"<=": Lteq,
		"~=": Matches,
		"&&": And,
		"||": Or,
	}
}
