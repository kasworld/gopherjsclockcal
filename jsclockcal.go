// Copyright 2015,2016,2017,2018 SeukWon Kang (kasworld@gmail.com)

package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("initPage", displayFrame)
}

var lasttime time.Time

func displayFrame() {
	js.Global.Call("requestAnimationFrame", displayFrame)
	thistime := time.Now()
	if lasttime.Second() == thistime.Second() {
		return
	}
	lasttime = thistime

	win := js.Global.Get("window")
	winW := win.Get("innerWidth").Int()
	winH := win.Get("innerHeight").Int()

	clockFontSize := winW / 5
	if winH/2 < clockFontSize {
		clockFontSize = winH / 3
	}
	updateClock(clockFontSize)

	dateFontSize := winW / 10
	if winH < dateFontSize*10 {
		dateFontSize = winH / 10
	}
	updateDate(dateFontSize)

	calendarFontSize := winW / 16
	if winH < calendarFontSize*12 {
		calendarFontSize = winH / 12
	}
	updateCalendar(calendarFontSize)
}

func updateClock(fontSize int) {
	jsObj := js.Global.Get("document").Call("getElementById", "clock")
	jsObj.Get("style").Set("font-size", fmt.Sprintf("%dpx", fontSize))
	var buf bytes.Buffer
	fmt.Fprintf(&buf, time.Now().Format("15:04:05"))
	jsObj.Set("innerHTML", buf.String())
}

func updateDate(fontSize int) {
	jsObj := js.Global.Get("document").Call("getElementById", "date")
	jsObj.Get("style").Set("font-size", fmt.Sprintf("%dpx", fontSize))
	var buf bytes.Buffer
	fmt.Fprintf(&buf, time.Now().Format("2006-01-02 Mon"))
	jsObj.Set("innerHTML", buf.String())
}

func updateCalendar(fontSize int) {
	jsObj := js.Global.Get("document").Call("getElementById", "calendar")
	jsObj.Get("style").Set("font-size", fmt.Sprintf("%dpx", fontSize))
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<table id=\"t01\">")

	for weekday := time.Sunday; weekday <= time.Saturday; weekday++ {
		fmt.Fprintf(&buf, "<td>%v</td>", weekday.String()[:3])
	}

	fmt.Fprintf(&buf, "<colgroup>")
	for weekday := time.Sunday; weekday <= time.Saturday; weekday++ {
		if weekday == 0 {
			fmt.Fprintf(&buf, "<col style=\"background-color:darkred; width:14%%; \">")
		} else if weekday == 6 {
			fmt.Fprintf(&buf, "<col style=\"background-color:darkblue; width:14%%; \">")
		} else {
			fmt.Fprintf(&buf, "<col style=\"background-color:gray; width:14%%; \">")
		}
	}
	fmt.Fprintf(&buf, "<colgroup>")

	today := time.Now()
	now := time.Now()
	now = now.AddDate(0, 0, -int(now.Weekday())-7)
	for week := 0; week < 5; week++ {
		fmt.Fprintf(&buf, "<tr>")
		for weekday := time.Sunday; weekday <= time.Saturday; weekday++ {
			if now.Month() != today.Month() {
				fmt.Fprintf(&buf, "<td style=\"color:darkgray;\">%d</td>", now.Day())
			} else {
				if now.Day() != today.Day() {
					fmt.Fprintf(&buf, "<td style=\"color:white;\">%d</td>", now.Day())
				} else {
					fmt.Fprintf(&buf, "<td style=\"color:orangered;\">%d</td>", now.Day())
				}
			}
			now = now.AddDate(0, 0, 1)
		}
		fmt.Fprintf(&buf, "</tr>")
	}
	fmt.Fprintf(&buf, "</table>")
	jsObj.Set("innerHTML", buf.String())
}
