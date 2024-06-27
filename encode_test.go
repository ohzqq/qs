//MIT License

//Copyright (c) 2020 Son Huynh

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package qs

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type basicVal struct {
	String     string     `query:"string"`
	Bool       bool       `query:"bool"`
	Int        int        `query:"int"`
	Int8       int8       `query:"int8"`
	Int16      int16      `query:"int16"`
	Int32      int32      `query:"int32"`
	Int64      int64      `query:"int64"`
	Uint       uint       `query:"uint"`
	Uint8      uint8      `query:"uint8"`
	Uint16     uint16     `query:"uint16"`
	Uint32     uint32     `query:"uint32"`
	Uint64     uint64     `query:"uint64"`
	Uintptr    uintptr    `query:"uintptr"`
	Float32    float32    `query:"float32"`
	Float64    float64    `query:"float64"`
	Complex64  complex64  `query:"complex64"`
	Complex128 complex128 `query:"complex128"`
	Time       time.Time  `query:"time"`
}

type basicValWithOmit struct {
	String     string     `query:"string,omitempty"`
	Bool       bool       `query:"bool,omitempty"`
	Int        int        `query:"int,omitempty"`
	Int8       int8       `query:"int8,omitempty"`
	Int16      int16      `query:"int16,omitempty"`
	Int32      int32      `query:"int32,omitempty"`
	Int64      int64      `query:"int64,omitempty"`
	Uint       uint       `query:"uint,omitempty"`
	Uint8      uint8      `query:"uint8,omitempty"`
	Uint16     uint16     `query:"uint16,omitempty"`
	Uint32     uint32     `query:"uint32,omitempty"`
	Uint64     uint64     `query:"uint64,omitempty"`
	Float32    float32    `query:"float32,omitempty"`
	Float64    float64    `query:"float64,omitempty"`
	Complex64  complex64  `query:"complex64,omitempty"`
	Complex128 complex128 `query:"complex128,omitempty"`
	Time       time.Time  `query:"time,omitempty"`
}

type basicPtr struct {
	String     *string     `query:"string"`
	Bool       *bool       `query:"bool"`
	Int        *int        `query:"int"`
	Int8       *int8       `query:"int8"`
	Int16      *int16      `query:"int16"`
	Int32      *int32      `query:"int32"`
	Int64      *int64      `query:"int64"`
	Uint       *uint       `query:"uint"`
	Uint8      *uint8      `query:"uint8"`
	Uint16     *uint16     `query:"uint16"`
	Uint32     *uint32     `query:"uint32"`
	Uint64     *uint64     `query:"uint64"`
	UinPtr     *uintptr    `query:"uintptr"`
	Float32    *float32    `query:"float32"`
	Float64    *float64    `query:"float64"`
	Complex64  *complex64  `query:"complex64"`
	Complex128 *complex128 `query:"complex128"`
	Time       *time.Time  `query:"time"`
}

type basicPtrWithOmit struct {
	String     *string     `query:"string,omitempty"`
	Bool       *bool       `query:"bool,omitempty"`
	Int        *int        `query:"int,omitempty"`
	Int8       *int8       `query:"int8,omitempty"`
	Int16      *int16      `query:"int16,omitempty"`
	Int32      *int32      `query:"int32,omitempty"`
	Int64      *int64      `query:"int64,omitempty"`
	Uint       *uint       `query:"uint,omitempty"`
	Uint8      *uint8      `query:"uint8,omitempty"`
	Uint16     *uint16     `query:"uint16,omitempty"`
	Uint32     *uint32     `query:"uint32,omitempty"`
	Uint64     *uint64     `query:"uint64,omitempty"`
	Float32    *float32    `query:"float32,omitempty"`
	Float64    *float64    `query:"float64,omitempty"`
	Complex64  *complex64  `query:"complex64,omitempty"`
	Complex128 *complex128 `query:"complex128,omitempty"`
	Time       *time.Time  `query:"time,omitempty"`
}

func TestIgnore(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	v := struct {
		anonymous string
		Test      string `query:"-"`
	}{}

	values, err := encoder.Values(v)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	assert.Equal(t, url.Values{}, values)
}

func TestWithTagAlias(t *testing.T) {
	test := assert.New(t)

	alias := `go`
	opt := WithTagAlias(alias)
	test.NotNil(opt)

	encoder := NewEncoder(opt)
	test.Equal(alias, encoder.tagAlias)
}

