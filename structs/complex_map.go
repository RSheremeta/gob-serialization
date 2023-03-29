package structs

// for Gob and JSON usage only!
type complexMap map[string]Huge

type ComplexAndHuge struct {
	ComplexMapsSlice []complexMap
}

func NewComplexAndHuge() ComplexAndHuge {
	return ComplexAndHuge{
		ComplexMapsSlice: makeComplexMapSlice(),
	}
}

func NewComplexAndHugeSlc() []ComplexAndHuge {
	size := 10
	var res []ComplexAndHuge

	for i := 0; i < size; i++ {
		res = append(res, NewComplexAndHuge())
	}

	return res
}

func NewComplexAndHugePtrSlc() []*ComplexAndHuge {
	slc := NewComplexAndHugeSlc()
	var res []*ComplexAndHuge

	for i := 0; i < len(slc); i++ {
		var item = slc[i]
		res = append(res, &item)
	}

	return res
}

func makeComplexMapSlice() []complexMap {
	size := 10
	var res []complexMap

	for i := 0; i < size; i++ {
		val := complexMap{
			"1": NewHuge(),
			"2": NewHuge(),
			"3": NewHuge(),
			"4": NewHuge(),
			"5": NewHuge(),
		}

		res = append(res, val)
	}

	return res
}
