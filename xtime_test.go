package goxt_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

// ==================== 时间格式化函数测试 ====================

func TestTimeToStringYmdHms(t *testing.T) {
	// 测试用例1: 标准时间格式化
	t1 := time.Date(2023, 6, 15, 14, 30, 45, 0, time.Local)
	result := goxt.TimeToStringYmdHms(t1)
	assert.Equal(t, goxt.XString("2023-06-15 14:30:45"), result)

	// 测试用例2: 零点时间
	t2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	result2 := goxt.TimeToStringYmdHms(t2)
	assert.Equal(t, goxt.XString("2023-01-01 00:00:00"), result2)

	// 测试用例3: 月末时间
	t3 := time.Date(2023, 12, 31, 23, 59, 59, 0, time.Local)
	result3 := goxt.TimeToStringYmdHms(t3)
	assert.Equal(t, goxt.XString("2023-12-31 23:59:59"), result3)
}

func TestTimeToStringYmdHmsS(t *testing.T) {
	// 测试用例1: 带毫秒的时间格式化
	t1 := time.Date(2023, 6, 15, 14, 30, 45, 123000000, time.Local)
	result := goxt.TimeToStringYmdHmsS(t1)
	assert.Equal(t, goxt.XString("2023-06-15 14:30:45.123"), result)

	// 测试用例2: 零毫秒
	t2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	result2 := goxt.TimeToStringYmdHmsS(t2)
	assert.Equal(t, goxt.XString("2023-01-01 00:00:00.000"), result2)

	// 测试用例3: 999毫秒
	t3 := time.Date(2023, 12, 31, 23, 59, 59, 999000000, time.Local)
	result3 := goxt.TimeToStringYmdHmsS(t3)
	assert.Equal(t, goxt.XString("2023-12-31 23:59:59.999"), result3)
}

func TestTimeToStringFormat(t *testing.T) {
	// 测试用例1: 自定义格式 - 年月日
	t1 := time.Date(2023, 6, 15, 14, 30, 45, 0, time.Local)
	result := goxt.TimeToStringFormat(t1, goxt.XString("2006年01月02日"))
	assert.Equal(t, goxt.XString("2023年06月15日"), result)

	// 测试用例2: 自定义格式 - 时分秒
	result2 := goxt.TimeToStringFormat(t1, goxt.XString("15时04分05秒"))
	assert.Equal(t, goxt.XString("14时30分45秒"), result2)

	// 测试用例3: 使用预定义格式
	result3 := goxt.TimeToStringFormat(t1, goxt.FmtYMd)
	assert.Equal(t, goxt.XString("2023-06-15"), result3)

	// 测试用例4: 仅年份
	result4 := goxt.TimeToStringFormat(t1, goxt.FmtY)
	assert.Equal(t, goxt.XString("2023"), result4)
}

// ==================== 时间解析函数测试 ====================

func TestParseTimeYmsHms(t *testing.T) {
	// 测试用例1: 标准时间字符串解析
	result, err := goxt.ParseTimeYmsHms(goxt.XString("2023-06-15 14:30:45"))
	assert.NoError(t, err)
	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, time.June, result.Month())
	assert.Equal(t, 15, result.Day())
	assert.Equal(t, 14, result.Hour())
	assert.Equal(t, 30, result.Minute())
	assert.Equal(t, 45, result.Second())

	// 测试用例2: 零点时间
	result2, err2 := goxt.ParseTimeYmsHms(goxt.XString("2023-01-01 00:00:00"))
	assert.NoError(t, err2)
	assert.Equal(t, 0, result2.Hour())

	// 测试用例3: 错误格式
	_, err3 := goxt.ParseTimeYmsHms(goxt.XString("invalid-time"))
	assert.Error(t, err3)
}

