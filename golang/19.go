package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://www.xarg.org/puzzle/project-euler/problem-19/

/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

type Week int

const (
	Sun Week = iota
	Mon
	Tue
	Wed
	Thur
	Fri
	Sat
)

func (day Week) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	if day < Sun || day > Sat {
		return "Unknown"
	}
	return names[day]
}

var MonNum = [2][13]int{
	{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}}

var WeekNum = [7]Week{Sun, Mon, Tue, Wed, Thur, Fri, Sat}

type Leap struct {
	days int
	leap bool
}

type Day struct {
	dom  int // day of week
	week Week
}

func (d *Day) setWeek(w Week) {
	d.week = w
}

type Month struct {
	m   int // month number [0, 11]
	num int // number of days
}

func (m *Month) setMonth(leap bool) {
	i := 0
	if leap {
		i = 1
	}
	m.num = MonNum[i][m.m]
}

type Year struct {
	y    int  // year number, e.g. 1990
	leap bool // leap or not
}

func (y *Year) setLeap() {
	y.leap = isLeap(y.y).leap
}

func (y Year) getDays() int {
	if y.leap {
		return 366
	}
	return 365
}

type Date struct {
	Year
	Month
	Day
	//	Week
}

func NewDate(y int, m int, d int) *Date {
	date := &Date{
		Year: Year{
			y: y,
		},
		Month: Month{
			m: m,
		},
		Day: Day{
			dom: d,
		},
	}
	date.Year.setLeap()
	date.Month.setMonth(date.Year.leap)
	return date
}

func NewDateStr(date string) *Date {
	var arr [3]int
	for i, v := range strings.Split(date, "-") {
		//		arr[i] = int(v)
		arr[i], _ = strconv.Atoi(v)
	}

	return NewDate(arr[0], arr[1], arr[2])
}

func (d *Date) CountDays() int {
	leap := 0
	if d.leap {
		leap = 1
	}

	days := 0
	for i := 1; i < d.Month.m; i++ {
		days += MonNum[leap][i]
	}

	return days + d.Day.dom
}

func (d *Date) DiffYearsInDays(t *Date) int {

	yd := 0
	for i := t.Year.y; i < d.Year.y; i++ {
		yd += isLeap(i).days
	}

	yd += d.CountDays() - t.CountDays()
	return yd
}

func (d *Date) WeekRelatedToDate(s *Date, w Week) Week {
	//fmt.Println(d.DiffYearsInDays(t))
	diff_days := d.DiffYearsInDays(s)
	pos := diff_days % 7
	sw := int(w)
	return Week((sw + pos) % 7)
}

func CountWeekDay(s *Date, e *Date, w Week) (interface{}, int) {
	diff_days := e.DiffYearsInDays(s)
	//	sum := 0

	pos := int(s.Day.week)
	count := 0
	ar := [6]int{-1, -1, -1, -1, -1, -1}
	// why ?
	for i := 0; i <= diff_days; i++ {
		t := (pos + i) % 7
		if t == 0 {
			ar[count] = i + 1
			count++
		}
	}

	return ar, count
}

func isLeap(n int) Leap {
	var l Leap
	if ((n%100 != 0) && (n%4 == 0)) || n%400 == 0 {
		l = Leap{366, true}
	} else {
		l = Leap{365, false}
	}
	return l
}

func main() {

	start := &Date{
		Year: Year{
			y: 1900,
		},
		Month: Month{
			m: 1,
		},
		Day: Day{
			dom: 1,
		},
	}
	start.Year.setLeap()
	start.Month.setMonth(start.Year.leap)
	start.Day.setWeek(Mon)

	w := start.Day.week
	sum := 0
	for pos := 1901; pos < 2001; pos++ {
		//fmt.Println(pos)
		s := NewDate(pos, 1, 1)
		sw := s.WeekRelatedToDate(start, w)
		s.Day.setWeek(sw)

		e := NewDate(pos, 1, 31)
		ew := e.WeekRelatedToDate(s, sw)
		e.Day.setWeek(ew)

		//ar, c := CountWeekDay(s, e, sw)
		_, c := CountWeekDay(s, e, sw)
		//fmt.Println(pos, ar, c)
		//
		start = e
		w = start.Day.week
		sum += c
		//	fmt.Println(sum)
	}

	fmt.Println("All the number of Sunday of frist Month: ", sum)

	start = NewDate(1900, 1, 1)
	start.Day.setWeek(Mon)

	w = start.Day.week
	first := 0
	for pos := 1901; pos < 2001; pos++ {
		for m := 1; m < 13; m++ {
			s := NewDate(pos, m, 1)
			sw := s.WeekRelatedToDate(start, w)
			if sw == Sun {
				first++
			}
		}

	}

	fmt.Println("Sundays fell on the first of the month during the twentieth century: ", first)

}
