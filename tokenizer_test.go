package gs_mkdown

import (
	"testing"
)

func TestSimpleScannerFromString(t *testing.T) {
	// Loop through TOKEN_TYPES map and a test that correct types are returned
	for key, value := range TOKEN_TYPES {
		t1, _ := SimpleScanner{}.FromString(key + "A Bunch of Text你好_")

		if t1.TokenType() != value {
			t.Errorf("Returned Token should be type of %s", value)
		}
	}

	// An empty string should return a nil type
	t2, _ := SimpleScanner{}.FromString("")

	if t2.TokenType() != NilType {
		t.Errorf("Returned Token should be NilType")
	}

	// An unknown first rune should return a NilType
	t3, _ := SimpleScanner{}.FromString("你好")
	if t3.TokenType() != NilType {
		t.Errorf("Returned Token should be NilType")
	}
}

func TestTextScannerFromString(t *testing.T) {
	// Texting a "normal" string
	t1, _ := TextScanner{}.FromString("A Long String of Text")

	if t1.TokenType() != TextType {
		t.Errorf("Returned Token should be TextType, not %s", t1.TokenType())
	}

	if t1.Value() != "A Long String of Text" {
		t.Errorf("Returned value should be original string, returned value was %s", t1.Value())
	}

	// Testing when the string has "special" characters in it
	// Should return nil, because it would pick up the special character
	t2, _ := TextScanner{}.FromString("_Not*asMuch\nText")

	if t2.TokenType() != NilType {
		t.Errorf("Returned Token should be NilType, not %s", t2.TokenType())
	}
}

func TestScanOneToken(t *testing.T) {
	t1, _ := ScanOneToken("_Boop")

	if t1.TokenType() != UnderscoreType {
		t.Errorf("Returned Token should be UnderscoreType, not %s", t1.TokenType())
	}

	t2, _ := ScanOneToken("Boop_")

	if t2.TokenType() != TextType {
		t.Errorf("Returned Token should be TextType, not %s", t2.TokenType())
	}

	if t2.Value() != "Boop" {
		t.Errorf("Returned value should be original string, returned value was %s", t2.Value())
	}
}

func TestTokenize(t *testing.T) {
	tokenList, _ := Tokenize("_Boop*Test")

	token, _ := tokenList.Get(0)
	if token.TokenType() != UnderscoreType {
		t.Errorf("Returned Token should be UnderscoreType, not %s", token.TokenType())
	}

	token, _ = tokenList.Get(1)
	if token.TokenType() != TextType {
		t.Errorf("Returned Token should be TextType, not %s", token.TokenType())
	}

	if token.Value() != "Boop" {
		t.Errorf("Returned value should be original string, returned value was %s", token.Value())
	}

	token, _ = tokenList.Get(2)
	if token.TokenType() != StarType {
		t.Errorf("Returned Token should be StarType, not %s", token.TokenType())
	}

	token, _ = tokenList.Get(3)
	if token.TokenType() != TextType {
		t.Errorf("Returned Token should be TextType, not %s", token.TokenType())
	}

	if token.Value() != "Test" {
		t.Errorf("Returned value should be original string, returned value was %s", token.Value())
	}
}

func TestTokenizeMultiLine(t *testing.T) {
	testString := `A Line
		With another line
		A ** a Final line **`

	tokenList, _ := Tokenize(testString)

	if tokenList.Length() != 11 {
		t.Errorf("Returned length should be 11, not %d", tokenList.Length())
	}

	token, _ := tokenList.Get(1)
	if token.TokenType() != NewlineType {
		t.Errorf("Returned Token should be NewlineType, not %s", token.TokenType())
	}

	token, _ = tokenList.Get(10)
	if token.TokenType() != EOFType {
		t.Errorf("Returned Token should be EOF, not %s", token.TokenType())
	}
}
