package goxt_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

// ==================== XFloat32 IsNaN 测试 ====================

func TestXFloat32IsNaN(t *testing.T) {
	t.Run("NaN值返回true", func(t *testing.T) {
		result := goxt.XFloat32(float32(math.NaN())).IsNaN()
		assert.True(t, bool(result))
	})

	t.Run("正常值返回false", func(t *testing.T) {
		result := goxt.XFloat32(3.14).IsNaN()
		assert.False(t, bool(result))
	})

	t.Run("零值返回false", func(t *testing.T) {
		result := goxt.XFloat32(0).IsNaN()
		assert.False(t, bool(result))
	})

	t.Run("负数返回false", func(t *testing.T) {
		result := goxt.XFloat32(-2.5).IsNaN()
		assert.False(t, bool(result))
	})
}

// ==================== XFloat32 IsInfinite 测试 ====================

func TestXFloat32IsInfinite(t *testing.T) {
	t.Run("正无穷返回true", func(t *testing.T) {
		result := goxt.XFloat32(float32(math.Inf(1))).IsInfinite()
		assert.True(t, bool(result))
	})

	t.Run("负无穷返回true", func(t *testing.T) {
		result := goxt.XFloat32(float32(math.Inf(-1))).IsInfinite()
		assert.True(t, bool(result))
	})

	t.Run("正常值返回false", func(t *testing.T) {
		result := goxt.XFloat32(1e30).IsInfinite()
		assert.False(t, bool(result))
	})

	t.Run("零值返回false", func(t *testing.T) {
		result := goxt.XFloat32(0).IsInfinite()
		assert.False(t, bool(result))
	})

	t.Run("NaN返回false", func(t *testing.T) {
		result := goxt.XFloat32(float32(math.NaN())).IsInfinite()
		assert.False(t, bool(result))
	})
}

// ==================== XFloat64 IsNaN 测试 ====================

func TestXFloat64IsNaN(t *testing.T) {
	t.Run("NaN值返回true", func(t *testing.T) {
		result := goxt.XFloat64(math.NaN()).IsNaN()
		assert.True(t, bool(result))
	})

	t.Run("正常值返回false", func(t *testing.T) {
		result := goxt.XFloat64(3.1415926).IsNaN()
		assert.False(t, bool(result))
	})

	t.Run("零值返回false", func(t *testing.T) {
		result := goxt.XFloat64(0).IsNaN()
		assert.False(t, bool(result))
	})

	t.Run("负数返回false", func(t *testing.T) {
		result := goxt.XFloat64(-2.5).IsNaN()
		assert.False(t, bool(result))
	})
}

// ==================== XFloat64 IsInfinite 测试 ====================

func TestXFloat64IsInfinite(t *testing.T) {
	t.Run("正无穷返回true", func(t *testing.T) {
		result := goxt.XFloat64(math.Inf(1)).IsInfinite()
		assert.True(t, bool(result))
	})

	t.Run("负无穷返回true", func(t *testing.T) {
		result := goxt.XFloat64(math.Inf(-1)).IsInfinite()
		assert.True(t, bool(result))
	})

	t.Run("正常值返回false", func(t *testing.T) {
		result := goxt.XFloat64(1e300).IsInfinite()
		assert.False(t, bool(result))
	})

	t.Run("零值返回false", func(t *testing.T) {
		result := goxt.XFloat64(0).IsInfinite()
		assert.False(t, bool(result))
	})

	t.Run("NaN返回false", func(t *testing.T) {
		result := goxt.XFloat64(math.NaN()).IsInfinite()
		assert.False(t, bool(result))
	})
}

// ==================== XFloat32 Pow 测试 ====================

func TestXFloat32Pow(t *testing.T) {
	t.Run("2的3次方", func(t *testing.T) {
		result := goxt.XFloat32(2).Pow(goxt.XInt(3))
		assert.InDelta(t, float64(8), float64(result), 1e-10)
	})

	t.Run("5的2次方", func(t *testing.T) {
		result := goxt.XFloat32(5).Pow(goxt.XInt(2))
		assert.InDelta(t, float64(25), float64(result), 1e-10)
	})

	t.Run("负数次方", func(t *testing.T) {
		result := goxt.XFloat32(2).Pow(goxt.XInt(-1))
		assert.InDelta(t, float64(0.5), float64(result), 1e-10)
	})

	t.Run("零次方", func(t *testing.T) {
		result := goxt.XFloat32(7).Pow(goxt.XInt(0))
		assert.InDelta(t, float64(1), float64(result), 1e-10)
	})

	t.Run("小数底数", func(t *testing.T) {
		result := goxt.XFloat32(4.0).Pow(goxt.XInt(2))
		assert.InDelta(t, float64(16), float64(result), 1e-10)
	})
}

// ==================== XFloat64 Pow 测试 ====================

func TestXFloat64Pow(t *testing.T) {
	t.Run("2的10次方", func(t *testing.T) {
		result := goxt.XFloat64(2).Pow(goxt.XInt(10))
		assert.InDelta(t, float64(1024), float64(result), 1e-10)
	})

	t.Run("3的4次方", func(t *testing.T) {
		result := goxt.XFloat64(3).Pow(goxt.XInt(4))
		assert.InDelta(t, float64(81), float64(result), 1e-10)
	})

	t.Run("负数次方", func(t *testing.T) {
		result := goxt.XFloat64(10).Pow(goxt.XInt(-2))
		assert.InDelta(t, float64(0.01), float64(result), 1e-10)
	})

	t.Run("零次方", func(t *testing.T) {
		result := goxt.XFloat64(99).Pow(goxt.XInt(0))
		assert.InDelta(t, float64(1), float64(result), 1e-10)
	})

	t.Run("使用XInt64指数", func(t *testing.T) {
		result := goxt.XFloat64(2).Pow(goxt.XInt64(3))
		assert.InDelta(t, float64(8), float64(result), 1e-10)
	})
}

// ==================== XFloat32 Absolute 测试 ====================

func TestXFloat32Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XFloat32(3.14).Absolute()
		assert.Equal(t, goxt.XFloat32(3.14), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XFloat32(-3.14).Absolute()
		assert.Equal(t, goxt.XFloat32(3.14), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XFloat32(0).Absolute()
		assert.Equal(t, goxt.XFloat32(0), result)
	})

	t.Run("负零的绝对值", func(t *testing.T) {
		result := goxt.XFloat32(-0.0).Absolute()
		assert.Equal(t, goxt.XFloat32(0.0), result)
	})
}

// ==================== XFloat64 Absolute 测试 ====================

func TestXFloat64Absolute(t *testing.T) {
	t.Run("正数绝对值", func(t *testing.T) {
		result := goxt.XFloat64(3.1415926).Absolute()
		assert.Equal(t, goxt.XFloat64(3.1415926), result)
	})

	t.Run("负数绝对值", func(t *testing.T) {
		result := goxt.XFloat64(-3.1415926).Absolute()
		assert.Equal(t, goxt.XFloat64(3.1415926), result)
	})

	t.Run("零的绝对值", func(t *testing.T) {
		result := goxt.XFloat64(0).Absolute()
		assert.Equal(t, goxt.XFloat64(0), result)
	})

	t.Run("负零的绝对值", func(t *testing.T) {
		result := goxt.XFloat64(-0.0).Absolute()
		assert.Equal(t, goxt.XFloat64(0.0), result)
	})

	t.Run("大数绝对值", func(t *testing.T) {
		result := goxt.XFloat64(-1e300).Absolute()
		assert.Equal(t, goxt.XFloat64(1e300), result)
	})
}
