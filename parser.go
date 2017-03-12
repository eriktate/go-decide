package decide

type Parser struct {
	Ops map[string]Operator
}

var parser *Parser
