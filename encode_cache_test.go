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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheStore(t *testing.T) {
	test := assert.New(t)

	s := &basicVal{}

	cacheStore := newCacheStore()
	test.NotNil(cacheStore)

	fields := cachedFields{&float64Field{}}
	cacheStore.Store(reflect.TypeOf(s), fields)
	cachedFlds := cacheStore.Retrieve(reflect.TypeOf(s))

	test.NotNil(cachedFlds)
	test.Len(cachedFlds, len(fields))
	test.True(&fields[0] == &cachedFlds[0])
}

func TestNewCacheField(t *testing.T) {
	test := assert.New(t)
	name := []byte(`abc`)
	opts := [][]byte{[]byte(`omitempty`)}

	cacheField := newCachedFieldByKind(reflect.ValueOf("").Kind(), name, opts)
	if stringField, ok := cacheField.(*stringField); ok {
		test.Equal(string(name), stringField.name)
		test.True(stringField.omitEmpty)
	} else {
		test.FailNow("")
	}
	test.IsType(&stringField{}, cacheField)
}

func TestNewCacheField2(t *testing.T) {
	test := assert.New(t)

	var strPtr *string
	cacheField := newCachedFieldByKind(reflect.ValueOf(strPtr).Kind(), nil, nil)
	test.Nil(cacheField)
}
