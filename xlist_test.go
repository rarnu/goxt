package goxt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

// ==================== 构造函数测试 ====================

func TestNewXList(t *testing.T) {
	list := goxt.NewXList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, goxt.XInt(0), list.Size())
	assert.True(t, bool(list.IsEmpty()))
}

func TestNewXListWithSize(t *testing.T) {
	t.Run("创建指定大小的列表", func(t *testing.T) {
		list := goxt.NewXListWithSize[string](5)
		assert.Equal(t, goxt.XInt(5), list.Size())
		assert.False(t, bool(list.IsEmpty()))
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := goxt.NewXListWithSize[int](0)
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})
}

func TestNewXListWithElements(t *testing.T) {
	t.Run("创建包含元素的列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		assert.Equal(t, goxt.XInt(5), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 5, list[4])
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := goxt.NewXListWithElements[int]()
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("创建单个元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements("hello")
		assert.Equal(t, goxt.XInt(1), list.Size())
		assert.Equal(t, "hello", list[0])
	})
}

func TestNewXListWithInit(t *testing.T) {
	t.Run("使用初始化函数创建列表", func(t *testing.T) {
		list := goxt.NewXListWithInit(5, func(i goxt.XInt) int {
			return int(i) * 2
		})
		assert.Equal(t, goxt.XInt(5), list.Size())
		assert.Equal(t, 0, list[0])
		assert.Equal(t, 2, list[1])
		assert.Equal(t, 8, list[4])
	})

	t.Run("创建空列表", func(t *testing.T) {
		list := goxt.NewXListWithInit(0, func(i goxt.XInt) int {
			return int(i)
		})
		assert.Equal(t, goxt.XInt(0), list.Size())
	})

	t.Run("初始化字符串列表", func(t *testing.T) {
		list := goxt.NewXListWithInit(3, func(i goxt.XInt) goxt.XString {
			return "item" + i.ToString()
		})
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.EqualValues(t, "item0", list[0])
		assert.EqualValues(t, "item2", list[2])
	})
}

func TestEmptyXList(t *testing.T) {
	list := goxt.EmptyXList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, goxt.XInt(0), list.Size())
	assert.True(t, bool(list.IsEmpty()))
}

// ==================== 基本操作测试 ====================

func TestSize(t *testing.T) {
	t.Run("空列表大小为0", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Equal(t, goxt.XInt(0), list.Size())
	})

	t.Run("添加元素后大小增加", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		assert.Equal(t, goxt.XInt(3), list.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("空列表返回true", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("非空列表返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1)
		assert.False(t, bool(list.IsEmpty()))
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("空列表返回false", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.False(t, bool(list.IsNotEmpty()))
	})

	t.Run("非空列表返回true", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2)
		assert.True(t, bool(list.IsNotEmpty()))
	})
}

// ==================== 查找和包含测试 ====================

func TestContains(t *testing.T) {
	t.Run("存在元素返回true", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		assert.True(t, bool(list.Contains(3)))
	})

	t.Run("不存在元素返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		assert.False(t, bool(list.Contains(10)))
	})

	t.Run("空列表返回false", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.False(t, bool(list.Contains(1)))
	})

	t.Run("查找第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements("a", "b", "c")
		assert.True(t, bool(list.Contains("a")))
	})

	t.Run("查找最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements("a", "b", "c")
		assert.True(t, bool(list.Contains("c")))
	})
}

func TestContainsAll(t *testing.T) {
	t.Run("所有元素都存在返回true", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		subset := goxt.NewXListWithElements(2, 3, 4)
		assert.True(t, bool(list.ContainsAll(subset)))
	})

	t.Run("部分元素不存在返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		subset := goxt.NewXListWithElements(2, 6, 4)
		assert.False(t, bool(list.ContainsAll(subset)))
	})

	t.Run("空子集返回true", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		subset := goxt.NewXList[int]()
		assert.True(t, bool(list.ContainsAll(subset)))
	})

	t.Run("空列表检查非空子集返回false", func(t *testing.T) {
		list := goxt.NewXList[int]()
		subset := goxt.NewXListWithElements(1, 2)
		assert.False(t, bool(list.ContainsAll(subset)))
	})
}

// ==================== 添加和删除测试 ====================

