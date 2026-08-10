package types

import "testing"

func TestParseJIDRejectsOutOfRangeDeviceParts(t *testing.T) {
	for _, invalid := range []string{
		"123.256:1@s.whatsapp.net",
		"123.1:65536@s.whatsapp.net",
		"123:65536@s.whatsapp.net",
	} {
		if _, err := ParseJID(invalid); err == nil {
			t.Fatalf("expected %q to be rejected", invalid)
		}
	}

	jid, err := ParseJID("123.255:65535@s.whatsapp.net")
	if err != nil {
		t.Fatalf("expected maximum values to parse: %v", err)
	}
	if jid.RawAgent != 255 || jid.Device != 65535 {
		t.Fatalf("unexpected parsed JID: %+v", jid)
	}
}
