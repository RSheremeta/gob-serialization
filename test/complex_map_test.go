package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-serialization/encode_decode"
	. "github.com/RSheremeta/gob-serialization/structs"
)

// for Gob and JSON usage only!

// Single

func BenchmarkEncodeSingleComplexMap(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHuge(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHuge(), encFn: ed.EncodeJSON},
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

func BenchmarkDecodeSingleComplexMap(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHuge(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHuge(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result ComplexAndHuge
				bm.decFn(bytes, &result)
			}
		})
	}
}

// Slice

func BenchmarkEncodeSliceComplexMap(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHugeSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHugeSlc(), encFn: ed.EncodeJSON},
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

func BenchmarkDecodeSliceComplexMap(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHugeSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHugeSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []ComplexAndHuge
				bm.decFn(bytes, &result)
			}
		})
	}
}

// Slice Pointers

func BenchmarkEncodePtrSliceComplexMap(b *testing.B) {
	bnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHugePtrSlc(), encFn: ed.EncodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHugePtrSlc(), encFn: ed.EncodeJSON},
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

func BenchmarkDecodePtrSliceComplexMap(b *testing.B) {
	type benchmark struct {
		name   string
		target any
		encFn  func(any) []byte
		decFn  func([]byte, any)
	}

	bnchmkrs := []benchmark{
		{name: "type=GOB struct_size=huge_complex_map", target: NewComplexAndHugePtrSlc(), encFn: ed.EncodeGob, decFn: ed.DecodeGob},
		{name: "type=JSON struct_size=huge_complex_map", target: NewComplexAndHugePtrSlc(), encFn: ed.EncodeJSON, decFn: ed.DecodeJSON},
	}

	for _, bm := range bnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			bytes := bm.encFn(strct)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*ComplexAndHuge
				bm.decFn(bytes, &result)
			}
		})
	}
}
