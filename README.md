# stime
包含两个函数

There are two functions

GetTime：获取一个时间范围的第一秒和最后一秒，可用于筛选查询，也可直接作为GetStringTime的参数

GetTime: Gets the first and last seconds of a time range, which can be used to filter queries, or can be used directly as a parameter to GetStringTime

```go
startTime, enTime := stime.GetTime(stime.SEVENDAYS)
```

GetStringTime：获取一个时间范围中每一天的字符串数据组成的数组，可用于前端展示图表

GetStringTime: Gets an array of string data for each day in the time range that can be used to display a chart at the front end

```go
var stringTimeList = make([]string, 0)
stringTimeList := stime.GetStringTime(stime.GetTime(stime.SEVENDAYS))
```

预设的时间范围有：今天【到现在】、今天【全天】、昨天、近七天、近30天、本月【到现在】、上月

The preset time range is: Today [to the present], today [all day], yesterday, nearly seven days, nearly 30 days, this month [to the present], last month

```go
const (
	_          = iota
	YESTERDAY  //昨天1
	SEVENDAYS  //近七天2
	THIRTYDAYS //近30天3
	THISMONTH  //本月4
	LASTMONTH  //上月5
	TODAY      //今天6
	FULLTODAY  //今天全天7
) //取的时间参数
```

代码很简单，如果想要修改输出或者调整字符串格式可以自己动手~

The code is simple and you can do it yourself if you want to modify the output or reformat the string~