func TestParseTimeYmsHmsS(t *testing.T) {
	// 测试用例1: 带毫秒的时间字符串解析
	result, err := goxt.ParseTimeYmsHmsS(goxt.XString("2023-06-15 14:30:45.123"))
	assert.NoError(t, err)
	assert.Equal(t, 123, result.Nanosecond()/1000000) // 转换为毫秒

	// 测试用例2: 零毫秒
	result2, err2 := goxt.ParseTimeYmsHmsS(goxt.XString("2023-01-01 00:00:00.000"))
	assert.NoError(t, err2)
	assert.Equal(t, 0, result2.Nanosecond())

	// 测试用例3: 错误格式
	_, err3 := goxt.ParseTimeYmsHmsS(goxt.XString("invalid-time"))
	assert.Error(t, err3)
}

func TestParseTimeYmsHmsLoc(t *testing.T) {
	// 测试用例1: 使用本地时区
	loc := time.Local
	result, err := goxt.ParseTimeYmsHmsLoc(goxt.XString("2023-06-15 14:30:45"), loc)
	assert.NoError(t, err)
	assert.Equal(t, 2023, result.Year())

	// 测试用例2: 使用UTC时区
	utcLoc := time.UTC
	result2, err2 := goxt.ParseTimeYmsHmsLoc(goxt.XString("2023-06-15 14:30:45"), utcLoc)
	assert.NoError(t, err2)
	assert.Equal(t, time.UTC, result2.Location())

	// 测试用例3: 错误格式
	_, err3 := goxt.ParseTimeYmsHmsLoc(goxt.XString("invalid"), loc)
	assert.Error(t, err3)
}

func TestParseTimeYmsHmsSLoc(t *testing.T) {
	// 测试用例1: 使用本地时区解析带毫秒时间
	loc := time.Local
	result, err := goxt.ParseTimeYmsHmsSLoc(goxt.XString("2023-06-15 14:30:45.123"), loc)
	assert.NoError(t, err)
	assert.Equal(t, 123, result.Nanosecond()/1000000)

	// 测试用例2: 使用UTC时区
	utcLoc := time.UTC
	result2, err2 := goxt.ParseTimeYmsHmsSLoc(goxt.XString("2023-06-15 14:30:45.456"), utcLoc)
	assert.NoError(t, err2)
	assert.Equal(t, time.UTC, result2.Location())

	// 测试用例3: 错误格式
	_, err3 := goxt.ParseTimeYmsHmsSLoc(goxt.XString("invalid"), loc)
	assert.Error(t, err3)
}

// ==================== 时间戳函数测试 ====================

func TestTimeInMillis(t *testing.T) {
	// 测试用例1: 返回13位时间戳
	result := goxt.TimeInMillis()
	assert.Greater(t, int64(result), int64(1000000000000))
	assert.Less(t, int64(result), int64(9999999999999))

	// 测试用例2: 时间戳应该是递增的
	time.Sleep(1 * time.Millisecond)
	result2 := goxt.TimeInMillis()
	assert.GreaterOrEqual(t, int64(result2), int64(result))
}

func TestTimeInSeconds(t *testing.T) {
	// 测试用例1: 返回10位时间戳
	result := goxt.TimeInSeconds()
	assert.Greater(t, int64(result), int64(1000000000))
	assert.Less(t, int64(result), int64(9999999999))

	// 测试用例2: 与Unix时间一致
	expected := time.Now().Unix()
	assert.InDelta(t, int64(result), expected, 1)
}

func TestTimeInMicro(t *testing.T) {
	// 测试用例1: 返回16位时间戳
	result := goxt.TimeInMicro()
	assert.Greater(t, int64(result), int64(1000000000000000))
	// 使用合理的上界，远小于int64最大值
	assert.Less(t, int64(result), int64(9999999999999999))

	// 测试用例2: 精度高于毫秒
	millis := goxt.TimeInMillis()
	micros := goxt.TimeInMicro()
	assert.Greater(t, int64(micros), int64(millis)*1000)
}

