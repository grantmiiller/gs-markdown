package gs_mkdown

import (
	"errors"
	"strings"
)

type SimpleScanner struct{}

const (
	NilType        = "NIL"
	EOFType        = "EOF"
	TextType       = "TEXT"
	UnderscoreType = "UNDERSCORE"
	StarType       = "STAR"
	NewlineType    = "NEWLINE"
	HashType       = "HASH"
	DashType       = "DASH"
)

var TOKEN_TYPES = map[string]string{
	"_":  UnderscoreType,
	"*":  StarType,
	"#":  HashType,
	"\n": NewlineType,
}

// FromString grabs the first character in string and attempts to find a
// known token from TOKEN_TYPES
func (scanner SimpleScanner) FromString(s string) (Token, error) {
	if len([]rune(s)) > 0 {
		value := string([]rune(s)[0])
		// If our map of values contains the character, return a Token

		if tType, ok := TOKEN_TYPES[value]; ok {
			return NewToken(tType, value)
		}
	}
	return NewNilToken()
}

type TextScanner struct{}

func (scanner TextScanner) FromString(s string) (Token, error) {
	var text []string

	// Loop through the runes/characters
	for _, rune := range s {
		// Make sure that the string we are using does not contain special characters
		t, err := SimpleScanner{}.FromString(string(rune))
		// If no special characters and no errors, append to string
		if err == nil && t.IsNil() {
			text = append(text, string(rune))
			// Else break out, because we may have a completed text token
		} else {
			break
		}
	}
	if len(text) > 0 {
		return NewToken(TextType, strings.Join(text, ""))
	}
	return NewNilToken()
}

var scanners = []TokenScanner{SimpleScanner{}, TextScanner{}}

func Tokenize(s string) (TokenList, error) {
	var tokenList TokenList
	if s == "" {
		t, err := NewToken(EOFType, "")
		if err != nil {
			return tokenList, err
		}
		return tokenList.Append(t), nil
	}
	t, err := ScanOneToken(s)
	tokenList = tokenList.Append(t)
	newList, err := Tokenize(string([]rune(s)[t.Length():]))
	return tokenList.Append(newList.All()...), err
}

func ScanOneToken(s string) (Token, error) {
	for _, scanner := range scanners {
		t, err := scanner.FromString(s)
		if err == nil && !t.IsNil() {
			return t, nil
		}
	}
	return BaseToken{}, errors.New("Could not match token")
}
