package goxt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/xjai/goxt"
)

// ==================== 构造函数测试 ====================

func TestNewXList(t *testing.T) {
	list := NewXList[XInt]()
	assert.NotNil(t, list)
	assert.Equal(t, XInt(0), list.Size())
	assert.True(t, bool(list.IsEmpty()))
}

func TestNewXListWithSize(t *testing.T) {
	t.Run("创建指定大小的列表", func(t *testing.T) {
		list := NewXListWithSize[XString](5)
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := NewXListWithSize[XInt](0)
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})
}

func TestNewXListWithElements(t *testing.T) {
	t.Run("创建包含元素的列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		assert.EqualValues(t, XInt(5), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 5, list[4])
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := NewXListWithElements[XInt]()
		assert.EqualValues(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("创建单个元素列表", func(t *testing.T) {
		list := NewXListWithElements[XString]("hello")
		assert.Equal(t, XInt(1), list.Size())
		assert.EqualValues(t, "hello", list[0])
	})
}

func TestNewXListWithInit(t *testing.T) {
	t.Run("使用初始化函数创建列表", func(t *testing.T) {
		list := NewXListWithInit[XInt](5, func(i XInt) XInt {
			return i * 2
		})
		assert.Equal(t, XInt(5), list.Size())
		assert.EqualValues(t, 0, list[0])
		assert.EqualValues(t, 2, list[1])
		assert.EqualValues(t, 8, list[4])
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := NewXListWithInit(0, func(i XInt) XInt {
			return i
		})
		assert.Equal(t, XInt(0), list.Size())
	})

	t.Run("初始化字符串列表", func(t *testing.T) {
		list := NewXListWithInit(3, func(i XInt) XString {
			return "item" + i.ToString()
		})
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, "item0", list[0])
		assert.EqualValues(t, "item2", list[2])
	})
}

func TestEmptyXList(t *testing.T) {
	list := EmptyXList[XInt]()
	assert.NotNil(t, list)
	assert.Equal(t, XInt(0), list.Size())
	assert.True(t, bool(list.IsEmpty()))
}

// ==================== 基本操作测试 ====================

func TestSize(t *testing.T) {
	t.Run("空列表大小为0", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Equal(t, XInt(0), list.Size())
	})

	t.Run("添加元素后大小增加", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		assert.Equal(t, XInt(3), list.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("空列表返回true", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("非空列表返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1)
		assert.False(t, bool(list.IsEmpty()))
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("空列表返回false", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.False(t, bool(list.IsNotEmpty()))
	})

	t.Run("非空列表返回true", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2)
		assert.True(t, bool(list.IsNotEmpty()))
	})
}

// ==================== 查找和包含测试 ====================

func TestContains(t *testing.T) {
	t.Run("存在元素返回true", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		assert.True(t, bool(list.Contains(3)))
	})

	t.Run("不存在元素返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		assert.False(t, bool(list.Contains(10)))
	})

	t.Run("空列表返回false", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.False(t, bool(list.Contains(1)))
	})

	t.Run("查找第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XString]("a", "b", "c")
		assert.True(t, bool(list.Contains("a")))
	})

	t.Run("查找最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XString]("a", "b", "c")
		assert.True(t, bool(list.Contains("c")))
	})
}

func TestContainsAll(t *testing.T) {
	t.Run("所有元素都存在返回true", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		subset := NewXListWithElements[XInt](2, 3, 4)
		assert.True(t, bool(list.ContainsAll(subset)))
	})

	t.Run("部分元素不存在返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		subset := NewXListWithElements[XInt](2, 6, 4)
		assert.False(t, bool(list.ContainsAll(subset)))
	})

	t.Run("空子集返回true", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		subset := NewXList[XInt]()
		assert.True(t, bool(list.ContainsAll(subset)))
	})

	t.Run("空列表检查非空子集返回false", func(t *testing.T) {
		list := NewXList[XInt]()
		subset := NewXListWithElements[XInt](1, 2)
		assert.False(t, bool(list.ContainsAll(subset)))
	})
}

// ==================== 添加和删除测试 ====================

func TestAdd(t *testing.T) {
	t.Run("添加元素到空列表", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.Add(1)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(1), list.Size())
		assert.EqualValues(t, 1, list[0])
	})

	t.Run("添加多个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2)
		list.Add(3)
		list.Add(4)
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 4, list[3])
	})

	t.Run("添加重复元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2)
		list.Add(2)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 2, list[2])
	})
}

