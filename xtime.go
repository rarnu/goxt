package goxt

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Year   = XString("2006")
	Month  = XString("01")
	Day    = XString("02")
	Hour   = XString("15")
	Minute = XString("04")
	Second = XString("05")

	FmtYMdHmsSSS = XString("2006-01-02 15:04:05.000")
	FmtYMdHms    = XString("2006-01-02 15:04:05")
	FmtYMdHm     = XString("2006-01-02 15:04")
	FmtYMdH      = XString("2006-01-02 15")
	FmtYMd       = XString("2006-01-02")
	FmtYM        = XString("2006-01")
	FmtY         = XString("2006")
	FmtYYYYMMdd  = XString("20060102")

	EmptyTime = time.Time{}

	yRegex        = regexp.MustCompile("^(\\d){4}$")
	yyyyMmDdRegex = regexp.MustCompile("^(\\d){4}(\\d){2}(\\d){2}$")
	ymRegex       = regexp.MustCompile("^(\\d){4}-(\\d){2}$")
	ymdRegex      = regexp.MustCompile("^(\\d){4}-(\\d){2}-(\\d){2}$")
	ymdHRegex     = regexp.MustCompile("^(\\d){4}-(\\d){2}-(\\d){2} (\\d){2}$")
	ymdHmRegex    = regexp.MustCompile("^(\\d){4}-(\\d){2}-(\\d){2} (\\d){2}:(\\d){2}$")
	ymdHmsRegex   = regexp.MustCompile("^(\\d){4}-(\\d){2}-(\\d){2} (\\d){2}:(\\d){2}:(\\d){2}$")
	ymdHmsSRegex  = regexp.MustCompile("^(\\d){4}-(\\d){2}-(\\d){2} (\\d){2}:(\\d){2}:(\\d){2}.(\\d){3}$")
)

const (
	DateDelta   = 693594 // Days between 1/1/0001 and 12/31/1899
	HoursPerDay = 24
	MinsPerHour = 60
	SecsPerMin  = 60
	MSecsPerSec = 1000

	MinsPerDay  = HoursPerDay * MinsPerHour
	SecsPerHour = SecsPerMin * MinsPerHour
	SecsPerDay  = MinsPerDay * SecsPerMin
	MSecsPerDay = SecsPerDay * MSecsPerSec

	OneMillisecond  = XFloat64(1) / MSecsPerDay
	HalfMilliSecond = OneMillisecond / 2

	JulianEpoch = XFloat64(-2415018.5)
	UnixEpoch   = JulianEpoch + XFloat64(2440587.5)

	ApproxDaysPerMonth = 30.4375
	ApproxDaysPerYear  = 365.25
)

func TimeToStringYmdHms(t time.Time) XString {
	return XString(t.Format(string(FmtYMdHms)))
}

func TimeToStringYmdHmsS(t time.Time) XString {
	return XString(t.Format(string(FmtYMdHmsSSS)))
}

func TimeToStringFormat(t time.Time, format XString) XString {
	return XString(t.Format(string(format)))
}

func ParseTimeYmsHms(timeStr XString) (time.Time, error) {
	return time.ParseInLocation(fmt.Sprintf("%s-%s-%s %s:%s:%s", Year, Month, Day, Hour, Minute, Second), string(timeStr), time.Local)
}

func ParseTimeYmsHmsS(timeStr XString) (time.Time, error) {
	return time.ParseInLocation(fmt.Sprintf("%s-%s-%s %s:%s:%s.000", Year, Month, Day, Hour, Minute, Second), string(timeStr), time.Local)
}

func ParseTimeYmsHmsLoc(timeStr XString, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(fmt.Sprintf("%s-%s-%s %s:%s:%s", Year, Month, Day, Hour, Minute, Second), string(timeStr), loc)
}

func ParseTimeYmsHmsSLoc(timeStr XString, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(fmt.Sprintf("%s-%s-%s %s:%s:%s.000", Year, Month, Day, Hour, Minute, Second), string(timeStr), loc)
}

// TimeInMillis 13位Unix时间戳
func TimeInMillis() XInt64 {
	return XInt64(time.Now().UnixMilli())
}

// TimeInSeconds 10位Unix时间戳
func TimeInSeconds() XInt64 {
	return XInt64(time.Now().Unix())
}

// TimeInMicro 16位Unix时间戳
func TimeInMicro() XInt64 {
	return XInt64(time.Now().UnixMicro())
}

// TimeInNano 19位Unix时间戳
func TimeInNano() XInt64 {
	return XInt64(time.Now().UnixNano())
}

func CurrentMinuteOfDay() XInt {
	t := time.Now()
	return XInt(t.Hour()*60 + t.Minute())
}

func CurrentSecondOfDay() XInt {
	t := time.Now()
	return XInt(t.Hour()*3600 + t.Minute()*60 + t.Second())
}

func MinuteOfDay(t time.Time) XInt {
	return XInt(t.Hour()*60 + t.Minute())
}

