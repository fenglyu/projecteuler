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

func (l Leap) ToString() string {
	return fmt.Sprintf("Leap:[ %d %v]", l.days, l.leap)
}

func (l Leap) String() string {
	return fmt.Sprintf("Leap: %d %v", l.days, l.leap)
}

type Month struct {
	m int
}

type Year struct {
	y int
}

type Date struct {
	Year
	Month
	Day
	Week
}

func (d Date) CountDays(s Date, w Week, d Date) {

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

	for i := 1889; i < 2001; i++ {
		l := isLeap(i)
		fmt.Println(l)
	}
}