func TestRemove(t *testing.T) {
	t.Run("删除存在的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4)
		result := list.Remove(3)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(3), list.Size())
		assert.False(t, bool(list.Contains(3)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4)
		result := list.Remove(10)
		assert.False(t, bool(result))
		assert.Equal(t, XInt(4), list.Size())
	})

	t.Run("删除第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		list.Remove(1)
		assert.Equal(t, XInt(2), list.Size())
		assert.EqualValues(t, 2, list[0])
	})

	t.Run("删除最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		list.Remove(3)
		assert.Equal(t, XInt(2), list.Size())
		assert.EqualValues(t, 2, list[1])
	})

	t.Run("从空列表删除", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.Remove(1)
		assert.False(t, bool(result))
	})
}

func TestAddAll(t *testing.T) {
	t.Run("添加另一个列表的所有元素", func(t *testing.T) {
		list1 := NewXListWithElements[XInt](1, 2)
		list2 := NewXListWithElements[XInt](3, 4, 5)
		result := list1.AddAll(list2)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(5), list1.Size())
		assert.EqualValues(t, 5, list1[4])
	})

	t.Run("添加空列表", func(t *testing.T) {
		list1 := NewXListWithElements[XInt](1, 2)
		list2 := NewXList[XInt]()
		result := list1.AddAll(list2)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(2), list1.Size())
	})

	t.Run("向空列表添加元素", func(t *testing.T) {
		list1 := NewXList[XInt]()
		list2 := NewXListWithElements[XInt](1, 2, 3)
		list1.AddAll(list2)
		assert.Equal(t, XInt(3), list1.Size())
	})
}

func TestRemoveAll(t *testing.T) {
	t.Run("删除多个存在的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		toRemove := NewXListWithElements[XInt](2, 4)
		list.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), list.Size())
		assert.False(t, bool(list.Contains(2)))
		assert.False(t, bool(list.Contains(4)))
		assert.True(t, bool(list.Contains(1)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRemove := NewXListWithElements[XInt](10, 20)
		list.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), list.Size())
	})

	t.Run("删除所有元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRemove := NewXListWithElements[XInt](1, 2, 3)
		list.RemoveAll(toRemove)
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("删除空集合", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRemove := NewXList[XInt]()
		list.RemoveAll(toRemove)
		assert.Equal(t, XInt(3), list.Size())
	})
}

func TestRetainAll(t *testing.T) {
	t.Run("保留指定元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		toRetain := NewXListWithElements[XInt](2, 4)
		list.RetainAll(toRetain)
		assert.Equal(t, XInt(2), list.Size())
		assert.True(t, bool(list.Contains(2)))
		assert.True(t, bool(list.Contains(4)))
		assert.False(t, bool(list.Contains(1)))
	})

	t.Run("保留所有元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRetain := NewXListWithElements[XInt](1, 2, 3)
		list.RetainAll(toRetain)
		assert.Equal(t, XInt(3), list.Size())
	})

	t.Run("保留空集合清空列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRetain := NewXList[XInt]()
		list.RetainAll(toRetain)
		assert.Equal(t, XInt(0), list.Size())
	})

	t.Run("保留不存在的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toRetain := NewXListWithElements[XInt](10, 20)
		list.RetainAll(toRetain)
		assert.Equal(t, XInt(0), list.Size())
	})
}

func TestClear(t *testing.T) {
	t.Run("清空非空列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		list.Clear()
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("清空空列表", func(t *testing.T) {
		list := NewXList[XInt]()
		list.Clear()
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})
}

// ==================== 索引查找测试 ====================

func TestIndexOf(t *testing.T) {
	t.Run("找到元素返回正确索引", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30, 40, 50)
		assert.Equal(t, XInt(2), list.IndexOf(30))
	})

	t.Run("元素不存在返回-1", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30)
		assert.Equal(t, XInt(-1), list.IndexOf(100))
	})

	t.Run("第一个元素索引为0", func(t *testing.T) {
		list := NewXListWithElements[XString]("a", "b", "c")
		assert.Equal(t, XInt(0), list.IndexOf("a"))
	})

	t.Run("最后一个元素索引", func(t *testing.T) {
		list := NewXListWithElements[XString]("a", "b", "c")
		assert.Equal(t, XInt(2), list.IndexOf("c"))
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Equal(t, XInt(-1), list.IndexOf(1))
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("找到元素返回最后出现的索引", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 2, 4)
		assert.Equal(t, XInt(3), list.LastIndexOf(2))
	})

	t.Run("元素只出现一次", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		assert.Equal(t, XInt(2), list.LastIndexOf(3))
	})

	t.Run("元素不存在返回-1", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		assert.Equal(t, XInt(-1), list.LastIndexOf(10))
	})

	t.Run("第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](5, 4, 3, 2, 1)
		assert.Equal(t, XInt(0), list.LastIndexOf(5))
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Equal(t, XInt(-1), list.LastIndexOf(1))
	})
}

