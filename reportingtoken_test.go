package whatsmeow

import "testing"

func TestExtractReportingTokenContentRejectsMalformedLengths(t *testing.T) {
	config := []reportingField{{FieldNumber: 1}}
	for _, malformed := range [][]byte{
		{0x0a, 0xff},
		{0x0a, 0xff, 0xff, 0xff, 0xff, 0x0f},
	} {
		if got := extractReportingTokenContent(malformed, config); got != nil {
			t.Fatalf("expected malformed input to be rejected, got %x", got)
		}
	}
}
