package structs

import "time"

type Big struct {
	BigID          int       `json:"big_id" xml:"big_id" yaml:"big_id"`
	Age            uint      `json:"age" xml:"age" yaml:"age"`
	Country        *string   `json:"country" xml:"country" yaml:"country"`
	CountryBytes   []byte    `json:"country_bytes" xml:"country_bytes" yaml:"country_bytes"`
	CreatedAt      time.Time `json:"created_at" xml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" xml:"updated_at" yaml:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" xml:"deleted_at" yaml:"deleted_at"`
	TinyField      Tiny      `json:"tiny_field" xml:"tiny_field" yaml:"tiny_field"`
	MediumField    Medium    `json:"medium_field" xml:"medium_field" yaml:"medium_field"`
	MediumFieldPtr *Medium   `json:"medium_field_ptr" xml:"medium_field_ptr" yaml:"medium_field_ptr"`
}

func NewBig() Big {
	ua := "Ukraine"
	uab := []byte(ua)
	t := time.Now()

	ln := "bar"
	mmb := 10.24

	return Big{
		BigID:        1,
		Age:          uint(26),
		Country:      &ua,
		CountryBytes: uab,
		CreatedAt:    t.AddDate(0, 0, -2),
		UpdatedAt:    t.AddDate(0, 0, -1),
		DeletedAt:    t,
		TinyField:    NewTiny(),
		MediumField:  NewMedium(),
		MediumFieldPtr: &Medium{
			MediumID:           1,
			FirstName:          "foo",
			LastName:           &ln,
			MediumMoneyBalance: &mmb,
			CreatedAt:          time.Now(),
			TinyField:          NewTiny(),
		},
	}
}

func NewBigSlc() []Big {
	return makeBigSlc(50)
}

func NewBigPtrsSlc() []*Big {
	size := 50
	var res []*Big

	for i := 0; i < size; i++ {
		ua := "Ukraine"
		uab := []byte(ua)
		t := time.Now()

		ln := "bar"
		mmb := 10.24

		var big = Big{
			BigID:        1,
			Age:          uint(26),
			Country:      &ua,
			CountryBytes: uab,
			CreatedAt:    t.AddDate(0, 0, -2),
			UpdatedAt:    t.AddDate(0, 0, -1),
			DeletedAt:    t,
			TinyField:    NewTiny(),
			MediumField:  NewMedium(),
			MediumFieldPtr: &Medium{
				MediumID:           1,
				FirstName:          "foo",
				LastName:           &ln,
				MediumMoneyBalance: &mmb,
				CreatedAt:          time.Now(),
				TinyField:          NewTiny(),
			},
		}

		res = append(res, &big)
	}

	return res
}

func makeBigSlc(size int) []Big {
	res := make([]Big, size)

	for i := 0; i < size; i++ {
		res = append(res, NewBig())
	}

	return res
}
