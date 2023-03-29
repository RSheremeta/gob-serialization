package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-serialization/encode_decode"
	. "github.com/RSheremeta/gob-serialization/structs"
)

func BenchmarkEncodeSingle(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTiny(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTiny(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTiny(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTiny(), encFn: ed.EncodeYAML},

		{name: "type=GOB struct_size=medium", target: NewMedium(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=medium", target: NewMedium(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=medium", target: NewMedium(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=medium", target: NewMedium(), encFn: ed.EncodeYAML},

		{name: "type=GOB struct_size=big", target: NewBig(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=big", target: NewBig(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=big", target: NewBig(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=big", target: NewBig(), encFn: ed.EncodeYAML},
		{name: "type=GOB struct_size=huge", target: NewHuge(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=huge", target: NewHuge(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=huge", target: NewHuge(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=huge", target: NewHuge(), encFn: ed.EncodeYAML},
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
		{name: "type=GOB struct_size=tiny", target: NewTiny(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTiny(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTiny(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTiny(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
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
		{name: "type=GOB struct_size=medium", target: NewMedium(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=medium", target: NewMedium(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=medium", target: NewMedium(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=medium", target: NewMedium(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
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
		{name: "type=GOB struct_size=big", target: NewBig(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=big", target: NewBig(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=big", target: NewBig(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=huge", target: NewBig(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
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
		{name: "type=GOB struct_size=huge", target: NewHuge(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=huge", target: NewHuge(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=huge", target: NewHuge(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=huge", target: NewHuge(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
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