func SecondOfDay(t time.Time) XInt {
	return XInt(t.Hour()*3600 + t.Minute()*60 + t.Second())
}

func MinutesToTime(minutes XInt) (hour XInt, minute XInt) {
	h0 := minutes / 60
	m0 := minutes % 60
	return h0, m0
}

func SecondsToTime(seconds XInt) (hour XInt, minute XInt, second XInt) {
	h0, m0 := MinutesToTime(seconds / 60)
	s0 := seconds % 60
	return h0, m0, s0
}

func IsLeapYear(year int) XBool {
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}

func YearsBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Hours() / 24 / ApproxDaysPerYear))
}

func MonthsBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Hours() / 24 / ApproxDaysPerMonth))
}

func DaysBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Hours() / 24))
}

func HoursBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Hours()))
}

func MinutesBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Minutes()))
}

func SecondsBetween(now, then time.Time) XInt {
	return XInt(math.Trunc(now.Sub(then).Seconds()))
}

func MilliSecondsBetween(now, then time.Time) XInt64 {
	return XInt64(now.Sub(then).Milliseconds())
}

func WithInPastYears(now, then time.Time, years XInt) XBool {
	return YearsBetween(now, then) <= years
}

func WithInPastMonths(now, then time.Time, months XInt) XBool {
	return MonthsBetween(now, then) <= months
}

func WithInPastDays(now, then time.Time, days XInt) XBool {
	return DaysBetween(now, then) <= days
}

func WithInPastHours(now, then time.Time, hours XInt) XBool {
	return HoursBetween(now, then) <= hours
}

func WithInPastMinutes(now, then time.Time, minutes XInt) XBool {
	return MinutesBetween(now, then) <= minutes
}

func WithInPastSeconds(now, then time.Time, seconds XInt) XBool {
	return SecondsBetween(now, then) <= seconds
}

func WithInPastMilliSeconds(now, then time.Time, milliSeconds XInt64) XBool {
	return MilliSecondsBetween(now, then) <= milliSeconds
}

func YearSpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Hours() / 24 / ApproxDaysPerYear)
}

func MonthSpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Hours() / 24 / ApproxDaysPerMonth)
}

func DaySpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Hours() / 24)
}

func HourSpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Hours())
}

func MinuteSpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Minutes())
}

func SecondSpan(now, then time.Time) XFloat64 {
	return XFloat64(now.Sub(then).Seconds())
}

func MilliSecondSpan(now, then time.Time) XInt64 {
	return XInt64(now.Sub(then).Milliseconds())
}

func Now() time.Time {
	return time.Now()
}

func AddHour(times time.Time, plusOrMinus XString, seconds XString) time.Time {
	h, _ := time.ParseDuration(fmt.Sprintf("%s%v", plusOrMinus, seconds))
	return times.Add(h)
}

func AddMinutes(times time.Time, plusOrMinus XString, minutes XString) time.Time {
	h, _ := time.ParseDuration(fmt.Sprintf("%s%v", plusOrMinus, minutes))
	return times.Add(h)
}

func AddSeconds(times time.Time, plusOrMinus XString, hours XString) time.Time {
	h, _ := time.ParseDuration(fmt.Sprintf("%s%v", plusOrMinus, hours))
	return times.Add(h)
}

func AddDays(times time.Time, days XInt) time.Time {
	return times.AddDate(0, 0, int(days))
}

func AddMonths(times time.Time, month XInt) time.Time {
	return times.AddDate(0, int(month), 0)
}

func AddYears(times time.Time, year XInt) time.Time {
	return times.AddDate(int(year), 0, 0)
}

func ParseTime(timeStr XString) time.Time {
	timeStr = XString(strings.TrimSpace(string(timeStr)))
	timeStr = XString(strings.TrimSpace(strings.ReplaceAll(string(timeStr), "\\'", " ")))

	if timeStr == "" {
		return EmptyTime
	}
	if yRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtY), string(timeStr)); err == nil {
			return times
		}
	} else if yyyyMmDdRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYYYYMMdd), string(timeStr)); err == nil {
			return times
		}
	} else if ymRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYM), string(timeStr)); err == nil {
			return times
		}
	} else if ymdRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYMd), string(timeStr)); err == nil {
			return times
		}
	} else if ymdHRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYMdH), string(timeStr)); err == nil {
			return times
		}
	} else if ymdHmRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYMdHm), string(timeStr)); err == nil {
			return times
		}
	} else if ymdHmsRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYMdHms), string(timeStr)); err == nil {
			return times
		}
	} else if ymdHmsSRegex.MatchString(string(timeStr)) {
		if times, err := time.Parse(string(FmtYMdHmsSSS), string(timeStr)); err == nil {
			return times
		}
	}
	return EmptyTime
}

func IsTimeEmpty(time time.Time) XBool {
	return time == EmptyTime
}

func NumToTimeDuration(num XInt, duration time.Duration) time.Duration {
	int64Num, _ := strconv.ParseInt(fmt.Sprintf("%v", num), 10, 64)
	return time.Duration(int64Num * duration.Nanoseconds())
}
