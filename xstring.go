package goxt

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	whiteSpaces = []rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}
	cutset      = "\t\n\v\f\r " + string(rune(0x85)) + string(rune(0xA0))
)

type XString string

func (s XString) Equal(other XString) XBool {
	return s == other
}

func (s XString) Length() XInt {
	return XInt(len(s))
}

func (s XString) RuneCount() XInt {
	return XInt(utf8.RuneCountInString(string(s)))
}

func (s XString) Uppercase() XString {
	return XString(strings.ToUpper(string(s)))
}

func (s XString) Lowercase() XString {
	return XString(strings.ToLower(string(s)))
}

func (s XString) Capitalize() XString {
	if s.RuneCount() == 0 {
		return ""
	}
	return XString(strings.ToUpper(string(s[:1]))) + s[1:]
}

func (s XString) Decapitalize() XString {
	if s.RuneCount() == 0 {
		return ""
	}
	return XString(strings.ToLower(string(s[:1]))) + s[1:]
}

func (s XString) Repeat(n XInt) XString {
	if n <= 0 {
		return ""
	}
	return XString(strings.Repeat(string(s), int(n)))
}

func (s XString) Trim() XString {
	return XString(strings.TrimSpace(string(s)))
}

func (s XString) TrimWithChars(chars ...XRune) XString {
	return s.TrimWithPredicate(func(r XRune) XBool {
		return XBool(slices.Contains(chars, r))
	})
}

func (s XString) TrimWithPredicate(predicate func(XRune) XBool) XString {
	runes := []XRune(s)
	if len(runes) == 0 {
		return ""
	}
	startIndex := XInt(0)
	endIndex := XInt(len(runes)) - 1
	startFound := XBool(false)

	for startIndex <= endIndex {
		index := IfThen(!startFound, startIndex, endIndex)
		match := predicate(runes[index])
		if !startFound {
			if !match {
				startFound = true
			} else {
				startIndex++
			}
		} else {
			if !match {
				break
			} else {
				endIndex--
			}
		}
	}
	if startIndex > endIndex {
		return ""
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	endByte := s.runeIndexToByteIndex(endIndex + 1)
	return s[startByte:endByte]
}

func (s XString) TrimStart() XString {
	return XString(strings.TrimLeft(string(s), cutset))
}

func (s XString) TrimEnd() XString {
	return XString(strings.TrimRight(string(s), cutset))
}

func (s XString) TrimStartWithChars(chars ...XRune) XString {
	return s.TrimStartWithPredicate(func(r XRune) XBool {
		return XBool(slices.Contains(chars, r))
	})
}

func (s XString) TrimEndWithChars(chars ...XRune) XString {
	return s.TrimEndWithPredicate(func(r XRune) XBool {
		return XBool(slices.Contains(chars, r))
	})
}

func (s XString) TrimStartWithPredicate(predicate func(XRune) XBool) XString {
	runes := []XRune(s)
	for index := range runes {
		if !predicate(runes[index]) {
			startByte := s.runeIndexToByteIndex(XInt(index))
			return s[startByte:]
		}
	}
	return ""
}

func (s XString) TrimEndWithPredicate(predicate func(XRune) XBool) XString {
	runes := []XRune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if !predicate(runes[i]) {
			endByte := s.runeIndexToByteIndex(XInt(i + 1))
			return s[:endByte]
		}
	}
	return ""
}

func (s XString) PadStart(length XInt, padChar XRune) XString {
	runeCount := s.RuneCount()
	if runeCount >= length {
		return s
	}
	padCount := length - runeCount
	padStr := XString(strings.Repeat(string(padChar), int(padCount)))
	return padStr + s
}

func (s XString) PadEnd(length XInt, padChar XRune) XString {
	runeCount := s.RuneCount()
	if runeCount >= length {
		return s
	}
	padCount := length - runeCount
	padStr := XString(strings.Repeat(string(padChar), int(padCount)))
	return s + padStr
}

func (s XString) IfEmpty(def func() XString) XString {
	if s.Length() == 0 {
		return def()
	}
	return s
}

func (s XString) IfBlank(def func() XString) XString {
	if s.Trim().Length() == 0 {
		return def()
	}
	return s
}

func (s XString) IsEmpty() XBool {
	return s.Length() == 0
}

func (s XString) IsNotEmpty() XBool {
	return !s.IsEmpty()
}

func (s XString) IsBlank() XBool {
	return s.Trim().Length() == 0
}

func (s XString) IsNotBlank() XBool {
	return !s.IsBlank()
}

func (s XString) LastIndex() XInt {
	return s.RuneCount() - 1
}

