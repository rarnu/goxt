package goxt_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xjai/goxt"
)

func TestXString_Length(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, 5, s1.Length())

	s2 := goxt.XString("你好")
	assert.EqualValues(t, 6, s2.Length())

	s3 := goxt.XString("")
	assert.EqualValues(t, 0, s3.Length())

	s4 := goxt.XString("🎉")
	assert.EqualValues(t, 4, s4.Length())
}

func TestXString_RuneCount(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, 5, s1.RuneCount())

	s2 := goxt.XString("你好世界")
	assert.EqualValues(t, 4, s2.RuneCount())

	s3 := goxt.XString("")
	assert.EqualValues(t, 0, s3.RuneCount())

	s4 := goxt.XString("a🎉b")
	assert.EqualValues(t, 3, s4.RuneCount())
}

func TestXString_Uppercase(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "HELLO", s1.Uppercase())

	s2 := goxt.XString("Hello World")
	assert.EqualValues(t, "HELLO WORLD", s2.Uppercase())

	s3 := goxt.XString("ALREADY UPPER")
	assert.EqualValues(t, "ALREADY UPPER", s3.Uppercase())

	s4 := goxt.XString("")
	assert.EqualValues(t, "", s4.Uppercase())
}

func TestXString_Lowercase(t *testing.T) {
	s1 := goxt.XString("HELLO")
	assert.EqualValues(t, "hello", s1.Lowercase())

	s2 := goxt.XString("Hello World")
	assert.EqualValues(t, "hello world", s2.Lowercase())

	s3 := goxt.XString("already lower")
	assert.EqualValues(t, "already lower", s3.Lowercase())

	s4 := goxt.XString("")
	assert.EqualValues(t, "", s4.Lowercase())
}

func TestXString_Capitalize(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "Hello", s1.Capitalize())

	s2 := goxt.XString("HELLO")
	assert.EqualValues(t, "HELLO", s2.Capitalize())

	s3 := goxt.XString("")
	assert.EqualValues(t, "", s3.Capitalize())
}

func TestXString_Decapitalize(t *testing.T) {
	s1 := goxt.XString("Hello")
	assert.EqualValues(t, "hello", s1.Decapitalize())

	s2 := goxt.XString("HELLO")
	assert.EqualValues(t, "hELLO", s2.Decapitalize())

	s3 := goxt.XString("")
	assert.EqualValues(t, "", s3.Decapitalize())
}

func TestXString_Repeat(t *testing.T) {
	s1 := goxt.XString("ab")
	assert.EqualValues(t, "ababab", s1.Repeat(3))

	s2 := goxt.XString("x")
	assert.EqualValues(t, "", s2.Repeat(0))

	s3 := goxt.XString("abc")
	assert.EqualValues(t, "", s3.Repeat(-1))

	s4 := goxt.XString("")
	assert.EqualValues(t, "", s4.Repeat(5))
}

func TestXString_Trim(t *testing.T) {
	s1 := goxt.XString("  hello  ")
	assert.EqualValues(t, "hello", s1.Trim())

	s2 := goxt.XString("\t\nhello\r\n")
	assert.EqualValues(t, "hello", s2.Trim())

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.Trim())

	s4 := goxt.XString("   ")
	assert.EqualValues(t, "", s4.Trim())
}

func TestXString_TrimWithChars(t *testing.T) {
	s1 := goxt.XString("xxxhelloxxx")
	assert.EqualValues(t, "hello", s1.TrimWithChars('x'))

	s2 := goxt.XString("  [abc]  ")
	assert.EqualValues(t, "abc", s2.TrimWithChars(' ', '[', ']'))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.TrimWithChars('x', 'y'))

	s4 := goxt.XString("xyzhellozyx")
	assert.EqualValues(t, "hello", s4.TrimWithChars('x', 'y', 'z'))
}

func TestXString_TrimWithPredicate(t *testing.T) {
	s1 := goxt.XString("123hello456")
	result1 := s1.TrimWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r >= '0' && r <= '9'
	})
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("   hello   ")
	result2 := s2.TrimWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r == ' '
	})
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.TrimWithPredicate(func(r goxt.XRune) goxt.XBool {
		return false
	})
	assert.EqualValues(t, "hello", result3)

	s4 := goxt.XString("aaa")
	result4 := s4.TrimWithPredicate(func(r goxt.XRune) goxt.XBool {
		return true
	})
	assert.EqualValues(t, "", result4)
}

func TestXString_TrimStart(t *testing.T) {
	s1 := goxt.XString("  hello  ")
	assert.EqualValues(t, "hello  ", s1.TrimStart())

	s2 := goxt.XString("\t\nhello")
	assert.EqualValues(t, "hello", s2.TrimStart())

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.TrimStart())

	s4 := goxt.XString("   ")
	assert.EqualValues(t, "", s4.TrimStart())
}

func TestXString_TrimEnd(t *testing.T) {
	s1 := goxt.XString("  hello  ")
	assert.EqualValues(t, "  hello", s1.TrimEnd())

	s2 := goxt.XString("hello\r\n")
	assert.EqualValues(t, "hello", s2.TrimEnd())

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.TrimEnd())

	s4 := goxt.XString("   ")
	assert.EqualValues(t, "", s4.TrimEnd())
}

func TestXString_TrimStartWithChars(t *testing.T) {
	s1 := goxt.XString("xxxhello")
	assert.EqualValues(t, "hello", s1.TrimStartWithChars('x'))

	s2 := goxt.XString("  [abc]")
	assert.EqualValues(t, "abc]", s2.TrimStartWithChars(' ', '['))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.TrimStartWithChars('x'))

	s4 := goxt.XString("xyzhello")
	assert.EqualValues(t, "hello", s4.TrimStartWithChars('x', 'y', 'z'))
}

