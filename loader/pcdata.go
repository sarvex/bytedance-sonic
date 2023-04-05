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

package loader

import (
	"encoding/binary"
	"fmt"
)

const (
    _N_PCDATA   = 4

    _PCDATA_UnsafePoint   = 0
    _PCDATA_StackMapIndex = 1
    _PCDATA_InlTreeIndex  = 2
    _PCDATA_ArgLiveIndex  = 3

    _PCDATA_INVALID_OFFSET = 0
)

const (
    // PCDATA_UnsafePoint values.
    PCDATA_UnsafePointSafe   = -1 // Safe for async preemption
    PCDATA_UnsafePointUnsafe = -2 // Unsafe for async preemption

    // PCDATA_Restart1(2) apply on a sequence of instructions, within
    // which if an async preemption happens, we should back off the PC
    // to the start of the sequence when resume.
    // We need two so we can distinguish the start/end of the sequence
    // in case that two sequences are next to each other.
    PCDATA_Restart1 = -3
    PCDATA_Restart2 = -4

    // Like PCDATA_RestartAtEntry, but back to function entry if async
    // preempted.
    PCDATA_RestartAtEntry = -5

    _PCDATA_START_VAL = -1
)

var emptyByte byte

// Pcvalue is the last program count corresponding to the value Val,
// which means that ** (previousPcvalue.PC, currentPcvalue.PC] **
// is the range where the value is effective.
type Pcvalue struct {
    PC  uint32 // the **LAST** PC corresponding to the Value
    Val int32  // the Value
}

// Pcdata represents pc->value mapping table
type Pcdata []Pcvalue

// Document: https://docs.google.com/document/d/1lyPIbmsYbXnpNj57a261hgOYVpNRcgydurVQIyZOz_o/pub
func (self Pcdata) MarshalBinary(maxpc uint32) (data []byte, err error) {
    buf := make([]byte, binary.MaxVarintLen32)

    // delta value always starts from -1
    sv := int32(_PCDATA_START_VAL)
    sp := uint32(0)

    // write value delta first, write pcdata last
    started := false
    for _, v := range self {
        dv := v.Val - sv
        if dv == 0 && started {
            continue
        }
        if v.PC < sp {
            panic(fmt.Sprintf("invalid pc %d, must be larger than the previous!", v.PC))
        }
        dp := v.PC - sp
        if dp == 0 {
            continue
        }

        if started {
            n := binary.PutUvarint(buf, uint64(dp))
            data = append(data, buf[:n]...)
            sp = v.PC
        }

        n := binary.PutVarint(buf, int64(dv))
        data = append(data, buf[:n]...)
        sv = v.Val
        started = true
    }

    if started {
        data = append(data, buf[:binary.PutUvarint(buf, uint64(maxpc - sp))]...)
        // add terminating varint-encoded 0, which is just 0
		data = append(data, 0)
    }
    return
}