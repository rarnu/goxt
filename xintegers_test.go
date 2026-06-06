package goxt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

// ==================== XInt RangeToOpen 测试 ====================

func TestXIntRangeToOpen(t *testing.T) {
	t.Run("基本范围 [0, 5)", func(t *testing.T) {
		result := goxt.XInt(0).RangeToOpen(5)
		assert.Equal(t, goxt.XInt(5), result.Size())
		assert.Equal(t, goxt.XInt(0), result[0])
		assert.Equal(t, goxt.XInt(4), result[4])
	})

	t.Run("负数范围 [-2, 3)", func(t *testing.T) {
		result := goxt.XInt(-2).RangeToOpen(3)
		assert.Equal(t, goxt.XInt(5), result.Size())
		assert.Equal(t, goxt.XInt(-2), result[0])
		assert.Equal(t, goxt.XInt(2), result[4])
	})

	t.Run("相同起点和终点返回空列表", func(t *testing.T) {
		result := goxt.XInt(5).RangeToOpen(5)
		assert.Equal(t, goxt.XInt(0), result.Size())
		assert.True(t, bool(result.IsEmpty()))
	})

	t.Run("单个元素范围", func(t *testing.T) {
		result := goxt.XInt(10).RangeToOpen(11)
		assert.Equal(t, goxt.XInt(1), result.Size())
		assert.Equal(t, goxt.XInt(10), result[0])
	})
}

// ==================== XInt RangeToClose 测试 ====================

func TestXIntRangeToClose(t *testing.T) {
	t.Run("基本范围 [0, 5]", func(t *testing.T) {
		result := goxt.XInt(0).RangeToClose(5)
		assert.Equal(t, goxt.XInt(6), result.Size())
		assert.Equal(t, goxt.XInt(0), result[0])
		assert.Equal(t, goxt.XInt(5), result[5])
	})

	t.Run("负数范围 [-2, 2]", func(t *testing.T) {
		result := goxt.XInt(-2).RangeToClose(2)
		assert.Equal(t, goxt.XInt(5), result.Size())
		assert.Equal(t, goxt.XInt(-2), result[0])
		assert.Equal(t, goxt.XInt(2), result[4])
	})

	t.Run("相同起点和终点返回单元素", func(t *testing.T) {
		result := goxt.XInt(5).RangeToClose(5)
		assert.Equal(t, goxt.XInt(1), result.Size())
		assert.Equal(t, goxt.XInt(5), result[0])
	})

	t.Run("单个元素范围", func(t *testing.T) {
		result := goxt.XInt(10).RangeToClose(11)
		assert.Equal(t, goxt.XInt(2), result.Size())
		assert.Equal(t, goxt.XInt(10), result[0])
		assert.Equal(t, goxt.XInt(11), result[1])
	})
}

// ==================== XInt ToStringRadix 测试 ====================

func TestXIntToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XInt(255).ToStringRadix(10)
		assert.EqualValues(t, "255", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XInt(10).ToStringRadix(2)
		assert.EqualValues(t, "1010", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XInt(255).ToStringRadix(16)
		assert.EqualValues(t, "ff", result)
	})

	t.Run("负数十进制转换", func(t *testing.T) {
		result := goxt.XInt(-42).ToStringRadix(10)
		assert.EqualValues(t, "-42", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XInt(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XInt Absolute 测试 ====================

func TestXIntAbsolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XInt(42).Absolute()
		assert.Equal(t, goxt.XInt(42), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XInt(-42).Absolute()
		assert.Equal(t, goxt.XInt(42), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XInt(0).Absolute()
		assert.Equal(t, goxt.XInt(0), result)
	})

}

// ==================== XInt8 ToStringRadix 测试 ====================

func TestXInt8ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XInt8(127).ToStringRadix(10)
		assert.EqualValues(t, "127", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XInt8(15).ToStringRadix(2)
		assert.EqualValues(t, "1111", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XInt8(127).ToStringRadix(16)
		assert.EqualValues(t, "7f", result)
	})

	t.Run("负数转换", func(t *testing.T) {
		result := goxt.XInt8(-50).ToStringRadix(10)
		assert.EqualValues(t, "-50", result)
	})
}

// ==================== XInt8 Absolute 测试 ====================

func TestXInt8Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XInt8(100).Absolute()
		assert.Equal(t, goxt.XInt8(100), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XInt8(-100).Absolute()
		assert.Equal(t, goxt.XInt8(100), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XInt8(0).Absolute()
		assert.Equal(t, goxt.XInt8(0), result)
	})
}

// ==================== XInt16 ToStringRadix 测试 ====================

func TestXInt16ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XInt16(1000).ToStringRadix(10)
		assert.EqualValues(t, "1000", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XInt16(255).ToStringRadix(2)
		assert.EqualValues(t, "11111111", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XInt16(4096).ToStringRadix(16)
		assert.EqualValues(t, "1000", result)
	})

	t.Run("负数转换", func(t *testing.T) {
		result := goxt.XInt16(-500).ToStringRadix(10)
		assert.EqualValues(t, "-500", result)
	})
}

