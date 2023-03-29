// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

var Stubs = []loader.GoC{
    {"_f32toa", &_subr__f32toa, &__f32toa},
    {"_f64toa", &_subr__f64toa, &__f64toa},
    {"_fsm_exec", &_subr__fsm_exec, &__fsm_exec},
    {"_get_by_path", &_subr__get_by_path, &__get_by_path},
    {"_html_escape", &_subr__html_escape, &__html_escape},
    {"_i64toa", &_subr__i64toa, &__i64toa},
    {"_lspace", &_subr__lspace, &__lspace},
    {"_quote", &_subr__quote, &__quote},
    {"_skip_array", &_subr__skip_array, &__skip_array},
    {"_skip_number", &_subr__skip_number, &__skip_number},
    {"_skip_object", &_subr__skip_object, &__skip_object},
    {"_skip_one", &_subr__skip_one, &__skip_one},
    {"_skip_one_fast", &_subr__skip_one_fast, &__skip_one_fast},
    {"_u64toa", &_subr__u64toa, &__u64toa},
    {"_unquote", &_subr__unquote, &__unquote},
    {"_validate_one", &_subr__validate_one, &__validate_one},
    {"_validate_utf8", &_subr__validate_utf8, &__validate_utf8},
    {"_validate_utf8_fast", &_subr__validate_utf8_fast, &__validate_utf8_fast},
    {"_value", &_subr__value, &__value},
    {"_vnumber", &_subr__vnumber, &__vnumber},
    {"_vsigned", &_subr__vsigned, &__vsigned},
    {"_vstring", &_subr__vstring, &__vstring},
    {"_vunsigned", &_subr__vunsigned, &__vunsigned},
}

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_advance_ns", _entry__advance_ns, _size__advance_ns, _stack__advance_ns, _pcsp__advance_ns},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
    {"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
    {"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}

var (
    _subr__f32toa             uintptr
    _subr__f64toa             uintptr
    _subr__fsm_exec           uintptr
    _subr__get_by_path        uintptr
    _subr__html_escape        uintptr
    _subr__i64toa             uintptr
    _subr__lspace             uintptr
    _subr__quote              uintptr
    _subr__skip_array         uintptr
    _subr__skip_number        uintptr
    _subr__skip_object        uintptr
    _subr__skip_one           uintptr
    _subr__skip_one_fast      uintptr
    _subr__u64toa             uintptr
    _subr__unquote            uintptr
    _subr__validate_one       uintptr
    _subr__validate_utf8      uintptr
    _subr__validate_utf8_fast uintptr
    _subr__value              uintptr
    _subr__vnumber            uintptr
    _subr__vsigned            uintptr
    _subr__vstring            uintptr
    _subr__vunsigned          uintptr
)

const (
    _stack__f32toa = 48
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__fsm_exec = 152
    _stack__advance_ns = 16
    _stack__advance_string = 64
    _stack__advance_string_default = 64
    _stack__do_skip_number = 48
    _stack__get_by_path = 280
    _stack__skip_one_fast = 192
    _stack__html_escape = 72
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 56
    _stack__skip_array = 160
    _stack__skip_number = 88
    _stack__skip_object = 160
    _stack__skip_one = 160
    _stack__unquote = 88
    _stack__validate_one = 160
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 24
    _stack__value = 328
    _stack__vnumber = 240
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 136
    _stack__decimal_to_f64 = 80
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__vsigned = 16
    _stack__vstring = 120
    _stack__vunsigned = 8
)

const (
    _entry__f32toa = 28800
    _entry__f64toa = 480
    _entry__format_significand = 32192
    _entry__format_integer = 3328
    _entry__fsm_exec = 18656
    _entry__advance_ns = 14192
    _entry__advance_string = 15056
    _entry__advance_string_default = 33584
    _entry__do_skip_number = 20944
    _entry__get_by_path = 25664
    _entry__skip_one_fast = 22576
    _entry__html_escape = 9200
    _entry__i64toa = 3760
    _entry__u64toa = 3872
    _entry__lspace = 80
    _entry__quote = 5152
    _entry__skip_array = 18608
    _entry__skip_number = 22208
    _entry__skip_object = 20608
    _entry__skip_one = 22352
    _entry__unquote = 6864
    _entry__validate_one = 22400
    _entry__validate_utf8 = 27552
    _entry__validate_utf8_fast = 28224
    _entry__value = 12608
    _entry__vnumber = 16384
    _entry__atof_eisel_lemire64 = 10448
    _entry__atof_native = 12000
    _entry__decimal_to_f64 = 10816
    _entry__right_shift = 33152
    _entry__left_shift = 32656
    _entry__vsigned = 17936
    _entry__vstring = 14880
    _entry__vunsigned = 18256
)

const (
    _size__f32toa = 3392
    _size__f64toa = 2848
    _size__format_significand = 464
    _size__format_integer = 432
    _size__fsm_exec = 1416
    _size__advance_ns = 688
    _size__advance_string = 1280
    _size__advance_string_default = 944
    _size__do_skip_number = 924
    _size__get_by_path = 1888
    _size__skip_one_fast = 2580
    _size__html_escape = 1248
    _size__i64toa = 48
    _size__u64toa = 1232
    _size__lspace = 368
    _size__quote = 1696
    _size__skip_array = 48
    _size__skip_number = 144
    _size__skip_object = 48
    _size__skip_one = 48
    _size__unquote = 2272
    _size__validate_one = 48
    _size__validate_utf8 = 672
    _size__validate_utf8_fast = 544
    _size__value = 1004
    _size__vnumber = 1552
    _size__atof_eisel_lemire64 = 368
    _size__atof_native = 608
    _size__decimal_to_f64 = 1184
    _size__right_shift = 400
    _size__left_shift = 496
    _size__vsigned = 320
    _size__vstring = 128
    _size__vunsigned = 336
)

var (
    _pcsp__f32toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {3350, 48},
        {3351, 40},
        {3353, 32},
        {3355, 24},
        {3357, 16},
        {3359, 8},
        {3363, 0},
        {3385, 48},
    }
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2788, 56},
        {2792, 48},
        {2793, 40},
        {2795, 32},
        {2797, 24},
        {2799, 16},
        {2801, 8},
        {2805, 0},
        {2843, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {452, 24},
        {453, 16},
        {455, 8},
        {457, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {1, 0},
        {4, 8},
        {412, 16},
        {413, 8},
        {414, 0},
        {423, 16},
        {424, 8},
        {426, 0},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1135, 88},
        {1139, 48},
        {1140, 40},
        {1142, 32},
        {1144, 24},
        {1146, 16},
        {1148, 8},
        {1149, 0},
        {1416, 88},
    }
    _pcsp__advance_ns = [][2]uint32{
        {1, 0},
        {4, 8},
        {645, 16},
        {646, 8},
        {647, 0},
        {671, 16},
        {672, 8},
        {674, 0},
    }
    _pcsp__advance_string = [][2]uint32{
        {14, 0},
        {18, 8},
        {20, 16},
        {22, 24},
        {24, 32},
        {26, 40},
        {27, 48},
        {557, 56},
        {561, 48},
        {562, 40},
        {564, 32},
        {566, 24},
        {568, 16},
        {570, 8},
        {571, 0},
        {1268, 56},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {552, 64},
        {556, 48},
        {557, 40},
        {559, 32},
        {561, 24},
        {563, 16},
        {565, 8},
        {566, 0},
        {931, 64},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {849, 48},
        {850, 40},
        {852, 32},
        {854, 24},
        {856, 16},
        {858, 8},
        {859, 0},
        {924, 48},
    }
    _pcsp__get_by_path = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1772, 88},
        {1776, 48},
        {1777, 40},
        {1779, 32},
        {1781, 24},
        {1783, 16},
        {1785, 8},
        {1786, 0},
        {1878, 88},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {316, 176},
        {317, 168},
        {319, 160},
        {321, 152},
        {323, 144},
        {325, 136},
        {329, 128},
        {2580, 176},
    }
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1224, 72},
        {1228, 48},
        {1229, 40},
        {1231, 32},
        {1233, 24},
        {1235, 16},
        {1237, 8},
        {1239, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {14, 0},
        {34, 8},
        {36, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {1, 0},
        {161, 8},
        {162, 0},
        {457, 8},
        {458, 0},
        {756, 8},
        {757, 0},
        {1221, 8},
        {1223, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {1, 0},
        {332, 8},
        {333, 0},
        {340, 8},
        {341, 0},
        {348, 8},
        {350, 0},
    }
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1649, 56},
        {1653, 48},
        {1654, 40},
        {1656, 32},
        {1658, 24},
        {1660, 16},
        {1662, 8},
        {1663, 0},
        {1690, 56},
    }
    _pcsp__skip_array = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {100, 40},
        {101, 32},
        {103, 24},
        {105, 16},
        {107, 8},
        {108, 0},
        {139, 40},
    }
    _pcsp__skip_object = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_one = [][2]uint32{
        {1, 0},
        {30, 8},
        {36, 0},
    }
    _pcsp__unquote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1684, 88},
        {1688, 48},
        {1689, 40},
        {1691, 32},
        {1693, 24},
        {1695, 16},
        {1697, 8},
        {1698, 0},
        {2270, 88},
    }
    _pcsp__validate_one = [][2]uint32{
        {1, 0},
        {35, 8},
        {41, 0},
    }
    _pcsp__validate_utf8 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {623, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {666, 48},
    }
    _pcsp__validate_utf8_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {5, 16},
        {247, 24},
        {251, 16},
        {252, 8},
        {253, 0},
        {527, 24},
        {531, 16},
        {532, 8},
        {534, 0},
    }
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {988, 88},
        {992, 48},
        {993, 40},
        {995, 32},
        {997, 24},
        {999, 16},
        {1001, 8},
        {1004, 0},
    }
    _pcsp__vnumber = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {803, 104},
        {807, 48},
        {808, 40},
        {810, 32},
        {812, 24},
        {814, 16},
        {816, 8},
        {817, 0},
        {1547, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {292, 32},
        {293, 24},
        {295, 16},
        {297, 8},
        {298, 0},
        {362, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {1, 0},
        {4, 8},
        {587, 56},
        {591, 8},
        {593, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1144, 56},
        {1148, 48},
        {1149, 40},
        {1151, 32},
        {1153, 24},
        {1155, 16},
        {1157, 8},
        {1158, 0},
        {1169, 56},
    }
    _pcsp__right_shift = [][2]uint32{
        {1, 0},
        {318, 8},
        {319, 0},
        {387, 8},
        {388, 0},
        {396, 8},
        {398, 0},
    }
    _pcsp__left_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {363, 24},
        {364, 16},
        {366, 8},
        {367, 0},
        {470, 24},
        {471, 16},
        {473, 8},
        {474, 0},
        {486, 24},
    }
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {112, 16},
        {113, 8},
        {114, 0},
        {125, 16},
        {126, 8},
        {127, 0},
        {260, 16},
        {261, 8},
        {262, 0},
        {266, 16},
        {267, 8},
        {268, 0},
        {306, 16},
        {307, 8},
        {308, 0},
        {316, 16},
        {317, 8},
        {319, 0},
    }
    _pcsp__vstring = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {105, 56},
        {109, 40},
        {110, 32},
        {112, 24},
        {114, 16},
        {116, 8},
        {118, 0},
    }
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {71, 8},
        {72, 0},
        {83, 8},
        {84, 0},
        {107, 8},
        {108, 0},
        {273, 8},
        {274, 0},
        {312, 8},
        {313, 0},
        {320, 8},
        {322, 0},
    }
)
