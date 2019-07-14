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
