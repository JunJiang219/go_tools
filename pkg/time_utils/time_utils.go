// Package time_utils 提供通用时间工具函数
package time_utils

import (
	"fmt"
	"time"
)

// 常用时间格式
const (
	DateTimeFormat     = "2006-01-02 15:04:05"
	DateFormat         = "2006-01-02"
	TimeFormat         = "15:04:05"
	DateTimeFormatCN   = "2006年01月02日 15:04:05"
	DateFormatCN       = "2006年01月02日"
	DateTimeFormatISO  = "2006-01-02T15:04:05Z07:00"
	DateTimeFormatUnix = "Mon Jan 02 15:04:05 MST 2006"
)

// Now 返回当前时间
func Now() time.Time {
	return time.Now()
}

// NowUnix 返回当前时间戳（秒）
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixMilli 返回当前时间戳（毫秒）
func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

// NowUnixNano 返回当前时间戳（纳秒）
func NowUnixNano() int64 {
	return time.Now().UnixNano()
}

// FormatTime 格式化时间为默认格式（YYYY-MM-DD HH:mm:ss）
func FormatTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// FormatDate 格式化时间为日期格式（YYYY-MM-DD）
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatTimeCustom 使用自定义格式格式化时间
func FormatTimeCustom(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime 解析时间字符串（默认格式：YYYY-MM-DD HH:mm:ss）
func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation(DateTimeFormat, s, time.Local)
}

// ParseDate 解析日期字符串（格式：YYYY-MM-DD）
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation(DateFormat, s, time.Local)
}

// ParseTimeCustom 使用自定义格式解析时间字符串
func ParseTimeCustom(s, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, s, time.Local)
}

// UnixToTime 时间戳（秒）转时间
func UnixToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// UnixMilliToTime 时间戳（毫秒）转时间
func UnixMilliToTime(msec int64) time.Time {
	return time.UnixMilli(msec)
}

// StartOfDay 获取某天的开始时间（00:00:00）
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 获取某天的结束时间（23:59:59.999999999）
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek 获取某周的开始时间（周一 00:00:00）
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return StartOfDay(t.AddDate(0, 0, -weekday+1))
}

// EndOfWeek 获取某周的结束时间（周日 23:59:59）
func EndOfWeek(t time.Time) time.Time {
	return EndOfDay(StartOfWeek(t).AddDate(0, 0, 6))
}

// StartOfMonth 获取某月的开始时间
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 获取某月的结束时间
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// StartOfYear 获取某年的开始时间
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 获取某年的结束时间
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// Today 返回今天的开始时间
func Today() time.Time {
	return StartOfDay(time.Now())
}

// Yesterday 返回昨天的开始时间
func Yesterday() time.Time {
	return StartOfDay(time.Now().AddDate(0, 0, -1))
}

// Tomorrow 返回明天的开始时间
func Tomorrow() time.Time {
	return StartOfDay(time.Now().AddDate(0, 0, 1))
}

// IsSameDay 判断两个时间是否是同一天
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameWeek 判断两个时间是否是同一周
func IsSameWeek(t1, t2 time.Time) bool {
	y1, w1 := t1.ISOWeek()
	y2, w2 := t2.ISOWeek()
	return y1 == y2 && w1 == w2
}

// IsSameMonth 判断两个时间是否是同一月
func IsSameMonth(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsToday 判断是否是今天
func IsToday(t time.Time) bool {
	return IsSameDay(t, time.Now())
}

// IsYesterday 判断是否是昨天
func IsYesterday(t time.Time) bool {
	return IsSameDay(t, time.Now().AddDate(0, 0, -1))
}

// IsTomorrow 判断是否是明天
func IsTomorrow(t time.Time) bool {
	return IsSameDay(t, time.Now().AddDate(0, 0, 1))
}

// IsWeekend 判断是否是周末
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// DaysBetween 计算两个时间之间相差的天数
func DaysBetween(t1, t2 time.Time) int {
	t1 = StartOfDay(t1)
	t2 = StartOfDay(t2)
	return int(t2.Sub(t1).Hours() / 24)
}

// HoursBetween 计算两个时间之间相差的小时数
func HoursBetween(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours()
}

// MinutesBetween 计算两个时间之间相差的分钟数
func MinutesBetween(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Minutes()
}

// SecondsBetween 计算两个时间之间相差的秒数
func SecondsBetween(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Seconds()
}

// AddDays 增加天数
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths 增加月数
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears 增加年数
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// AddHours 增加小时
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddMinutes 增加分钟
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSeconds 增加秒
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// IsLeapYear 判断是否是闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth 获取某月的天数
func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// DaysInYear 获取某年的天数
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// Age 计算年龄（根据生日）
func Age(birthday time.Time) int {
	now := time.Now()
	age := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		age--
	}
	return age
}

