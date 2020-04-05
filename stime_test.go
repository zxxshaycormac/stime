package stime

import (
	"navi/kvp-mps/pkg/define"
	"testing"
	"time"
)

func TestMpsTime_GetList(t *testing.T) {
	s_time, e_time := GetTime(6)
	t.Logf("GetTime(6) : %v / %v", s_time, e_time)
	t.Logf("GetTime(6) : %v / %v", time.Unix(s_time, 0), time.Unix(e_time, 0))
	s_time, e_time = GetTime(7)
	t.Logf("GetTime(7) : %v / %v", s_time,  time.Unix(e_time, 0))
	s_time, e_time = GetTime(3)
	//t.Logf("GetTime(3) : %v / %v", s_time, e_time)
	s_time, e_time = GetTime(4)
	//t.Logf("GetTime(4) : %v / %v", s_time, e_time)
	s_time, e_time = GetTime(5)
	//t.Logf("GetTime(5) : %v / %v", time.Unix(s_time, 0).Format(define.TIME_LAYOUT), time.Unix(e_time, 0).Format(define.TIME_LAYOUT))

	//t.Logf("GetStringTime(GetTime(1)) : %v", GetStringTime(GetTime(1)))
	//t.Logf("GetStringTime(GetTime(2)) : %v", GetStringTime(GetTime(2)))
}

func TestGetStringTime(t *testing.T) {
	s, e := getLastMonth()
	d := GetTimeList(s, e)
	t.Log(d)
}

func TestGetTime(t *testing.T) {
	//d := time.Now()
	firstDay := time.Now().AddDate(0, 0, -1)
	t.Log(GetZeroTime(firstDay).Format(define.TIME_LAYOUT))
	//lastDay := d.AddDate(0, -1, -d.Day())
	t.Log(GetZeroTime(time.Now()).Format(define.TIME_LAYOUT))
}