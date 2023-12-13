package shorting

import (
	"strings"
	"testing"
)

func TestGenerateIdentifier(t *testing.T) {
	identifier := GenerateIdentifier()
	if len(identifier) != Length {
		t.Errorf("identifier len must be %v, got %v", Length, len(identifier))
	}
	if len(strings.Trim(identifier, CharSet)) != 0 {
		t.Errorf("identifier must contain only this symbols: %v", CharSet)
	}
}
