package structs

import "time"

type Medium struct {
	MediumID           int       `json:"medium_id" xml:"medium_id" yaml:"medium_id"`
	FirstName          string    `json:"first_name" xml:"first_name"  yaml:"first_name"`
	LastName           *string   `json:"last_name" xml:"last_name"  yaml:"last_name"`
	MediumMoneyBalance *float64  `json:"medium_money_balance" xml:"medium_money_balance"  yaml:"medium_money_balance"`
	CreatedAt          time.Time `json:"created_at" xml:"created_at"  yaml:"created_at"`
	TinyField          Tiny      `json:"tiny_field" xml:"tiny_field"  yaml:"tiny_field"`
}

func NewMedium() Medium {
	ln := "bar"
	mmb := 10.24

	return Medium{
		MediumID:           1,
		FirstName:          "foo",
		LastName:           &ln,
		MediumMoneyBalance: &mmb,
		CreatedAt:          time.Now(),
		TinyField:          NewTiny(),
	}
}

func NewMediumSlc() []Medium {
	size := 10
	res := make([]Medium, size)

	for i := 0; i < size; i++ {
		res = append(res, NewMedium())
	}

	return res
}

func NewMediumPtrsSlc() []*Medium {
	size := 10
	var res []*Medium

	for i := 0; i < size; i++ {
		ln := "bar"
		mmb := 10.24

		var medium = Medium{
			MediumID:           1,
			FirstName:          "foo",
			LastName:           &ln,
			MediumMoneyBalance: &mmb,
			CreatedAt:          time.Now(),
			TinyField:          NewTiny(),
		}

		res = append(res, &medium)
	}

	return res
}
