package goxt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

// ==================== NewXMap 测试 ====================

func TestNewXMap(t *testing.T) {
	t.Run("创建空map", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.NotNil(t, m)
		assert.Equal(t, goxt.XInt(0), m.Size())
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("创建int-key map", func(t *testing.T) {
		m := goxt.NewXMap[int, string]()
		assert.NotNil(t, m)
		assert.True(t, bool(m.IsEmpty()))
	})
}

// ==================== NewXMapWithSize 测试 ====================

func TestNewXMapWithSize(t *testing.T) {
	t.Run("创建指定大小的空map", func(t *testing.T) {
		m := goxt.NewXMapWithSize[string, int](10)
		assert.NotNil(t, m)
		assert.Equal(t, goxt.XInt(0), m.Size())
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("创建大小为0的map", func(t *testing.T) {
		m := goxt.NewXMapWithSize[int, int](0)
		assert.NotNil(t, m)
		assert.True(t, bool(m.IsEmpty()))
	})
}

// ==================== NewXMapWithElements 测试 ====================

func TestNewXMapWithElements(t *testing.T) {
	t.Run("从多个元素创建map", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
			goxt.XPair[string, int]{First: "c", Second: 3},
		)
		assert.Equal(t, goxt.XInt(3), m.Size())
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
		assert.Equal(t, 3, m["c"])
	})

	t.Run("单个元素", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, string]{First: "key", Second: "value"})
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, "value", m["key"])
	})

	t.Run("无元素", func(t *testing.T) {
		m := goxt.NewXMapWithElements[string, int]()
		assert.Equal(t, goxt.XInt(0), m.Size())
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("重复key覆盖", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "a", Second: 2},
		)
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, 2, m["a"])
	})
}

// ==================== NewXMapWithInit 测试 ====================

func TestNewXMapWithInit(t *testing.T) {
	t.Run("使用init函数创建map", func(t *testing.T) {
		m := goxt.NewXMapWithInit(3, func(i goxt.XInt) goxt.XPair[int, string] {
			return goxt.XPair[int, string]{First: int(i), Second: "v" + string(rune('0'+i))}
		})
		assert.Equal(t, goxt.XInt(3), m.Size())
		assert.Equal(t, "v0", m[0])
		assert.Equal(t, "v1", m[1])
		assert.Equal(t, "v2", m[2])
	})

	t.Run("init大小为0", func(t *testing.T) {
		m := goxt.NewXMapWithInit(0, func(i goxt.XInt) goxt.XPair[int, string] {
			return goxt.XPair[int, string]{First: 0, Second: ""}
		})
		assert.Equal(t, goxt.XInt(0), m.Size())
	})

	t.Run("init单个元素", func(t *testing.T) {
		m := goxt.NewXMapWithInit(1, func(i goxt.XInt) goxt.XPair[string, int] {
			return goxt.XPair[string, int]{First: "only", Second: 42}
		})
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, 42, m["only"])
	})
}

// ==================== EmptyXMap 测试 ====================

func TestEmptyXMap(t *testing.T) {
	t.Run("返回空map", func(t *testing.T) {
		m := goxt.EmptyXMap[string, int]()
		assert.NotNil(t, m)
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("返回的map可修改", func(t *testing.T) {
		m := goxt.EmptyXMap[string, int]()
		m["key"] = 1
		assert.Equal(t, goxt.XInt(1), m.Size())
	})
}

// ==================== Size 测试 ====================

func TestXMapSize(t *testing.T) {
	t.Run("空map大小", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.Equal(t, goxt.XInt(0), m.Size())
	})

	t.Run("非空map大小", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		assert.Equal(t, goxt.XInt(2), m.Size())
	})

	t.Run("添加后大小变化", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m["a"] = 1
		assert.Equal(t, goxt.XInt(1), m.Size())
		m["b"] = 2
		assert.Equal(t, goxt.XInt(2), m.Size())
	})
}

// ==================== IsEmpty 测试 ====================