func TestGetTag(t *testing.T) {
	test := assert.New(t)

	e := NewEncoder().dataPool.Get().(*encoder)

	s := struct {
		A string `query:"abc"`
	}{}

	field := reflect.TypeOf(s).Field(0)
	e.getTagNameAndOpts(field)

	test.Len(e.tags, 1)
	test.Equal("abc", string(e.tags[0]))
}

func TestGetTag2(t *testing.T) {
	test := assert.New(t)

	e := NewEncoder().dataPool.Get().(*encoder)

	s := struct {
		ABC string
	}{}

	field := reflect.TypeOf(s).Field(0)
	e.getTagNameAndOpts(field)

	test.Len(e.tags, 1)
	test.Equal("ABC", string(e.tags[0]))
}

func TestGetTag3(t *testing.T) {
	test := assert.New(t)

	e := NewEncoder().dataPool.Get().(*encoder)

	s := struct {
		ABC string `query:",omitempty"`
	}{}

	field := reflect.TypeOf(s).Field(0)
	e.getTagNameAndOpts(field)

	test.Len(e.tags, 2)
	test.Equal("ABC", string(e.tags[0]))
	test.Equal("omitempty", string(e.tags[1]))
}

func TestEncodeInvalidValue(t *testing.T) {
	t.Parallel()

	var ptr *string

	testCases := []struct {
		name  string
		input interface{}
	}{
		{
			name:  "string",
			input: "abc",
		},
		{
			name:  "nil pointer",
			input: ptr,
		},
		{
			name:  "nil",
			input: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			test := assert.New(t)

			encoder := NewEncoder()
			_, err := encoder.Values(testCase.input)
			test.Error(err)

			values := make(url.Values)
			err = encoder.Encode(testCase.input, values)
		})
	}

}

func TestEncodeBasicVal(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := basicVal{
		String:     "abc",
		Bool:       true,
		Int:        12,
		Int8:       int8(8),
		Int16:      int16(16),
		Int32:      int32(32),
		Int64:      int64(64),
		Uint:       24,
		Uint8:      uint8(8),
		Uint16:     uint16(16),
		Uint32:     uint32(32),
		Uint64:     uint64(64),
		Uintptr:    uintptr(72),
		Float32:    float32(0.1234),
		Float64:    1.2345,
		Complex64:  complex64(64),
		Complex128: complex128(128),
		Time:       tm,
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"string":     []string{"abc"},
		"bool":       []string{"true"},
		"int":        []string{"12"},
		"int8":       []string{"8"},
		"int16":      []string{"16"},
		"int32":      []string{"32"},
		"int64":      []string{"64"},
		"uint":       []string{"24"},
		"uint8":      []string{"8"},
		"uint16":     []string{"16"},
		"uint32":     []string{"32"},
		"uint64":     []string{"64"},
		"uintptr":    []string{"72"},
		"float32":    []string{"0.1234"},
		"float64":    []string{"1.2345"},
		"complex64":  []string{complex128ToStr(complex128(64))},
		"complex128": []string{complex128ToStr(complex128(128))},
		"time":       []string{tm.Format(time.RFC3339)},
	}
	assert.Equal(t, expected, values)
}

func TestEncodeBasicPtr(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := basicPtr{
		String:     withStr("abc"),
		Bool:       withBool(true),
		Int:        withInt(12),
		Int8:       withInt8(int8(8)),
		Int16:      withInt16(int16(16)),
		Int32:      withInt32(int32(32)),
		Int64:      withInt64(int64(64)),
		Uint:       withUint(uint(24)),
		Uint8:      withUint8(uint8(8)),
		Uint16:     withUint16(uint16(16)),
		Uint32:     withUint32(uint32(32)),
		Uint64:     withUint64(uint64(64)),
		UinPtr:     withUintPtr(uintptr(72)),
		Float32:    withFloat32(float32(0.1234)),
		Float64:    withFloat64(1.2345),
		Complex64:  withComplex64(complex64(64)),
		Complex128: withComplex128(complex128(128)),
		Time:       withTime(tm),
	}
	actualValues1, err := encoder.Values(s)
	test.NoError(err)

	actualValues2 := make(url.Values)
	err = encoder.Encode(&s, actualValues2)
	test.NoError(err)

	expected := url.Values{
		"string":     []string{"abc"},
		"bool":       []string{"true"},
		"int":        []string{"12"},
		"int8":       []string{"8"},
		"int16":      []string{"16"},
		"int32":      []string{"32"},
		"int64":      []string{"64"},
		"uint":       []string{"24"},
		"uint8":      []string{"8"},
		"uint16":     []string{"16"},
		"uint32":     []string{"32"},
		"uint64":     []string{"64"},
		"uintptr":    []string{"72"},
		"float32":    []string{"0.1234"},
		"float64":    []string{"1.2345"},
		"complex64":  []string{complex128ToStr(complex128(64))},
		"complex128": []string{complex128ToStr(complex128(128))},
		"time":       []string{tm.Format(time.RFC3339)},
	}

	test.Equal(expected, actualValues1)
	test.Equal(expected, actualValues2)
}