func TestXString_TrimEndWithChars(t *testing.T) {
	s1 := goxt.XString("helloxxx")
	assert.EqualValues(t, "hello", s1.TrimEndWithChars('x'))

	s2 := goxt.XString("[abc]  ")
	assert.EqualValues(t, "[abc", s2.TrimEndWithChars(' ', ']'))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.TrimEndWithChars('x'))

	s4 := goxt.XString("helloxyz")
	assert.EqualValues(t, "hello", s4.TrimEndWithChars('x', 'y', 'z'))
}

func TestXString_TrimStartWithPredicate(t *testing.T) {
	s1 := goxt.XString("123hello")
	result1 := s1.TrimStartWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r >= '0' && r <= '9'
	})
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("   hello")
	result2 := s2.TrimStartWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r == ' '
	})
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.TrimStartWithPredicate(func(r goxt.XRune) goxt.XBool {
		return false
	})
	assert.EqualValues(t, "hello", result3)

	s4 := goxt.XString("123")
	result4 := s4.TrimStartWithPredicate(func(r goxt.XRune) goxt.XBool {
		return true
	})
	assert.EqualValues(t, "", result4)
}

func TestXString_TrimEndWithPredicate(t *testing.T) {
	s1 := goxt.XString("hello456")
	result1 := s1.TrimEndWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r >= '0' && r <= '9'
	})
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("hello   ")
	result2 := s2.TrimEndWithPredicate(func(r goxt.XRune) goxt.XBool {
		return r == ' '
	})
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.TrimEndWithPredicate(func(r goxt.XRune) goxt.XBool {
		return false
	})
	assert.EqualValues(t, "hello", result3)

	s4 := goxt.XString("456")
	result4 := s4.TrimEndWithPredicate(func(r goxt.XRune) goxt.XBool {
		return true
	})
	assert.EqualValues(t, "", result4)
}

func TestXString_PadStart(t *testing.T) {
	s1 := goxt.XString("abc")
	assert.EqualValues(t, "00abc", s1.PadStart(5, '0'))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s2.PadStart(3, 'x'))

	s3 := goxt.XString("x")
	assert.EqualValues(t, "xxxxx", s3.PadStart(5, 'x'))

	s4 := goxt.XString("")
	assert.EqualValues(t, "   ", s4.PadStart(3, ' '))
}

func TestXString_PadEnd(t *testing.T) {
	s1 := goxt.XString("abc")
	assert.EqualValues(t, "abc00", s1.PadEnd(5, '0'))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s2.PadEnd(3, 'x'))

	s3 := goxt.XString("x")
	assert.EqualValues(t, "xxxxx", s3.PadEnd(5, 'x'))

	s4 := goxt.XString("")
	assert.EqualValues(t, "   ", s4.PadEnd(3, ' '))
}

func TestXString_IfEmpty(t *testing.T) {
	s1 := goxt.XString("hello")
	result1 := s1.IfEmpty(func() goxt.XString { return "default" })
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("")
	result2 := s2.IfEmpty(func() goxt.XString { return "default" })
	assert.EqualValues(t, "default", result2)

	s3 := goxt.XString("   ")
	result3 := s3.IfEmpty(func() goxt.XString { return "default" })
	assert.EqualValues(t, "   ", result3)
}

func TestXString_IfBlank(t *testing.T) {
	s1 := goxt.XString("hello")
	result1 := s1.IfBlank(func() goxt.XString { return "default" })
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("")
	result2 := s2.IfBlank(func() goxt.XString { return "default" })
	assert.EqualValues(t, "default", result2)

	s3 := goxt.XString("   ")
	result3 := s3.IfBlank(func() goxt.XString { return "default" })
	assert.EqualValues(t, "default", result3)
}

func TestXString_IsEmpty(t *testing.T) {
	s1 := goxt.XString("")
	assert.True(t, bool(s1.IsEmpty()))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.IsEmpty()))

	s3 := goxt.XString("   ")
	assert.False(t, bool(s3.IsEmpty()))
}

func TestXString_IsNotEmpty(t *testing.T) {
	s1 := goxt.XString("")
	assert.False(t, bool(s1.IsNotEmpty()))

	s2 := goxt.XString("hello")
	assert.True(t, bool(s2.IsNotEmpty()))

	s3 := goxt.XString("   ")
	assert.True(t, bool(s3.IsNotEmpty()))
}

func TestXString_IsBlank(t *testing.T) {
	s1 := goxt.XString("")
	assert.True(t, bool(s1.IsBlank()))

	s2 := goxt.XString("   ")
	assert.True(t, bool(s2.IsBlank()))

	s3 := goxt.XString("hello")
	assert.False(t, bool(s3.IsBlank()))

	s4 := goxt.XString("  hello  ")
	assert.False(t, bool(s4.IsBlank()))
}

func TestXString_IsNotBlank(t *testing.T) {
	s1 := goxt.XString("")
	assert.False(t, bool(s1.IsNotBlank()))

	s2 := goxt.XString("   ")
	assert.False(t, bool(s2.IsNotBlank()))

	s3 := goxt.XString("hello")
	assert.True(t, bool(s3.IsNotBlank()))

	s4 := goxt.XString("  hello  ")
	assert.True(t, bool(s4.IsNotBlank()))
}

