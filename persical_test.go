package persical

import (
	"bufio"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// from: https://en.wikipedia.org/wiki/Iranian_calendars
// we should add even bigger raw calendar tables
var input = `
1354*	21 3 1975
1355	21 3 1976
1356	21 3 1977
1357	21 3 1978
1358*	21 3 1979
1359	21 3 1980
1360	21 3 1981
1361	21 3 1982
1362*	21 3 1983
1363	21 3 1984
1364	21 3 1985
1365	21 3 1986
1366*	21 3 1987
1367	21 3 1988
1368	21 3 1989
1369	21 3 1990
1370*	21 3 1991
1371	21 3 1992
1372	21 3 1993
1373	21 3 1994
1374	21 3 1995
1375*	20 3 1996
1376	21 3 1997
1377	21 3 1998
1378	21 3 1999
1379*	20 3 2000
1380	21 3 2001
1381	21 3 2002
1382	21 3 2003
1383*	20 3 2004
1384	21 3 2005
1385	21 3 2006
1386	21 3 2007
1387*	20 3 2008
1388	21 3 2009
1389	21 3 2010
1390	21 3 2011
1391*	20 3 2012
1392	21 3 2013
1393	21 3 2014
1394	21 3 2015
1395*	20 3 2016
1396	21 3 2017
1397	21 3 2018
1398	21 3 2019
1399*	20 3 2020
1400	21 3 2021
1401	21 3 2022
1402	21 3 2023
1403*	20 3 2024
1404	21 3 2025
1405	21 3 2026
1406	21 3 2027
1407	20 3 2028
1408*	20 3 2029
1409	21 3 2030
1410	21 3 2031
1411	20 3 2032
1412*	20 3 2033
1413	21 3 2034
1414	21 3 2035
1415	20 3 2036
1416*	20 3 2037
1417	21 3 2038
1418	21 3 2039
1419	20 3 2040
`

type inputRecord struct {
	py, gd, gm, gy int
	flag           bool
}

func parseInput(t *testing.T) (res []inputRecord) {
	assert := assert.New(t)

	sr := strings.NewReader(input)

	r := bufio.NewReader(sr)
	for l, err := "", error(nil); err == nil; l, err = r.ReadString('\n') {
		l := strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		l = strings.Replace(l, "\t", "  ", -1)
		l = strings.Replace(l, "*  ", "* ", -1)

		py, err := strconv.Atoi(l[:4])
		assert.NoError(err)
		flag := false
		if l[4] == '*' {
			flag = true
		}
		gd, err := strconv.Atoi(l[6:8])
		assert.NoError(err)
		gm, err := strconv.Atoi(l[9:10])
		assert.NoError(err)
		gy, err := strconv.Atoi(l[11:])
		assert.NoError(err)

		res = append(res, inputRecord{
			py:   py,
			gd:   gd,
			gm:   gm,
			gy:   gy,
			flag: flag,
		})
	}
	return
}

func TestTable(t *testing.T) {
	assert := assert.New(t)
	records := parseInput(t)
	var prev inputRecord
	for _, rec := range records {
		py, pm, pd := rec.py, 1, 1

		gy, gm, gd := PersianToGregorian(py, pm, pd)
		assert.Equal(rec.gd, gd)
		assert.Equal(rec.gm, gm)
		assert.Equal(rec.gy, gy)

		py, pm, pd = GregorianToPersian(gy, gm, gd)
		assert.Equal(rec.py, py)
		assert.Equal(1, pm)
		assert.Equal(1, pd)

		if prev.gd != 0 {
			py, pm, pd = GregorianToPersian(gy, gm, gd-1)
			if prev.flag {
				assert.True(pd > 29)
			}
		}

		prev = rec
	}
}

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