func TestZeroVal(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	values, err := encoder.Values(basicVal{})
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"string":     []string{""},
		"bool":       []string{"false"},
		"int":        []string{"0"},
		"int8":       []string{"0"},
		"int16":      []string{"0"},
		"int32":      []string{"0"},
		"int64":      []string{"0"},
		"uint":       []string{"0"},
		"uint8":      []string{"0"},
		"uint16":     []string{"0"},
		"uint32":     []string{"0"},
		"uint64":     []string{"0"},
		"uintptr":    []string{"0"},
		"float32":    []string{"0"},
		"float64":    []string{"0"},
		"complex64":  []string{complexZeroValStr()},
		"complex128": []string{complexZeroValStr()},
		"time":       []string{time.Time{}.Format(time.RFC3339)},
	}
	test.Equal(expected, values)
}

func TestZeroPtr(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	values, err := encoder.Values(basicPtr{})
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"string":     []string{""},
		"bool":       []string{""},
		"int":        []string{""},
		"int8":       []string{""},
		"int16":      []string{""},
		"int32":      []string{""},
		"int64":      []string{""},
		"uint":       []string{""},
		"uint8":      []string{""},
		"uint16":     []string{""},
		"uint32":     []string{""},
		"uint64":     []string{""},
		"uintptr":    []string{""},
		"float32":    []string{""},
		"float64":    []string{""},
		"complex64":  []string{""},
		"complex128": []string{""},
		"time":       []string{""},
	}
	assert.Equal(t, expected, values)
}

func TestOmitZeroVal(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	values, err := encoder.Values(basicValWithOmit{})
	test.NoError(err)
	test.Equal(url.Values{}, values)
}

func TestOmitZeroPtr(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	values, err := encoder.Values(basicPtrWithOmit{})
	test.NoError(err)
	test.Equal(url.Values{}, values)
}

func TestIgnoreEmptySlice(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		A []string  `query:"a"`
		B []string  `query:"b"`
		C *[]string `query:"c"`
	}{
		A: nil,
		B: []string{},
		C: nil,
	}

	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	test.Equal(url.Values{}, values)
}

func TestSliceValWithBasicVal(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		StringList []string `query:"str_list"`
		BoolList   []bool   `query:"bool_list"`
		IntList    []int    `query:"int_list"`
	}{
		StringList: []string{"", "a", "b", "c"},
		BoolList:   []bool{true, false},
		IntList:    []int{0, 1, 2, 3},
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list":  []string{"", "a", "b", "c"},
		"bool_list": []string{"true", "false"},
		"int_list":  []string{"0", "1", "2", "3"},
	}
	assert.Equal(t, expected, values)
}

func TestSliceValWithBasicPtr(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		StringList []*string `query:"str_list"`
		BoolList   []*bool   `query:"bool_list"`
		IntList    []*int    `query:"int_list"`
	}{
		StringList: []*string{withStr(""), withStr("a"), withStr("b"), withStr("c")},
		BoolList:   []*bool{withBool(true), withBool(false)},
		IntList:    []*int{withInt(0), withInt(1), withInt(2), withInt(3)},
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list":  []string{"", "a", "b", "c"},
		"bool_list": []string{"true", "false"},
		"int_list":  []string{"0", "1", "2", "3"},
	}
	assert.Equal(t, expected, values)
}