func TestAdd(t *testing.T) {
	t.Run("添加元素到空列表", func(t *testing.T) {
		list := goxt.NewXList[int]()
		result := list.Add(1)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(1), list.Size())
		assert.Equal(t, 1, list[0])
	})

	t.Run("添加多个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2)
		list.Add(3)
		list.Add(4)
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 4, list[3])
	})

	t.Run("添加重复元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2)
		list.Add(2)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 2, list[2])
	})
}

func TestRemove(t *testing.T) {
	t.Run("删除存在的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4)
		result := list.Remove(3)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.False(t, bool(list.Contains(3)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4)
		result := list.Remove(10)
		assert.False(t, bool(result))
		assert.Equal(t, goxt.XInt(4), list.Size())
	})

	t.Run("删除第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		list.Remove(1)
		assert.Equal(t, goxt.XInt(2), list.Size())
		assert.Equal(t, 2, list[0])
	})

	t.Run("删除最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		list.Remove(3)
		assert.Equal(t, goxt.XInt(2), list.Size())
		assert.Equal(t, 2, list[1])
	})

	t.Run("从空列表删除", func(t *testing.T) {
		list := goxt.NewXList[int]()
		result := list.Remove(1)
		assert.False(t, bool(result))
	})
}

func TestAddAll(t *testing.T) {
	t.Run("添加另一个列表的所有元素", func(t *testing.T) {
		list1 := goxt.NewXListWithElements(1, 2)
		list2 := goxt.NewXListWithElements(3, 4, 5)
		result := list1.AddAll(list2)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(5), list1.Size())
		assert.Equal(t, 5, list1[4])
	})

	t.Run("添加空列表", func(t *testing.T) {
		list1 := goxt.NewXListWithElements(1, 2)
		list2 := goxt.NewXList[int]()
		result := list1.AddAll(list2)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(2), list1.Size())
	})

	t.Run("向空列表添加元素", func(t *testing.T) {
		list1 := goxt.NewXList[int]()
		list2 := goxt.NewXListWithElements(1, 2, 3)
		list1.AddAll(list2)
		assert.Equal(t, goxt.XInt(3), list1.Size())
	})
}

func TestRemoveAll(t *testing.T) {
	t.Run("删除多个存在的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		toRemove := goxt.NewXListWithElements(2, 4)
		list.RemoveAll(toRemove)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.False(t, bool(list.Contains(2)))
		assert.False(t, bool(list.Contains(4)))
		assert.True(t, bool(list.Contains(1)))
	})

	t.Run("删除不存在的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRemove := goxt.NewXListWithElements(10, 20)
		list.RemoveAll(toRemove)
		assert.Equal(t, goxt.XInt(3), list.Size())
	})

	t.Run("删除所有元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRemove := goxt.NewXListWithElements(1, 2, 3)
		list.RemoveAll(toRemove)
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("删除空集合", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRemove := goxt.NewXList[int]()
		list.RemoveAll(toRemove)
		assert.Equal(t, goxt.XInt(3), list.Size())
	})
}

func TestRetainAll(t *testing.T) {
	t.Run("保留指定元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		toRetain := goxt.NewXListWithElements(2, 4)
		list.RetainAll(toRetain)
		assert.Equal(t, goxt.XInt(2), list.Size())
		assert.True(t, bool(list.Contains(2)))
		assert.True(t, bool(list.Contains(4)))
		assert.False(t, bool(list.Contains(1)))
	})

	t.Run("保留所有元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRetain := goxt.NewXListWithElements(1, 2, 3)
		list.RetainAll(toRetain)
		assert.Equal(t, goxt.XInt(3), list.Size())
	})

	t.Run("保留空集合清空列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRetain := goxt.NewXList[int]()
		list.RetainAll(toRetain)
		assert.Equal(t, goxt.XInt(0), list.Size())
	})

	t.Run("保留不存在的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toRetain := goxt.NewXListWithElements(10, 20)
		list.RetainAll(toRetain)
		assert.Equal(t, goxt.XInt(0), list.Size())
	})
}