// ==================== 子列表和插入测试 ====================

func TestSubList(t *testing.T) {
	t.Run("获取中间子列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		sub := list.SubList(1, 4)
		assert.Equal(t, XInt(3), sub.Size())
		assert.EqualValues(t, 2, sub[0])
		assert.EqualValues(t, 4, sub[2])
	})

	t.Run("获取从头开始的子列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		sub := list.SubList(0, 2)
		assert.Equal(t, XInt(2), sub.Size())
		assert.EqualValues(t, 1, sub[0])
		assert.EqualValues(t, 2, sub[1])
	})

	t.Run("获取到末尾的子列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		sub := list.SubList(3, 5)
		assert.Equal(t, XInt(2), sub.Size())
		assert.EqualValues(t, 4, sub[0])
		assert.EqualValues(t, 5, sub[1])
	})

	t.Run("获取空子列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		sub := list.SubList(1, 1)
		assert.Equal(t, XInt(0), sub.Size())
	})
}

func TestInsert(t *testing.T) {
	t.Run("在开头插入", func(t *testing.T) {
		list := NewXListWithElements[XInt](2, 3, 4)
		result := list.Insert(0, 1)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 2, list[1])
	})

	t.Run("在中间插入", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 3, 4)
		result := list.Insert(1, 2)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 2, list[1])
		assert.EqualValues(t, 3, list[2])
	})

	t.Run("在末尾插入", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.Insert(3, 4)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 4, list[3])
	})

	t.Run("索引越界返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.Insert(5, 4)
		assert.False(t, bool(result))
		assert.Equal(t, XInt(3), list.Size())
	})

	t.Run("负数索引返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.Insert(-1, 4)
		assert.False(t, bool(result))
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("删除指定索引的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30, 40)
		removed := list.RemoveAt(2)
		assert.NotNil(t, removed)
		assert.EqualValues(t, 30, *removed)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 40, list[2])
	})

	t.Run("删除第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveAt(0)
		assert.NotNil(t, removed)
		assert.EqualValues(t, 1, *removed)
		assert.Equal(t, XInt(2), list.Size())
	})

	t.Run("删除最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveAt(2)
		assert.NotNil(t, removed)
		assert.EqualValues(t, 3, *removed)
		assert.Equal(t, XInt(2), list.Size())
	})

	t.Run("索引越界返回nil", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveAt(5)
		assert.Nil(t, removed)
		assert.Equal(t, XInt(3), list.Size())
	})

	t.Run("负数索引返回nil", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveAt(-1)
		assert.Nil(t, removed)
	})
}

// ==================== 首尾元素操作测试 ====================

func TestFirst(t *testing.T) {
	t.Run("获取第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30)
		first := list.First()
		assert.EqualValues(t, 10, first)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](42)
		first := list.First()
		assert.EqualValues(t, 42, first)
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Panics(t, func() {
			list.First()
		})
	})
}

func TestFirstOrNull(t *testing.T) {
	t.Run("获取第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30)
		first := list.FirstOrNull()
		assert.NotNil(t, first)
		assert.EqualValues(t, 10, *first)
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := NewXList[XInt]()
		first := list.FirstOrNull()
		assert.Nil(t, first)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XString]("hello")
		first := list.FirstOrNull()
		assert.NotNil(t, first)
		assert.EqualValues(t, "hello", *first)
	})
}

func TestLast(t *testing.T) {
	t.Run("获取最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30)
		last := list.Last()
		assert.EqualValues(t, 30, last)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](42)
		last := list.Last()
		assert.EqualValues(t, 42, last)
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Panics(t, func() {
			list.Last()
		})
	})
}

