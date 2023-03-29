package main

import (
	"testing"

	. "github.com/RSheremeta/gob-serialization/structs"
)

func BenchmarkEncodeSlice(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTinySlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinySlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinySlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinySlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=medium", target: NewMediumSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumSlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=big", target: NewBigSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=big", target: NewBigSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=big", target: NewBigSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=big", target: NewBigSlc(), encFn: encodeYAML},

		{name: "type=GOB struct_size=huge", target: NewHugeSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugeSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugeSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugeSlc(), encFn: encodeYAML},
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

func BenchmarkDecodeSlice(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	// Tiny
	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=tiny", target: NewTinySlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=tiny", target: NewTinySlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=tiny", target: NewTinySlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=tiny", target: NewTinySlc(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []Tiny
				bm.decFn(bytes, &result)
			}
		})
	}

	// Medium
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=medium", target: NewMediumSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=medium", target: NewMediumSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=medium", target: NewMediumSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=medium", target: NewMediumSlc(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []Medium
				bm.decFn(bytes, &result)
			}
		})
	}

	// Big
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=big", target: NewBigSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=big", target: NewBigSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=big", target: NewBigSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=big", target: NewBigSlc(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []Big
				bm.decFn(bytes, &result)
			}
		})
	}

	// Huge
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=huge", target: NewHugeSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=huge", target: NewHugeSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=huge", target: NewHugeSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: NewHugeSlc(), encFn: encodeYAML, decFn: decodeYAML},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []Huge
				bm.decFn(bytes, &result)
			}
		})
	}
}