func TestClear(t *testing.T) {
	t.Run("清空非空列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		list.Clear()
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("清空空列表", func(t *testing.T) {
		list := goxt.NewXList[int]()
		list.Clear()
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})
}

// ==================== 索引查找测试 ====================

func TestIndexOf(t *testing.T) {
	t.Run("找到元素返回正确索引", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30, 40, 50)
		assert.Equal(t, goxt.XInt(2), list.IndexOf(30))
	})

	t.Run("元素不存在返回-1", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30)
		assert.Equal(t, goxt.XInt(-1), list.IndexOf(100))
	})

	t.Run("第一个元素索引为0", func(t *testing.T) {
		list := goxt.NewXListWithElements("a", "b", "c")
		assert.Equal(t, goxt.XInt(0), list.IndexOf("a"))
	})

	t.Run("最后一个元素索引", func(t *testing.T) {
		list := goxt.NewXListWithElements("a", "b", "c")
		assert.Equal(t, goxt.XInt(2), list.IndexOf("c"))
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Equal(t, goxt.XInt(-1), list.IndexOf(1))
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("找到元素返回最后出现的索引", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 2, 4)
		assert.Equal(t, goxt.XInt(3), list.LastIndexOf(2))
	})

	t.Run("元素只出现一次", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		assert.Equal(t, goxt.XInt(2), list.LastIndexOf(3))
	})

	t.Run("元素不存在返回-1", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		assert.Equal(t, goxt.XInt(-1), list.LastIndexOf(10))
	})

	t.Run("第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(5, 4, 3, 2, 1)
		assert.Equal(t, goxt.XInt(0), list.LastIndexOf(5))
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Equal(t, goxt.XInt(-1), list.LastIndexOf(1))
	})
}

// ==================== 子列表和插入测试 ====================

func TestSubList(t *testing.T) {
	t.Run("获取中间子列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		sub := list.SubList(1, 4)
		assert.Equal(t, goxt.XInt(3), sub.Size())
		assert.Equal(t, 2, sub[0])
		assert.Equal(t, 4, sub[2])
	})

	t.Run("获取从头开始的子列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		sub := list.SubList(0, 2)
		assert.Equal(t, goxt.XInt(2), sub.Size())
		assert.Equal(t, 1, sub[0])
		assert.Equal(t, 2, sub[1])
	})

	t.Run("获取到末尾的子列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		sub := list.SubList(3, 5)
		assert.Equal(t, goxt.XInt(2), sub.Size())
		assert.Equal(t, 4, sub[0])
		assert.Equal(t, 5, sub[1])
	})

	t.Run("获取空子列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		sub := list.SubList(1, 1)
		assert.Equal(t, goxt.XInt(0), sub.Size())
	})
}

func TestInsert(t *testing.T) {
	t.Run("在开头插入", func(t *testing.T) {
		list := goxt.NewXListWithElements(2, 3, 4)
		result := list.Insert(0, 1)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 2, list[1])
	})

	t.Run("在中间插入", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 3, 4)
		result := list.Insert(1, 2)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 2, list[1])
		assert.Equal(t, 3, list[2])
	})

	t.Run("在末尾插入", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		result := list.Insert(3, 4)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 4, list[3])
	})

	t.Run("索引越界返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		result := list.Insert(5, 4)
		assert.False(t, bool(result))
		assert.Equal(t, goxt.XInt(3), list.Size())
	})

	t.Run("负数索引返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		result := list.Insert(-1, 4)
		assert.False(t, bool(result))
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("删除指定索引的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30, 40)
		removed := list.RemoveAt(2)
		assert.NotNil(t, removed)
		assert.Equal(t, 30, *removed)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 40, list[2])
	})

	t.Run("删除第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveAt(0)
		assert.NotNil(t, removed)
		assert.Equal(t, 1, *removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
	})

	t.Run("删除最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveAt(2)
		assert.NotNil(t, removed)
		assert.Equal(t, 3, *removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
	})

	t.Run("索引越界返回nil", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveAt(5)
		assert.Nil(t, removed)
		assert.Equal(t, goxt.XInt(3), list.Size())
	})

	t.Run("负数索引返回nil", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveAt(-1)
		assert.Nil(t, removed)
	})
}

// ==================== 首尾元素操作测试 ====================

func TestFirst(t *testing.T) {
	t.Run("获取第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30)
		first := list.First()
		assert.Equal(t, 10, first)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(42)
		first := list.First()
		assert.Equal(t, 42, first)
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Panics(t, func() {
			list.First()
		})
	})
}

func TestFirstOrNull(t *testing.T) {
	t.Run("获取第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30)
		first := list.FirstOrNull()
		assert.NotNil(t, first)
		assert.Equal(t, 10, *first)
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := goxt.NewXList[int]()
		first := list.FirstOrNull()
		assert.Nil(t, first)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements("hello")
		first := list.FirstOrNull()
		assert.NotNil(t, first)
		assert.Equal(t, "hello", *first)
	})
}

