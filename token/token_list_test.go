package token

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var tl TokenList
	token, _ := NewToken(UNDERSCORE_TYPE, "_")
	tl = tl.Append(token)

	if tl.Length() != 1 {
		t.Errorf("Token should have been appending to token list")
	}

	if tt, _ := tl.Get(0); tt.TokenType() != UNDERSCORE_TYPE {
		t.Errorf("Token at index 1 should be UNDERSCORE_TYPE")
	}

	token, _ = NewToken(STAR_TYPE, "*")
	tl = tl.Append(token)

	if tl.Length() != 2 {
		t.Errorf("New token should have been appending to token list")
	}

	if tt, _ := tl.Get(1); tt.TokenType() != STAR_TYPE {
		t.Errorf("Token at index 1 should be STAR_TYPE")
	}

	var newList TokenList
	token, _ = NewToken(TEXT_TYPE, "Hello")
	newList = newList.Append(token)
	newList = newList.Append(tl.All()...)

	if newList.Length() != 3 {
		t.Errorf("Append should be able to handle multiple tokens")
	}
}

func TestSlice(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: TEXT_TYPE, value: "A Silly String"},
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: STAR_TYPE, value: "*"},
	)

	newList := tl.Slice(2, 4)
	if newList.Length() != 2 {
		t.Errorf("The returns list should be 2")
	}

	if token, _ := newList.Get(1); token.TokenType() != STAR_TYPE {
		t.Errorf("Index 1 should be STAR_TYPE")
	}
}

func TestPeek(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: TEXT_TYPE, value: "A Silly String"},
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: STAR_TYPE, value: "*"},
	)

	if tl.Peek([]string{UNDERSCORE_TYPE, UNDERSCORE_TYPE, TEXT_TYPE}) {
		t.Errorf("Should be false due to types not matching")
	}

	if !tl.Peek([]string{UNDERSCORE_TYPE, TEXT_TYPE, UNDERSCORE_TYPE}) {
		t.Errorf("Should be true due to matching types")
	}
}

func TestFindTokenType(t *testing.T) {
	var tl TokenList
	tl = tl.Append(
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: TEXT_TYPE, value: "A Silly String"},
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: NEWLINE_TYPE, value: "\n"},
		BaseToken{t_type: STAR_TYPE, value: "*"},
	)

	// Should find the newline token at index 3
	if i, _ := tl.FindTokenType(NEWLINE_TYPE, 0); i != 3 {
		t.Errorf("NEWLINE_TYPE should be found at index of 3, not %d", i)
	}

	// Should return -1 and error if index is out of range
	if i, err := tl.FindTokenType(NEWLINE_TYPE, 5); i != -1 || err == nil {
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
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: TEXT_TYPE, value: "A Silly String"},
		BaseToken{t_type: UNDERSCORE_TYPE, value: "_"},
		BaseToken{t_type: NEWLINE_TYPE, value: "\n"},
		BaseToken{t_type: STAR_TYPE, value: "*"},
	)

	// Should find the newline token at index 3
	if !tl.PeekAt(UNDERSCORE_TYPE, 2) {
		t.Errorf("Should return true")
	}

	// Should return false if index is out of range
	if tl.PeekAt(UNDERSCORE_TYPE, 5) {
		t.Errorf("Should return false")
	}

	if !tl.PeekAt(TEXT_TYPE, 1) {
		t.Errorf("Should return true")
	}
}
