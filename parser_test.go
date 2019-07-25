package gs_mkdown

import "testing"

// Tests that "orphan" newlines are handled
// TODO: add test if newline is orphaned in middle of processing
func bodyParserLoneNewLineTest(t *testing.T) {
	var tl TokenList

	tl.Append(
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: EOFType, value: ""},
	)

	nl, _ := BodyParser(tl)
	if nl.consumed != 2 {
		t.Errorf("consmed should be 2, not %d", nl.consumed)
	}

	if nl.nType != BodyNode {
		t.Errorf("Should have returned a BodyType node")
	}
}

// Test that BodyParser errors when there is no EOF token at the end
func bodyParserNoEOFTest(t *testing.T) {
	var tl TokenList

	tl.Append(
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: NewlineType, value: "\n"},
	)

	_, err := BodyParser(tl)
	if err == nil {
		t.Errorf("Should have thrown an error")
	}
}

func TextParserTest(t *testing.T) {
	var tl TokenList

	tl.Append(
		BaseToken{tType: TextType, value: "This story is"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: TextType, value: "EPIC"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: StarType, value: "*"},
	)

	nl, err := TextParser(tl)

	if err != nil {
		t.Errorf("Should not have errored")
	}

	if nl[0].nType != TextNode {
		t.Errorf("Should equal TextNode, not %s", nl[0].nType)
	}
	if nl[1].nType != BoldNode && nl[1].value != "EPIC" {
		t.Errorf("Node should equal BoldNode, not %s, and have value of EPIC, not %s", nl[1].nType, nl[1].value)
	}
}
