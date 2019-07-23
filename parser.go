package gs_mkdown

import "errors"

// Theory:
// Split the tokens based on newlines into paragraph and line items
// Then parse the new lines for the "styling" markdown

type Node struct {
	nType    string
	value    string
	consumed int
	nodes    []Node
}

const (
	NullNode = "NULL"
	BoldNode = "BOLD"
	BodyNode = "BODY"

	NoEOFTokenError = "No EOF token found, bailing early"
)

func NewNode(nType string, value string, consumed int, nodes []Node) Node {
	return Node{nType: nType, value: value, consumed: consumed}
}

func BodyParser(tl TokenList) (Node, error) {
	// If last token isn't EOF, throw an error
	if t, _ := tl.Get(tl.Length() - 1); t.TokenType() != EOFType {
		return Node{}, errors.New(NoEOFTokenError)
	}

	// Set up consumed and Node List
	tokenLength := tl.Length()
	consumed := 0
	var nl []Node

	for {
		// If we have consumed all the tokens, exit
		if consumed == tokenLength {
			break
		}

		// We are checking a couple things with the first tokn
		t1, err := tl.Get(0)
		if err != nil {
			return Node{}, err
		}

		// If newline is the first token, get rid of it, add one to consumed
		// May do something more elegant in the future
		if t1.TokenType() == NewlineType {
			tl = tl.Slice(1, tl.Length())
			consumed++
		}

		// If first token is EOFType, we are done. Return what we have as a
		// body node
		if t1.TokenType() == EOFType {
			consumed++
			break
		}

		// Process Newline for Ordered List or Un-Ordered List
		// If first node is Dash or Number
		// Return new nodes, non-processed tokens, and consumed amount
		// Plus consumed to consumed amount

		// Process Paragraph Nodes
		// Make a paragraph Node
		// Return new nodes, non-processed tokens, and consumed amount
		// Plus to consumed amount

	}
	return NewNode(BodyNode, "", consumed, nl), nil
}