func TestTimeInNano(t *testing.T) {
	// 测试用例1: 返回19位时间戳
	result := goxt.TimeInNano()
	assert.Greater(t, int64(result), int64(1000000000000000000))
	// int64最大值是9223372036854775807，使用合理的上界
	assert.Less(t, int64(result), int64(9223372036854775807))

	// 测试用例2: 精度高于微秒，添加延迟确保时间差异
	micros := goxt.TimeInMicro()
	time.Sleep(1 * time.Microsecond)
	nanos := goxt.TimeInNano()
	assert.GreaterOrEqual(t, int64(nanos), int64(micros)*1000)
}

// ==================== 时间计算函数测试 ====================

func TestCurrentMinuteOfDay(t *testing.T) {
	// 测试用例1: 函数应该返回当前分钟的合理值
	result := goxt.CurrentMinuteOfDay()
	assert.GreaterOrEqual(t, int(result), 0)
	assert.LessOrEqual(t, int(result), 1439)

	// 测试用例2: 零点
	t1 := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	result1 := goxt.MinuteOfDay(t1)
	assert.Equal(t, goxt.XInt(0), result1)

	// 测试用例3: 中午12点
	t2 := time.Date(2023, 6, 15, 12, 0, 0, 0, time.Local)
	result2 := goxt.MinuteOfDay(t2)
	assert.Equal(t, goxt.XInt(720), result2)

	// 测试用例4: 23:59
	t3 := time.Date(2023, 6, 15, 23, 59, 0, 0, time.Local)
	result3 := goxt.MinuteOfDay(t3)
	assert.Equal(t, goxt.XInt(1439), result3)
}

func TestCurrentSecondOfDay(t *testing.T) {
	// 测试用例1: 函数应该返回当前秒的合理值
	result := goxt.CurrentSecondOfDay()
	assert.GreaterOrEqual(t, int(result), 0)
	assert.LessOrEqual(t, int(result), 86399)

	// 测试用例2: 零点
	t1 := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	result1 := goxt.SecondOfDay(t1)
	assert.Equal(t, goxt.XInt(0), result1)

	// 测试用例3: 1小时
	t2 := time.Date(2023, 6, 15, 1, 0, 0, 0, time.Local)
	result2 := goxt.SecondOfDay(t2)
	assert.Equal(t, goxt.XInt(3600), result2)

	// 测试用例4: 23:59:59
	t3 := time.Date(2023, 6, 15, 23, 59, 59, 0, time.Local)
	result3 := goxt.SecondOfDay(t3)
	assert.Equal(t, goxt.XInt(86399), result3)
}

func TestMinutesToTime(t *testing.T) {
	// 测试用例1: 0分钟
	h, m := goxt.MinutesToTime(0)
	assert.Equal(t, goxt.XInt(0), h)
	assert.Equal(t, goxt.XInt(0), m)

	// 测试用例2: 90分钟 = 1小时30分钟
	h2, m2 := goxt.MinutesToTime(90)
	assert.Equal(t, goxt.XInt(1), h2)
	assert.Equal(t, goxt.XInt(30), m2)

	// 测试用例3: 1439分钟 = 23小时59分钟
	h3, m3 := goxt.MinutesToTime(1439)
	assert.Equal(t, goxt.XInt(23), h3)
	assert.Equal(t, goxt.XInt(59), m3)

	// 测试用例4: 1440分钟 = 24小时0分钟
	h4, m4 := goxt.MinutesToTime(1440)
	assert.Equal(t, goxt.XInt(24), h4)
	assert.Equal(t, goxt.XInt(0), m4)
}