func TestLastOrNull(t *testing.T) {
	t.Run("获取最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](10, 20, 30)
		last := list.LastOrNull()
		assert.NotNil(t, last)
		assert.EqualValues(t, 30, *last)
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := NewXList[XInt]()
		last := list.LastOrNull()
		assert.Nil(t, last)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XString]("world")
		last := list.LastOrNull()
		assert.NotNil(t, last)
		assert.EqualValues(t, "world", *last)
	})
}

func TestAddFirst(t *testing.T) {
	t.Run("向空列表添加第一个元素", func(t *testing.T) {
		list := NewXList[XInt]()
		list.AddFirst(1)
		assert.Equal(t, XInt(1), list.Size())
		assert.EqualValues(t, 1, list[0])
	})

	t.Run("向非空列表头部添加", func(t *testing.T) {
		list := NewXListWithElements[XInt](2, 3, 4)
		list.AddFirst(1)
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 2, list[1])
	})

	t.Run("多次添加到头部", func(t *testing.T) {
		list := NewXList[XInt]()
		list.AddFirst(3)
		list.AddFirst(2)
		list.AddFirst(1)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 2, list[1])
		assert.EqualValues(t, 3, list[2])
	})
}

func TestAddLast(t *testing.T) {
	t.Run("向空列表添加最后一个元素", func(t *testing.T) {
		list := NewXList[XInt]()
		list.AddLast(1)
		assert.Equal(t, XInt(1), list.Size())
		assert.EqualValues(t, 1, list[0])
	})

	t.Run("向非空列表尾部添加", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		list.AddLast(4)
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 4, list[3])
	})

	t.Run("多次添加到尾部", func(t *testing.T) {
		list := NewXList[XInt]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 3, list[2])
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("删除第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveFirst()
		assert.EqualValues(t, 1, removed)
		assert.Equal(t, XInt(2), list.Size())
		assert.EqualValues(t, 2, list[0])
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](42)
		removed := list.RemoveFirst()
		assert.EqualValues(t, 42, removed)
		assert.Equal(t, XInt(0), list.Size())
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Panics(t, func() {
			list.RemoveFirst()
		})
	})
}

func TestRemoveFirstOrNull(t *testing.T) {
	t.Run("删除第一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveFirstOrNull()
		assert.NotNil(t, removed)
		assert.EqualValues(t, 1, *removed)
		assert.Equal(t, XInt(2), list.Size())
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := NewXList[XInt]()
		removed := list.RemoveFirstOrNull()
		assert.Nil(t, removed)
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("删除最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveLast()
		assert.EqualValues(t, 3, removed)
		assert.Equal(t, XInt(2), list.Size())
		assert.EqualValues(t, 2, list[1])
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](42)
		removed := list.RemoveLast()
		assert.EqualValues(t, 42, removed)
		assert.Equal(t, XInt(0), list.Size())
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Panics(t, func() {
			list.RemoveLast()
		})
	})
}

func TestRemoveLastOrNull(t *testing.T) {
	t.Run("删除最后一个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		removed := list.RemoveLastOrNull()
		assert.NotNil(t, removed)
		assert.EqualValues(t, 3, *removed)
		assert.Equal(t, XInt(2), list.Size())
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := NewXList[XInt]()
		removed := list.RemoveLastOrNull()
		assert.Nil(t, removed)
	})
}

// ==================== 批量插入和范围删除测试 ====================

func TestInsertAll(t *testing.T) {
	t.Run("在开头插入多个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](4, 5, 6)
		toInsert := NewXListWithElements[XInt](1, 2, 3)
		result := list.InsertAll(0, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(6), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 3, list[2])
		assert.EqualValues(t, 4, list[3])
	})

	t.Run("在中间插入多个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 5, 6)
		toInsert := NewXListWithElements[XInt](3, 4)
		result := list.InsertAll(2, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(6), list.Size())
		assert.EqualValues(t, 3, list[2])
		assert.EqualValues(t, 4, list[3])
		assert.EqualValues(t, 5, list[4])
	})

	t.Run("在末尾插入多个元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toInsert := NewXListWithElements[XInt](4, 5, 6)
		result := list.InsertAll(3, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(6), list.Size())
		assert.EqualValues(t, 4, list[3])
		assert.EqualValues(t, 6, list[5])
	})

	t.Run("插入空列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toInsert := NewXList[XInt]()
		result := list.InsertAll(1, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, XInt(3), list.Size())
	})

	t.Run("索引越界返回false", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		toInsert := NewXListWithElements[XInt](4, 5)
		result := list.InsertAll(10, toInsert)
		assert.False(t, bool(result))
	})
}

