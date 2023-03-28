package main

import "testing"

func BenchmarkEncodeSingle(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: newTiny(), encFn: encodeGob},
		{name: "type=JSON struct_size=tiny", target: newTiny(), encFn: encodeJSON},
		{name: "type=XML struct_size=tiny", target: newTiny(), encFn: encodeXML},
		{name: "type=YAML struct_size=tiny", target: newTiny(), encFn: encodeYAML},
		{name: "type=GOB struct_size=medium", target: newMedium(), encFn: encodeGob},
		{name: "type=JSON struct_size=medium", target: newMedium(), encFn: encodeJSON},
		{name: "type=XML struct_size=medium", target: newMedium(), encFn: encodeXML},
		{name: "type=YAML struct_size=medium", target: newMedium(), encFn: encodeYAML},
		{name: "type=GOB struct_size=huge", target: newHuge(), encFn: encodeGob},
		{name: "type=JSON struct_size=huge", target: newHuge(), encFn: encodeJSON},
		{name: "type=XML struct_size=huge", target: newHuge(), encFn: encodeXML},
		{name: "type=YAML struct_size=huge", target: newHuge(), encFn: encodeYAML},
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

func BenchmarkDecodeSingle(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	// Tiny
	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=tiny", target: newTiny(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=tiny", target: newTiny(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=tiny", target: newTiny(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=tiny", target: newTiny(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=medium", target: newMedium(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=medium", target: newMedium(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=medium", target: newMedium(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=medium", target: newMedium(), encFn: encodeYAML, decFn: decodeYAML},
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
		{name: "type=GOB struct_size=huge", target: newHuge(), encFn: encodeGob, decFn: decodeGob},
		{name: "type=JSON struct_size=huge", target: newHuge(), encFn: encodeJSON, decFn: decodeJSON},
		{name: "type=XML struct_size=huge", target: newHuge(), encFn: encodeXML, decFn: decodeXML},
		{name: "type=YAML struct_size=huge", target: newHuge(), encFn: encodeYAML, decFn: decodeYAML},
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