func TestXString_LastIndex(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, 4, s1.LastIndex())

	s2 := goxt.XString("a")
	assert.EqualValues(t, 0, s2.LastIndex())

	s3 := goxt.XString("")
	assert.EqualValues(t, -1, s3.LastIndex())

	s4 := goxt.XString("你好")
	assert.EqualValues(t, 1, s4.LastIndex())
}

func TestXString_Substring(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "ell", s1.Substring(1, 4))

	s2 := goxt.XString("hello world")
	assert.EqualValues(t, "hello", s2.Substring(0, 5))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.Substring(-1, 10))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "", s4.Substring(3, 2))

	s5 := goxt.XString("你好世界")
	assert.EqualValues(t, "好世", s5.Substring(1, 3))
}

func TestXString_SubstringWithStartIndex(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "ello", s1.SubstringWithStartIndex(1))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s2.SubstringWithStartIndex(0))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "", s3.SubstringWithStartIndex(10))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s4.SubstringWithStartIndex(-1))

	s5 := goxt.XString("你好世界")
	assert.EqualValues(t, "世界", s5.SubstringWithStartIndex(2))
}

func TestXString_SubstringBefore(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.SubstringBefore("-", nil)
	assert.EqualValues(t, "hello", result1)

	s2 := goxt.XString("hello")
	result2 := s2.SubstringBefore("-", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.SubstringBefore("-", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.SubstringBefore("-", nil)
	assert.EqualValues(t, "a", result4)
}

func TestXString_SubstringAfter(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.SubstringAfter("-", nil)
	assert.EqualValues(t, "world-foo", result1)

	s2 := goxt.XString("hello")
	result2 := s2.SubstringAfter("-", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.SubstringAfter("-", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.SubstringAfter("-", nil)
	assert.EqualValues(t, "b-c", result4)
}

func TestXString_SubstringBeforeLast(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.SubstringBeforeLast("-", nil)
	assert.EqualValues(t, "hello-world", result1)

	s2 := goxt.XString("hello")
	result2 := s2.SubstringBeforeLast("-", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.SubstringBeforeLast("-", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.SubstringBeforeLast("-", nil)
	assert.EqualValues(t, "a-b", result4)
}

func TestXString_SubstringAfterLast(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.SubstringAfterLast("-", nil)
	assert.EqualValues(t, "foo", result1)

	s2 := goxt.XString("hello")
	result2 := s2.SubstringAfterLast("-", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.SubstringAfterLast("-", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.SubstringAfterLast("-", nil)
	assert.EqualValues(t, "c", result4)
}

func TestXString_ReplaceRange(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "heXYZ", s1.ReplaceRange(2, 5, "XYZ"))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, "heXYZo", s2.ReplaceRange(2, 4, "XYZ"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.ReplaceRange(3, 2, "XYZ"))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "XYZ", s4.ReplaceRange(0, 5, "XYZ"))

	s5 := goxt.XString("你好世界")
	assert.EqualValues(t, "你XYZ", s5.ReplaceRange(1, 4, "XYZ"))
}

func TestXString_RemoveRange(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, "he", s1.RemoveRange(2, 5))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, "ho", s2.RemoveRange(1, 4))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.RemoveRange(3, 2))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "", s4.RemoveRange(0, 5))
}

func TestXString_RemovePrefix(t *testing.T) {
	s1 := goxt.XString("hello-world")
	assert.EqualValues(t, "-world", s1.RemovePrefix("hello"))

	s2 := goxt.XString("hello-world")
	assert.EqualValues(t, "hello-world", s2.RemovePrefix("xyz"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "", s3.RemovePrefix("hello"))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s4.RemovePrefix(""))
}

func TestXString_RemoveSuffix(t *testing.T) {
	s1 := goxt.XString("hello-world")
	assert.EqualValues(t, "hello-", s1.RemoveSuffix("world"))

	s2 := goxt.XString("hello-world")
	assert.EqualValues(t, "hello-world", s2.RemoveSuffix("xyz"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "", s3.RemoveSuffix("hello"))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s4.RemoveSuffix(""))
}

func TestXString_RemoveSurrounding(t *testing.T) {
	s1 := goxt.XString("[hello]")
	assert.EqualValues(t, "hello", s1.RemoveSurrounding("[", "]"))

	s2 := goxt.XString("[hello]")
	assert.EqualValues(t, "[hello]", s2.RemoveSurrounding("(", ")"))

	s3 := goxt.XString("((hello))")
	assert.EqualValues(t, "hello", s3.RemoveSurrounding("((", "))"))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s4.RemoveSurrounding("[", "]"))
}

func TestXString_ReplaceBefore(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.ReplaceBefore("-", "XXX", nil)
	assert.EqualValues(t, "XXX-world-foo", result1)

	s2 := goxt.XString("hello")
	result2 := s2.ReplaceBefore("-", "XXX", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.ReplaceBefore("-", "XXX", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.ReplaceBefore("-", "X", nil)
	assert.EqualValues(t, "X-b-c", result4)
}

func TestXString_ReplaceAfter(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.ReplaceAfter("-", "XXX", nil)
	assert.EqualValues(t, "hello-XXX", result1)

	s2 := goxt.XString("hello")
	result2 := s2.ReplaceAfter("-", "XXX", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.ReplaceAfter("-", "XXX", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.ReplaceAfter("-", "X", nil)
	assert.EqualValues(t, "a-X", result4)
}

func TestXString_ReplaceAfterLast(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.ReplaceAfterLast("-", "XXX", nil)
	assert.EqualValues(t, "hello-world-XXX", result1)

	s2 := goxt.XString("hello")
	result2 := s2.ReplaceAfterLast("-", "XXX", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.ReplaceAfterLast("-", "XXX", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.ReplaceAfterLast("-", "X", nil)
	assert.EqualValues(t, "a-b-X", result4)
}

func TestXString_ReplaceBeforeLast(t *testing.T) {
	s1 := goxt.XString("hello-world-foo")
	result1 := s1.ReplaceBeforeLast("-", "XXX", nil)
	assert.EqualValues(t, "XXX-foo", result1)

	s2 := goxt.XString("hello")
	result2 := s2.ReplaceBeforeLast("-", "XXX", nil)
	assert.EqualValues(t, "hello", result2)

	s3 := goxt.XString("hello")
	result3 := s3.ReplaceBeforeLast("-", "XXX", new(goxt.XString("not-found")))
	assert.EqualValues(t, "not-found", result3)

	s4 := goxt.XString("a-b-c")
	result4 := s4.ReplaceBeforeLast("-", "X", nil)
	assert.EqualValues(t, "X-c", result4)
}

func TestXString_Replace(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.EqualValues(t, "hello universe", s1.Replace("world", "universe"))

	s2 := goxt.XString("aaa")
	assert.EqualValues(t, "bbb", s2.Replace("a", "b"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.Replace("xyz", "abc"))

	s4 := goxt.XString("")
	assert.EqualValues(t, "", s4.Replace("a", "b"))
}

func TestXString_ReplaceIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.EqualValues(t, "Hello Universe", s1.ReplaceIgnoreCase("WORLD", "Universe"))

	s2 := goxt.XString("AAA BBB AAA")
	assert.EqualValues(t, "XXX BBB XXX", s2.ReplaceIgnoreCase("aaa", "XXX"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.ReplaceIgnoreCase("xyz", "abc"))

	s4 := goxt.XString("HeLLo")
	assert.EqualValues(t, "Hi", s4.ReplaceIgnoreCase("HELLO", "Hi"))
}

func TestXString_ReplaceFirst(t *testing.T) {
	s1 := goxt.XString("aaa")
	assert.EqualValues(t, "baa", s1.ReplaceFirst("a", "b"))

	s2 := goxt.XString("hello world")
	assert.EqualValues(t, "hi world", s2.ReplaceFirst("hello", "hi"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.ReplaceFirst("xyz", "abc"))

	s4 := goxt.XString("aaa")
	assert.EqualValues(t, "aaa", s4.ReplaceFirst("", "b"))
}

func TestXString_ReplaceFirstIgnoreCase(t *testing.T) {
	s1 := goxt.XString("AAA aaa")
	assert.EqualValues(t, "BBB aaa", s1.ReplaceFirstIgnoreCase("aaa", "BBB"))

	s2 := goxt.XString("Hello WORLD")
	assert.EqualValues(t, "Hi WORLD", s2.ReplaceFirstIgnoreCase("hello", "Hi"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.ReplaceFirstIgnoreCase("xyz", "abc"))

	s4 := goxt.XString("AaA")
	assert.EqualValues(t, "X", s4.ReplaceFirstIgnoreCase("AAA", "X"))
}

func TestXString_StartsWith(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.True(t, bool(s1.StartsWith("hello")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.StartsWith("world")))

	s3 := goxt.XString("")
	assert.True(t, bool(s3.StartsWith("")))

	s4 := goxt.XString("hello")
	assert.True(t, bool(s4.StartsWith("")))
}

func TestXString_StartsWithIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.True(t, bool(s1.StartsWithIgnoreCase("HELLO")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.StartsWithIgnoreCase("WORLD")))

	s3 := goxt.XString("HELLO")
	assert.True(t, bool(s3.StartsWithIgnoreCase("hello")))

	s4 := goxt.XString("Hi")
	assert.False(t, bool(s4.StartsWithIgnoreCase("Hello")))
}

func TestXString_StartsWithStartIndex(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.True(t, bool(s1.StartsWithStartIndex("world", 6)))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.StartsWithStartIndex("hello", 1)))

	s3 := goxt.XString("hello")
	assert.False(t, bool(s3.StartsWithStartIndex("hello", 1)))

	s4 := goxt.XString("hello")
	assert.False(t, bool(s4.StartsWithStartIndex("hello", -1)))

	s5 := goxt.XString("hello")
	assert.True(t, bool(s5.StartsWithStartIndex("ello", 1)))
}

func TestXString_StartsWithStartIndexIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.True(t, bool(s1.StartsWithStartIndexIgnoreCase("WORLD", 6)))

	s2 := goxt.XString("Hello")
	assert.False(t, bool(s2.StartsWithStartIndexIgnoreCase("HELLO", 1)))

	s3 := goxt.XString("hello")
	assert.False(t, bool(s3.StartsWithStartIndexIgnoreCase("hello", 10)))

	s4 := goxt.XString("Hi")
	assert.False(t, bool(s4.StartsWithStartIndexIgnoreCase("Hello", 0)))

	s5 := goxt.XString("Hello")
	assert.True(t, bool(s5.StartsWithStartIndexIgnoreCase("ELLO", 1)))
}

func TestXString_EndsWith(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.True(t, bool(s1.EndsWith("world")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.EndsWith("world")))

	s3 := goxt.XString("")
	assert.True(t, bool(s3.EndsWith("")))

	s4 := goxt.XString("hello")
	assert.True(t, bool(s4.EndsWith("")))
}

func TestXString_EndsWithIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.True(t, bool(s1.EndsWithIgnoreCase("WORLD")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.EndsWithIgnoreCase("WORLD")))

	s3 := goxt.XString("WORLD")
	assert.True(t, bool(s3.EndsWithIgnoreCase("world")))

	s4 := goxt.XString("Hi")
	assert.False(t, bool(s4.EndsWithIgnoreCase("Hello")))
}

func TestXString_IndexOf(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.EqualValues(t, 0, s1.IndexOf("hello"))

	s2 := goxt.XString("hello world")
	assert.EqualValues(t, 6, s2.IndexOf("world"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.IndexOf("xyz"))

	s4 := goxt.XString("aaa")
	assert.EqualValues(t, 0, s4.IndexOf("a"))
}

func TestXString_LastIndexOf(t *testing.T) {
	s1 := goxt.XString("hello world hello")
	assert.EqualValues(t, 12, s1.LastIndexOf("hello"))

	s2 := goxt.XString("aaa")
	assert.EqualValues(t, 2, s2.LastIndexOf("a"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.LastIndexOf("xyz"))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, 0, s4.LastIndexOf("hello"))
}

func TestXString_IndexOfIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.EqualValues(t, 0, s1.IndexOfIgnoreCase("HELLO"))

	s2 := goxt.XString("Hello WORLD")
	assert.EqualValues(t, 6, s2.IndexOfIgnoreCase("world"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.IndexOfIgnoreCase("xyz"))

	s4 := goxt.XString("HeLLo")
	assert.EqualValues(t, 0, s4.IndexOfIgnoreCase("hello"))
}

func TestXString_LastIndexOfIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World HELLO")
	assert.EqualValues(t, 12, s1.LastIndexOfIgnoreCase("hello"))

	s2 := goxt.XString("AAA aaa")
	assert.EqualValues(t, 6, s2.LastIndexOfIgnoreCase("a"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.LastIndexOfIgnoreCase("xyz"))

	s4 := goxt.XString("HeLLo")
	assert.EqualValues(t, 0, s4.LastIndexOfIgnoreCase("HELLO"))
}

func TestXString_Contains(t *testing.T) {
	s1 := goxt.XString("hello world")
	assert.True(t, bool(s1.Contains("hello")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.Contains("xyz")))

	s3 := goxt.XString("")
	assert.True(t, bool(s3.Contains("")))

	s4 := goxt.XString("hello")
	assert.True(t, bool(s4.Contains("")))
}

func TestXString_ContainsIgnoreCase(t *testing.T) {
	s1 := goxt.XString("Hello World")
	assert.True(t, bool(s1.ContainsIgnoreCase("HELLO")))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.ContainsIgnoreCase("XYZ")))

	s3 := goxt.XString("HELLO")
	assert.True(t, bool(s3.ContainsIgnoreCase("hello")))

	s4 := goxt.XString("HeLLo WoRLd")
	assert.True(t, bool(s4.ContainsIgnoreCase("WORLD")))
}

func TestXString_ContainsRegex(t *testing.T) {
	s1 := goxt.XString("hello123")
	assert.True(t, bool(s1.ContainsRegex(`\d+`)))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.ContainsRegex(`\d+`)))

	s3 := goxt.XString("abc123def")
	assert.True(t, bool(s3.ContainsRegex(`[0-9]+`)))

	s4 := goxt.XString("")
	assert.False(t, bool(s4.ContainsRegex(`.+`)))
}

func TestXString_Matches(t *testing.T) {
	s1 := goxt.XString("hello123")
	assert.True(t, bool(s1.Matches(`^[a-z]+\d+$`)))

	s2 := goxt.XString("hello")
	assert.False(t, bool(s2.Matches(`^\d+$`)))

	s3 := goxt.XString("12345")
	assert.True(t, bool(s3.Matches(`^\d+$`)))

	s4 := goxt.XString("hello world")
	assert.False(t, bool(s4.Matches(`^hello$`)))
}

func TestXString_ReplaceRegex(t *testing.T) {
	s1 := goxt.XString("hello123world456")
	assert.EqualValues(t, "helloXworldX", s1.ReplaceRegex(`\d+`, "X"))

	s2 := goxt.XString("aaa bbb ccc")
	assert.EqualValues(t, "XXX XXX XXX", s2.ReplaceRegex(`\w+`, "XXX"))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, "hello", s3.ReplaceRegex(`\d+`, "X"))

	s4 := goxt.XString("abc123def")
	assert.EqualValues(t, "abcdef", s4.ReplaceRegex(`\d+`, ""))
}

func TestXString_ReplaceRegexFunc(t *testing.T) {
	s1 := goxt.XString("hello123")
	result1 := s1.ReplaceRegexFunc(`\d+`, func(match goxt.XString) goxt.XString {
		return "NUM"
	})
	assert.EqualValues(t, "helloNUM", result1)

	s2 := goxt.XString("aa bb cc")
	result2 := s2.ReplaceRegexFunc(`\w+`, func(match goxt.XString) goxt.XString {
		return match.Uppercase()
	})
	assert.EqualValues(t, "AA BB CC", result2)

	s3 := goxt.XString("hello")
	result3 := s3.ReplaceRegexFunc(`\d+`, func(match goxt.XString) goxt.XString {
		return "X"
	})
	assert.EqualValues(t, "hello", result3)

	s4 := goxt.XString("123")
	result4 := s4.ReplaceRegexFunc(`\d+`, func(match goxt.XString) goxt.XString {
		return "456"
	})
	assert.EqualValues(t, "456", result4)
}

func TestXString_Split(t *testing.T) {
	s1 := goxt.XString("a,b,c")
	result1 := s1.Split(",")
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result1)

	s2 := goxt.XString("hello world")
	result2 := s2.Split(" ")
	assert.EqualValues(t, []goxt.XString{"hello", "world"}, result2)

	s3 := goxt.XString("a,b;c")
	result3 := s3.Split(",", ";")
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result3)

	s4 := goxt.XString("hello")
	result4 := s4.Split(",")
	assert.EqualValues(t, []goxt.XString{"hello"}, result4)
}

