package structs

type Huge struct {
	BigSlice        []Big  `json:"big_slice" xml:"big_slice" yaml:"big_slice"`
	BigPtrSlice     []*Big `json:"big_ptr_slice" xml:"big_ptr_slice" yaml:"big_ptr_slice"`
	BiggerSlice     []Big  `json:"bigger_slice" xml:"bigger_slice" yaml:"bigger_slice"`
	TheBiggestSlice []Big  `json:"the_biggest_slice" xml:"the_biggest_slice" yaml:"the_biggest_slice"`
}

func NewHuge() Huge {
	return Huge{
		BigSlice:        NewBigSlc(),
		BigPtrSlice:     NewBigPtrsSlc(),
		BiggerSlice:     makeBigSlc(256),
		TheBiggestSlice: makeBigSlc(1024),
	}
}

func NewHugeSlc() []Huge {
	size := 50
	res := make([]Huge, size)

	for i := 0; i < size; i++ {
		res = append(res, NewHuge())
	}

	return res
}

func NewHugePtrsSlc() []*Huge {
	size := 50
	var res []*Huge

	for i := 0; i < size; i++ {
		var huge = NewHuge()
		res = append(res, &huge)
	}

	return res
}
