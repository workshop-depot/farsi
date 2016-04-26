package persical

import "testing"

func TestBulk(t *testing.T) {
	testCases := []testCase{
		testCase{
			[3]int{1395, 1, 1},
			[3]int{2016, 3, 20},
		},
		testCase{
			[3]int{1395, 1, 31},
			[3]int{2016, 4, 19},
		},
		testCase{
			[3]int{1395, 12, 1},
			[3]int{2017, 2, 19},
		},
		testCase{
			[3]int{1395, 12, 30},
			[3]int{2017, 3, 20},
		},
		testCase{
			[3]int{1396, 12, 1},
			[3]int{2018, 2, 20},
		},
		testCase{
			[3]int{1396, 12, 29},
			[3]int{2018, 3, 20},
		},
	}

	for _, v := range testCases {
		if !runTestCase(v.p, v.g) {
			t.Logf("%+v", v)
			t.Fail()
		}
	}
}

type testCase struct{ p, g [3]int }

func runTestCase(p, g [3]int) bool {
	py, pm, pd := p[0], p[1], p[2]

	gy, gm, gd := PersianToGregorian(py, pm, pd)
	if gy != g[0] || gm != g[1] || gd != g[2] {
		return false
	}

	py, pm, pd = GregorianToPersian(gy, gm, gd)
	if py != p[0] || pm != p[1] || pd != p[2] {
		return false
	}

	return true
}

func TestSingle(t *testing.T) {
	py, pm, pd := 1393, 10, 16

	gy, gm, gd := PersianToGregorian(py, pm, pd)
	if gy != 2015 || gm != 1 || gd != 6 {
		t.Log(gy, gm, gd)
		t.Fail()
	}

	py, pm, pd = GregorianToPersian(gy, gm, gd)
	if py != 1393 || pm != 10 || pd != 16 {
		t.Log(py, pm, pd)
		t.Fail()
	}
}