func TestXString_SplitIgnoreCase(t *testing.T) {
	s1 := goxt.XString("aHELLObhelloc")
	result1 := s1.SplitIgnoreCase("hello")
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result1)

	s2 := goxt.XString("A, B, C")
	result2 := s2.SplitIgnoreCase(", ")
	assert.EqualValues(t, []goxt.XString{"A", "B", "C"}, result2)

	s3 := goxt.XString("hello")
	result3 := s3.SplitIgnoreCase("WORLD")
	assert.EqualValues(t, []goxt.XString{"hello"}, result3)

	s4 := goxt.XString("XxYyZz")
	result4 := s4.SplitIgnoreCase("XY")
	assert.EqualValues(t, []goxt.XString{"X", "yZz"}, result4)
}

func TestXString_SplitWithLimit(t *testing.T) {
	s1 := goxt.XString("a,b,c,d")
	result1 := s1.SplitWithLimit(3, ",")
	assert.EqualValues(t, []goxt.XString{"a", "b", "c,d"}, result1)

	s2 := goxt.XString("hello world foo bar")
	result2 := s2.SplitWithLimit(2, " ")
	assert.EqualValues(t, []goxt.XString{"hello", "world foo bar"}, result2)

	s3 := goxt.XString("a,b,c")
	result3 := s3.SplitWithLimit(10, ",")
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result3)

	s4 := goxt.XString("a,b,c")
	result4 := s4.SplitWithLimit(1, ",")
	assert.EqualValues(t, []goxt.XString{"a,b,c"}, result4)
}