func TestSlicePtrWithBasicVal(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	strList := []string{"", "a", "b", "c"}
	boolList := []bool{true, false}
	intList := []int{0, 1, 2, 3}

	s := struct {
		StringList *[]string `query:"str_list"`
		BoolList   *[]bool   `query:"bool_list"`
		IntList    *[]int    `query:"int_list"`
	}{
		StringList: &strList,
		BoolList:   &boolList,
		IntList:    &intList,
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list":  []string{"", "a", "b", "c"},
		"bool_list": []string{"true", "false"},
		"int_list":  []string{"0", "1", "2", "3"},
	}
	assert.Equal(t, expected, values)
}

func TestSlicePtrWithBasicPtr(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	strList := []*string{withStr(""), withStr("a"), withStr("b"), withStr("c")}
	boolList := []*bool{withBool(true), withBool(false)}
	intList := []*int{withInt(0), withInt(1), withInt(2), withInt(3)}

	s := struct {
		StringList *[]*string `query:"str_list"`
		BoolList   *[]*bool   `query:"bool_list"`
		IntList    *[]*int    `query:"int_list"`
	}{
		StringList: &strList,
		BoolList:   &boolList,
		IntList:    &intList,
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list":  []string{"", "a", "b", "c"},
		"bool_list": []string{"true", "false"},
		"int_list":  []string{"0", "1", "2", "3"},
	}
	assert.Equal(t, expected, values)
}

func TestTimeFormat(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	times := struct {
		Rfc3339    time.Time  `query:"default_fmt"`
		Second     time.Time  `query:"default_second,second"`
		Millis     time.Time  `query:"default_millis,millis"`
		Rfc3339Ptr *time.Time `query:"default_fmt_ptr"`
		SecondPtr  *time.Time `query:"default_second_ptr,second"`
		MillisPtr  *time.Time `query:"default_millis_ptr,millis"`
	}{
		tm,
		tm,
		tm,
		&tm,
		&tm,
		&tm,
	}
	values, err := encoder.Values(times)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"default_fmt":        []string{"1970-01-01T00:10:00Z"},
		"default_second":     []string{"600"},
		"default_millis":     []string{"600000"},
		"default_fmt_ptr":    []string{"1970-01-01T00:10:00Z"},
		"default_second_ptr": []string{"600"},
		"default_millis_ptr": []string{"600000"},
	}
	assert.Equal(t, expected, values)
}

func TestBoolFormat(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		Bool1   bool  `query:"bool_1,int"`
		Bool2   bool  `query:"bool_2,int"`
		NilBool *bool `query:",omitempty"`
	}{
		Bool2: true,
	}

	values, err := encoder.Values(&s)
	test.NoError(err)

	expected := url.Values{
		"bool_1": []string{"0"},
		"bool_2": []string{"1"},
	}
	test.Equal(expected, values)
}

func TestArrayFormat_Comma(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := struct {
		EmptyList  []string     `query:"empty_list,comma"`
		StringList []string     `query:"str_list,comma"`
		Times      []*time.Time `query:"times,comma"`
	}{
		StringList: []string{"a", "b", "c"},
		Times:      []*time.Time{&tm, nil},
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"empty_list": []string{""},
		"str_list":   []string{"a,b,c"},
		"times":      []string{tm.Format(time.RFC3339)},
	}
	assert.Equal(t, expected, values)
}

func TestArrayFormat_Repeat(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := struct {
		StringList []string     `query:"str_list"`
		Times      []*time.Time `query:"times"`
	}{
		StringList: []string{"a", "b", "c"},
		Times:      []*time.Time{&tm, nil},
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list": []string{"a", "b", "c"},
		"times":    []string{tm.Format(time.RFC3339)},
	}
	assert.Equal(t, expected, values)
}

func TestArrayFormat_Bracket(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := struct {
		StringList []string     `query:"str_list,bracket"`
		Times      []*time.Time `query:"times,bracket"`
	}{
		StringList: []string{"a", "b", "c"},
		Times:      []*time.Time{&tm, nil},
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list[]": []string{"a", "b", "c"},
		"times[]":    []string{tm.Format(time.RFC3339)},
	}
	assert.Equal(t, expected, values)
}

