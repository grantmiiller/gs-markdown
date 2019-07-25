package gs_mkdown

import "errors"

// TODO:
// - Create Style Parser, should contain:
// -- Create Parser for bold
// -- Create Parser for italics
// - Create Parser for line items

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
	NilNode  = "NIL"
	TextNode = "TEXT"

	UnorderedNode = "UNORDERED"
	OrderedNode   = "ORDERED"

	NoEOFTokenError      = "No EOF token found, bailing early"
	UnexpectedTokenError = "Ran into a token not able to handle"
)

func NewNode(nType string, value string, consumed int, nodes []Node) Node {
	return Node{nType: nType, value: value, consumed: consumed}
}

// UnorderedListParser parsers tokens for list nodes
// func UnorderedListParser(tl TokenList) (Node, error) {
// 	t1, err := tl.Get(0)
// 	if err != nil || t1.TokenType() != DashType {
// 		// If we don't make a list, return a nil node
// 		return Node{nType: NilNode, value: "", consumed: 0}, nil
// 	}
// 	node := Node{nType: UnorderedNode, value: "", consumed: 0}
// 	if i, _ := tl.FindTokenType(NewlineType); i != -1 {

// 	}

// }

func TextParser(tl TokenList) ([]Node, error) {
	var nl []Node

	for {
		if tl.Length() <= 0 {
			break
		}

		t, _ := tl.Get(0)
		tType := t.TokenType()

		if tType == TextType {
			nl = append(nl, Node{nType: TextNode, value: t.Value(), consumed: 1})
			tl = tl.Slice(1, tl.Length())
			continue
		}

		if tType == StarType || tType == DashType {
			if tl.Peek([]string{tType, tType, TextType, tType, tType}) {
				valToken, _ := tl.Get(2)
				nl = append(nl, Node{nType: BoldNode, value: valToken.Value(), consumed: 5})
				tl = tl.Slice(5, tl.Length())
				continue
			} else {
				nl = append(nl, Node{nType: TextNode, value: t.Value(), consumed: 1})
				tl = tl.Slice(1, tl.Length())
				continue
			}
		}
		return nl, errors.New(UnexpectedTokenError)
	}
	return nl, nil
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
