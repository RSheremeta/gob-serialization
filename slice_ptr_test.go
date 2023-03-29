package main

import (
	"testing"

	. "github.com/RSheremeta/gob-serialization/structs"
)

func BenchmarkEncodePtrSlice(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=big", target: NewBigPtrsSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=big", target: NewBigPtrsSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=big", target: NewBigPtrsSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=big", target: NewBigPtrsSlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeYAML},
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
		{name: "type=GOB struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinyPtrsSlc(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumPtrsSlc(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=big", target: NewBigPtrsSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=big", target: NewBigPtrsSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=big", target: NewBigPtrsSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=big", target: NewBigPtrsSlc(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugePtrsSlc(), encFn: encodeYAML, decFn: decodeYAML},
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
