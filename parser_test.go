package gs_mkdown

import (
	"testing"
)

func TestBodyParserLoneNewLine(t *testing.T) {
	var tl TokenList

	tl = tl.Append(
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: EOFType, value: ""},
	)

	nl, _ := BodyParser(tl)
	if nl.consumed != 2 {
		t.Errorf("consumed should be 2, not %d", nl.consumed)
	}

	if nl.nType != BodyNode {
		t.Errorf("Should have returned a BodyType node")
	}
}

func TestBodyParserNoEOF(t *testing.T) {
	var tl TokenList

	tl = tl.Append(
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: NewlineType, value: "\n"},
	)

	_, err := BodyParser(tl)
	if err == nil {
		t.Errorf("Should have thrown an error")
	}
}

func TestTextParser(t *testing.T) {
	var tl TokenList

	tl = tl.Append(
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
	if nl[1].nType != BoldNode || nl[1].value != "EPIC" {
		t.Errorf("Node should equal BoldNode, not %s, and have value of EPIC, not %s", nl[1].nType, nl[1].value)
	}
}

func TestListParserUnordered(t *testing.T) {
	var tl TokenList

	tl = tl.Append(
		BaseToken{tType: DashType, value: "-"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: TextType, value: "EPIC"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: StarType, value: "*"},
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: DashType, value: "-"},
		BaseToken{tType: TextNode, value: "Testing"},
		BaseToken{tType: NewlineType, value: "\n"},
		BaseToken{tType: DashType, value: "-"},
		BaseToken{tType: TextNode, value: "Testing"},
		BaseToken{tType: EOFType, value: ""},
	)
	node, _ := ListParser(tl)
	if len(node.nodes) != 3 {
		t.Errorf("There should be 3 nodes, not %d", len(node.nodes))
	}

	if node.nodes[0].nodes[0].nType != BoldNode {
		t.Errorf("Should be a bold node, not %s", node.nodes[0].nodes[0].value)
	}
}