func TestXString_SplitWithLimitIgnoreCase(t *testing.T) {
	s1 := goxt.XString("aHELLObHELLOcHELLod")
	result1 := s1.SplitWithLimitIgnoreCase(3, "hello")
	assert.EqualValues(t, []goxt.XString{"a", "b", "cHELLod"}, result1)

	s2 := goxt.XString("A, B, C, D")
	result2 := s2.SplitWithLimitIgnoreCase(2, ", ")
	assert.EqualValues(t, []goxt.XString{"A", "B, C, D"}, result2)

	s3 := goxt.XString("hello")
	result3 := s3.SplitWithLimitIgnoreCase(5, "WORLD")
	assert.EqualValues(t, []goxt.XString{"hello"}, result3)

	s4 := goxt.XString("XxYyZz")
	result4 := s4.SplitWithLimitIgnoreCase(2, "xy")
	assert.EqualValues(t, []goxt.XString{"X", "yZz"}, result4)
}

func TestXString_SplitRegex(t *testing.T) {
	s1 := goxt.XString("a1b2c3")
	result1 := s1.SplitRegex(regexp.MustCompile(`\d`))
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result1)

	s2 := goxt.XString("hello, world; foo")
	result2 := s2.SplitRegex(regexp.MustCompile(`[,;]`))
	assert.EqualValues(t, []goxt.XString{"hello", " world", " foo"}, result2)

	s3 := goxt.XString("hello")
	result3 := s3.SplitRegex(regexp.MustCompile(`\d`))
	assert.EqualValues(t, []goxt.XString{"hello"}, result3)

	s4 := goxt.XString("a b c")
	result4 := s4.SplitRegex(regexp.MustCompile(`\s+`))
	assert.EqualValues(t, []goxt.XString{"a", "b", "c"}, result4)
}