func TestLast(t *testing.T) {
	t.Run("获取最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30)
		last := list.Last()
		assert.Equal(t, 30, last)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(42)
		last := list.Last()
		assert.Equal(t, 42, last)
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Panics(t, func() {
			list.Last()
		})
	})
}

func TestLastOrNull(t *testing.T) {
	t.Run("获取最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(10, 20, 30)
		last := list.LastOrNull()
		assert.NotNil(t, last)
		assert.Equal(t, 30, *last)
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := goxt.NewXList[int]()
		last := list.LastOrNull()
		assert.Nil(t, last)
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements("world")
		last := list.LastOrNull()
		assert.NotNil(t, last)
		assert.Equal(t, "world", *last)
	})
}

func TestAddFirst(t *testing.T) {
	t.Run("向空列表添加第一个元素", func(t *testing.T) {
		list := goxt.NewXList[int]()
		list.AddFirst(1)
		assert.Equal(t, goxt.XInt(1), list.Size())
		assert.Equal(t, 1, list[0])
	})

	t.Run("向非空列表头部添加", func(t *testing.T) {
		list := goxt.NewXListWithElements(2, 3, 4)
		list.AddFirst(1)
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 2, list[1])
	})

	t.Run("多次添加到头部", func(t *testing.T) {
		list := goxt.NewXList[int]()
		list.AddFirst(3)
		list.AddFirst(2)
		list.AddFirst(1)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 2, list[1])
		assert.Equal(t, 3, list[2])
	})
}

func TestAddLast(t *testing.T) {
	t.Run("向空列表添加最后一个元素", func(t *testing.T) {
		list := goxt.NewXList[int]()
		list.AddLast(1)
		assert.Equal(t, goxt.XInt(1), list.Size())
		assert.Equal(t, 1, list[0])
	})

	t.Run("向非空列表尾部添加", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		list.AddLast(4)
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 4, list[3])
	})

	t.Run("多次添加到尾部", func(t *testing.T) {
		list := goxt.NewXList[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 3, list[2])
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("删除第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveFirst()
		assert.Equal(t, 1, removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
		assert.Equal(t, 2, list[0])
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(42)
		removed := list.RemoveFirst()
		assert.Equal(t, 42, removed)
		assert.Equal(t, goxt.XInt(0), list.Size())
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Panics(t, func() {
			list.RemoveFirst()
		})
	})
}

func TestRemoveFirstOrNull(t *testing.T) {
	t.Run("删除第一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveFirstOrNull()
		assert.NotNil(t, removed)
		assert.Equal(t, 1, *removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := goxt.NewXList[int]()
		removed := list.RemoveFirstOrNull()
		assert.Nil(t, removed)
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("删除最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveLast()
		assert.Equal(t, 3, removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
		assert.Equal(t, 2, list[1])
	})

	t.Run("单元素列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(42)
		removed := list.RemoveLast()
		assert.Equal(t, 42, removed)
		assert.Equal(t, goxt.XInt(0), list.Size())
	})

	t.Run("空列表panic", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Panics(t, func() {
			list.RemoveLast()
		})
	})
}

func TestRemoveLastOrNull(t *testing.T) {
	t.Run("删除最后一个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		removed := list.RemoveLastOrNull()
		assert.NotNil(t, removed)
		assert.Equal(t, 3, *removed)
		assert.Equal(t, goxt.XInt(2), list.Size())
	})

	t.Run("空列表返回nil", func(t *testing.T) {
		list := goxt.NewXList[int]()
		removed := list.RemoveLastOrNull()
		assert.Nil(t, removed)
	})
}

// ==================== 批量插入和范围删除测试 ====================

