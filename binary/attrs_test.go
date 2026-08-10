package binary

import "testing"

func TestAttrUtilityRejectsOutOfRangeIntegers(t *testing.T) {
	tests := []struct {
		name  string
		value string
		get   func(*AttrUtility) bool
	}{
		{name: "uint16 overflow", value: "65536", get: func(au *AttrUtility) bool { _, ok := au.GetUint16("value", true); return ok }},
		{name: "uint32 overflow", value: "4294967296", get: func(au *AttrUtility) bool { _, ok := au.GetUint32("value", true); return ok }},
		{name: "int overflow", value: "999999999999999999999999999999999", get: func(au *AttrUtility) bool { _, ok := au.GetInt("value", true); return ok }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &AttrUtility{Attrs: Attrs{"value": tt.value}}
			if tt.get(au) {
				t.Fatal("expected the conversion to fail")
			}
			if len(au.Errors) != 1 {
				t.Fatalf("expected one conversion error, got %d", len(au.Errors))
			}
		})
	}
}