// IsBefore 判断 t1 是否在 t2 之前
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsAfter 判断 t1 是否在 t2 之后
func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// IsBetween 判断时间是否在两个时间之间
func IsBetween(t, start, end time.Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

// FormatDuration 格式化时间间隔为人类可读格式
func FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%d秒", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%d分钟", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		hours := int(d.Hours())
		minutes := int(d.Minutes()) % 60
		if minutes > 0 {
			return fmt.Sprintf("%d小时%d分钟", hours, minutes)
		}
		return fmt.Sprintf("%d小时", hours)
	}
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	if hours > 0 {
		return fmt.Sprintf("%d天%d小时", days, hours)
	}
	return fmt.Sprintf("%d天", days)
}

// RelativeTime 返回相对时间描述（如：3分钟前、2小时后）
func RelativeTime(t time.Time) string {
	now := time.Now()
	d := t.Sub(now)

	if d < 0 {
		d = -d
		switch {
		case d < time.Minute:
			return "刚刚"
		case d < time.Hour:
			return fmt.Sprintf("%d分钟前", int(d.Minutes()))
		case d < 24*time.Hour:
			return fmt.Sprintf("%d小时前", int(d.Hours()))
		case d < 48*time.Hour:
			return "昨天"
		case d < 7*24*time.Hour:
			return fmt.Sprintf("%d天前", int(d.Hours()/24))
		case d < 30*24*time.Hour:
			return fmt.Sprintf("%d周前", int(d.Hours()/(24*7)))
		case d < 365*24*time.Hour:
			return fmt.Sprintf("%d个月前", int(d.Hours()/(24*30)))
		default:
			return fmt.Sprintf("%d年前", int(d.Hours()/(24*365)))
		}
	}

	switch {
	case d < time.Minute:
		return "马上"
	case d < time.Hour:
		return fmt.Sprintf("%d分钟后", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%d小时后", int(d.Hours()))
	case d < 48*time.Hour:
		return "明天"
	case d < 7*24*time.Hour:
		return fmt.Sprintf("%d天后", int(d.Hours()/24))
	case d < 30*24*time.Hour:
		return fmt.Sprintf("%d周后", int(d.Hours()/(24*7)))
	case d < 365*24*time.Hour:
		return fmt.Sprintf("%d个月后", int(d.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d年后", int(d.Hours()/(24*365)))
	}
}

// GetWeekday 获取星期几（中文）
func GetWeekday(t time.Time) string {
	weekdays := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	return weekdays[t.Weekday()]
}

// GetWeekdayShort 获取星期几（简写）
func GetWeekdayShort(t time.Time) string {
	weekdays := []string{"日", "一", "二", "三", "四", "五", "六"}
	return "周" + weekdays[t.Weekday()]
}

// GetQuarter 获取季度（1-4）
func GetQuarter(t time.Time) int {
	return (int(t.Month())-1)/3 + 1
}

// StartOfQuarter 获取某季度的开始时间
func StartOfQuarter(t time.Time) time.Time {
	quarter := GetQuarter(t)
	month := time.Month((quarter-1)*3 + 1)
	return time.Date(t.Year(), month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter 获取某季度的结束时间
func EndOfQuarter(t time.Time) time.Time {
	return StartOfQuarter(t).AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// GetDateRange 获取从开始日期到结束日期的所有日期列表
func GetDateRange(start, end time.Time) []time.Time {
	var dates []time.Time
	start = StartOfDay(start)
	end = StartOfDay(end)

	for current := start; !current.After(end); current = current.AddDate(0, 0, 1) {
		dates = append(dates, current)
	}
	return dates
}

// SetTimezone 设置时区
func SetTimezone(t time.Time, tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t, err
	}
	return t.In(loc), nil
}

// ToUTC 转换为 UTC 时间
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// ToLocal 转换为本地时间
func ToLocal(t time.Time) time.Time {
	return t.Local()
}
