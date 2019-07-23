package gs_mkdown

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
	tType string
	value string
}

// NewToken returns a new BaseToken or NilToken, depending on passed value of
// tType
func NewToken(tType string, value string) (Token, error) {
	if tType == "" {
		return BaseToken{}, errors.New("A token type is required")
	}

	if tType == NilType {
		t, err := NewNilToken()
		if err != nil {
			return BaseToken{}, errors.New("A token type is required")
		}
		return t, nil
	}

	t := BaseToken{tType, value}
	return t, nil
}

// ToString prints the properties of the token
func (t BaseToken) ToString() {
	fmt.Printf("<type: %s, value: %s>\n", t.tType, t.value)
}

// Length returns how many characters are in the value property string
func (t BaseToken) Length() int {
	return len([]rune(t.value))
}

// TokenType returns the value in tType property
func (t BaseToken) TokenType() string {
	return t.tType
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
	return nilToken{BaseToken{NilType, ""}}, nil
}

func (t nilToken) IsNil() bool {
	return true
}
