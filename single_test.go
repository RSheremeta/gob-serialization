package main

import (
	"testing"

	. "github.com/RSheremeta/gob-serialization/structs"
)

func BenchmarkEncodeSingle(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTiny(), encFn: encodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTiny(), encFn: encodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTiny(), encFn: encodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTiny(), encFn: encodeYAML},

		{name: "type=GOB struct_size=medium", target: NewMedium(), encFn: encodeGob},
		{name: "type=JSON struct_size=medium", target: NewMedium(), encFn: encodeJSON},
		{name: "type=XML struct_size=medium", target: NewMedium(), encFn: encodeXML},
		{name: "type=YAML struct_size=medium", target: NewMedium(), encFn: encodeYAML},

		{name: "type=GOB struct_size=big", target: NewBig(), encFn: encodeGob},
		{name: "type=JSON struct_size=big", target: NewBig(), encFn: encodeJSON},
		{name: "type=XML struct_size=big", target: NewBig(), encFn: encodeXML},
		{name: "type=YAML struct_size=big", target: NewBig(), encFn: encodeYAML},

		{name: "type=GOB struct_size=huge", target: NewHuge(), encFn: encodeGob},
		{name: "type=JSON struct_size=huge", target: NewHuge(), encFn: encodeJSON},
		{name: "type=XML struct_size=huge", target: NewHuge(), encFn: encodeXML},
		{name: "type=YAML struct_size=huge", target: NewHuge(), encFn: encodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			b.ResetTimer()
			// b.N is a built-in var which specifies num of iterations run (?)
			for i := 0; i < b.N; i++ {
				res := bm.encFn(strct)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSingle(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	// Tiny
	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=tiny", target: NewTiny(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTiny(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTiny(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTiny(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result Tiny
				bm.decFn(bytes, &result)
			}
		})
	}

	// Medium
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=medium", target: NewMedium(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=medium", target: NewMedium(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=medium", target: NewMedium(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=medium", target: NewMedium(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result Medium
				bm.decFn(bytes, &result)
			}
		})
	}

	// Huge
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=big", target: NewBig(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=big", target: NewBig(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=big", target: NewBig(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: NewBig(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result Big
				bm.decFn(bytes, &result)
			}
		})
	}

	// Huge
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=huge", target: NewHuge(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=huge", target: NewHuge(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=huge", target: NewHuge(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: NewHuge(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result Huge
				bm.decFn(bytes, &result)
			}
		})
	}
}