// ==================== XInt16 Absolute 测试 ====================

func TestXInt16Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XInt16(500).Absolute()
		assert.Equal(t, goxt.XInt16(500), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XInt16(-500).Absolute()
		assert.Equal(t, goxt.XInt16(500), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XInt16(0).Absolute()
		assert.Equal(t, goxt.XInt16(0), result)
	})
}

// ==================== XInt32 ToStringRadix 测试 ====================

func TestXInt32ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XInt32(100000).ToStringRadix(10)
		assert.EqualValues(t, "100000", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XInt32(1024).ToStringRadix(2)
		assert.EqualValues(t, "10000000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XInt32(65535).ToStringRadix(16)
		assert.EqualValues(t, "ffff", result)
	})

	t.Run("负数转换", func(t *testing.T) {
		result := goxt.XInt32(-12345).ToStringRadix(10)
		assert.EqualValues(t, "-12345", result)
	})
}

// ==================== XInt32 Absolute 测试 ====================

func TestXInt32Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XInt32(12345).Absolute()
		assert.Equal(t, goxt.XInt32(12345), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XInt32(-12345).Absolute()
		assert.Equal(t, goxt.XInt32(12345), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XInt32(0).Absolute()
		assert.Equal(t, goxt.XInt32(0), result)
	})
}

// ==================== XInt64 ToStringRadix 测试 ====================

func TestXInt64ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XInt64(9223372036854775807).ToStringRadix(10)
		assert.EqualValues(t, "9223372036854775807", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XInt64(1024).ToStringRadix(2)
		assert.EqualValues(t, "10000000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XInt64(281474976710655).ToStringRadix(16)
		assert.EqualValues(t, "ffffffffffff", result)
	})

	t.Run("负数转换", func(t *testing.T) {
		result := goxt.XInt64(-9223372036854775808).ToStringRadix(10)
		assert.EqualValues(t, "-9223372036854775808", result)
	})
}

// ==================== XInt64 Absolute 测试 ====================

func TestXInt64Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XInt64(9223372036854775807).Absolute()
		assert.Equal(t, goxt.XInt64(9223372036854775807), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XInt64(0).Absolute()
		assert.Equal(t, goxt.XInt64(0), result)
	})
}

// ==================== XUint ToStringRadix 测试 ====================

func TestXUintToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XUint(4294967295).ToStringRadix(10)
		assert.EqualValues(t, "4294967295", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XUint(255).ToStringRadix(2)
		assert.EqualValues(t, "11111111", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XUint(4294967295).ToStringRadix(16)
		assert.EqualValues(t, "ffffffff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XUint(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XUint8 ToStringRadix 测试 ====================

func TestXUint8ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XUint8(255).ToStringRadix(10)
		assert.EqualValues(t, "255", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XUint8(128).ToStringRadix(2)
		assert.EqualValues(t, "10000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XUint8(255).ToStringRadix(16)
		assert.EqualValues(t, "ff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XUint8(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XUint16 ToStringRadix 测试 ====================

func TestXUint16ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XUint16(65535).ToStringRadix(10)
		assert.EqualValues(t, "65535", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XUint16(1024).ToStringRadix(2)
		assert.EqualValues(t, "10000000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XUint16(65535).ToStringRadix(16)
		assert.EqualValues(t, "ffff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XUint16(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XUint32 ToStringRadix 测试 ====================

func TestXUint32ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XUint32(4294967295).ToStringRadix(10)
		assert.EqualValues(t, "4294967295", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XUint32(2048).ToStringRadix(2)
		assert.EqualValues(t, "100000000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XUint32(4294967295).ToStringRadix(16)
		assert.EqualValues(t, "ffffffff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XUint32(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XUint64 ToStringRadix 测试 ====================

func TestXUint64ToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XUint64(18446744073709551615).ToStringRadix(10)
		assert.EqualValues(t, "18446744073709551615", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XUint64(4096).ToStringRadix(2)
		assert.EqualValues(t, "1000000000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XUint64(18446744073709551615).ToStringRadix(16)
		assert.EqualValues(t, "ffffffffffffffff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XUint64(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})
}

// ==================== XByte ToStringRadix 测试 ====================

func TestXByteToStringRadix(t *testing.T) {
	t.Run("十进制转换", func(t *testing.T) {
		result := goxt.XByte(255).ToStringRadix(10)
		assert.EqualValues(t, "255", result)
	})

	t.Run("二进制转换", func(t *testing.T) {
		result := goxt.XByte(128).ToStringRadix(2)
		assert.EqualValues(t, "10000000", result)
	})

	t.Run("十六进制转换", func(t *testing.T) {
		result := goxt.XByte(255).ToStringRadix(16)
		assert.EqualValues(t, "ff", result)
	})

	t.Run("零值转换", func(t *testing.T) {
		result := goxt.XByte(0).ToStringRadix(10)
		assert.EqualValues(t, "0", result)
	})

	t.Run("八进制转换", func(t *testing.T) {
		result := goxt.XByte(64).ToStringRadix(8)
		assert.EqualValues(t, "100", result)
	})
}
