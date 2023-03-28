package main

import "testing"

func BenchmarkEncodeSlice(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: newTinySlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=tiny", target: newTinySlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=tiny", target: newTinySlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=tiny", target: newTinySlc(), encFn: encodeYAML},
		{name: "type=GOB struct_size=medium", target: newMediumSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=medium", target: newMediumSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=medium", target: newMediumSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=medium", target: newMediumSlc(), encFn: encodeYAML},
		{name: "type=GOB struct_size=huge", target: newHugeSlc(), encFn: encodeGob},
		{name: "type=JSON struct_size=huge", target: newHugeSlc(), encFn: encodeJSON},
		{name: "type=XML struct_size=huge", target: newHugeSlc(), encFn: encodeXML},
		{name: "type=YAML struct_size=huge", target: newHugeSlc(), encFn: encodeYAML},
	}

	/*	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			b.ResetTimer()
			res := bm.encFn(strct)
			_ = res
		})
	}*/

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
		{name: "type=GOB struct_size=tiny", target: newTinySlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=tiny", target: newTinySlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=tiny", target: newTinySlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=tiny", target: newTinySlc(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=medium", target: newMediumSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=medium", target: newMediumSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=medium", target: newMediumSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=medium", target: newMediumSlc(), encFn: encodeYAML, decFn: decodeYAML},
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

	// Huge
	bnchmkrs = []benchmark{
		{name: "type=GOB struct_size=huge", target: newHugeSlc(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=huge", target: newHugeSlc(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=huge", target: newHugeSlc(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: newHugeSlc(), encFn: encodeYAML, decFn: decodeYAML},
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