func TestSecondsToTime(t *testing.T) {
	// 测试用例1: 0秒
	h, m, s := goxt.SecondsToTime(0)
	assert.Equal(t, goxt.XInt(0), h)
	assert.Equal(t, goxt.XInt(0), m)
	assert.Equal(t, goxt.XInt(0), s)

	// 测试用例2: 3661秒 = 1小时1分钟1秒
	h2, m2, s2 := goxt.SecondsToTime(3661)
	assert.Equal(t, goxt.XInt(1), h2)
	assert.Equal(t, goxt.XInt(1), m2)
	assert.Equal(t, goxt.XInt(1), s2)

	// 测试用例3: 86399秒 = 23小时59分钟59秒
	h3, m3, s3 := goxt.SecondsToTime(86399)
	assert.Equal(t, goxt.XInt(23), h3)
	assert.Equal(t, goxt.XInt(59), m3)
	assert.Equal(t, goxt.XInt(59), s3)
}

// ==================== 闰年和间隔函数测试 ====================

func TestIsLeapYear(t *testing.T) {
	// 测试用例1: 普通闰年(能被4整除)
	assert.True(t, bool(goxt.IsLeapYear(2024)))

	// 测试用例2: 世纪年但不是闰年(能被100整除但不能被400整除)
	assert.False(t, bool(goxt.IsLeapYear(1900)))

	// 测试用例3: 世纪闰年(能被400整除)
	assert.True(t, bool(goxt.IsLeapYear(2000)))

	// 测试用例4: 普通非闰年
	assert.False(t, bool(goxt.IsLeapYear(2023)))

	// 测试用例5: 边界情况
	assert.True(t, bool(goxt.IsLeapYear(2004)))
	assert.False(t, bool(goxt.IsLeapYear(2001)))
}

func TestYearsBetween(t *testing.T) {
	// 测试用例1: 相差1年
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	then := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	result := goxt.YearsBetween(now, then)
	assert.Equal(t, goxt.XInt(1), result)

	// 测试用例2: 相差2年
	then2 := time.Date(2022, 6, 15, 0, 0, 0, 0, time.Local)
	result2 := goxt.YearsBetween(now, then2)
	assert.Equal(t, goxt.XInt(2), result2)

	// 测试用例3: 同一年
	then3 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	result3 := goxt.YearsBetween(now, then3)
	assert.GreaterOrEqual(t, int(result3), 0)
}

func TestMonthsBetween(t *testing.T) {
	// 测试用例1: 相差1个月
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	then := time.Date(2024, 5, 15, 0, 0, 0, 0, time.Local)
	result := goxt.MonthsBetween(now, then)
	assert.GreaterOrEqual(t, int(result), 1)

	// 测试用例2: 相差12个月
	then2 := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	result2 := goxt.MonthsBetween(now, then2)
	assert.GreaterOrEqual(t, int(result2), 12)
}

