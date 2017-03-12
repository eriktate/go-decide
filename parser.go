package decide

type Parser struct {
	Ops map[string]Operator
}

var parser *Parser

// GetParser returns the default go-decide parser.
func GetParser() *Parser {
	return parser
}

// Parse is the primary method of parsing tokens into an Expression that can be
// evaluated. It provides a nicer API than subParse which does all of the actual work.
func (p *Parser) Parse(tokens []string) (*Expr, error) {
	expr, idx, err := p.subParse(tokens, 0)
	return expr, err
}

// subParse handles all of the heavy lifting involved with parsing tokens into Expressions.
// It recursively explores a given slice of tokens and returns the resulting expression,
// index it left off at and any error that may have occurred during the process.
func (p *Parser) subParse(tokens []string, idx int) {

}
