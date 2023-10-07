package main

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"log"
)

/*
*
 */
func yesterdayTodayTomorrow() {

	// Return datetime of today
	fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
	carbon.Now().String()           // 2020-08-05 13:14:15
	carbon.Now().ToString()         // 2020-08-05 13:14:15 +0800 CST
	carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	// Return date of today
	carbon.Now().ToDateString() // 2020-08-05
	// Return time of today
	carbon.Now().ToTimeString() // 13:14:15
	// Return datetime of today in a given timezone
	carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
	// Return timestamp with second of today
	carbon.Now().Timestamp() // 1596604455
	// Return timestamp with millisecond of today
	carbon.Now().TimestampMilli() // 1596604455999
	// Return timestamp with microsecond of today
	carbon.Now().TimestampMicro() // 1596604455999999
	// Return timestamp with nanosecond of today
	carbon.Now().TimestampNano() // 1596604455999999999

	// Return datetime of yesterday
	fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
	carbon.Yesterday().String()           // 2020-08-04 13:14:15
	carbon.Yesterday().ToString()         // 2020-08-04 13:14:15 +0800 CST
	carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
	// Return date of yesterday
	carbon.Yesterday().ToDateString() // 2020-08-04
	// Return time of yesterday
	carbon.Yesterday().ToTimeString() // 13:14:15
	// Return datetime of yesterday on a given day
	carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
	// Return datetime of yesterday in a given timezone
	carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
	// Return timestamp with second of yesterday
	carbon.Yesterday().Timestamp() // 1596518055
	// Return timestamp with millisecond of yesterday
	carbon.Yesterday().TimestampMilli() // 1596518055999
	// Return timestamp with microsecond of yesterday
	carbon.Yesterday().TimestampMicro() // 1596518055999999
	// Return timestamp with nanosecond of yesterday
	carbon.Yesterday().TimestampNano() // 1596518055999999999

	// Return datetime of tomorrow
	fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
	carbon.Tomorrow().String()           // 2020-08-06 13:14:15
	carbon.Tomorrow().ToString()         // 2020-08-06 13:14:15 +0800 CST
	carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
	// Return date of tomorrow
	carbon.Tomorrow().ToDateString() // 2020-08-06
	// Return time of tomorrow
	carbon.Tomorrow().ToTimeString() // 13:14:15
	// Return datetime of tomorrow on a given day
	carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
	// Return datetime of tomorrow in a given timezone
	carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
	// Return timestamp with second of tomorrow
	carbon.Tomorrow().Timestamp() // 1596690855
	// Return timestamp with millisecond of tomorrow
	carbon.Tomorrow().TimestampMilli() // 1596690855999
	// Return timestamp with microsecond of tomorrow
	carbon.Tomorrow().TimestampMicro() // 1596690855999999
	// Return timestamp with nanosecond of tomorrow
	carbon.Tomorrow().TimestampNano() // 1596690855999999999

}

/*
*
 */
func parseATimeString() {
	carbon.Parse("").ToDateTimeString()                    // empty string
	carbon.Parse("0").ToDateTimeString()                   // empty string
	carbon.Parse("00:00:00").ToDateTimeString()            // empty string
	carbon.Parse("0000-00-00").ToDateTimeString()          // empty string
	carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // empty string

	carbon.Parse("now").ToString()       // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0800 CST
	carbon.Parse("tomorrow").ToString()  // 2020-08-06 13:14:15 +0800 CST

	carbon.Parse("2020").ToString()                 // 2020-01-01 00:00:00 +0800 CST
	carbon.Parse("2020-8").ToString()               // 2020-08-01 00:00:00 +0800 CST
	carbon.Parse("2020-08").ToString()              // 2020-08-01 00:00:00 +0800 CST
	carbon.Parse("2020-8-5").ToString()             // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("2020-8-05").ToString()            // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("2020-08-05").ToString()           // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("2020-08-05.999").ToString()       // 2020-08-05 00:00:00.999 +0800 CST
	carbon.Parse("2020-08-05.999999").ToString()    // 2020-08-05 00:00:00.999999 +0800 CST
	carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

	carbon.Parse("2020-8-5 13:14:15").ToString()             // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-8-05 13:14:15").ToString()            // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-5 13:14:15").ToString()            // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05 13:14:15").ToString()           // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("2020-8-5T13:14:15+08:00").ToString()             // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-8-05T13:14:15+08:00").ToString()            // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05T13:14:15+08:00").ToString()           // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("20200805").ToString()                       // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("20200805131415").ToString()                 // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("20200805131415.999").ToString()             // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("20200805131415.999999").ToString()          // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("20200805131415.999999999").ToString()       // 2020-08-05 13:14:15.999999999 +0800 CST
	carbon.Parse("20200805131415.999+08:00").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("20200805131415.999999+08:00").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
}

func demo() {
	lang := carbon.NewLanguage()
	lang.SetLocale("zh-CN")

	c := carbon.SetLanguage(lang)
	if c.Error != nil {
		// 错误处理
		log.Fatal(c.Error)
	}

	c.Now().AddHours(1).DiffForHumans()      // 1 小时后
	c.Now().AddHours(1).ToMonthString()      // 八月
	c.Now().AddHours(1).ToShortMonthString() // 8月
	c.Now().AddHours(1).ToWeekString()       // 星期二
	c.Now().AddHours(1).ToShortWeekString()  // 周二
	c.Now().AddHours(1).Constellation()      // 狮子座
	c.Now().AddHours(1).Season()             // 夏季
}

func main() {
	//yesterdayTodayTomorrow()
	//parseATimeString()
	demo()
}