func TestXString_SplitRegexWithLimit(t *testing.T) {
	s1 := goxt.XString("a1b2c3d")
	result1 := s1.SplitRegexWithLimit(regexp.MustCompile(`\d`), 3)
	assert.EqualValues(t, []goxt.XString{"a", "b", "c3d"}, result1)

	s2 := goxt.XString("hello, world; foo")
	result2 := s2.SplitRegexWithLimit(regexp.MustCompile(`[,;]`), 2)
	assert.EqualValues(t, []goxt.XString{"hello", " world; foo"}, result2)

	s3 := goxt.XString("a b c")
	result3 := s3.SplitRegexWithLimit(regexp.MustCompile(`\s+`), 1)
	assert.EqualValues(t, []goxt.XString{"a b c"}, result3)

	s4 := goxt.XString("1a2b3c")
	result4 := s4.SplitRegexWithLimit(regexp.MustCompile(`\d`), 10)
	assert.EqualValues(t, []goxt.XString{"", "a", "b", "c"}, result4)
}

func TestXString_Lines(t *testing.T) {
	s1 := goxt.XString("abc\ndef\nxyz")
	result1 := s1.Lines()
	assert.EqualValues(t, []goxt.XString{"abc", "def", "xyz"}, result1)

	s2 := goxt.XString("line1\r\nline2\r\nline3")
	result2 := s2.Lines()
	assert.EqualValues(t, []goxt.XString{"line1", "line2", "line3"}, result2)

	s3 := goxt.XString("single line")
	result3 := s3.Lines()
	assert.EqualValues(t, []goxt.XString{"single line"}, result3)

	s4 := goxt.XString("")
	result4 := s4.Lines()
	assert.EqualValues(t, []goxt.XString{}, result4)

	s5 := goxt.XString("line1\nline2\n")
	result5 := s5.Lines()
	assert.EqualValues(t, []goxt.XString{"line1", "line2"}, result5)
}

