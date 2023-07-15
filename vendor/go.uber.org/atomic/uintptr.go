// @generated Code generated by gen-atomicint.

// Copyright (c) 2020-2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
)

// Uintptr is an atomic wrapper around uintptr.
type Uintptr struct {
	_ nocmp // disallow non-atomic comparison

	v uintptr
}

// NewUintptr creates a new Uintptr.
func NewUintptr(val uintptr) *Uintptr {
	return &Uintptr{v: val}
}

// Load atomically loads the wrapped value.
func (i *Uintptr) Load() uintptr {
	return atomic.LoadUintptr(&i.v)
}

// Add atomically adds to the wrapped uintptr and returns the new value.
func (i *Uintptr) Add(delta uintptr) uintptr {
	return atomic.AddUintptr(&i.v, delta)
}

// Sub atomically subtracts from the wrapped uintptr and returns the new value.
func (i *Uintptr) Sub(delta uintptr) uintptr {
	return atomic.AddUintptr(&i.v, ^(delta - 1))
}

// Inc atomically increments the wrapped uintptr and returns the new value.
func (i *Uintptr) Inc() uintptr {
	return i.Add(1)
}

// Dec atomically decrements the wrapped uintptr and returns the new value.
func (i *Uintptr) Dec() uintptr {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
//
// Deprecated: Use CompareAndSwap.
func (i *Uintptr) CAS(old, new uintptr) (swapped bool) {
	return i.CompareAndSwap(old, new)
}

// CompareAndSwap is an atomic compare-and-swap.
func (i *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *Uintptr) Store(val uintptr) {
	atomic.StoreUintptr(&i.v, val)
}

// Swap atomically swaps the wrapped uintptr and returns the old value.
func (i *Uintptr) Swap(val uintptr) (old uintptr) {
	return atomic.SwapUintptr(&i.v, val)
}

// MarshalJSON encodes the wrapped uintptr into JSON.
func (i *Uintptr) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Load())
}

// UnmarshalJSON decodes JSON into the wrapped uintptr.
func (i *Uintptr) UnmarshalJSON(b []byte) error {
	var v uintptr
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	i.Store(v)
	return nil
}

// String encodes the wrapped value as a string.
func (i *Uintptr) String() string {
	v := i.Load()
	return strconv.FormatUint(uint64(v), 10)
}