func TestArrayFormat_Index(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0).UTC()

	s := struct {
		StringList []string     `query:"str_list,index"`
		Times      []*time.Time `query:"times,index"`
		NilSlice   *[]int       `query:",omitempty"`
	}{
		StringList: []string{"a", "b", "c"},
		Times:      []*time.Time{&tm, nil},
		NilSlice:   nil,
	}
	values, err := encoder.Values(s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"str_list[0]": []string{"a"},
		"str_list[1]": []string{"b"},
		"str_list[2]": []string{"c"},
		"times[0]":    []string{tm.Format(time.RFC3339)},
	}
	assert.Equal(t, expected, values)
}

func TestNestedStruct(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(600, 0)

	type newTime time.Time

	type Nested struct {
		Time   time.Time `query:"time,second"`
		Name   *string   `query:"name,omitempty"`
		NewStr newTime   `query:"new_time,omitempty"`
	}

	s := struct {
		Nested           Nested   `query:"nested"`
		NestedOmitNilPtr *Nested  `query:"nested_omit_nil_ptr,omitempty"`
		NestedNilPtr     *Nested  `query:"nested_ptr"`
		NestedPtr        *Nested  `query:"nested_ptr"`
		NestedList       []Nested `query:"nest_list,index"`
	}{
		Nested: Nested{
			Time: tm,
		},
		NestedPtr: &Nested{
			Time: tm,
		},
		NestedList: []Nested{
			{
				Time: tm,
				Name: withStr("abc"),
			},
			{
				Time: tm,
				Name: withStr("def"),
			},
		},
	}

	values, err := encoder.Values(&s)
	if err != nil {
		test.FailNow(err.Error())
		return
	}
	expected := url.Values{
		"nested[time]":       []string{"600"},
		"nested_ptr[time]":   []string{"600"},
		"nested_ptr":         []string{""},
		"nest_list[0][time]": []string{"600"},
		"nest_list[0][name]": []string{"abc"},
		"nest_list[1][time]": []string{"600"},
		"nest_list[1][name]": []string{"def"},
	}
	assert.Equal(t, expected, values)
}

func TestEncodeInterface(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := &struct {
		String     interface{} `query:"string"`
		Bool       interface{} `query:"bool"`
		Int        interface{} `query:"int"`
		EmptyFloat interface{} `query:"float,omitempty"`
		NilPtr     interface{} `query:"nil_ptr"`
		OmitNilPtr interface{} `query:"omit_nil_ptr,omitempty"`
	}{
		String: "abc",
		Bool:   true,
		Int:    withInt(5),
		NilPtr: nil,
	}

	values, err := encoder.Values(&s)
	test.NoError(err)

	expected := url.Values{
		"string":  []string{"abc"},
		"bool":    []string{"true"},
		"int":     []string{"5"},
		"nil_ptr": []string{""},
	}
	test.Equal(expected, values)
}

func TestEncodeMap(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		Map       map[string]bool    `query:"map,int,omitempty"`
		PtrMap    map[string]*bool   `query:"ptr_map"`
		NilMap    map[string]int     `query:"nil_map"`
		NilPtrMap *map[string]int    `query:"nil_ptr_map"`
		EmptyMap  map[*string]string `query:"empty_map"`
	}{
		Map: map[string]bool{
			"abc": true,
		},
		PtrMap: map[string]*bool{
			"xyz": withBool(false),
		},
		EmptyMap: make(map[*string]string),
	}
	values, err := encoder.Values(s)
	test.NoError(err)

	expected := url.Values{
		"map[abc]":     []string{"true"},
		"ptr_map[xyz]": []string{"false"},
	}
	test.Equal(expected, values)
}

type Timestamp struct {
	time.Time
}

func (t Timestamp) EncodeParam() (string, error) {
	return t.Format(time.RFC3339), nil
}

func (t Timestamp) IsZero() bool {
	return t.Time.IsZero()
}

type TimestampPtr struct {
	time.Time
}

func (t *TimestampPtr) EncodeParam() (string, error) {
	return t.Format(time.RFC3339), nil
}

func (t *TimestampPtr) IsZero() bool {
	return t.Time.IsZero()
}