func (s XString) Substring(startIndex, endIndex XInt) XString {
	l := s.RuneCount()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > l {
		endIndex = l
	}
	if startIndex > endIndex {
		return ""
	}
	if startIndex == 0 && endIndex == l {
		return s
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	endByte := s.runeIndexToByteIndex(endIndex)
	return s[startByte:endByte]
}

func (s XString) SubstringWithStartIndex(startIndex XInt) XString {
	l := s.RuneCount()
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex > l {
		return ""
	}
	if startIndex == 0 {
		return s
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	return s[startByte:]
}

func (s XString) SubstringBefore(delimiter XString, missingDelimiterValue *XString) XString {
	byteIdx := strings.Index(string(s), string(delimiter))
	if byteIdx == -1 {
		if missingDelimiterValue != nil {
			return *missingDelimiterValue
		}
		return s
	}
	runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	return s.Substring(0, runeIdx)
}

func (s XString) SubstringAfter(delimiter XString, missingDelimiterValue *XString) XString {
	byteIdx := strings.Index(string(s), string(delimiter))
	if byteIdx == -1 {
		if missingDelimiterValue != nil {
			return *missingDelimiterValue
		}
		return s
	}
	delimiterRuneLen := delimiter.RuneCount()
	startRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx)) + delimiterRuneLen
	return s.SubstringWithStartIndex(startRuneIdx)
}

func (s XString) SubstringBeforeLast(delimiter XString, missingDelimiterValue *XString) XString {
	byteIdx := strings.LastIndex(string(s), string(delimiter))
	if byteIdx == -1 {
		if missingDelimiterValue != nil {
			return *missingDelimiterValue
		}
		return s
	}
	runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	return s.Substring(0, runeIdx)
}

func (s XString) SubstringAfterLast(delimiter XString, missingDelimiterValue *XString) XString {
	byteIdx := strings.LastIndex(string(s), string(delimiter))
	if byteIdx == -1 {
		if missingDelimiterValue != nil {
			return *missingDelimiterValue
		}
		return s
	}
	delimiterRuneLen := delimiter.RuneCount()
	startRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx)) + delimiterRuneLen
	return s.SubstringWithStartIndex(startRuneIdx)
}

func (s XString) ReplaceRange(startIndex XInt, endIndex XInt, replacement XString) XString {
	l := s.RuneCount()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > l {
		endIndex = l
	}
	if startIndex > endIndex {
		return s
	}

	startByte := s.runeIndexToByteIndex(startIndex)
	endByte := s.runeIndexToByteIndex(endIndex)
	return s[:startByte] + replacement + s[endByte:]
}

func (s XString) RemoveRange(startIndex XInt, endIndex XInt) XString {
	return s.ReplaceRange(startIndex, endIndex, "")
}

func (s XString) RemovePrefix(prefix XString) XString {
	if strings.HasPrefix(string(s), string(prefix)) {
		return s[len(prefix):]
	}
	return s
}

func (s XString) RemoveSuffix(suffix XString) XString {
	if strings.HasSuffix(string(s), string(suffix)) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

func (s XString) RemoveSurrounding(prefix XString, suffix XString) XString {
	if strings.HasPrefix(string(s), string(prefix)) && strings.HasSuffix(string(s), string(suffix)) {
		return s[len(prefix) : len(s)-len(suffix)]
	}
	return s
}

func (s XString) ReplaceBefore(delimiter, replacement XString, missingDelimiterValue *XString) XString {
	missing := s
	if missingDelimiterValue != nil {
		missing = *missingDelimiterValue
	}
	byteIdx := strings.Index(string(s), string(delimiter))
	if byteIdx == -1 {
		return missing
	}
	runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	return replacement + s.SubstringWithStartIndex(runeIdx)
}

func (s XString) ReplaceAfter(delimiter, replacement XString, missingDelimiterValue *XString) XString {
	missing := s
	if missingDelimiterValue != nil {
		missing = *missingDelimiterValue
	}
	byteIdx := strings.Index(string(s), string(delimiter))
	if byteIdx == -1 {
		return missing
	}
	delimiterRuneLen := delimiter.RuneCount()
	endRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx)) + delimiterRuneLen
	return s.Substring(0, endRuneIdx) + replacement
}

func (s XString) ReplaceAfterLast(delimiter, replacement XString, missingDelimiterValue *XString) XString {
	missing := s
	if missingDelimiterValue != nil {
		missing = *missingDelimiterValue
	}
	byteIdx := strings.LastIndex(string(s), string(delimiter))
	if byteIdx == -1 {
		return missing
	}
	delimiterRuneLen := delimiter.RuneCount()
	endRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx)) + delimiterRuneLen
	return s.Substring(0, endRuneIdx) + replacement
}

func (s XString) ReplaceBeforeLast(delimiter, replacement XString, missingDelimiterValue *XString) XString {
	missing := s
	if missingDelimiterValue != nil {
		missing = *missingDelimiterValue
	}
	byteIdx := strings.LastIndex(string(s), string(delimiter))
	if byteIdx == -1 {
		return missing
	}
	runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	return replacement + s.SubstringWithStartIndex(runeIdx)
}

func (s XString) Replace(old XString, new XString) XString {
	return XString(strings.ReplaceAll(string(s), string(old), string(new)))
}

