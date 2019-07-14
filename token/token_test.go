package token

import (
	"reflect"
	"testing"
)

func TestNewToken(t *testing.T) {
	_, err := NewToken("", "derp")
	if err == nil {
		t.Errorf("An empty t_type should throw an error")
	}

	t1, err := NewToken(NIL_TYPE, "")
	name1 := reflect.TypeOf(t1).Name()
	if name1 != "nilToken" {
		t.Errorf("Token with passed NIL_TYPE should be nilToken")
	}

}

func TestLength(t *testing.T) {
	t1, _ := NewToken("EOF", "derp")
	if t1.Length() != 4 {
		t.Errorf("Length is broken, should be 4")
	}

	t2, _ := NewToken("NEWLINE", "")
	if t2.Length() != 0 {
		t.Errorf("Length is broken, should be 2")
	}
}

func TestTokenIsNil(t *testing.T) {
	t1, _ := NewToken("EOF", "")
	if t1.IsNil() {
		t.Errorf("Token should not be nil")
	}
}

func TestTokenIsPresent(t *testing.T) {
	t1, _ := NewToken("EOF", "")
	if !t1.IsPresent() {
		t.Errorf("Token should be present")
	}
}

func TestTokenValue(t *testing.T) {
	t1, _ := NewToken(TEXT_TYPE, "Boop")
	if t1.Value() != "Boop" {
		t.Errorf("Should return value of original test string")
	}
}

func TestNilTokenIsTypeNil(t *testing.T) {
	t1, _ := NewNilToken()
	if t1.t_type != NIL_TYPE {
		t.Errorf("Token should be t_type of NIL_TYPE")
	}
}

func TestNilTokenLength(t *testing.T) {
	t1, _ := NewNilToken()
	if t1.Length() != 0 {
		t.Errorf("Length is broken, should be 4")
	}
}

func TestNilTokenIsNil(t *testing.T) {
	t1, _ := NewNilToken()
	if !t1.IsNil() {
		t.Errorf("Token should not be nil")
	}
}

