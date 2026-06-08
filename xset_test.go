package goxt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/xjai/goxt"
)

// ==================== 构造函数测试 ====================

func TestNewXSet(t *testing.T) {
	t.Run("创建空集合", func(t *testing.T) {
		set := NewXSet[XInt]()
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("创建字符串集合", func(t *testing.T) {
		set := NewXSet[XString]()
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size())
	})
}

func TestNewXSetWithSize(t *testing.T) {
	t.Run("创建指定大小的集合", func(t *testing.T) {
		set := NewXSetWithSize[XInt](10)
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size()) // 容量为10，但大小为0
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("创建大小为0的集合", func(t *testing.T) {
		set := NewXSetWithSize[XString](0)
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size())
	})

	t.Run("创建大容量集合", func(t *testing.T) {
		set := NewXSetWithSize[XInt](1000)
		assert.NotNil(t, set)
		assert.True(t, bool(set.IsEmpty()))
	})
}

func TestNewXSetWithElements(t *testing.T) {
	t.Run("创建包含元素的集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		assert.Equal(t, XInt(5), set.Size())
		assert.True(t, bool(set.Contains(3)))
		assert.True(t, bool(set.Contains(5)))
	})

	t.Run("创建空集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt]()
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("创建单个元素集合", func(t *testing.T) {
		set := NewXSetWithElements[XString]("hello")
		assert.Equal(t, XInt(1), set.Size())
		assert.True(t, bool(set.Contains("hello")))
	})

	t.Run("自动去重", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 2, 3, 3, 3)
		assert.Equal(t, XInt(3), set.Size())
		assert.True(t, bool(set.Contains(1)))
		assert.True(t, bool(set.Contains(2)))
		assert.True(t, bool(set.Contains(3)))
	})

	t.Run("创建字符串集合", func(t *testing.T) {
		set := NewXSetWithElements[XString]("a", "b", "c")
		assert.Equal(t, XInt(3), set.Size())
		assert.True(t, bool(set.Contains("b")))
	})
}

func TestNewXSetWithInit(t *testing.T) {
	t.Run("使用初始化函数创建集合", func(t *testing.T) {
		set := NewXSetWithInit(5, func(i XInt) XInt {
			return i * 2
		})
		assert.Equal(t, XInt(5), set.Size())
		assert.True(t, bool(set.Contains(0)))
		assert.True(t, bool(set.Contains(8)))
	})

	t.Run("创建空集合", func(t *testing.T) {
		set := NewXSetWithInit(0, func(i XInt) XInt {
			return i
		})
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("初始化字符串集合", func(t *testing.T) {
		set := NewXSetWithInit(3, func(i XInt) XString {
			return "item" + i.ToString()
		})
		assert.Equal(t, XInt(3), set.Size())
		assert.True(t, bool(set.Contains("item0")))
		assert.True(t, bool(set.Contains("item2")))
	})

	t.Run("生成连续数字", func(t *testing.T) {
		set := NewXSetWithInit(5, func(i XInt) XInt {
			return i + 1
		})
		assert.Equal(t, XInt(5), set.Size())
		for i := 1; i <= 5; i++ {
			assert.True(t, bool(set.Contains(XInt(i))))
		}
	})
}

func TestEmptyXSet(t *testing.T) {
	t.Run("创建空集合", func(t *testing.T) {
		set := EmptyXSet[XInt]()
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("创建空字符串集合", func(t *testing.T) {
		set := EmptyXSet[XString]()
		assert.NotNil(t, set)
		assert.Equal(t, XInt(0), set.Size())
	})
}

// ==================== 基本操作测试 ====================

func TestXSetSize(t *testing.T) {
	t.Run("空集合大小为0", func(t *testing.T) {
		set := NewXSet[XInt]()
		assert.Equal(t, XInt(0), set.Size())
	})

	t.Run("添加元素后大小增加", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		assert.Equal(t, XInt(3), set.Size())
	})

	t.Run("重复添加不增加大小", func(t *testing.T) {
		set := NewXSet[XInt]()
		set.Add(1)
		set.Add(1)
		set.Add(1)
		assert.Equal(t, XInt(1), set.Size())
	})
}

func TestXSetIsEmpty(t *testing.T) {
	t.Run("空集合返回true", func(t *testing.T) {
		set := NewXSet[XInt]()
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("非空集合返回false", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1)
		assert.False(t, bool(set.IsEmpty()))
	})

	t.Run("清空后返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		set.Clear()
		assert.True(t, bool(set.IsEmpty()))
	})
}

func TestXSetIsNotEmpty(t *testing.T) {
	t.Run("空集合返回false", func(t *testing.T) {
		set := NewXSet[XInt]()
		assert.False(t, bool(set.IsNotEmpty()))
	})

	t.Run("非空集合返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		assert.True(t, bool(set.IsNotEmpty()))
	})

	t.Run("单元素集合返回true", func(t *testing.T) {
		set := NewXSetWithElements[XString]("test")
		assert.True(t, bool(set.IsNotEmpty()))
	})
}

// ==================== 查找和包含测试 ====================

func TestXSetContains(t *testing.T) {
	t.Run("存在元素返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		assert.True(t, bool(set.Contains(3)))
	})

	t.Run("不存在元素返回false", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		assert.False(t, bool(set.Contains(10)))
	})

	t.Run("空集合返回false", func(t *testing.T) {
		set := NewXSet[XInt]()
		assert.False(t, bool(set.Contains(1)))
	})

	t.Run("查找第一个添加的元素", func(t *testing.T) {
		set := NewXSetWithElements[XString]("a", "b", "c")
		assert.True(t, bool(set.Contains("a")))
	})

	t.Run("查找最后一个添加的元素", func(t *testing.T) {
		set := NewXSetWithElements[XString]("x", "y", "z")
		assert.True(t, bool(set.Contains("z")))
	})
}