func (s XString) ReplaceIgnoreCase(old XString, new XString) XString {
	re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(string(old)))
	return XString(re.ReplaceAllString(string(s), string(new)))
}

func (s XString) ReplaceFirst(oldValue XString, newValue XString) XString {
	if oldValue == "" {
		return s
	}
	byteIdx := strings.Index(string(s), string(oldValue))
	if byteIdx == -1 {
		return s
	}
	startRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	oldRuneLen := oldValue.RuneCount()
	endRuneIdx := startRuneIdx + oldRuneLen
	return s.Substring(0, startRuneIdx) + newValue + s.SubstringWithStartIndex(endRuneIdx)
}

func (s XString) ReplaceFirstIgnoreCase(oldValue XString, newValue XString) XString {
	if oldValue == "" {
		return s
	}
	lowerS := s.Lowercase()
	lowerOld := oldValue.Lowercase()
	byteIdx := strings.Index(string(lowerS), string(lowerOld))
	if byteIdx == -1 {
		return s
	}
	startRuneIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
	oldRuneLen := oldValue.RuneCount()
	endRuneIdx := startRuneIdx + oldRuneLen
	return s.Substring(0, startRuneIdx) + newValue + s.SubstringWithStartIndex(endRuneIdx)
}

func (s XString) StartsWith(prefix XString) XBool {
	return XBool(strings.HasPrefix(string(s), string(prefix)))
}

func (s XString) StartsWithIgnoreCase(prefix XString) XBool {
	return XBool(strings.HasPrefix(string(s.Uppercase()), string(prefix.Uppercase())))
}

