package token

import "errors"

type TokenList struct {
	tokens []Token
}

type TokenScanner interface {
	FromString(s string) (Token, error)
}

// Append appends any number of tokens to the token list and
// returns the TokenList
func (t TokenList) Append(tokens ...Token) TokenList {
	for _, token := range tokens {
		t.tokens = append(t.tokens, token)
	}
	return t
}

// Get returns the token at a given index, or returns an erro
func (t TokenList) Get(i int) (Token, error) {
	if i < len(t.tokens) {
		return t.tokens[i], nil
	}
	return BaseToken{}, errors.New("Index out of range")
}

// All returns the tokens list
func (t TokenList) All() []Token {
	return t.tokens
}

// Length returns the number of tokens in the token list
func (t TokenList) Length() int {
	return len(t.tokens)
}

// Peek checks to see if token in list match the passed list of types
// and returns true if they match, or false if they do not
func (t TokenList) Peek(types []string) bool {
	if t.Length() < len(types) {
		return false
	}
	for i, typeTest := range types {
		if typeTest != t.tokens[i].TokenType() {
			return false
		}
	}
	return true
}
