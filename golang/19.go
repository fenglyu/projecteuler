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

type Leap struct {
	days int
	leap bool
}

//func (l Leap) ToString() string {
//	return fmt.Sprintf("Leap:[ %d %v]", l.days, l.leap)
//}
//
//func (l Leap) String() string {
//	return fmt.Sprintf("Leap: %d %v", l.days, l.leap)
//}

type Day struct {
	dom  int // day of week
	week Week
}

type Month struct {
	m   int // month number [0, 11]
	num int // number of days
}

type Year struct {
	y    int  // year number, e.g. 1990
	leap bool // leap or not
}

type Date struct {
	Year
	Month
	Day
	//	Week
}

func (d *Date) CountDays(w Week, t *Date) {
	y := d.Year.y - t.Year.y
	m := d.Month.m - t.Month.m
	day := d.Day.dom - t.Day.dom
	fmt.Println(y)
	fmt.Println(m)
	fmt.Println(day)
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
			y:    1992,
			leap: true,
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

	t := &Date{
		Year: Year{
			y:    2041,
			leap: false,
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

	f := &Date{
		Year: Year{
			y:    2019,
			leap: false,
		},
		Month: Month{
			m:   9,
			num: 30,
		},
		Day: Day{
			dom:  26,
			week: Thur,
		},
	}

	t.CountDays(Tue, d)
	f.CountDays(Thur, t)
}
