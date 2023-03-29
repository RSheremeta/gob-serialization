package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-serialization/encode_decode"
	. "github.com/RSheremeta/gob-serialization/structs"
)

func BenchmarkEncodePtrSlice(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeYAML},

		{name: "type=GOB struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeYAML},

		{name: "type=GOB struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeYAML},

		{name: "type=GOB struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeYAML},
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

func BenchmarkDecodePtrSlice(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	// Tiny
	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*Tiny
				bm.decFn(bytes, &result)
			}
		})
	}

	// Medium
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumPtrsSlc(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*Medium
				bm.decFn(bytes, &result)
			}
		})
	}

	// Big
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=big", target: NewBigPtrsSlc(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*Big
				bm.decFn(bytes, &result)
			}
		})
	}

	// Huge
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeXML, decFn: ed.DecodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugePtrsSlc(), encFn: ed.EncodeYAML, decFn: ed.DecodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*Big
				bm.decFn(bytes, &result)
			}
		})
	}
}
