package decide

import "log"

type Lexer struct {
	Specials map[string]bool
	Delims   map[string]bool
	Scopes   map[string]bool
}

var lexer *Lexer

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
}

func GetLexer() *Lexer {
	return lexer
}

func (l *Lexer) Lex(target string) []string {
	var haveSpecial bool
	var inString bool
	length := len(target)
	if length == 0 {
		return nil
	}
	var start int
	var index int
	result := make([]string, length)

	for i := 0; i < len(target); i++ {
		log.Printf("Evaluating char: %s", string(target[i]))
		if string(target[i]) == "\"" {
			if inString {
				if isEscaped(target, i) {
					continue
				}

				inString = false
				result[index] = string(target[start : i+1])
				index++
				start = i + 1
				continue
			}
			inString = true
			start = i
		}

		if inString {
			continue
		}

		if isDelim(string(target[i])) {
			if start != i {
				log.Printf("Saving word: %s$", string(target[start:i]))
				result[index] = string(target[start:i])
				index++
			}
			start = i + 1
			continue
		}

		if isScope(string(target[i])) {
			log.Println("FOUND PAREN")
			log.Printf("INDEX IS: %d", index)
			if start != i {
				log.Printf("Saving word: %s$", string(target[start:i]))
				result[index] = string(target[start:i])
				index++
			}
			log.Printf("SAVING PAREN: %s")
			result[index] = string(target[i])
			start = i + 1
			index++
			continue
		}

		if isSpecial(string(target[i])) {
			if !haveSpecial {
				if start != i {
					log.Printf("Saving word: %s$", string(target[start:i]))
					result[index] = string(target[start:i])
					start = i
					index++
				}

				if i == len(target)-1 {
					result[index] = string(target[i])
				}

				haveSpecial = true
			}
			continue
		}

		if haveSpecial {
			if start != i {
				log.Printf("Saving word: %s$", string(target[start:i]))
				result[index] = string(target[start:i])
				start = i
				index++
			}
		}

		if i == len(target)-1 {
			result[index] = string(target[start:len(target)])
		}
		haveSpecial = false
	}

	// It's faster to grab a subslice at the end then continuously append.
	return result[0:index]
}

func isSpecial(char string) bool {
	_, ok := lexer.Specials[char]
	return ok
}

func isDelim(char string) bool {
	_, ok := lexer.Delims[char]
	return ok
}

func isScope(char string) bool {
	_, ok := lexer.Scopes[char]
	return ok
}

func isEscaped(source string, index int) bool {
	if string(source[index-1]) == "\\" {
		return true
	}
	return false
}
