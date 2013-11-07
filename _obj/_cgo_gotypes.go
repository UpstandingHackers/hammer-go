// Created by cgo - DO NOT EDIT

package hammer

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int32) { *dst = syscall.Errno(x) }
type _Ctype_FILE _Ctype_struct___sFILE

type _Ctype_HAction *[0]byte

type _Ctype_HParseResult _Ctype_struct_HParseResult_

type _Ctype_HParsedToken _Ctype_struct_HParsedToken_

type _Ctype_HParser _Ctype_struct_HParser_

type _Ctype_HParserBackend uint32

type _Ctype_HPredicate *[0]byte

type _Ctype_HTokenType uint32

type _Ctype___darwin_off_t _Ctype___int64_t

type _Ctype___darwin_size_t _Ctype_ulong

type _Ctype___int64_t _Ctype_longlong

type _Ctype_bool _Ctype_int

type _Ctype_char int8

type _Ctype_fpos_t _Ctype___darwin_off_t

type _Ctype_int int32

type _Ctype_int64_t _Ctype_longlong

type _Ctype_longlong int64

type _Ctype_short int16

type _Ctype_size_t _Ctype___darwin_size_t

type _Ctype_struct_HParseResult_ struct {
//line :1
	ast	*_Ctype_HParsedToken
//line :1
	bit_length	_Ctype_longlong
//line :1
	arena	*[0]byte
//line :1
}

type _Ctype_struct_HParsedToken_ struct {
//line :1
	token_type	_Ctype_HTokenType
//line :1
	_	[4]byte
//line :1
	anon0	[16]byte
//line :1
	index	_Ctype_size_t
//line :1
	bit_offset	_Ctype_char
//line :1
	_	[7]byte
//line :1
}

type _Ctype_struct_HParser_ struct {
//line :1
	vtable	*[0]byte
//line :1
	backend	_Ctype_HParserBackend
//line :1
	_	[4]byte
//line :1
	backend_data	unsafe.Pointer
//line :1
	env	unsafe.Pointer
//line :1
	desugared	*[0]byte
//line :1
}

type _Ctype_struct___sFILE struct {
//line :1
	_p	*_Ctype_unsignedchar
//line :1
	_r	_Ctype_int
//line :1
	_w	_Ctype_int
//line :1
	_flags	_Ctype_short
//line :1
	_file	_Ctype_short
//line :1
	_	[4]byte
//line :1
	_bf	_Ctype_struct___sbuf
//line :1
	_lbfsize	_Ctype_int
//line :1
	_	[4]byte
//line :1
	_cookie	unsafe.Pointer
//line :1
	_close	*[0]byte
//line :1
	_read	*[0]byte
//line :1
	_seek	*[0]byte
//line :1
	_write	*[0]byte
//line :1
	_ub	_Ctype_struct___sbuf
//line :1
	_extra	*[0]byte
//line :1
	_ur	_Ctype_int
//line :1
	_ubuf	[3]_Ctype_unsignedchar
//line :1
	_nbuf	[1]_Ctype_unsignedchar
//line :1
	_lb	_Ctype_struct___sbuf
//line :1
	_blksize	_Ctype_int
//line :1
	_	[4]byte
//line :1
	_offset	_Ctype_fpos_t
//line :1
}

type _Ctype_struct___sbuf struct {
//line :1
	_base	*_Ctype_unsignedchar
//line :1
	_size	_Ctype_int
//line :1
	_	[4]byte
//line :1
}

type _Ctype_uint8_t _Ctype_unsignedchar

type _Ctype_ulong uint64

type _Ctype_union___0 [16]byte

type _Ctype_unsignedchar uint8

type _Ctype_void [0]byte

func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_GoString(*_Ctype_char) string
func _Cfunc_h_action(*_Ctype_HParser, _Ctype_HAction) *_Ctype_HParser
func _Cfunc_h_and(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_attr_bool(*_Ctype_HParser, *[0]byte) *_Ctype_HParser
func _Cfunc_h_bind_indirect(*_Ctype_HParser, *_Ctype_HParser) _Ctype_void
func _Cfunc_h_bits(_Ctype_size_t, _Ctype_bool) *_Ctype_HParser
func _Cfunc_h_butnot(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_ch(_Ctype_uint8_t) *_Ctype_HParser
func _Cfunc_h_ch_range(_Ctype_uint8_t, _Ctype_uint8_t) *_Ctype_HParser
func _Cfunc_h_choice(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_compile(*_Ctype_HParser, _Ctype_HParserBackend, unsafe.Pointer) _Ctype_int
func _Cfunc_h_difference(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_end_p() *_Ctype_HParser
func _Cfunc_h_epsilon_p() *_Ctype_HParser
func _Cfunc_h_ignore(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_in(*_Ctype_uint8_t, _Ctype_size_t) *_Ctype_HParser
func _Cfunc_h_indirect() *_Ctype_HParser
func _Cfunc_h_int16() *_Ctype_HParser
func _Cfunc_h_int32() *_Ctype_HParser
func _Cfunc_h_int64() *_Ctype_HParser
func _Cfunc_h_int8() *_Ctype_HParser
func _Cfunc_h_int_range(*_Ctype_HParser, _Ctype_int64_t, _Ctype_int64_t) *_Ctype_HParser
func _Cfunc_h_left(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_length_value(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_many(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_many1(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_middle(*_Ctype_HParser, *_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_not(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_not_in(*_Ctype_uint8_t, _Ctype_size_t) *_Ctype_HParser
func _Cfunc_h_nothing_p() *_Ctype_HParser
func _Cfunc_h_optional(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_parse(*_Ctype_HParser, *_Ctype_uint8_t, _Ctype_size_t) *_Ctype_HParseResult
func _Cfunc_h_parse_result_free(*_Ctype_HParseResult) _Ctype_void
func _Cfunc_h_pprint(*_Ctype_FILE, *_Ctype_HParsedToken, _Ctype_int, _Ctype_int) _Ctype_void
func _Cfunc_h_repeat_n(*_Ctype_HParser, _Ctype_size_t) *_Ctype_HParser
func _Cfunc_h_right(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_sepBy(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_sepBy1(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_sequence(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_token(*_Ctype_uint8_t, _Ctype_size_t) *_Ctype_HParser
func _Cfunc_h_uint16() *_Ctype_HParser
func _Cfunc_h_uint32() *_Ctype_HParser
func _Cfunc_h_uint64() *_Ctype_HParser
func _Cfunc_h_uint8() *_Ctype_HParser
func _Cfunc_h_whitespace(*_Ctype_HParser) *_Ctype_HParser
func _Cfunc_h_write_result_unamb(*_Ctype_HParsedToken) *_Ctype_char
func _Cfunc_h_xor(*_Ctype_HParser, *_Ctype_HParser) *_Ctype_HParser
