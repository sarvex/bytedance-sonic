//go:build amd64
// +build amd64

// go:build amd64

/**
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package encoder

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/option"
)

// EncodeInto is like Encode but uses a user-supplied buffer instead of allocating
// a new one.
func EncodeInto(buf *[]byte, val interface{}, opts Options) error {
    stk := newStack()
    efv := rt.UnpackEface(val)
    err := encodeTypedPointer(buf, efv.Type, &efv.Value, stk, uint64(opts))

    /* return the stack into pool */
    if err != nil {
        resetStack(stk)
    }
    freeStack(stk)

    /* EscapeHTML needs to allocate a new buffer*/
    if opts & EscapeHTML != 0 {
        dest := HTMLEscape(nil, *buf)
        freeBytes(*buf) // free origin used buffer
        *buf = dest
    }

    /* avoid GC ahead */
    runtime.KeepAlive(buf)
    runtime.KeepAlive(efv)
    return err
}

// Quote returns the JSON-quoted version of s.
func Quote(s string) string {
    var n int
    var p []byte

    /* check for empty string */
    if s == "" {
        return `""`
    }

    /* allocate space for result */
    n = len(s) + 2
    p = make([]byte, 0, n)

    /* call the encoder */
    _ = encodeString(&p, s)
    return rt.Mem2Str(p)
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
// so that the JSON will be safe to embed inside HTML <script> tags.
// For historical reasons, web browsers don't honor standard HTML
// escaping within <script> tags, so an alternative JSON encoding must
// be used.
func HTMLEscape(dest []byte, src []byte) []byte {
    nb := len(src)

    // initilize dest buffer
    cap := nb * 6 / 5
    if dest == nil {
        dest = make([]byte, 0, cap)
    }
    ds := (*rt.GoSlice)(unsafe.Pointer(&dest))
    sp := (*rt.GoSlice)(unsafe.Pointer(&src)).Ptr
    ds.Len = 0
    if (ds.Cap < cap) {
        *ds = growslice(typeByte, *ds, cap)
    }

    for nb > 0 {
        dp := unsafe.Pointer(uintptr(ds.Ptr) + uintptr(ds.Len))
        dn := ds.Cap - ds.Len

        ret := native.HTMLEscape(sp, nb, dp, &dn)
        ds.Len += dn

        if ret >= 0 {
            break
        }
        ret = ^ret
        nb -= ret

        *ds = growslice(typeByte, *ds, ds.Cap * 2)
        sp = unsafe.Pointer(uintptr(sp) + uintptr(ret))
    }
    return dest
}


// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
    cfg := option.DefaultCompileOptions()
    for _, opt := range opts {
        opt(&cfg)
        break
    }
    return pretouchRec(map[reflect.Type]uint8{vt: 0}, cfg)
}

// Valid validates json and returns first non-blank character position,
// if it is only one valid json value.
// Otherwise returns invalid character position using start.
func Valid(data []byte) (ok bool, start int) {
    n := len(data)
    if n == 0 {
        return false, -1
    }
    s := rt.Mem2Str(data)
    p := 0
    m := types.NewStateMachine()
    ret := native.ValidateOne(&s, &p, m)
    types.FreeStateMachine(m) 
    if ret < 0 {
        return false, p-1
    }
    for ;p < n; p++ {
        if (types.SPACE_MASK & (1 << data[p])) == 0 {
            return false, p
        }
    }
    return true, ret
}