func TestInsertAll(t *testing.T) {
	t.Run("在开头插入多个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(4, 5, 6)
		toInsert := goxt.NewXListWithElements(1, 2, 3)
		result := list.InsertAll(0, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(6), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 3, list[2])
		assert.Equal(t, 4, list[3])
	})

	t.Run("在中间插入多个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 5, 6)
		toInsert := goxt.NewXListWithElements(3, 4)
		result := list.InsertAll(2, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(6), list.Size())
		assert.Equal(t, 3, list[2])
		assert.Equal(t, 4, list[3])
		assert.Equal(t, 5, list[4])
	})

	t.Run("在末尾插入多个元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toInsert := goxt.NewXListWithElements(4, 5, 6)
		result := list.InsertAll(3, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(6), list.Size())
		assert.Equal(t, 4, list[3])
		assert.Equal(t, 6, list[5])
	})

	t.Run("插入空列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toInsert := goxt.NewXList[int]()
		result := list.InsertAll(1, toInsert)
		assert.True(t, bool(result))
		assert.Equal(t, goxt.XInt(3), list.Size())
	})

	t.Run("索引越界返回false", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		toInsert := goxt.NewXListWithElements(4, 5)
		result := list.InsertAll(10, toInsert)
		assert.False(t, bool(result))
	})
}

func TestRemoveRange(t *testing.T) {
	t.Run("删除中间范围的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5, 6)
		list.RemoveRange(2, 4)
		assert.Equal(t, goxt.XInt(4), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 2, list[1])
		assert.Equal(t, 5, list[2])
		assert.Equal(t, 6, list[3])
	})

	t.Run("删除开头的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		list.RemoveRange(0, 2)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 3, list[0])
		assert.Equal(t, 5, list[2])
	})

	t.Run("删除末尾的元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		list.RemoveRange(3, 5)
		assert.Equal(t, goxt.XInt(3), list.Size())
		assert.Equal(t, 1, list[0])
		assert.Equal(t, 3, list[2])
	})

	t.Run("删除所有元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		list.RemoveRange(0, 3)
		assert.Equal(t, goxt.XInt(0), list.Size())
		assert.True(t, bool(list.IsEmpty()))
	})

	t.Run("无效范围不操作", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		list.RemoveRange(3, 2) // fromIndex >= toIndex
		assert.Equal(t, goxt.XInt(5), list.Size())
	})
}

// ==================== 辅助方法测试 ====================

func TestLastIndex(t *testing.T) {
	t.Run("非空列表返回最后索引", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		assert.Equal(t, goxt.XInt(4), list.LastIndex())
	})

	t.Run("单元素列表返回0", func(t *testing.T) {
		list := goxt.NewXListWithElements(42)
		assert.Equal(t, goxt.XInt(0), list.LastIndex())
	})

	t.Run("空列表返回-1", func(t *testing.T) {
		list := goxt.NewXList[int]()
		assert.Equal(t, goxt.XInt(-1), list.LastIndex())
	})
}

func TestForEach(t *testing.T) {
	t.Run("遍历所有元素", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3, 4, 5)
		sum := 0
		list.ForEach(func(item int) {
			sum += item
		})
		assert.Equal(t, 15, sum)
	})

	t.Run("修改元素值", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		results := make([]int, 0)
		list.ForEach(func(item int) {
			results = append(results, item*2)
		})
		assert.Equal(t, []int{2, 4, 6}, results)
	})

	t.Run("空列表不执行操作", func(t *testing.T) {
		list := goxt.NewXList[int]()
		called := false
		list.ForEach(func(item int) {
			called = true
		})
		assert.False(t, called)
	})

	t.Run("遍历字符串列表", func(t *testing.T) {
		list := goxt.NewXListWithElements("a", "b", "c")
		result := ""
		list.ForEach(func(item string) {
			result += item
		})
		assert.Equal(t, "abc", result)
	})
}

func TestIfEmpty(t *testing.T) {
	t.Run("空列表返回默认值", func(t *testing.T) {
		list := goxt.NewXList[int]()
		result := list.IfEmpty(func() goxt.XList[int] {
			return goxt.NewXListWithElements(1, 2, 3)
		})
		assert.Equal(t, goxt.XInt(3), result.Size())
		assert.Equal(t, 1, result[0])
	})

	t.Run("非空列表返回原列表", func(t *testing.T) {
		list := goxt.NewXListWithElements(1, 2, 3)
		result := list.IfEmpty(func() goxt.XList[int] {
			return goxt.NewXListWithElements(4, 5, 6)
		})
		assert.Equal(t, goxt.XInt(3), result.Size())
		assert.Equal(t, 1, result[0])
		assert.Equal(t, 2, result[1])
	})

	t.Run("空列表返回空默认值", func(t *testing.T) {
		list := goxt.NewXList[int]()
		result := list.IfEmpty(func() goxt.XList[int] {
			return goxt.NewXList[int]()
		})
		assert.Equal(t, goxt.XInt(0), result.Size())
	})
}
