package gs_mkdown

import (
	"errors"
)

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
	if i < t.Length() && i >= 0 {
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

// Slice returns a new TokenList with the underlying slice of tokens
func (t TokenList) Slice(start int, end int) TokenList {
	return TokenList{tokens: t.All()[start:end]}
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

const (
	INDEX_ERROR = "Index out of range"
	TYPE_ERROR  = "Could not find token of that type"
)

// FindTokenType returns the index of the first instance of tType passed
// or returns -1 and an error if index is out of range or token cannot be found
func (t TokenList) FindTokenType(tType string, index int) (int, error) {
	if t.Length() <= index {
		return -1, errors.New(INDEX_ERROR)
	}
	for i, token := range t.All()[index:] {
		if token.TokenType() == tType {
			return i + index, nil
		}
	}
	return -1, errors.New(TYPE_ERROR)
}

// PeekAt returns true or false if the token type matches the token at index
func (t TokenList) PeekAt(tType string, index int) bool {
	if t.Length() <= index {
		return false
	}

	if t.All()[index].TokenType() == tType {
		return true
	}

	return false
}
