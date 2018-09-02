# farsi
This module will contain utilities to help with anything Farsi.

# farsi/calendar
This package provides Persian Calendar Calculations for Go (golang). It separated from [Roozh](https://github.com/dc0d/Roozh) to make it `go get`table. Roozh contains implementations for converting between Persian dates and Gregorian dates for _JavaScript_, _Java_ and _C#_ too.

Basically we just use two functions, `PersianToGregorian` and `GregorianToPersian` for conversion between calendars, because Go does not provide a way to describe a calendar other than Gregorian calendar, the default expected/assumed one by package `time`. It is based on the work of [Kazimierz M. Borkowski](http://www.astro.uni.torun.pl/%7Ekb/Papers/EMP/PersianC-EMP.htm).