func TestEncodeCustomType(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(0, 0).UTC()

	var zeroPtrTs *TimestampPtr
	s := struct {
		OmitTimestamp    Timestamp     `query:"zero_ts,omitempty"`
		ZeroTimestamp    Timestamp     `query:"zero_ts"`
		Timestamp        Timestamp     `query:"ts"`
		InterfaceTs      interface{}   `query:"interface_ts"`
		ZeroInterfaceTs  interface{}   `query:"zero_interface_ts,omitempty"`
		OmitPtrTimestamp *TimestampPtr `query:"omit_ptr_ts,omitempty"`
		NilPtrTimestamp  *TimestampPtr `query:"zero_ptr_ts"`
		TimestampPtr     *TimestampPtr `query:"timestamp_ptr"`
		TsList           []Timestamp   `query:"ts_list"`
		TsCommaList      []Timestamp   `query:"ts_comma_list,comma"`
		TsBracketList    []Timestamp   `query:"ts_bracket_list,bracket"`
		TsIndexList      []Timestamp   `query:"ts_index_list,index"`
		OmitTsList       []*Timestamp  `query:"omit_ts_list"`
		NilTsList        []*Timestamp  `query:"nil_ts_list"`
		TsPtrList        []*Timestamp  `query:"ts_ptr_list"`
		TsPtrCommaList   []*Timestamp  `query:"ts_ptr_comma_list,comma"`
		TsPtrBracketList []*Timestamp  `query:"ts_ptr_bracket_list,bracket"`
		TsPtrIndexList   []*Timestamp  `query:"ts_ptr_index_list,index"`
	}{
		OmitTimestamp:   Timestamp{time.Time{}.UTC()},
		ZeroTimestamp:   Timestamp{time.Time{}.UTC()},
		Timestamp:       Timestamp{tm},
		ZeroInterfaceTs: zeroPtrTs,
		InterfaceTs:     Timestamp{tm},
		TimestampPtr:    &TimestampPtr{tm},
		TsList: []Timestamp{
			{tm}, {tm},
		},
		TsCommaList: []Timestamp{
			{tm}, {tm},
		},
		TsBracketList: []Timestamp{
			{tm}, {tm},
		},
		TsIndexList: []Timestamp{
			{tm}, {tm},
		},
		NilTsList: []*Timestamp{
			nil,
		},
		TsPtrList: []*Timestamp{
			nil, {tm}, {tm},
		},
		TsPtrCommaList: []*Timestamp{
			nil, {tm}, {tm},
		},
		TsPtrBracketList: []*Timestamp{
			nil, {tm}, {tm},
		},
		TsPtrIndexList: []*Timestamp{
			nil, {tm}, {tm},
		},
	}

	values, err := encoder.Values(&s)
	test.NoError(err)

	expected := url.Values{
		"zero_ts":               []string{""},
		"ts":                    []string{"1970-01-01T00:00:00Z"},
		"zero_ptr_ts":           []string{""},
		"timestamp_ptr":         []string{"1970-01-01T00:00:00Z"},
		"interface_ts":          []string{"1970-01-01T00:00:00Z"},
		"ts_list":               []string{"1970-01-01T00:00:00Z", "1970-01-01T00:00:00Z"},
		"ts_comma_list":         []string{"1970-01-01T00:00:00Z,1970-01-01T00:00:00Z"},
		"ts_bracket_list[]":     []string{"1970-01-01T00:00:00Z", "1970-01-01T00:00:00Z"},
		"ts_index_list[0]":      []string{"1970-01-01T00:00:00Z"},
		"ts_index_list[1]":      []string{"1970-01-01T00:00:00Z"},
		"ts_ptr_list":           []string{"1970-01-01T00:00:00Z", "1970-01-01T00:00:00Z"},
		"ts_ptr_comma_list":     []string{"1970-01-01T00:00:00Z,1970-01-01T00:00:00Z"},
		"ts_ptr_bracket_list[]": []string{"1970-01-01T00:00:00Z", "1970-01-01T00:00:00Z"},
		"ts_ptr_index_list[0]":  []string{"1970-01-01T00:00:00Z"},
		"ts_ptr_index_list[1]":  []string{"1970-01-01T00:00:00Z"},
	}
	test.Equal(expected, values)
}

type ErrTimestamp struct {
	time.Time
}

func (t *ErrTimestamp) EncodeParam() (string, error) {
	return "", fmt.Errorf("failed to encode param")
}

func (t *ErrTimestamp) IsZero() bool {
	return t.Time.IsZero()
}