func TestXSetContainsAll(t *testing.T) {
	t.Run("所有元素都存在返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		subset := NewXListWithElements[XInt](2, 3, 4)
		assert.True(t, bool(set.ContainsAll(subset)))
	})

	t.Run("部分元素不存在返回false", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		subset := NewXListWithElements[XInt](2, 6, 4)
		assert.False(t, bool(set.ContainsAll(subset)))
	})

	t.Run("空子集返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		subset := NewXList[XInt]()
		assert.True(t, bool(set.ContainsAll(subset)))
	})

	t.Run("空集合检查非空子集返回false", func(t *testing.T) {
		set := NewXSet[XInt]()
		subset := NewXListWithElements[XInt](1, 2)
		assert.False(t, bool(set.ContainsAll(subset)))
	})

	t.Run("检查单个元素", func(t *testing.T) {
		set := NewXSetWithElements[XString]("a", "b", "c")
		subset := NewXListWithElements[XString]("b")
		assert.True(t, bool(set.ContainsAll(subset)))
	})
}

// ==================== 添加和删除测试 ====================

func TestXSetAdd(t *testing.T) {
	t.Run("添加元素到空集合", func(t *testing.T) {
		set := NewXSet[XInt]()
		result := set.Add(1)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(1), set.Size())
		assert.True(t, bool(set.Contains(1)))
	})

	t.Run("添加多个元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		set.Add(3)
		set.Add(4)
		assert.Equal(t, XInt(4), set.Size())
		assert.True(t, bool(set.Contains(4)))
	})

	t.Run("添加重复元素返回true", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		result := set.Add(2)
		assert.False(t, bool(result))
		assert.Equal(t, XInt(2), set.Size()) // 大小不变
	})

	t.Run("添加字符串元素", func(t *testing.T) {
		set := NewXSet[XString]()
		set.Add("hello")
		assert.Equal(t, XInt(1), set.Size())
		assert.True(t, bool(set.Contains("hello")))
	})
}

func TestXSetRemove(t *testing.T) {
	t.Run("删除存在的元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4)
		result := set.Remove(3)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(3), set.Size())
		assert.False(t, bool(set.Contains(3)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4)
		result := set.Remove(10)
		assert.False(t, bool(result))
		assert.Equal(t, XInt(4), set.Size())
	})

	t.Run("删除第一个元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		set.Remove(1)
		assert.Equal(t, XInt(2), set.Size())
		assert.False(t, bool(set.Contains(1)))
	})

	t.Run("从空集合删除", func(t *testing.T) {
		set := NewXSet[XInt]()
		result := set.Remove(1)
		assert.False(t, bool(result))
	})

	t.Run("删除字符串元素", func(t *testing.T) {
		set := NewXSetWithElements[XString]("a", "b", "c")
		result := set.Remove("b")
		assert.True(t, bool(result))
		assert.Equal(t, XInt(2), set.Size())
	})
}

