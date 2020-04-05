package stime

import (
	"time"
)

const (
	_          = iota
	YESTERDAY  //昨天1
	SEVENDAYS  //近七天2
	THIRTYDAYS //近30天3
	THISMONTH  //本月4
	LASTMONTH  //上月5
	TODAY      //今天6
	FULLTODAY  //完整的今天7
) //取的时间参数

//取头尾时间戳
func GetTime(typ int) (int64, int64) {
	switch typ {
	case TODAY:
		return getToday()
	case FULLTODAY:
		return getFullToday()
	case YESTERDAY:
		return getYesterday()
	case SEVENDAYS:
		return getSevenDays()
	case THIRTYDAYS:
		return getThirtyDays()
	case THISMONTH:
		return getThisMonth()
	case LASTMONTH:
		return getLastMonth()
	default:
		return getThirtyDays()
	}
}

//取时间范围内的每天的数组
func GetTimeList(s_time, e_time int64) []string {
	var dayList = make([]string, 0)

	for {
		dayList = append(dayList, time.Unix(s_time, 0).Format("2006-01-02"))
		s_time = s_time + 86400
		if s_time > e_time {
			break
		}
	}

	return dayList
}

//今天
func getToday() (int64, int64) {
	return GetZeroTime(time.Now()).Unix(), time.Now().Unix()
}

//完整的今天
func getFullToday() (int64, int64) {
	return GetZeroTime(time.Now()).Unix(), GetZeroTime(time.Now()).AddDate(0, 0, 1).Unix() - 1
}

//昨天
func getYesterday() (int64, int64) {
	return GetZeroTime(time.Now()).AddDate(0, 0, -1).Unix(), GetZeroTime(time.Now()).Unix() - 1
}

//近七天
func getSevenDays() (int64, int64) {
	return GetZeroTime(time.Now()).AddDate(0, 0, -7).Unix(), GetZeroTime(time.Now()).Unix() - 1
}

//近30天
func getThirtyDays() (int64, int64) {
	return GetZeroTime(time.Now()).AddDate(0, 0, -30).Unix(), GetZeroTime(time.Now()).Unix() - 1
}

//本月
func getThisMonth() (int64, int64) {
	d := time.Now()
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d).Unix(), GetZeroTime(time.Now()).Unix() - 1
}

//上月
func getLastMonth() (int64, int64) {
	d := time.Now()
	firstDay := d.AddDate(0, -1, -d.Day()+1)
	lastDay := d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(firstDay).Unix(), GetZeroTime(lastDay).Unix() - 1
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