func TestDaysBetween(t *testing.T) {
	// 测试用例1: 相差1天
	now := time.Date(2024, 6, 16, 0, 0, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	result := goxt.DaysBetween(now, then)
	assert.Equal(t, goxt.XInt(1), result)

	// 测试用例2: 相差30天
	then2 := time.Date(2024, 5, 17, 0, 0, 0, 0, time.Local)
	result2 := goxt.DaysBetween(now, then2)
	assert.Equal(t, goxt.XInt(30), result2)

	// 测试用例3: 同一天
	then3 := time.Date(2024, 6, 16, 12, 0, 0, 0, time.Local)
	result3 := goxt.DaysBetween(now, then3)
	assert.Equal(t, goxt.XInt(0), result3)
}

func TestHoursBetween(t *testing.T) {
	// 测试用例1: 相差1小时
	now := time.Date(2024, 6, 15, 13, 0, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.HoursBetween(now, then)
	assert.Equal(t, goxt.XInt(1), result)

	// 测试用例2: 相差24小时
	then2 := time.Date(2024, 6, 14, 13, 0, 0, 0, time.Local)
	result2 := goxt.HoursBetween(now, then2)
	assert.Equal(t, goxt.XInt(24), result2)
}

func TestMinutesBetween(t *testing.T) {
	// 测试用例1: 相差1分钟
	now := time.Date(2024, 6, 15, 12, 1, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.MinutesBetween(now, then)
	assert.Equal(t, goxt.XInt(1), result)

	// 测试用例2: 相差60分钟
	then2 := time.Date(2024, 6, 15, 11, 1, 0, 0, time.Local)
	result2 := goxt.MinutesBetween(now, then2)
	assert.Equal(t, goxt.XInt(60), result2)
}

func TestSecondsBetween(t *testing.T) {
	// 测试用例1: 相差1秒
	now := time.Date(2024, 6, 15, 12, 0, 1, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.SecondsBetween(now, then)
	assert.Equal(t, goxt.XInt(1), result)

	// 测试用例2: 相差60秒
	then2 := time.Date(2024, 6, 15, 11, 59, 1, 0, time.Local)
	result2 := goxt.SecondsBetween(now, then2)
	assert.Equal(t, goxt.XInt(60), result2)
}

func TestMilliSecondsBetween(t *testing.T) {
	// 测试用例1: 相差1毫秒
	now := time.Date(2024, 6, 15, 12, 0, 0, 1000000, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.MilliSecondsBetween(now, then)
	assert.Equal(t, goxt.XInt64(1), result)

	// 测试用例2: 相差1000毫秒
	then2 := time.Date(2024, 6, 15, 11, 59, 59, 0, time.Local)
	result2 := goxt.MilliSecondsBetween(now, then2)
	assert.Equal(t, goxt.XInt64(1001), result2)
}

// ==================== WithinPast 系列函数测试 ====================

func TestWithInPastYears(t *testing.T) {
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 在1年内（相差半年）
	then1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, bool(goxt.WithInPastYears(now, then1, 1)))

	// 测试用例2: 超过1年（相差约2年，确保YearsBetween返回>=2）
	then2 := time.Date(2022, 6, 1, 0, 0, 0, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastYears(now, then2, 1)))

	// 测试用例3: 在2年内
	assert.True(t, bool(goxt.WithInPastYears(now, then2, 2)))
}

func TestWithInPastMonths(t *testing.T) {
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 在1个月内（相差半个月）
	then1 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)
	assert.True(t, bool(goxt.WithInPastMonths(now, then1, 1)))

	// 测试用例2: 超过1个月（相差约3个月，确保MonthsBetween返回>=2）
	then2 := time.Date(2024, 3, 10, 0, 0, 0, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastMonths(now, then2, 1)))

	// 测试用例3: 在3个月内
	assert.True(t, bool(goxt.WithInPastMonths(now, then2, 3)))
}

func TestWithInPastDays(t *testing.T) {
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 在1天内
	then1 := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	assert.True(t, bool(goxt.WithInPastDays(now, then1, 1)))

	// 测试用例2: 超过1天
	then2 := time.Date(2024, 6, 13, 0, 0, 0, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastDays(now, then2, 1)))

	// 测试用例3: 在7天内
	assert.True(t, bool(goxt.WithInPastDays(now, then2, 7)))
}

func TestWithInPastHours(t *testing.T) {
	now := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 在1小时内
	then1 := time.Date(2024, 6, 15, 11, 30, 0, 0, time.Local)
	assert.True(t, bool(goxt.WithInPastHours(now, then1, 1)))

	// 测试用例2: 超过1小时
	then2 := time.Date(2024, 6, 15, 10, 0, 0, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastHours(now, then2, 1)))

	// 测试用例3: 在3小时内
	assert.True(t, bool(goxt.WithInPastHours(now, then2, 3)))
}

func TestWithInPastMinutes(t *testing.T) {
	now := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 在1分钟内
	then1 := time.Date(2024, 6, 15, 11, 59, 30, 0, time.Local)
	assert.True(t, bool(goxt.WithInPastMinutes(now, then1, 1)))

	// 测试用例2: 超过1分钟
	then2 := time.Date(2024, 6, 15, 11, 58, 0, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastMinutes(now, then2, 1)))

	// 测试用例3: 在5分钟内
	assert.True(t, bool(goxt.WithInPastMinutes(now, then2, 5)))
}

