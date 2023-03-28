package main

import "time"

type Tiny struct {
	TinyID  int    `json:"tiny_id"`
	SomeStr string `json:"some_str"`
}

type Medium struct {
	MediumID           int       `json:"medium_id"`
	FirstName          string    `json:"first_name"`
	LastName           *string   `json:"last_name"`
	MediumMoneyBalance *float64  `json:"medium_money_balance"`
	CreatedAt          time.Time `json:"created_at"`
	TinyField          Tiny      `json:"tiny_field"`
}

type Huge struct {
	HugeID         int       `json:"huge_id"`
	Age            uint      `json:"age"`
	Country        *string   `json:"country"`
	CountryBytes   []byte    `json:"country_bytes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	TinyField      Tiny      `json:"tiny_field"`
	MediumField    Medium    `json:"medium_field"`
	MediumFieldPtr *Medium   `json:"medium_field_ptr"`
}

// Tiny

func newTiny() Tiny {
	return Tiny{
		TinyID:  1,
		SomeStr: "foo bar",
	}
}

func newTinySlc() []Tiny {
	var res []Tiny
	for i := 0; i < 5; i++ {
		res = append(res, newTiny())
	}

	return res
}

func newTinyPtrsSlc() []*Tiny {
	var res []*Tiny

	for i := 0; i < 5; i++ {
		var tiny = &Tiny{
			TinyID:  1,
			SomeStr: "foo bar",
		}

		res = append(res, tiny)
	}

	return res
}

// Medium

func newMedium() Medium {
	ln := "bar"
	mmb := 10.24

	return Medium{
		MediumID:           1,
		FirstName:          "foo",
		LastName:           &ln,
		MediumMoneyBalance: &mmb,
		CreatedAt:          time.Now(),
		TinyField:          newTiny(),
	}
}

func newMediumSlc() []Medium {
	var res []Medium
	for i := 0; i < 5; i++ {
		res = append(res, newMedium())
	}

	return res
}

func newMediumPtrsSlc() []*Medium {
	var res []*Medium

	for i := 0; i < 5; i++ {
		ln := "bar"
		mmb := 10.24

		var medium = Medium{
			MediumID:           1,
			FirstName:          "foo",
			LastName:           &ln,
			MediumMoneyBalance: &mmb,
			CreatedAt:          time.Now(),
			TinyField:          newTiny(),
		}

		res = append(res, &medium)
	}

	return res
}

// Huge

func newHuge() Huge {
	ua := "Ukraine"
	uab := []byte(ua)
	t := time.Now()

	ln := "bar"
	mmb := 10.24

	return Huge{
		HugeID:       1,
		Age:          uint(26),
		Country:      &ua,
		CountryBytes: uab,
		CreatedAt:    t.AddDate(0, 0, -2),
		UpdatedAt:    t.AddDate(0, 0, -1),
		DeletedAt:    t,
		TinyField:    newTiny(),
		MediumField:  newMedium(),
		MediumFieldPtr: &Medium{
			MediumID:           1,
			FirstName:          "foo",
			LastName:           &ln,
			MediumMoneyBalance: &mmb,
			CreatedAt:          time.Now(),
			TinyField:          newTiny(),
		},
	}
}

func newHugeSlc() []Huge {
	var res []Huge
	for i := 0; i < 5; i++ {
		res = append(res, newHuge())
	}

	return res
}

func newHugePtrsSlc() []*Huge {
	var res []*Huge

	for i := 0; i < 5; i++ {
		ua := "Ukraine"
		uab := []byte(ua)
		t := time.Now()

		ln := "bar"
		mmb := 10.24

		var huge = Huge{
			HugeID:       1,
			Age:          uint(26),
			Country:      &ua,
			CountryBytes: uab,
			CreatedAt:    t.AddDate(0, 0, -2),
			UpdatedAt:    t.AddDate(0, 0, -1),
			DeletedAt:    t,
			TinyField:    newTiny(),
			MediumField:  newMedium(),
			MediumFieldPtr: &Medium{
				MediumID:           1,
				FirstName:          "foo",
				LastName:           &ln,
				MediumMoneyBalance: &mmb,
				CreatedAt:          time.Now(),
				TinyField:          newTiny(),
			},
		}

		res = append(res, &huge)
	}

	return res
}