func TestXMapIsEmpty(t *testing.T) {
	t.Run("空map返回true", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("非空map返回false", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m["a"] = 1
		assert.False(t, bool(m.IsEmpty()))
	})
}

// ==================== IsNotEmpty 测试 ====================

func TestXMapIsNotEmpty(t *testing.T) {
	t.Run("空map返回false", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.False(t, bool(m.IsNotEmpty()))
	})

	t.Run("非空map返回true", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m["a"] = 1
		assert.True(t, bool(m.IsNotEmpty()))
	})
}

// ==================== ContainsKey 测试 ====================

func TestXMapContainsKey(t *testing.T) {
	t.Run("包含的key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		assert.True(t, bool(m.ContainsKey("a")))
	})

	t.Run("不包含的key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		assert.False(t, bool(m.ContainsKey("b")))
	})

	t.Run("空map不包含任何key", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.False(t, bool(m.ContainsKey("a")))
	})
}

// ==================== ContainsValue 测试 ====================

func TestXMapContainsValue(t *testing.T) {
	t.Run("包含的value", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		assert.True(t, bool(m.ContainsValue(1)))
	})

	t.Run("不包含的value", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		assert.False(t, bool(m.ContainsValue(2)))
	})

	t.Run("空map不包含任何value", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		assert.False(t, bool(m.ContainsValue(1)))
	})

	t.Run("多value查找", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
			goxt.XPair[string, int]{First: "c", Second: 3},
		)
		assert.True(t, bool(m.ContainsValue(2)))
		assert.False(t, bool(m.ContainsValue(4)))
	})
}

// ==================== Keys 测试 ====================

func TestXMapKeys(t *testing.T) {
	t.Run("获取所有key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		keys := m.Keys()
		assert.Equal(t, goxt.XInt(2), keys.Size())
		assert.True(t, bool(keys.Contains("a")))
		assert.True(t, bool(keys.Contains("b")))
	})

	t.Run("空map的keys", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		keys := m.Keys()
		assert.Equal(t, goxt.XInt(0), keys.Size())
	})

	t.Run("单个元素的keys", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[int, string]{First: 42, Second: "answer"})
		keys := m.Keys()
		assert.Equal(t, goxt.XInt(1), keys.Size())
		assert.True(t, bool(keys.Contains(42)))
	})
}

// ==================== Values 测试 ====================

func TestXMapValues(t *testing.T) {
	t.Run("获取所有value", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		values := m.Values()
		assert.True(t, bool(values.Contains(1)))
		assert.True(t, bool(values.Contains(2)))
	})

	t.Run("空map的values", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		values := m.Values()
		assert.Equal(t, goxt.XInt(0), values.Size())
	})

	t.Run("单个元素的values", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[int, string]{First: 1, Second: "hello"})
		values := m.Values()
		assert.True(t, bool(values.Contains("hello")))
	})
}

// ==================== Entries 测试 ====================

func TestXMapEntries(t *testing.T) {
	t.Run("获取所有entries", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		entries := m.Entries()
		assert.Equal(t, goxt.XInt(2), entries.Size())
		assert.True(t, bool(entries.Contains(goxt.XPair[string, int]{First: "a", Second: 1})))
		assert.True(t, bool(entries.Contains(goxt.XPair[string, int]{First: "b", Second: 2})))
	})

	t.Run("空map的entries", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		entries := m.Entries()
		assert.Equal(t, goxt.XInt(0), entries.Size())
	})

	t.Run("单个元素的entries", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[int, string]{First: 1, Second: "v"})
		entries := m.Entries()
		assert.Equal(t, goxt.XInt(1), entries.Size())
		assert.True(t, bool(entries.Contains(goxt.XPair[int, string]{First: 1, Second: "v"})))
	})
}

// ==================== Remove 测试 ====================