func TestWithInPastSeconds(t *testing.T) {
	now := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 在1秒内
	then1 := time.Date(2024, 6, 15, 11, 59, 59, 500000000, time.Local)
	assert.True(t, bool(goxt.WithInPastSeconds(now, then1, 1)))

	// 测试用例2: 超过1秒
	then2 := time.Date(2024, 6, 15, 11, 59, 58, 0, time.Local)
	assert.False(t, bool(goxt.WithInPastSeconds(now, then2, 1)))

	// 测试用例3: 在10秒内
	assert.True(t, bool(goxt.WithInPastSeconds(now, then2, 10)))
}

func TestWithInPastMilliSeconds(t *testing.T) {
	now := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 在100毫秒内
	then1 := time.Date(2024, 6, 15, 11, 59, 59, 950000000, time.Local)
	assert.True(t, bool(goxt.WithInPastMilliSeconds(now, then1, 100)))

	// 测试用例2: 超过100毫秒
	then2 := time.Date(2024, 6, 15, 11, 59, 59, 800000000, time.Local)
	assert.False(t, bool(goxt.WithInPastMilliSeconds(now, then2, 100)))

	// 测试用例3: 在500毫秒内
	assert.True(t, bool(goxt.WithInPastMilliSeconds(now, then2, 500)))
}

// ==================== Span 系列函数测试 ====================

func TestYearSpan(t *testing.T) {
	// 测试用例1: 相差1年
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	then := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	result := goxt.YearSpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.1)

	// 测试用例2: 相差0.5年
	then2 := time.Date(2023, 12, 15, 0, 0, 0, 0, time.Local)
	result2 := goxt.YearSpan(now, then2)
	assert.InDelta(t, float64(result2), 0.5, 0.1)
}

func TestMonthSpan(t *testing.T) {
	// 测试用例1: 相差1个月
	now := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	then := time.Date(2024, 5, 15, 0, 0, 0, 0, time.Local)
	result := goxt.MonthSpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.1)

	// 测试用例2: 相差6个月
	then2 := time.Date(2023, 12, 15, 0, 0, 0, 0, time.Local)
	result2 := goxt.MonthSpan(now, then2)
	assert.InDelta(t, float64(result2), 6.0, 0.5)
}

func TestDaySpan(t *testing.T) {
	// 测试用例1: 相差1天
	now := time.Date(2024, 6, 16, 0, 0, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	result := goxt.DaySpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.01)

	// 测试用例2: 相差半天
	then2 := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result2 := goxt.DaySpan(now, then2)
	assert.InDelta(t, float64(result2), 0.5, 0.01)
}

func TestHourSpan(t *testing.T) {
	// 测试用例1: 相差1小时
	now := time.Date(2024, 6, 15, 13, 0, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.HourSpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.01)

	// 测试用例2: 相差1.5小时
	then2 := time.Date(2024, 6, 15, 11, 30, 0, 0, time.Local)
	result2 := goxt.HourSpan(now, then2)
	assert.InDelta(t, float64(result2), 1.5, 0.01)
}

func TestMinuteSpan(t *testing.T) {
	// 测试用例1: 相差1分钟
	now := time.Date(2024, 6, 15, 12, 1, 0, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.MinuteSpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.01)

	// 测试用例2: 相差30分钟
	then2 := time.Date(2024, 6, 15, 11, 31, 0, 0, time.Local)
	result2 := goxt.MinuteSpan(now, then2)
	assert.InDelta(t, float64(result2), 30.0, 0.01)
}

