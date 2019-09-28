package main

import (
	"fmt"
)

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

func (d *Date) GetWeekOfDay(s *Date, w Week) Week {
	//fmt.Println(d.DiffYearsInDays(t))
	diff_days := d.DiffYearsInDays(s)
	pos := diff_days % 7
	sw := int(w)
	return Week((sw + pos) % 7)
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

	//	for i := 1889; i < 2001; i++ {
	//		l := isLeap(i)
	//		fmt.Println(l)
	//	}

	d := &Date{
		Year: Year{
			y: 1992,
			//		leap: true,
		},
		Month: Month{
			m:   2,
			num: 29,
		},
		Day: Day{
			dom:  12,
			week: Tue,
		},
	}
	d.Year.setLeap()

	t := &Date{
		Year: Year{
			y: 2041,
			//		leap: false,
		},
		Month: Month{
			m:   11,
			num: 30,
		},
		Day: Day{
			dom:  28,
			week: Thur,
		},
	}

	t.Year.setLeap()

	f := &Date{
		Year: Year{
			y: 2018,
			//			leap: false,
		},
		Month: Month{
			m:   2,
			num: 28,
		},
		Day: Day{
			dom: 21,
			//			week: Thur,
		},
	}

	f.Year.setLeap()

	g := &Date{
		Year: Year{
			y: 2019,
			//			leap: false,
		},
		Month: Month{
			m: 9,
			//			num: 30,
		},
		Day: Day{
			dom: 29,
			//	week: Thur,
		},
	}

	f.Year.setLeap()
	fmt.Println(t.CountDays())
	fmt.Println(f.CountDays())
	fmt.Println(f.DiffYearsInDays(t))
	fmt.Println(g.GetWeekOfDay(f, Wed))
}
