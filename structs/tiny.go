package structs

type Tiny struct {
	TinyID  int    `json:"tiny_id" xml:"tiny_id" yaml:"tiny_id"`
	SomeStr string `json:"some_str" xml:"some_str" yaml:"some_str"`
}

func NewTiny() Tiny {
	return Tiny{
		TinyID:  1,
		SomeStr: "foo bar",
	}
}

func NewTinySlc() []Tiny {
	size := 5
	res := make([]Tiny, size)

	for i := 0; i < size; i++ {
		res = append(res, NewTiny())
	}

	return res
}

func NewTinyPtrsSlc() []*Tiny {
	size := 5
	var res []*Tiny

	for i := 0; i < size; i++ {
		var tiny = &Tiny{
			TinyID:  1,
			SomeStr: "foo bar",
		}

		res = append(res, tiny)
	}

	return res
}
