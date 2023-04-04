//go:build !go1.16 || !amd64
// +build !go1.16 !amd64

/*
 * Copyright 2021 ByteDance Inc.
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

package native

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    MaxFrameSize   uintptr = 400
    BufPaddingSize int     = 64
)

var (
    S_f64toa uintptr
    S_f32toa uintptr
    S_i64toa uintptr
    S_u64toa uintptr
    S_lspace uintptr
)

var (
    S_quote       uintptr
    S_unquote     uintptr
    S_html_escape uintptr
)

var (
    S_value     uintptr
    S_vstring   uintptr
    S_vnumber   uintptr
    S_vsigned   uintptr
    S_vunsigned uintptr
)

var (
    S_skip_one      uintptr
    S_skip_one_fast uintptr
    S_get_by_path   uintptr
    S_skip_array    uintptr
    S_skip_object   uintptr
    S_skip_number   uintptr
)

var (
    S_validate_one       uintptr
    S_validate_utf8      uintptr
    S_validate_utf8_fast uintptr
)

var (
    __Quote func(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer, flags uint64) int

    __Unquote func(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep unsafe.Pointer, flags uint64) int

    __HTMLEscape func(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer) int

    __Value func(s unsafe.Pointer, n int, p int, v unsafe.Pointer, flags uint64) int

    __SkipOne func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) int

    __SkipOneFast func(s unsafe.Pointer, p unsafe.Pointer) int

    __GetByPath func(s unsafe.Pointer, p unsafe.Pointer, path unsafe.Pointer, m unsafe.Pointer) int

    __ValidateOne func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer) int

    __I64toa func(out unsafe.Pointer, val int64) (ret int)

    __U64toa func(out unsafe.Pointer, val uint64) (ret int)

    __F64toa func(out unsafe.Pointer, val float64) (ret int)

    __ValidateUTF8 func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer) (ret int)

    __ValidateUTF8Fast func(s unsafe.Pointer) (ret int)
)

//go:nosplit
func Quote(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int {
    return __Quote(rt.NoEscape(unsafe.Pointer(s)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(dn)), flags)
}

//go:nosplit
func Unquote(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) int {
    return __Unquote(rt.NoEscape(unsafe.Pointer(s)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(ep)), flags)
}

//go:nosplit
func HTMLEscape(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int {
    return __HTMLEscape(rt.NoEscape(unsafe.Pointer(s)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(dn)))
}

//go:nosplit
func Value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) int {
    return __Value(rt.NoEscape(unsafe.Pointer(s)), n, p, rt.NoEscape(unsafe.Pointer(v)), flags)
}

//go:nosplit
func SkipOne(s *string, p *int, m *types.StateMachine, flags uint64) int {
    return __SkipOne(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}

//go:nosplit
func SkipOneFast(s *string, p *int) int {
    return __SkipOneFast(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}

//go:nosplit
func GetByPath(s *string, p *int, path *[]interface{}, m *types.StateMachine) int {
    return __GetByPath(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(path)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func ValidateOne(s *string, p *int, m *types.StateMachine) int {
    return __ValidateOne(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func I64toa(out *byte, val int64) (ret int) {
    return __I64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func U64toa(out *byte, val uint64) (ret int) {
    return __U64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func F64toa(out *byte, val float64) (ret int) {
    return __F64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func ValidateUTF8(s *string, p *int, m *types.StateMachine) (ret int) {
    return __ValidateUTF8(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func ValidateUTF8Fast(s *string) (ret int) {
    return __ValidateUTF8Fast(rt.NoEscape(unsafe.Pointer(s)))
}