func TestXSetAddAll(t *testing.T) {
	t.Run("添加另一个列表的所有元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		list := NewXListWithElements[XInt](3, 4, 5)
		result := set.AddAll(list)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(5), set.Size())
		assert.True(t, bool(set.Contains(5)))
	})

	t.Run("添加空列表", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		list := NewXList[XInt]()
		result := set.AddAll(list)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(2), set.Size())
	})

	t.Run("向空集合添加元素", func(t *testing.T) {
		set := NewXSet[XInt]()
		list := NewXListWithElements[XInt](1, 2, 3)
		set.AddAll(list)
		assert.Equal(t, XInt(3), set.Size())
	})

	t.Run("添加包含重复元素的列表", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2)
		list := NewXListWithElements[XInt](2, 3, 3, 4)
		set.AddAll(list)
		assert.Equal(t, XInt(4), set.Size()) // 2已存在，3重复
	})
}

func TestXSetRemoveAll(t *testing.T) {
	t.Run("删除多个存在的元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		toRemove := NewXListWithElements[XInt](2, 4)
		set.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), set.Size())
		assert.False(t, bool(set.Contains(2)))
		assert.False(t, bool(set.Contains(4)))
		assert.True(t, bool(set.Contains(1)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRemove := NewXListWithElements[XInt](10, 20)
		set.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), set.Size())
	})

	t.Run("删除所有元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRemove := NewXListWithElements[XInt](1, 2, 3)
		set.RemoveAll(toRemove)
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("删除空集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRemove := NewXList[XInt]()
		set.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), set.Size())
	})
}

func TestXSetRetainAll(t *testing.T) {
	t.Run("保留指定元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		toRetain := NewXListWithElements[XInt](2, 4)
		result := set.RetainAll(toRetain)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(2), set.Size())
		assert.True(t, bool(set.Contains(2)))
		assert.True(t, bool(set.Contains(4)))
		assert.False(t, bool(set.Contains(1)))
	})

	t.Run("保留所有元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRetain := NewXListWithElements[XInt](1, 2, 3)
		result := set.RetainAll(toRetain)
		assert.False(t, bool(result)) // 没有变化
		assert.Equal(t, XInt(3), set.Size())
	})

	t.Run("保留空集合清空集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRetain := NewXList[XInt]()
		result := set.RetainAll(toRetain)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(0), set.Size())
	})

	t.Run("保留不存在的元素", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		toRetain := NewXListWithElements[XInt](10, 20)
		result := set.RetainAll(toRetain)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(0), set.Size())
	})

	t.Run("空集合调用RetainAll", func(t *testing.T) {
		set := NewXSet[XInt]()
		toRetain := NewXListWithElements[XInt](1, 2, 3)
		result := set.RetainAll(toRetain)
		assert.False(t, bool(result))
		assert.Equal(t, XInt(0), set.Size())
	})
}

func TestXSetClear(t *testing.T) {
	t.Run("清空非空集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3, 4, 5)
		set.Clear()
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("清空空集合", func(t *testing.T) {
		set := NewXSet[XInt]()
		set.Clear()
		assert.Equal(t, XInt(0), set.Size())
		assert.True(t, bool(set.IsEmpty()))
	})

	t.Run("清空后可以继续添加", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		set.Clear()
		set.Add(10)
		assert.Equal(t, XInt(1), set.Size())
		assert.True(t, bool(set.Contains(10)))
	})
}

// ==================== 辅助方法测试 ====================

func TestXSetIfEmpty(t *testing.T) {
	t.Run("空集合返回默认值", func(t *testing.T) {
		set := NewXSet[XInt]()
		result := set.IfEmpty(func() XSet[XInt] {
			return NewXSetWithElements[XInt](1, 2, 3)
		})
		assert.Equal(t, XInt(3), result.Size())
		assert.True(t, bool(result.Contains(1)))
		assert.True(t, bool(result.Contains(3)))
	})

	t.Run("非空集合返回原集合", func(t *testing.T) {
		set := NewXSetWithElements[XInt](1, 2, 3)
		result := set.IfEmpty(func() XSet[XInt] {
			return NewXSetWithElements[XInt](4, 5, 6)
		})
		assert.Equal(t, XInt(3), result.Size())
		assert.True(t, bool(result.Contains(1)))
		assert.True(t, bool(result.Contains(2)))
		assert.False(t, bool(result.Contains(4)))
	})

	t.Run("空集合返回空默认值", func(t *testing.T) {
		set := NewXSet[XInt]()
		result := set.IfEmpty(func() XSet[XInt] {
			return NewXSet[XInt]()
		})
		assert.Equal(t, XInt(0), result.Size())
	})

	t.Run("空集合返回非空默认值", func(t *testing.T) {
		set := NewXSet[XString]()
		result := set.IfEmpty(func() XSet[XString] {
			return NewXSetWithElements[XString]("default")
		})
		assert.Equal(t, XInt(1), result.Size())
		assert.True(t, bool(result.Contains("default")))
	})
}