func TestSecondSpan(t *testing.T) {
	// 测试用例1: 相差1秒
	now := time.Date(2024, 6, 15, 12, 0, 1, 0, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.SecondSpan(now, then)
	assert.InDelta(t, float64(result), 1.0, 0.01)

	// 测试用例2: 相差30秒
	then2 := time.Date(2024, 6, 15, 11, 59, 31, 0, time.Local)
	result2 := goxt.SecondSpan(now, then2)
	assert.InDelta(t, float64(result2), 30.0, 0.01)
}

func TestMilliSecondSpan(t *testing.T) {
	// 测试用例1: 相差1毫秒
	now := time.Date(2024, 6, 15, 12, 0, 0, 1000000, time.Local)
	then := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)
	result := goxt.MilliSecondSpan(now, then)
	assert.Equal(t, goxt.XInt64(1), result)

	// 测试用例2: 相差500毫秒
	then2 := time.Date(2024, 6, 15, 11, 59, 59, 500000000, time.Local)
	result2 := goxt.MilliSecondSpan(now, then2)
	assert.Equal(t, goxt.XInt64(501), result2)
}

// ==================== 时间加减函数测试 ====================

func TestAddHour(t *testing.T) {
	base := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 加1小时
	result := goxt.AddHour(base, "+", goxt.XString("1h"))
	assert.Equal(t, 13, result.Hour())

	// 测试用例2: 减1小时
	result2 := goxt.AddHour(base, "-", goxt.XString("1h"))
	assert.Equal(t, 11, result2.Hour())

	// 测试用例3: 加2小时
	result3 := goxt.AddHour(base, "+", goxt.XString("2h"))
	assert.Equal(t, 14, result3.Hour())
}

func TestAddMinutes(t *testing.T) {
	base := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 加30分钟
	result := goxt.AddMinutes(base, "+", goxt.XString("30m"))
	assert.Equal(t, 30, result.Minute())

	// 测试用例2: 减30分钟
	result2 := goxt.AddMinutes(base, "-", goxt.XString("30m"))
	assert.Equal(t, 30, result2.Minute())
	assert.Equal(t, 11, result2.Hour())

	// 测试用例3: 加90分钟
	result3 := goxt.AddMinutes(base, "+", goxt.XString("90m"))
	assert.Equal(t, 30, result3.Minute())
	assert.Equal(t, 13, result3.Hour())
}

func TestAddSeconds(t *testing.T) {
	base := time.Date(2024, 6, 15, 12, 0, 0, 0, time.Local)

	// 测试用例1: 加30秒
	result := goxt.AddSeconds(base, "+", goxt.XString("30s"))
	assert.Equal(t, 30, result.Second())

	// 测试用例2: 减30秒
	result2 := goxt.AddSeconds(base, "-", goxt.XString("30s"))
	assert.Equal(t, 30, result2.Second())
	assert.Equal(t, 59, result2.Minute())

	// 测试用例3: 加90秒
	result3 := goxt.AddSeconds(base, "+", goxt.XString("90s"))
	assert.Equal(t, 30, result3.Second())
	assert.Equal(t, 1, result3.Minute())
}

func TestAddDays(t *testing.T) {
	base := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 加1天
	result := goxt.AddDays(base, 1)
	assert.Equal(t, 16, result.Day())

	// 测试用例2: 减1天
	result2 := goxt.AddDays(base, -1)
	assert.Equal(t, 14, result2.Day())

	// 测试用例3: 加30天
	result3 := goxt.AddDays(base, 30)
	assert.Equal(t, 15, result3.Day())
	assert.Equal(t, 7, int(result3.Month()))
}

func TestAddMonths(t *testing.T) {
	base := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 加1个月
	result := goxt.AddMonths(base, 1)
	assert.Equal(t, 7, int(result.Month()))

	// 测试用例2: 减1个月
	result2 := goxt.AddMonths(base, -1)
	assert.Equal(t, 5, int(result2.Month()))

	// 测试用例3: 加12个月
	result3 := goxt.AddMonths(base, 12)
	assert.Equal(t, 6, int(result3.Month()))
	assert.Equal(t, 2025, result3.Year())
}

