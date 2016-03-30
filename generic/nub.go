// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generic

import "reflect"

var trueVal = reflect.ValueOf(true)

// Nub returns v with duplicates removed. It keeps the first instance
// of each distinct value and preserves their order.
func Nub(v Slice) Slice {
	rv := reflectSlice(v)
	indexes := make([]int, 0)
	set := reflect.MakeMap(reflect.MapOf(rv.Type().Elem(), reflect.TypeOf(false)))
	for i := 0; i < rv.Len(); i++ {
		x := rv.Index(i)
		if set.MapIndex(x).IsValid() {
			continue
		}
		set.SetMapIndex(x, trueVal)
		indexes = append(indexes, i)
	}
	return MultiIndex(v, indexes)
}

// NubAppend is equivalent to appending all of the slices in vs and
// then calling Nub on the result, but more efficient.
func NubAppend(vs ...Slice) Slice {
	if len(vs) == 0 {
		return nil
	}

	rv := reflectSlice(vs[0])
	set := reflect.MakeMap(reflect.MapOf(rv.Type().Elem(), reflect.TypeOf(false)))
	out := reflect.MakeSlice(rv.Type(), 0, 0)

	for _, v := range vs {
		rv := reflectSlice(v)
		for i := 0; i < rv.Len(); i++ {
			x := rv.Index(i)
			if set.MapIndex(x).IsValid() {
				continue
			}
			set.SetMapIndex(x, trueVal)
			out = reflect.Append(out, x)
		}
	}

	return out.Interface()
}
