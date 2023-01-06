// +build !amd64
// go:build !amd64

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
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/bytedance/sonic/internal/quote"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/option"
)

// Encode returns the JSON encoding of val, encoded with opts.
func EncodeInto(buf *[]byte, val interface{}, opts Options) (error) {
	w := bytes.NewBuffer(*buf)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(opts&EscapeHTML != 0)

	err := encoder.Encode(val)
	*buf = w.Bytes()
	return err
}

// Quote returns the JSON-quoted version of s.
func Quote(s string) string {
	out := make([]byte, 0, len(s)+2)
	quote.QuoteString(&out, s, false)
	return rt.Mem2Str(out)
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
// so that the JSON will be safe to embed inside HTML <script> tags.
// For historical reasons, web browsers don't honor standard HTML
// escaping within <script> tags, so an alternative JSON encoding must
// be used.
func HTMLEscape(s string) string {
	out := make([]byte, 0, len(s))
	quote.EscapeHTML(&out, s)
	return rt.Mem2Str(out)
}

// Valid validates json and returns first non-blank character position,
// if it is only one valid json value.
// Otherwise returns invalid character position using start.
func Valid(data []byte) (ok bool, start int) {
	pos := quote.SkipBlank(rt.Mem2Str(data), 0)
	if pos < 0 {
		return false, 0
	}
	ok = json.Valid(data)
	if !ok {
		return false, 0
	}
	return 
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
    return nil
}