func TestRemoveRange(t *testing.T) {
	t.Run("删除中间范围的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5, 6)
		list.RemoveRange(2, 4)
		assert.Equal(t, XInt(4), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 2, list[1])
		assert.EqualValues(t, 5, list[2])
		assert.EqualValues(t, 6, list[3])
	})

	t.Run("删除开头的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		list.RemoveRange(0, 2)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 3, list[0])
		assert.EqualValues(t, 5, list[2])
	})

	t.Run("删除末尾的元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		list.RemoveRange(3, 5)
		assert.Equal(t, XInt(3), list.Size())
		assert.EqualValues(t, 1, list[0])
		assert.EqualValues(t, 3, list[2])
	})

	t.Run("删除所有元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		list.RemoveRange(0, 3)
		assert.Equal(t, XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("无效范围不操作", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		list.RemoveRange(3, 2) // fromIndex >= toIndex
		assert.Equal(t, XInt(5), list.Size())
	})
}

// ==================== 辅助方法测试 ====================

func TestLastIndex(t *testing.T) {
	t.Run("非空列表返回最后索引", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		assert.Equal(t, XInt(4), list.LastIndex())
	})

	t.Run("单元素列表返回0", func(t *testing.T) {
		list := NewXListWithElements[XInt](42)
		assert.Equal(t, XInt(0), list.LastIndex())
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := NewXList[XInt]()
		assert.Equal(t, XInt(-1), list.LastIndex())
	})
}

func TestForEach(t *testing.T) {
	t.Run("遍历所有元素", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3, 4, 5)
		sum := XInt(0)
		list.ForEach(func(item XInt) {
			sum += item
		})
		assert.EqualValues(t, 15, sum)
	})

	t.Run("修改元素值", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		results := make([]XInt, 0)
		list.ForEach(func(item XInt) {
			results = append(results, item*2)
		})
		assert.EqualValues(t, []XInt{2, 4, 6}, results)
	})

	t.Run("空列表不执行操作", func(t *testing.T) {
		list := NewXList[XInt]()
		called := false
		list.ForEach(func(item XInt) {
			called = true
		})
		assert.False(t, called)
	})

	t.Run("遍历字符串列表", func(t *testing.T) {
		list := NewXListWithElements[XString]("a", "b", "c")
		result := XString("")
		list.ForEach(func(item XString) {
			result += item
		})
		assert.EqualValues(t, "abc", result)
	})
}

func TestIfEmpty(t *testing.T) {
	t.Run("空列表返回默认值", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.IfEmpty(func() XList[XInt] {
			return NewXListWithElements[XInt](1, 2, 3)
		})
		assert.Equal(t, XInt(3), result.Size())
		assert.EqualValues(t, 1, result[0])
	})

	t.Run("非空列表返回原列表", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.IfEmpty(func() XList[XInt] {
			return NewXListWithElements[XInt](4, 5, 6)
		})
		assert.Equal(t, XInt(3), result.Size())
		assert.EqualValues(t, 1, result[0])
		assert.EqualValues(t, 2, result[1])
	})

	t.Run("空列表返回空默认值", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.IfEmpty(func() XList[XInt] {
			return NewXList[XInt]()
		})
		assert.Equal(t, XInt(0), result.Size())
	})
}

func TestXList_JoinToString(t *testing.T) {
	t.Run("列表为空返回空字符串", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.JoinToStringWithAllDefault(",")
		assert.Equal(t, XString(""), result)
	})
	t.Run("列表不为空返回字符串", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.JoinToStringWithAllDefault(",")
		assert.Equal(t, XString("1,2,3"), result)
	})
	t.Run("列表为空返回默认值", func(t *testing.T) {
		list := NewXList[XInt]()
		result := list.JoinToStringWithDefaultTransform(",", "default", "")
		assert.Equal(t, XString("default"), result)
	})
	t.Run("列表不为空返回字符串", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.JoinToStringWithDefaultTransform(",", "default", "")
		assert.Equal(t, XString("default1,2,3"), result)
	})
	t.Run("列表不为空返回转换字符串", func(t *testing.T) {
		list := NewXListWithElements[XInt](1, 2, 3)
		result := list.JoinToString(",", "{", "}", func(item XInt) XString {
			return XString(fmt.Sprintf("item %d", item))
		})
		assert.Equal(t, XString("{item 1,item 2,item 3}"), result)
	})

}
