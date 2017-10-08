package persical

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

var persianNames = map[PersianMonth]string{
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

func (m PersianMonth) String() string {
	return persianNames[m]
}