func TestEncodeErrCustomType(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	tm := time.Unix(0, 0).UTC()

	s1 := struct {
		ErrTimestamp *ErrTimestamp
	}{
		ErrTimestamp: &ErrTimestamp{tm},
	}
	s2 := struct {
		ErrTimestamps []*ErrTimestamp
	}{
		ErrTimestamps: []*ErrTimestamp{
			{tm},
		},
	}
	s3 := struct {
		ErrBracketList []*ErrTimestamp `query:",bracket"`
	}{
		ErrBracketList: []*ErrTimestamp{
			{tm},
		},
	}
	s4 := struct {
		ErrIndexList []*ErrTimestamp `query:",index"`
	}{
		ErrIndexList: []*ErrTimestamp{
			{tm},
		},
	}
	s5 := struct {
		ErrIndexList []*ErrTimestamp `query:",comma"`
	}{
		ErrIndexList: []*ErrTimestamp{
			{tm},
		},
	}
	s6 := struct {
		ErrTimestampMap map[string]*ErrTimestamp
	}{
		ErrTimestampMap: map[string]*ErrTimestamp{
			"test": {tm},
		},
	}
	s7 := struct {
		ErrTimestampMap map[*ErrTimestamp]bool
	}{
		ErrTimestampMap: map[*ErrTimestamp]bool{
			{tm}: true,
		},
	}
	s8 := struct {
		Embed struct {
			ErrTimestampMap *ErrTimestamp
		}
	}{
		Embed: struct {
			ErrTimestampMap *ErrTimestamp
		}{
			ErrTimestampMap: &ErrTimestamp{tm},
		},
	}
	s9 := struct {
		ErrTimestamp interface{}
	}{
		ErrTimestamp: &ErrTimestamp{tm},
	}

	values, err := encoder.Values(&s1)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s1, values)
	test.Error(err)

	values, err = encoder.Values(&s2)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s2, values)
	test.Error(err)

	values, err = encoder.Values(&s3)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s3, values)
	test.Error(err)

	values, err = encoder.Values(&s4)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s4, values)
	test.Error(err)

	values, err = encoder.Values(&s5)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s5, values)
	test.Error(err)

	values, err = encoder.Values(&s6)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s6, values)
	test.Error(err)

	values, err = encoder.Values(&s7)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s7, values)
	test.Error(err)

	values, err = encoder.Values(&s8)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s8, values)
	test.Error(err)

	values, err = encoder.Values(&s9)
	test.Error(err)
	values = url.Values{}
	err = encoder.Encode(&s9, values)
	test.Error(err)
}

func TestEncoderIgnoreUnregisterType(t *testing.T) {
	test := assert.New(t)
	encoder := NewEncoder()

	s := struct {
		Fn  []func()             `query:"fn,bracket"`
		Ch  []chan struct{}      `query:"chan,comma"`
		Ch2 []chan struct{}      `query:"chan2,index"`
		Ch3 []chan struct{}      `query:"chan3"`
		Ch4 map[chan bool]func() `query:"chan4"`
	}{
		Fn:  []func(){func() {}},
		Ch:  []chan struct{}{make(chan struct{})},
		Ch2: []chan struct{}{make(chan struct{})},
	}

	values, err := encoder.Values(s)
	test.NoError(err)

	test.Equal(url.Values{}, values)
}

//------------------------------------------------

func withStr(v string) *string {
	return &v
}

func withBool(v bool) *bool {
	return &v
}

func withInt(v int) *int {
	return &v
}

func withInt8(v int8) *int8 {
	return &v
}

func withInt16(v int16) *int16 {
	return &v
}

func withInt32(v int32) *int32 {
	return &v
}

func withInt64(v int64) *int64 {
	return &v
}

func withUint(v uint) *uint {
	return &v
}

func withUint8(v uint8) *uint8 {
	return &v
}

func withUint16(v uint16) *uint16 {
	return &v
}

func withUint32(v uint32) *uint32 {
	return &v
}

func withUint64(v uint64) *uint64 {
	return &v
}

func withUintPtr(v uintptr) *uintptr {
	return &v
}

func withFloat32(v float32) *float32 {
	return &v
}

func withFloat64(v float64) *float64 {
	return &v
}

func withComplex64(v complex64) *complex64 {
	return &v
}

func withComplex128(v complex128) *complex128 {
	return &v
}

func complexZeroValStr() string {
	return strconv.FormatComplex(complex128(0), 'f', -1, 128)
}

func complex128ToStr(v complex128) string {
	return strconv.FormatComplex(v, 'f', -1, 128)
}

func withTime(v time.Time) *time.Time {
	return &v
}