func TestAddYears(t *testing.T) {
	base := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)

	// 测试用例1: 加1年
	result := goxt.AddYears(base, 1)
	assert.Equal(t, 2025, result.Year())

	// 测试用例2: 减1年
	result2 := goxt.AddYears(base, -1)
	assert.Equal(t, 2023, result2.Year())

	// 测试用例3: 加5年
	result3 := goxt.AddYears(base, 5)
	assert.Equal(t, 2029, result3.Year())
}

// ==================== ParseTime 和 IsTimeEmpty 测试 ====================

func TestParseTime(t *testing.T) {
	// 测试用例1: 解析完整日期时间
	result := goxt.ParseTime(goxt.XString("2023-06-15 14:30:45"))
	assert.NotEqual(t, goxt.EmptyTime, result)
	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, 14, result.Hour())

	// 测试用例2: 解析仅日期
	result2 := goxt.ParseTime(goxt.XString("2023-06-15"))
	assert.NotEqual(t, goxt.EmptyTime, result2)
	assert.Equal(t, 15, result2.Day())

	// 测试用例3: 解析年月
	result3 := goxt.ParseTime(goxt.XString("2023-06"))
	assert.NotEqual(t, goxt.EmptyTime, result3)
	assert.Equal(t, 6, int(result3.Month()))

	// 测试用例4: 解析纯年份
	result4 := goxt.ParseTime(goxt.XString("2023"))
	assert.NotEqual(t, goxt.EmptyTime, result4)
	assert.Equal(t, 2023, result4.Year())

	// 测试用例5: 解析紧凑格式
	result5 := goxt.ParseTime(goxt.XString("20230615"))
	assert.NotEqual(t, goxt.EmptyTime, result5)
	assert.Equal(t, 15, result5.Day())

	// 测试用例6: 解析带毫秒
	result6 := goxt.ParseTime(goxt.XString("2023-06-15 14:30:45.123"))
	assert.NotEqual(t, goxt.EmptyTime, result6)
	assert.Equal(t, 123, result6.Nanosecond()/1000000)

	// 测试用例7: 空字符串
	result7 := goxt.ParseTime(goxt.XString(""))
	assert.Equal(t, goxt.EmptyTime, result7)

	// 测试用例8: 无效格式
	result8 := goxt.ParseTime(goxt.XString("invalid"))
	assert.Equal(t, goxt.EmptyTime, result8)

	// 测试用例9: 带空格的时间字符串
	result9 := goxt.ParseTime(goxt.XString("  2023-06-15  "))
	assert.NotEqual(t, goxt.EmptyTime, result9)
}

func TestIsTimeEmpty(t *testing.T) {
	// 测试用例1: 空时间
	assert.True(t, bool(goxt.IsTimeEmpty(goxt.EmptyTime)))

	// 测试用例2: 零值时间
	var zeroTime time.Time
	assert.True(t, bool(goxt.IsTimeEmpty(zeroTime)))

	// 测试用例3: 非空时间
	nonEmpty := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)
	assert.False(t, bool(goxt.IsTimeEmpty(nonEmpty)))
}

func TestNumToTimeDuration(t *testing.T) {
	// 测试用例1: 转换1秒
	result := goxt.NumToTimeDuration(1, time.Second)
	assert.Equal(t, time.Second, result)

	// 测试用例2: 转换5分钟
	result2 := goxt.NumToTimeDuration(5, time.Minute)
	assert.Equal(t, 5*time.Minute, result2)

	// 测试用例3: 转换100毫秒
	result3 := goxt.NumToTimeDuration(100, time.Millisecond)
	assert.Equal(t, 100*time.Millisecond, result3)

	// 测试用例4: 转换0
	result4 := goxt.NumToTimeDuration(0, time.Second)
	assert.Equal(t, time.Duration(0), result4)
}

func TestNow(t *testing.T) {
	// 测试用例1: 返回当前时间
	result := goxt.Now()
	assert.NotEqual(t, goxt.EmptyTime, result)

	// 测试用例2: 时间在合理范围内
	now := time.Now()
	diff := now.Sub(result)
	assert.Less(t, math.Abs(diff.Seconds()), 1.0)
}