func TestXString_ToBoolean(t *testing.T) {
	s1 := goxt.XString("true")
	assert.True(t, bool(s1.ToBoolean()))

	s2 := goxt.XString("TRUE")
	assert.True(t, bool(s2.ToBoolean()))

	s3 := goxt.XString("false")
	assert.False(t, bool(s3.ToBoolean()))

	s4 := goxt.XString("anything")
	assert.False(t, bool(s4.ToBoolean()))
}

func TestXString_ToBooleanOrNil(t *testing.T) {
	s1 := goxt.XString("true")
	result1 := s1.ToBooleanOrNil()
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.True(t, bool(*result1))
	}
	s2 := goxt.XString("false")
	result2 := s2.ToBooleanOrNil()
	assert.NotNil(t, result2)
	if result2 != nil {
		assert.False(t, bool(*result2))
	}

	s3 := goxt.XString("invalid")
	result3 := s3.ToBooleanOrNil()
	assert.Nil(t, result3)

	s4 := goxt.XString("TRUE")
	result4 := s4.ToBooleanOrNil()
	assert.NotNil(t, result4)
	if result4 != nil {
		assert.True(t, bool(*result4))
	}
}

func TestXString_ToInt(t *testing.T) {
	s1 := goxt.XString("123")
	assert.EqualValues(t, 123, s1.ToInt())

	s2 := goxt.XString("-456")
	assert.EqualValues(t, -456, s2.ToInt())

	s3 := goxt.XString("0")
	assert.EqualValues(t, 0, s3.ToInt())

	defer func() {
		recover()
	}()
	goxt.XString("abc").ToInt()
}

func TestXString_ToIntOrNil(t *testing.T) {
	s1 := goxt.XString("123")
	result1 := s1.ToIntOrNil()
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.EqualValues(t, 123, *result1)
	}

	s2 := goxt.XString("invalid")
	result2 := s2.ToIntOrNil()
	assert.Nil(t, result2)

	s3 := goxt.XString("-456")
	result3 := s3.ToIntOrNil()
	assert.NotNil(t, result3)
	if result3 != nil {
		assert.EqualValues(t, -456, *result3)
	}
}

func TestXString_ToInt64(t *testing.T) {
	s1 := goxt.XString("123456789")
	assert.EqualValues(t, 123456789, s1.ToInt64())

	s2 := goxt.XString("-987654321")
	assert.EqualValues(t, -987654321, s2.ToInt64())

	s3 := goxt.XString("0")
	assert.EqualValues(t, 0, s3.ToInt64())

	defer func() {
		recover()
	}()
	goxt.XString("abc").ToInt64()
}

func TestXString_ToInt64OrNil(t *testing.T) {
	s1 := goxt.XString("123456789")
	result1 := s1.ToInt64OrNil()
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.EqualValues(t, 123456789, *result1)
	}

	s2 := goxt.XString("invalid")
	result2 := s2.ToInt64OrNil()
	assert.Nil(t, result2)

	s3 := goxt.XString("-987654321")
	result3 := s3.ToInt64OrNil()
	assert.NotNil(t, result3)
	if result3 != nil {
		assert.EqualValues(t, -987654321, *result3)
	}
}

func TestXString_ToFloat32(t *testing.T) {
	s1 := goxt.XString("123.45")
	assert.InDelta(t, 123.45, float32(s1.ToFloat32()), 0.01)

	s2 := goxt.XString("-67.89")
	assert.InDelta(t, -67.89, float32(s2.ToFloat32()), 0.01)

	s3 := goxt.XString("0.0")
	assert.InDelta(t, 0.0, float32(s3.ToFloat32()), 0.01)

	defer func() {
		recover()
	}()
	goxt.XString("abc").ToFloat32()
}

func TestXString_ToFloat32OrNil(t *testing.T) {
	s1 := goxt.XString("123.45")
	result1 := s1.ToFloat32OrNil()
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.InDelta(t, 123.45, float32(*result1), 0.01)
	}

	s2 := goxt.XString("invalid")
	result2 := s2.ToFloat32OrNil()
	assert.Nil(t, result2)

	s3 := goxt.XString("-67.89")
	result3 := s3.ToFloat32OrNil()
	assert.NotNil(t, result3)
	if result3 != nil {
		assert.InDelta(t, -67.89, float32(*result3), 0.01)
	}
}

func TestXString_ToFloat64(t *testing.T) {
	s1 := goxt.XString("123.456789")
	assert.InDelta(t, 123.456789, float64(s1.ToFloat64()), 0.0001)

	s2 := goxt.XString("-987.654321")
	assert.InDelta(t, -987.654321, float64(s2.ToFloat64()), 0.0001)

	s3 := goxt.XString("0.0")
	assert.InDelta(t, 0.0, float64(s3.ToFloat64()), 0.0001)

	defer func() {
		recover()
	}()
	goxt.XString("abc").ToFloat64()
}

func TestXString_ToFloat64OrNil(t *testing.T) {
	s1 := goxt.XString("123.456789")
	result1 := s1.ToFloat64OrNil()
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.InDelta(t, 123.456789, float64(*result1), 0.0001)
	}

	s2 := goxt.XString("invalid")
	result2 := s2.ToFloat64OrNil()
	assert.Nil(t, result2)

	s3 := goxt.XString("-987.654321")
	result3 := s3.ToFloat64OrNil()
	assert.NotNil(t, result3)
	if result3 != nil {
		assert.InDelta(t, -987.654321, float64(*result3), 0.0001)
	}
}

