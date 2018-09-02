package text

import (
	"fmt"
	"strings"
	"time"
)

//-----------------------------------------------------------------------------

// PersianSeason .
type PersianSeason int

// PersianSeason values
const (
	Bahar PersianSeason = iota + 1
	Tabestan
	Paiiz
	Zemestan
)

var persianSeasonFa = map[PersianSeason]string{
	Bahar:    "بهار",
	Tabestan: "تابستان",
	Paiiz:    "پاییز",
	Zemestan: "زمستان",
}

var persianSeasonEn = map[PersianSeason]string{
	Bahar:    "Bahar",
	Tabestan: "Tabestan",
	Paiiz:    "Paiiz",
	Zemestan: "Zemestan",
}

func (m PersianSeason) String() string {
	return persianSeasonFa[m]
}

// StringEn .
func (m PersianSeason) StringEn() string {
	return persianSeasonEn[m]
}

// Of .
func (m PersianSeason) Of(pm PersianMonth) PersianSeason {
	switch pm {
	case Farvardin:
		fallthrough
	case Ordibehesht:
		fallthrough
	case Khordad:
		return Bahar

	case Tir:
		fallthrough
	case Mordad:
		fallthrough
	case Shahrivar:
		return Tabestan

	case Mehr:
		fallthrough
	case Aban:
		fallthrough
	case Azar:
		return Paiiz
	}

	return Zemestan
}

//-----------------------------------------------------------------------------

// PersianWeekday .
type PersianWeekday time.Weekday

// valid values for PersianWeekday
const (
	YekShanbeh PersianWeekday = iota
	DoShanbeh
	SeShanbeh
	ChaharShanbeh
	PanjShanbeh
	Adineh
	Shanbeh
)

var persianWeekdayFa = map[PersianWeekday]string{
	YekShanbeh:    "یکشنبه",
	DoShanbeh:     "دوشنبه",
	SeShanbeh:     "سه‌شنبه",
	ChaharShanbeh: "چهارشنبه",
	PanjShanbeh:   "پنجشنبه",
	Adineh:        "آدینه",
	Shanbeh:       "شنبه",
}

var persianWeekdayEn = map[PersianWeekday]string{
	YekShanbeh:    "YekShanbeh",
	DoShanbeh:     "DoShanbeh",
	SeShanbeh:     "SeShanbeh",
	ChaharShanbeh: "ChaharShanbeh",
	PanjShanbeh:   "PanjShanbeh",
	Adineh:        "Adineh",
	Shanbeh:       "Shanbeh",
}

func (m PersianWeekday) String() string {
	return persianWeekdayFa[m]
}

// StringEn .
func (m PersianWeekday) StringEn() string {
	return persianWeekdayEn[m]
}

//-----------------------------------------------------------------------------

//PersianMonth represents a Persian month
type PersianMonth int32

//List of Persian months
const (
	Farvardin PersianMonth = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

var persianMonthEn = map[PersianMonth]string{
	Farvardin:   "Farvardin",
	Ordibehesht: "Ordibehesht",
	Khordad:     "Khordad",
	Tir:         "Tir",
	Mordad:      "Mordad",
	Shahrivar:   "Shahrivar",
	Mehr:        "Mehr",
	Aban:        "Aban",
	Azar:        "Azar",
	Dey:         "Dey",
	Bahman:      "Bahman",
	Esfand:      "Esfand",
}

var persianMonthFa = map[PersianMonth]string{
	Farvardin:   "فروردین",
	Ordibehesht: "اردیبهشت",
	Khordad:     "خرداد",
	Tir:         "تیر",
	Mordad:      "مرداد",
	Shahrivar:   "شهریور",
	Mehr:        "مهر",
	Aban:        "آبان",
	Azar:        "آذر",
	Dey:         "دی",
	Bahman:      "بهمن",
	Esfand:      "اسفند",
}

func (m PersianMonth) String() string {
	return persianMonthFa[m]
}

// StringEn .
func (m PersianMonth) StringEn() string {
	return persianMonthEn[m]
}

//-----------------------------------------------------------------------------

const (
	persianNumbers = "۰۱۲۳۴۵۶۷۸۹"
	latinNumbers   = "0123456789"
)

var p2l = make(map[rune]rune)
var l2p = make(map[rune]rune)

func init() {
	var p = []rune(persianNumbers)
	var l = []rune(latinNumbers)
	if len(l) != len(p) {
		panic(fmt.Errorf("BOTH []rune MUST BE OF EQUAL LENGTH"))
	}
	for i := 0; i < 10; i++ {
		p2l[p[i]] = l[i]
		l2p[l[i]] = p[i]
	}
}

// NumberString contains only Persian or only English digits
type NumberString string

// ToPersian .
func (nu NumberString) ToPersian() string {
	q := []rune(string(nu))
	for k, v := range q {
		q[k] = l2p[v]
	}
	return string(q)
}

// ToEnglish .
func (nu NumberString) ToEnglish() string {
	q := []rune(string(nu))
	for k, v := range q {
		q[k] = p2l[v]
	}
	return string(q)
}

//-----------------------------------------------------------------------------

// PolishYeKaf replaces characters by their persian equivalent (for keyboards or OSs that suck).
func PolishYeKaf(s string) (res string) {
	res = strings.Replace(
		s,
		"ي",
		"ی",
		-1)
	res = strings.Replace(
		res,
		"ك",
		"ک",
		-1)

	return
}

//-----------------------------------------------------------------------------
