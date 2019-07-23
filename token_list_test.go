package gs_mkdown

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var tl TokenList
	token, _ := NewToken(UnderscoreType, "_")
	tl = tl.Append(token)

	if tl.Length() != 1 {
		t.Errorf("Token should have been appending to token list")
	}

	if tt, _ := tl.Get(0); tt.TokenType() != UnderscoreType {
		t.Errorf("Token at index 1 should be UnderscoreType")
	}

	token, _ = NewToken(StarType, "*")
	tl = tl.Append(token)

	if tl.Length() != 2 {
		t.Errorf("New token should have been appending to token list")
	}

	if tt, _ := tl.Get(1); tt.TokenType() != StarType {
		t.Errorf("Token at index 1 should be StarType")
	}

	var newList TokenList
	token, _ = NewToken(TextType, "Hello")
	newList = newList.Append(token)
	newList = newList.Append(tl.All()...)

	if newList.Length() != 3 {
		t.Errorf("Append should be able to handle multiple tokens")
	}
}

func TestSlice(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: TextType, value: "A Silly String"},
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: StarType, value: "*"},
	)

	newList := tl.Slice(2, 4)
	if newList.Length() != 2 {
		t.Errorf("The returns list should be 2")
	}

	if token, _ := newList.Get(1); token.TokenType() != StarType {
		t.Errorf("Index 1 should be StarType")
	}
}

func TestPeek(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: TextType, value: "A Silly String"},
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: StarType, value: "*"},
	)

	if tl.Peek([]string{UnderscoreType, UnderscoreType, TextType}) {
		t.Errorf("Should be false due to types not matching")
	}

	if !tl.Peek([]string{UnderscoreType, TextType, UnderscoreType}) {
		t.Errorf("Should be true due to matching types")
	}
}

func TestFindTokenType(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: TextType, value: "A Silly String"},
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: StarType, value: "*"},
	)

	// Should find the newline token at index 3
	if i, _ := tl.FindTokenType(NewlineType, 0); i != 3 {
		t.Errorf("NewlineType should be found at index of 3, not %d", i)
	}

	// Should return -1 and error if index is out of range
	if i, err := tl.FindTokenType(NewlineType, 5); i != -1 || err == nil {
		t.Errorf("Should return -1 and an error when index is out of range")
	}

	// Should return -1 and error if index is out of range
	if i, err := tl.FindTokenType("ALYINGTYPETHATISFAKE", 0); i != -1 || err == nil {
		t.Errorf("Should return -1 and an error when type isn't found")
	}
}

func TestPeekAt(t *testing.T) {
	var tl TokenList

	tl = tl.Append(
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: TextType, value: "A Silly String"},
		BaseToken{tType: UnderscoreType, value: "_"},
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: StarType, value: "*"},
	)

	// Should find the newline token at index 3
	if !tl.PeekAt(UnderscoreType, 2) {
		t.Errorf("Should return true")
	}

	// Should return false if index is out of range
	if tl.PeekAt(UnderscoreType, 5) {
		t.Errorf("Should return false")
	}

	if !tl.PeekAt(TextType, 1) {
		t.Errorf("Should return true")
	}
}
