package parser

// Theory:
// Split the tokens based on newlines into paragraph and line items
// Then parse the new lines for the "styling" markdown

type Node struct {
	n_type   string
	value    string
	consumed int
}

func NewNode(n_type string, value string, consumed int) (Node, error) {
	return Node{n_type: n_type, value: value, consumed: consumed}, nil
}