func (s XString) StartsWithStartIndex(prefix XString, startIndex XInt) XBool {
	if startIndex < 0 || startIndex >= s.RuneCount() {
		return false
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	substr := s[startByte:]
	return XBool(strings.HasPrefix(string(substr), string(prefix)))
}

func (s XString) StartsWithStartIndexIgnoreCase(prefix XString, startIndex XInt) XBool {
	if startIndex < 0 || startIndex >= s.RuneCount() {
		return false
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	substr := s[startByte:]
	prefixLen := prefix.RuneCount()
	if substr.RuneCount() < prefixLen {
		return false
	}
	prefixByteLen := s.runeIndexToByteIndex(prefixLen)
	return XBool(strings.EqualFold(string(substr[:prefixByteLen]), string(prefix)))
}

func (s XString) EndsWith(prefix XString) XBool {
	return XBool(strings.HasSuffix(string(s), string(prefix)))
}

func (s XString) EndsWithIgnoreCase(prefix XString) XBool {
	return XBool(strings.HasSuffix(string(s.Uppercase()), string(prefix.Uppercase())))
}

func (s XString) IndexOf(str XString) XInt {
	byteIdx := strings.Index(string(s), string(str))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) LastIndexOf(str XString) XInt {
	byteIdx := strings.LastIndex(string(s), string(str))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) IndexOfIgnoreCase(str XString) XInt {
	byteIdx := strings.Index(string(s.Uppercase()), string(str.Uppercase()))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) LastIndexOfIgnoreCase(str XString) XInt {
	byteIdx := strings.LastIndex(string(s.Uppercase()), string(str.Uppercase()))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) IndexOfWithStartIndex(substr XString, startIndex XInt) XInt {
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= s.RuneCount() {
		return -1
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	searchRange := s[startByte:]
	byteIdx := strings.Index(string(searchRange), string(substr))
	if byteIdx == -1 {
		return -1
	}
	// 将 byte 索引转换为 rune 索引
	return startIndex + s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) LastIndexOfWithStartIndex(substr XString, startIndex XInt) XInt {
	maxIndex := s.RuneCount() - 1
	if startIndex > maxIndex {
		startIndex = maxIndex
	}
	if startIndex < 0 {
		return -1
	}
	endByte := s.runeIndexToByteIndex(startIndex + 1)
	searchRange := s[:endByte]
	byteIdx := strings.LastIndex(string(searchRange), string(substr))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) IndexOfWithStartIndexIgnoreCase(substr XString, startIndex XInt) XInt {
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= s.RuneCount() {
		return -1
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	searchRange := s[startByte:]
	lowerSearch := searchRange.Lowercase()
	lowerSubstr := substr.Lowercase()
	byteIdx := strings.Index(string(lowerSearch), string(lowerSubstr))
	if byteIdx == -1 {
		return -1
	}
	return startIndex + s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) LastIndexOfWithStartIndexIgnoreCase(substr XString, startIndex XInt) XInt {
	maxIndex := s.RuneCount() - 1
	if startIndex > maxIndex {
		startIndex = maxIndex
	}
	if startIndex < 0 {
		return -1
	}
	endByte := s.runeIndexToByteIndex(startIndex + 1)
	searchRange := s[:endByte]
	lowerSearch := searchRange.Lowercase()
	lowerSubstr := substr.Lowercase()
	byteIdx := strings.LastIndex(string(lowerSearch), string(lowerSubstr))
	if byteIdx == -1 {
		return -1
	}
	return s.byteIndexToRuneIndex(XInt(byteIdx))
}

func (s XString) IndexOfAny(chars []XRune) XInt {
	runes := []XRune(s)
	for i, currentChar := range runes {
		for _, targetChar := range chars {
			if currentChar == targetChar {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) LastIndexOfAny(chars []XRune) XInt {
	runes := []XRune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if currentChar == targetChar {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) IndexOfAnyIgnoreCase(chars []XRune) XInt {
	runes := []XRune(s)
	for i, currentChar := range runes {
		for _, targetChar := range chars {
			if unicode.ToLower(rune(currentChar)) == unicode.ToLower(rune(targetChar)) {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) LastIndexOfAnyIgnoreCase(chars []XRune) XInt {
	runes := []XRune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if unicode.ToLower(rune(currentChar)) == unicode.ToLower(rune(targetChar)) {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) IndexOfAnyWithStartIndex(chars []XRune, startIndex XInt) XInt {
	runes := []XRune(s)
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= XInt(len(runes)) {
		return -1
	}
	for i := int(startIndex); i < len(runes); i++ {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if currentChar == targetChar {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) LastIndexOfAnyWithStartIndex(chars []XRune, startIndex XInt) XInt {
	runes := []XRune(s)
	maxIndex := XInt(len(runes)) - 1
	if startIndex > maxIndex {
		startIndex = maxIndex
	}
	if startIndex < 0 {
		return -1
	}
	for i := int(startIndex); i >= 0; i-- {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if currentChar == targetChar {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) IndexOfAnyWithStartIndexIgnoreCase(chars []XRune, startIndex XInt) XInt {
	runes := []XRune(s)
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= XInt(len(runes)) {
		return -1
	}
	for i := int(startIndex); i < len(runes); i++ {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if unicode.ToLower(rune(currentChar)) == unicode.ToLower(rune(targetChar)) {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) LastIndexOfAnyWithStartIndexIgnoreCase(chars []XRune, startIndex XInt) XInt {
	runes := []XRune(s)
	maxIndex := XInt(len(runes)) - 1
	if startIndex > maxIndex {
		startIndex = maxIndex
	}
	if startIndex < 0 {
		return -1
	}
	for i := int(startIndex); i >= 0; i-- {
		currentChar := runes[i]
		for _, targetChar := range chars {
			if unicode.ToLower(rune(currentChar)) == unicode.ToLower(rune(targetChar)) {
				return XInt(i)
			}
		}
	}
	return -1
}

func (s XString) IndexOfAnyStrings(strs []XString) XInt {
	ret := s.FindAnyOf(strs, 0)
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) IndexOfAnyStringsIgnoreCase(strs []XString) XInt {
	ret := s.FindAnyOfIgnoreCase(strs, 0)
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) IndexOfAnyStringsWithStartIndex(strs []XString, startIndex XInt) XInt {
	ret := s.FindAnyOf(strs, startIndex)
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) IndexOfAnyStringsWithStartIndexIgnoreCase(strs []XString, startIndex XInt) XInt {
	ret := s.FindAnyOfIgnoreCase(strs, startIndex)
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) LastIndexOfAnyStrings(strs []XString) XInt {
	ret := s.FindLastAnyOf(strs, s.LastIndex())
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) LastIndexOfAnyStringsIgnoreCase(strs []XString) XInt {
	ret := s.FindLastAnyOfIgnoreCase(strs, s.LastIndex())
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) LastIndexOfAnyStringsWithStartIndex(strs []XString, startIndex XInt) XInt {
	ret := s.FindLastAnyOf(strs, startIndex)
	if ret == nil {
		return -1
	}
	return ret.Index
}
func (s XString) LastIndexOfAnyStringsWithStartIndexIgnoreCase(strs []XString, startIndex XInt) XInt {
	ret := s.FindLastAnyOfIgnoreCase(strs, startIndex)
	if ret == nil {
		return -1
	}
	return ret.Index
}

func (s XString) FindAnyOf(strs []XString, startIndex XInt) *struct {
	Index XInt
	Value XString
} {
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= s.RuneCount() {
		return nil
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	searchStr := s[startByte:]
	minIndex := XInt(-1)
	var matchedStr XString
	for _, sub := range strs {
		if sub == "" {
			continue
		}
		byteIdx := strings.Index(string(searchStr), string(sub))
		if byteIdx != -1 {
			runeOffset := s.byteIndexToRuneIndex(XInt(byteIdx))
			actualIdx := startIndex + runeOffset
			if minIndex == -1 || actualIdx < minIndex {
				minIndex = actualIdx
				matchedStr = sub
			}
		}
	}
	if minIndex == -1 {
		return nil
	}
	return &struct {
		Index XInt
		Value XString
	}{Index: minIndex, Value: matchedStr}
}

func (s XString) FindAnyOfIgnoreCase(strs []XString, startIndex XInt) *struct {
	Index XInt
	Value XString
} {
	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= s.RuneCount() {
		return nil
	}
	startByte := s.runeIndexToByteIndex(startIndex)
	searchStr := s[startByte:]
	searchStrLower := searchStr.Lowercase()
	minIndex := XInt(-1)
	var matchedStr XString
	for _, sub := range strs {
		if sub == "" {
			continue
		}
		subLower := sub.Lowercase()
		byteIdx := strings.Index(string(searchStrLower), string(subLower))
		if byteIdx != -1 {
			runeOffset := s.byteIndexToRuneIndex(XInt(byteIdx))
			actualIdx := startIndex + runeOffset
			if minIndex == -1 || actualIdx < minIndex {
				minIndex = actualIdx
				matchedStr = sub
			}
		}
	}
	if minIndex == -1 {
		return nil
	}
	return &struct {
		Index XInt
		Value XString
	}{Index: minIndex, Value: matchedStr}
}

func (s XString) FindLastAnyOf(strs []XString, startIndex XInt) *struct {
	Index XInt
	Value XString
} {
	runeLen := s.RuneCount()
	if startIndex >= runeLen {
		startIndex = runeLen - 1
	}
	if startIndex < 0 {
		return nil
	}
	endByte := s.runeIndexToByteIndex(startIndex + 1)
	searchStr := s[:endByte]
	maxIndex := XInt(-1)
	var matchedStr XString
	for _, sub := range strs {
		if sub == "" {
			continue
		}
		byteIdx := strings.LastIndex(string(searchStr), string(sub))
		if byteIdx != -1 {
			runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
			if runeIdx > maxIndex {
				maxIndex = runeIdx
				matchedStr = sub
			}
		}
	}
	if maxIndex == -1 {
		return nil
	}
	return &struct {
		Index XInt
		Value XString
	}{Index: maxIndex, Value: matchedStr}
}

func (s XString) FindLastAnyOfIgnoreCase(strs []XString, startIndex XInt) *struct {
	Index XInt
	Value XString
} {
	runeLen := s.RuneCount()
	if startIndex >= runeLen {
		startIndex = runeLen - 1
	}
	if startIndex < 0 {
		return nil
	}
	endByte := s.runeIndexToByteIndex(startIndex + 1)
	searchStr := s[:endByte]
	searchStrLower := searchStr.Lowercase()
	maxIndex := XInt(-1)
	var matchedStr XString
	for _, sub := range strs {
		if sub == "" {
			continue
		}
		subLower := sub.Lowercase()
		byteIdx := strings.LastIndex(string(searchStrLower), string(subLower))
		if byteIdx != -1 {
			runeIdx := s.byteIndexToRuneIndex(XInt(byteIdx))
			if runeIdx > maxIndex {
				maxIndex = runeIdx
				matchedStr = sub
			}
		}
	}
	if maxIndex == -1 {
		return nil
	}
	return &struct {
		Index XInt
		Value XString
	}{Index: maxIndex, Value: matchedStr}
}

func (s XString) Contains(sub XString) XBool {
	return XBool(strings.Contains(string(s), string(sub)))
}

func (s XString) ContainsIgnoreCase(sub XString) XBool {
	return XBool(strings.Contains(string(s.Uppercase()), string(sub.Uppercase())))
}

func (s XString) ContainsRegex(pattern XString) XBool {
	re := regexp.MustCompile(string(pattern))
	return XBool(re.MatchString(string(s)))
}

func (s XString) Matches(pattern XString) XBool {
	re := regexp.MustCompile(string(pattern))
	matches := re.MatchString(string(s)) && XString(re.FindString(string(s))) == s
	return XBool(matches)
}

func (s XString) ReplaceRegex(pattern XString, replacement XString) XString {
	re := regexp.MustCompile(string(pattern))
	return XString(re.ReplaceAllString(string(s), string(replacement)))
}

func (s XString) ReplaceRegexFunc(pattern XString, transform func(XString) XString) XString {
	re := regexp.MustCompile(string(pattern))
	retStr := re.ReplaceAllStringFunc(string(s), func(it string) string {
		return string(transform(XString(it)))
	})
	return XString(retStr)
}

func (s XString) ReplaceRegexFirst(pattern XString, replacement XString) XString {
	re := regexp.MustCompile(string(pattern))
	return XString(re.ReplaceAllString(string(s), string(replacement)))
}

func (s XString) Split(delimiters ...XString) XList[XString] {
	if len(delimiters) == 0 {
		return []XString{s}
	}
	// 构建正则表达式模式：将各个分隔符进行转义并用|连接
	escapedDelimiters := make([]string, len(delimiters))
	for i, d := range delimiters {
		escapedDelimiters[i] = regexp.QuoteMeta(string(d))
	}
	pattern := strings.Join(escapedDelimiters, "|")
	var regex = regexp.MustCompile(pattern)
	return s.splitByRegex(regex, 0)
}

func (s XString) SplitIgnoreCase(delimiters ...XString) XList[XString] {
	if len(delimiters) == 0 {
		return []XString{s}
	}
	// 构建正则表达式模式：将各个分隔符进行转义并用|连接
	escapedDelimiters := make([]string, len(delimiters))
	for i, d := range delimiters {
		escapedDelimiters[i] = regexp.QuoteMeta(string(d))
	}
	pattern := strings.Join(escapedDelimiters, "|")
	var regex = regexp.MustCompile("(?i)" + pattern)
	return s.splitByRegex(regex, 0)
}

func (s XString) SplitWithLimit(limit XInt, delimiters ...XString) XList[XString] {
	if len(delimiters) == 0 {
		return []XString{s}
	}
	// 构建正则表达式模式：将各个分隔符进行转义并用|连接
	escapedDelimiters := make([]string, len(delimiters))
	for i, d := range delimiters {
		escapedDelimiters[i] = regexp.QuoteMeta(string(d))
	}
	pattern := strings.Join(escapedDelimiters, "|")
	var regex = regexp.MustCompile(pattern)
	return s.splitByRegex(regex, limit)
}

func (s XString) SplitWithLimitIgnoreCase(limit XInt, delimiters ...XString) XList[XString] {
	if len(delimiters) == 0 {
		return []XString{s}
	}
	// 构建正则表达式模式：将各个分隔符进行转义并用|连接
	escapedDelimiters := make([]string, len(delimiters))
	for i, d := range delimiters {
		escapedDelimiters[i] = regexp.QuoteMeta(string(d))
	}
	pattern := strings.Join(escapedDelimiters, "|")
	var regex = regexp.MustCompile("(?i)" + pattern)
	return s.splitByRegex(regex, limit)
}

func (s XString) SplitRegex(regex *regexp.Regexp) XList[XString] {
	return s.splitByRegex(regex, 0)
}

func (s XString) SplitRegexWithLimit(regex *regexp.Regexp, limit XInt) XList[XString] {
	return s.splitByRegex(regex, limit)
}

func (s XString) Lines() XList[XString] {
	if s == "" {
		return []XString{}
	}
	// 使用正则表达式处理各种换行符
	regex := regexp.MustCompile(`\r\n|\n|\r`)
	parts := regex.Split(string(s), -1)
	// 移除末尾的空行
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	ret := make(XList[XString], len(parts))
	for i, part := range parts {
		ret[i] = XString(part)
	}
	return ret
}

func (s XString) ToBoolean() XBool {
	if strings.ToLower(string(s)) == "true" {
		return true
	}
	return false
}

func (s XString) ToBooleanOrNil() *XBool {
	if strings.ToLower(string(s)) == "true" {
		return new(XBool(true))
	}
	if strings.ToLower(string(s)) == "false" {
		return new(XBool(false))
	}
	return nil
}

func (s XString) ToByte() XByte {
	v, e := strconv.ParseUint(string(s), 10, 8)
	if e != nil {
		panic(e)
	}
	return XByte(v)
}

func (s XString) ToByteOrNil() *XByte {
	v, e := strconv.ParseUint(string(s), 10, 8)
	if e != nil {
		return nil
	}
	return new(XByte(v))
}

func (s XString) ToByteWithRadix(radix XInt) XByte {
	v, e := strconv.ParseUint(string(s), int(radix), 8)
	if e != nil {
		panic(e)
	}
	return XByte(v)
}

func (s XString) ToByteWithRadixOrNil(radix XInt) *XByte {
	v, e := strconv.ParseUint(string(s), int(radix), 8)
	if e != nil {
		return nil
	}
	return new(XByte(v))
}

func (s XString) ToInt8() XInt8 {
	v, e := strconv.ParseInt(string(s), 10, 8)
	if e != nil {
		panic(e)
	}
	return XInt8(v)
}

func (s XString) ToUInt8() XUint8 {
	v, e := strconv.ParseUint(string(s), 10, 8)
	if e != nil {
		panic(e)
	}
	return XUint8(v)
}

func (s XString) ToInt8OrNil() *XInt8 {
	v, e := strconv.ParseInt(string(s), 10, 8)
	if e != nil {
		return nil
	}
	return new(XInt8(v))
}

func (s XString) ToUInt8OrNil() *XUint8 {
	v, e := strconv.ParseUint(string(s), 10, 8)
	if e != nil {
		return nil
	}
	return new(XUint8(v))
}

func (s XString) ToInt8WithRadix(radix XInt) XInt8 {
	v, e := strconv.ParseInt(string(s), int(radix), 8)
	if e != nil {
		panic(e)
	}
	return XInt8(v)
}

func (s XString) ToUInt8WithRadix(radix XInt) XUint8 {
	v, e := strconv.ParseUint(string(s), int(radix), 8)
	if e != nil {
		panic(e)
	}
	return XUint8(v)
}

func (s XString) ToInt8WithRadixOrNil(radix XInt) *XInt8 {
	v, e := strconv.ParseInt(string(s), int(radix), 8)
	if e != nil {
		return nil
	}
	return new(XInt8(v))
}

func (s XString) ToUInt8WithRadixOrNil(radix XInt) *XUint8 {
	v, e := strconv.ParseUint(string(s), int(radix), 8)
	if e != nil {
		return nil
	}
	return new(XUint8(v))
}

func (s XString) ToInt16() XInt16 {
	v, e := strconv.ParseInt(string(s), 10, 16)
	if e != nil {
		panic(e)
	}
	return XInt16(v)
}

func (s XString) ToUInt16() XUint16 {
	v, e := strconv.ParseUint(string(s), 10, 16)
	if e != nil {
		panic(e)
	}
	return XUint16(v)
}

func (s XString) ToInt16OrNil() *XInt16 {
	v, e := strconv.ParseInt(string(s), 10, 16)
	if e != nil {
		return nil
	}
	return new(XInt16(v))
}

func (s XString) ToUInt16OrNil() *XUint16 {
	v, e := strconv.ParseUint(string(s), 10, 16)
	if e != nil {
		return nil
	}
	return new(XUint16(v))
}

func (s XString) ToInt16WithRadix(radix XInt) XInt16 {
	v, e := strconv.ParseInt(string(s), int(radix), 16)
	if e != nil {
		panic(e)
	}
	return XInt16(v)
}

func (s XString) ToUInt16WithRadix(radix XInt) XUint16 {
	v, e := strconv.ParseUint(string(s), int(radix), 16)
	if e != nil {
		panic(e)
	}
	return XUint16(v)
}

func (s XString) ToInt16WithRadixOrNil(radix XInt) *XInt16 {
	v, e := strconv.ParseInt(string(s), int(radix), 16)
	if e != nil {
		return nil
	}
	return new(XInt16(v))
}

func (s XString) ToUInt16WithRadixOrNil(radix XInt) *XUint16 {
	v, e := strconv.ParseUint(string(s), int(radix), 16)
	if e != nil {
		return nil
	}
	return new(XUint16(v))
}

func (s XString) ToInt32() XInt32 {
	v, e := strconv.ParseInt(string(s), 10, 32)
	if e != nil {
		panic(e)
	}
	return XInt32(v)
}

func (s XString) ToUInt32() XUint32 {
	v, e := strconv.ParseUint(string(s), 10, 32)
	if e != nil {
		panic(e)
	}
	return XUint32(v)
}

func (s XString) ToInt32OrNil() *XInt32 {
	v, e := strconv.ParseInt(string(s), 10, 32)
	if e != nil {
		return nil
	}
	return new(XInt32(v))
}

func (s XString) ToUInt32OrNil() *XUint32 {
	v, e := strconv.ParseUint(string(s), 10, 32)
	if e != nil {
		return nil
	}
	return new(XUint32(v))
}

func (s XString) ToInt32WithRadix(radix XInt) XInt32 {
	v, e := strconv.ParseInt(string(s), int(radix), 32)
	if e != nil {
		panic(e)
	}
	return XInt32(v)
}

func (s XString) ToUInt32WithRadix(radix XInt) XUint32 {
	v, e := strconv.ParseUint(string(s), int(radix), 32)
	if e != nil {
		panic(e)
	}
	return XUint32(v)
}

func (s XString) ToInt32WithRadixOrNil(radix XInt) *XInt32 {
	v, e := strconv.ParseInt(string(s), int(radix), 32)
	if e != nil {
		return nil
	}
	return new(XInt32(v))
}

func (s XString) ToInt() XInt {
	v, e := strconv.ParseInt(string(s), 10, 32)
	if e != nil {
		panic(e)
	}
	return XInt(v)
}

func (s XString) ToIntOrNil() *XInt {
	v, e := strconv.ParseInt(string(s), 10, 32)
	if e != nil {
		return nil
	}
	return new(XInt(v))
}

func (s XString) ToIntWithRadix(radix XInt) XInt {
	v, e := strconv.ParseInt(string(s), int(radix), 32)
	if e != nil {
		panic(e)
	}
	return XInt(v)
}

func (s XString) ToIntWithRadixOrNil(radix XInt) *XInt {
	v, e := strconv.ParseInt(string(s), int(radix), 32)
	if e != nil {
		return nil
	}
	return new(XInt(v))
}

func (s XString) ToInt64() XInt64 {
	v, e := strconv.ParseInt(string(s), 10, 64)
	if e != nil {
		panic(e)
	}
	return XInt64(v)
}

func (s XString) ToInt64OrNil() *XInt64 {
	v, e := strconv.ParseInt(string(s), 10, 64)
	if e != nil {
		return nil
	}
	return new(XInt64(v))
}

func (s XString) ToInt64WithRadix(radix XInt) XInt64 {
	v, e := strconv.ParseInt(string(s), int(radix), 64)
	if e != nil {
		panic(e)
	}
	return XInt64(v)
}

func (s XString) ToUInt64WithRadix(radix XInt) XUint64 {
	v, e := strconv.ParseUint(string(s), int(radix), 64)
	if e != nil {
		panic(e)
	}
	return XUint64(v)
}

func (s XString) ToInt64WithRadixOrNil(radix XInt) *XInt64 {
	v, e := strconv.ParseInt(string(s), int(radix), 64)
	if e != nil {
		return nil
	}
	return new(XInt64(v))
}

func (s XString) ToUInt64WithRadixOrNil(radix XInt) *XUint64 {
	v, e := strconv.ParseUint(string(s), int(radix), 64)
	if e != nil {
		return nil
	}
	return new(XUint64(v))
}

func (s XString) ToFloat32() XFloat32 {
	v, e := strconv.ParseFloat(string(s), 32)
	if e != nil {
		panic(e)
	}
	return XFloat32(v)
}

func (s XString) ToFloat32OrNil() *XFloat32 {
	v, e := strconv.ParseFloat(string(s), 32)
	if e != nil {
		return nil
	}
	return new(XFloat32(v))
}

func (s XString) ToFloat64() XFloat64 {
	v, e := strconv.ParseFloat(string(s), 64)
	if e != nil {
		panic(e)
	}
	return XFloat64(v)
}

func (s XString) ToFloat64OrNil() *XFloat64 {
	v, e := strconv.ParseFloat(string(s), 64)
	if e != nil {
		return nil
	}
	return new(XFloat64(v))
}

func (s XString) ToRuneArray() XList[XRune] {
	return []XRune(s)
}

func (s XString) ToRuneArrayWithRange(startIndex XInt, endIndex XInt) XList[XRune] {
	if startIndex < 0 {
		startIndex = 0
	}
	runeLen := s.RuneCount()
	if endIndex > runeLen {
		endIndex = runeLen
	}
	if startIndex > endIndex {
		startIndex = endIndex
	}
	runes := []XRune(s)
	return runes[startIndex:endIndex]
}

func (s XString) ToByteArray() XList[XByte] {
	return []XByte(s)
}

func (s XString) splitByRegex(regex *regexp.Regexp, limit XInt) XList[XString] {
	if limit == 1 {
		return []XString{s}
	}
	// 查找所有匹配的分隔符位置
	matches := regex.FindAllStringIndex(string(s), -1)
	if matches == nil {
		return []XString{s}
	}
	result := make([]XString, 0)
	lastIndex := 0
	for i, match := range matches {
		// 添加分隔符前的子串
		result = append(result, s[lastIndex:match[0]])
		// 检查limit限制
		if limit > 0 && XInt(len(result)) == limit-1 {
			// 将剩余部分作为最后一个元素
			lastIndex = match[1]
			// 继续添加后续所有内容
			remaining := s[lastIndex:]
			result = append(result, remaining)
			return result
		}
		lastIndex = match[1]
		// 如果这是最后一个匹配，添加剩余部分
		if i == len(matches)-1 && lastIndex < len(s) {
			result = append(result, s[lastIndex:])
		}
	}
	// 处理没有匹配的情况
	if len(matches) == 0 {
		return []XString{s}
	}
	// 根据limit处理结果
	if limit < 0 {
		// 移除末尾的空字符串
		for len(result) > 0 && result[len(result)-1] == "" {
			result = result[:len(result)-1]
		}
	} else if limit == 0 {
		// 移除所有末尾的空字符串
		for len(result) > 0 && result[len(result)-1] == "" {
			result = result[:len(result)-1]
		}
	}
	return result
}

func (s XString) runeIndexToByteIndex(runeIndex XInt) XInt {
	if runeIndex <= 0 {
		return 0
	}
	runeCount := s.RuneCount()
	if runeCount == 0 {
		return 0
	}
	if runeIndex >= runeCount {
		return XInt(len(s))
	}
	byteIndex := 0
	for i := XInt(0); i < runeIndex; i++ {
		_, size := utf8.DecodeRuneInString(string(s[byteIndex:]))
		if size == 0 {
			break
		}
		byteIndex += size
	}
	return XInt(byteIndex)
}

func (s XString) byteIndexToRuneIndex(byteIndex XInt) XInt {
	if byteIndex <= 0 {
		return 0
	}
	strLen := XInt(len(s))
	if strLen == 0 {
		return 0
	}
	if byteIndex >= strLen {
		return s.RuneCount()
	}
	runeIndex := 0
	for i := 0; i < int(byteIndex); {
		if i >= len(s) {
			break
		}
		_, size := utf8.DecodeRuneInString(string(s[i:]))
		if size == 0 {
			break
		}
		i += size
		runeIndex++
	}
	return XInt(runeIndex)
}