func TestXMapRemove(t *testing.T) {
	t.Run("移除存在的key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		v := m.Remove("a")
		assert.NotNil(t, v)
		assert.Equal(t, 1, *v)
		assert.Equal(t, goxt.XInt(0), m.Size())
	})

	t.Run("移除不存在的key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		v := m.Remove("b")
		assert.Nil(t, v)
		assert.Equal(t, goxt.XInt(1), m.Size())
	})

	t.Run("从空map移除", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		v := m.Remove("a")
		assert.Nil(t, v)
	})

	t.Run("移除后key不存在", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		m.Remove("a")
		assert.False(t, bool(m.ContainsKey("a")))
	})
}

// ==================== PutAll 测试 ====================

func TestXMapPutAll(t *testing.T) {
	t.Run("合并两个map", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		from := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "b", Second: 2})
		m.PutAll(from)
		assert.Equal(t, goxt.XInt(2), m.Size())
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
	})

	t.Run("合并覆盖已有key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		from := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 99})
		m.PutAll(from)
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, 99, m["a"])
	})

	t.Run("合并空map", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		from := goxt.NewXMap[string, int]()
		m.PutAll(from)
		assert.Equal(t, goxt.XInt(1), m.Size())
	})

	t.Run("向空map合并", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		from := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		m.PutAll(from)
		assert.Equal(t, goxt.XInt(2), m.Size())
	})
}

// ==================== Clear 测试 ====================

func TestXMapClear(t *testing.T) {
	t.Run("清空非空map", func(t *testing.T) {
		m := goxt.NewXMapWithElements(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
		)
		m.Clear()
		assert.Equal(t, goxt.XInt(0), m.Size())
		assert.True(t, bool(m.IsEmpty()))
	})

	t.Run("清空空map", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m.Clear()
		assert.Equal(t, goxt.XInt(0), m.Size())
	})
}

// ==================== IfEmpty 测试 ====================

func TestXMapIfEmpty(t *testing.T) {
	t.Run("空map返回默认值", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		result := m.IfEmpty(func() goxt.XMap[string, int] {
			return goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "default", Second: 0})
		})
		assert.Equal(t, goxt.XInt(1), result.Size())
		assert.Equal(t, 0, result["default"])
	})

	t.Run("非空map返回自身", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		result := m.IfEmpty(func() goxt.XMap[string, int] {
			return goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "default", Second: 0})
		})
		assert.Equal(t, goxt.XInt(1), result.Size())
		assert.Equal(t, 1, result["a"])
	})

	t.Run("非空map不调用默认函数", func(t *testing.T) {
		called := false
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		_ = m.IfEmpty(func() goxt.XMap[string, int] {
			called = true
			return goxt.NewXMap[string, int]()
		})
		assert.False(t, called)
	})
}

// ==================== PutAllPairs 测试 ====================

func TestXMapPutAllPairs(t *testing.T) {
	t.Run("添加多个pair", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m.PutAllPairs(
			goxt.XPair[string, int]{First: "a", Second: 1},
			goxt.XPair[string, int]{First: "b", Second: 2},
			goxt.XPair[string, int]{First: "c", Second: 3},
		)
		assert.Equal(t, goxt.XInt(3), m.Size())
		assert.Equal(t, 1, m["a"])
		assert.Equal(t, 2, m["b"])
		assert.Equal(t, 3, m["c"])
	})

	t.Run("添加单个pair", func(t *testing.T) {
		m := goxt.NewXMap[string, int]()
		m.PutAllPairs(goxt.XPair[string, int]{First: "a", Second: 1})
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, 1, m["a"])
	})

	t.Run("无pair", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		m.PutAllPairs()
		assert.Equal(t, goxt.XInt(1), m.Size())
	})

	t.Run("覆盖已有key", func(t *testing.T) {
		m := goxt.NewXMapWithElements(goxt.XPair[string, int]{First: "a", Second: 1})
		m.PutAllPairs(goxt.XPair[string, int]{First: "a", Second: 99})
		assert.Equal(t, goxt.XInt(1), m.Size())
		assert.Equal(t, 99, m["a"])
	})
}