func TestXString_ToRuneArray(t *testing.T) {
	s1 := goxt.XString("hello")
	result1 := s1.ToRuneArray()
	assert.EqualValues(t, []goxt.XRune{'h', 'e', 'l', 'l', 'o'}, result1)

	s2 := goxt.XString("你好")
	result2 := s2.ToRuneArray()
	assert.EqualValues(t, []goxt.XRune{'你', '好'}, result2)

	s3 := goxt.XString("")
	result3 := s3.ToRuneArray()
	assert.EqualValues(t, []goxt.XRune{}, result3)
}

func TestXString_ToByteArray(t *testing.T) {
	s1 := goxt.XString("abc")
	result1 := s1.ToByteArray()
	assert.EqualValues(t, []goxt.XByte{97, 98, 99}, result1)

	s2 := goxt.XString("")
	result2 := s2.ToByteArray()
	assert.EqualValues(t, []goxt.XByte{}, result2)

	s3 := goxt.XString("A")
	result3 := s3.ToByteArray()
	assert.EqualValues(t, []goxt.XByte{65}, result3)
}

func TestXString_IndexOfWithStartIndex(t *testing.T) {
	s1 := goxt.XString("hello hello")
	assert.EqualValues(t, 6, s1.IndexOfWithStartIndex("hello", 1))

	s2 := goxt.XString("aaa")
	assert.EqualValues(t, 1, s2.IndexOfWithStartIndex("a", 1))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.IndexOfWithStartIndex("hello", 10))

	s4 := goxt.XString("hello")
	assert.EqualValues(t, 0, s4.IndexOfWithStartIndex("hello", -1))
}

func TestXString_LastIndexOfWithStartIndex(t *testing.T) {
	s1 := goxt.XString("hello hello hello")
	assert.EqualValues(t, 6, s1.LastIndexOfWithStartIndex("hello", 10))

	s2 := goxt.XString("aaa")
	assert.EqualValues(t, 1, s2.LastIndexOfWithStartIndex("a", 1))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.LastIndexOfWithStartIndex("hello", -1))

	s4 := goxt.XString("hello hello")
	assert.EqualValues(t, 0, s4.LastIndexOfWithStartIndex("hello", 5))
}

func TestXString_IndexOfAny(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, 0, s1.IndexOfAny([]goxt.XRune{'h', 'e', 'l'}))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, 2, s2.IndexOfAny([]goxt.XRune{'l', 'o'}))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.IndexOfAny([]goxt.XRune{'x', 'y', 'z'}))

	s4 := goxt.XString("")
	assert.EqualValues(t, -1, s4.IndexOfAny([]goxt.XRune{'a'}))
}

func TestXString_LastIndexOfAny(t *testing.T) {
	s1 := goxt.XString("hello")
	assert.EqualValues(t, 4, s1.LastIndexOfAny([]goxt.XRune{'o', 'l'}))

	s2 := goxt.XString("hello")
	assert.EqualValues(t, 3, s2.LastIndexOfAny([]goxt.XRune{'l'}))

	s3 := goxt.XString("hello")
	assert.EqualValues(t, -1, s3.LastIndexOfAny([]goxt.XRune{'x', 'y', 'z'}))

	s4 := goxt.XString("")
	assert.EqualValues(t, -1, s4.LastIndexOfAny([]goxt.XRune{'a'}))
}

func TestXString_FindAnyOf(t *testing.T) {
	s1 := goxt.XString("hello world")
	result1 := s1.FindAnyOf([]goxt.XString{"hello", "world"}, 0)
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.EqualValues(t, 0, result1.Index)
		assert.EqualValues(t, "hello", result1.Value)
	}

	s2 := goxt.XString("hello world")
	result2 := s2.FindAnyOf([]goxt.XString{"world", "hello"}, 0)
	assert.NotNil(t, result2)
	if result2 != nil {
		assert.EqualValues(t, 0, result2.Index)
		assert.EqualValues(t, "hello", result2.Value)
	}

	s3 := goxt.XString("hello")
	result3 := s3.FindAnyOf([]goxt.XString{"xyz"}, 0)
	assert.Nil(t, result3)

	s4 := goxt.XString("abc def")
	result4 := s4.FindAnyOf([]goxt.XString{"def"}, 0)
	assert.NotNil(t, result4)
	if result4 != nil {
		assert.EqualValues(t, 4, result4.Index)
		assert.EqualValues(t, "def", result4.Value)
	}
}

func TestXString_FindLastAnyOf(t *testing.T) {
	s1 := goxt.XString("hello world hello")
	result1 := s1.FindLastAnyOf([]goxt.XString{"hello", "world"}, 16)
	assert.NotNil(t, result1)
	if result1 != nil {
		assert.EqualValues(t, 12, result1.Index)
		assert.EqualValues(t, "hello", result1.Value)
	}

	s2 := goxt.XString("abc def abc")
	result2 := s2.FindLastAnyOf([]goxt.XString{"abc", "def"}, 8)
	assert.NotNil(t, result2)
	if result2 != nil {
		assert.EqualValues(t, 4, result2.Index)
		assert.EqualValues(t, "def", result2.Value)
	}

	s3 := goxt.XString("hello")
	result3 := s3.FindLastAnyOf([]goxt.XString{"xyz"}, 4)
	assert.Nil(t, result3)

	s4 := goxt.XString("a b c")
	result4 := s4.FindLastAnyOf([]goxt.XString{"b", "c"}, 4)
	assert.NotNil(t, result4)
	if result4 != nil {
		assert.EqualValues(t, 4, result4.Index)
		assert.EqualValues(t, "c", result4.Value)
	}
}