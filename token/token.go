package token

import (
	"errors"
	"fmt"
)

type Token interface {
	ToString()
	Length() int
	IsNil() bool
	IsPresent() bool
	TokenType() string
	Value() string
}

type BaseToken struct {
	t_type string
	value  string
}

// NewToken returns a new BaseToken or NilToken, depending on passed value of
// t_type
func NewToken(t_type string, value string) (Token, error) {
	if t_type == "" {
		return BaseToken{}, errors.New("A token type is required")
	}

	if t_type == NIL_TYPE {
		t, err := NewNilToken()
		if err != nil {
			return BaseToken{}, errors.New("A token type is required")
		}
		return t, nil
	}

	t := BaseToken{t_type, value}
	return t, nil
}

// ToString prints the properties of the token
func (t BaseToken) ToString() {
	fmt.Printf("<type: %s, value: %s>\n", t.t_type, t.value)
}

// Length returns how many characters are in the value property string
func (t BaseToken) Length() int {
	return len([]rune(t.value))
}

// TokenType returns the value in t_type property
func (t BaseToken) TokenType() string {
	return t.t_type
}

// Value returns the value in value property
func (t BaseToken) Value() string {
	return t.value
}

// IsNil returns if the Token is Nil
func (t BaseToken) IsNil() bool {
	return false
}

func (t BaseToken) IsPresent() bool {
	return true
}

type nilToken struct {
	BaseToken
}

func NewNilToken() (nilToken, error) {
	return nilToken{BaseToken{NIL_TYPE, ""}}, nil
}

func (t nilToken) IsNil() bool {
	return true
}
