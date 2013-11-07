
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void
·_Cfunc__CMalloc(uintptr n, int8 *p)
{
	p = runtime·cmalloc(n);
	FLUSH(&p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_action
void _cgo_69e7774c0d1c_Cfunc_h_action(void*);

void
·_Cfunc_h_action(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_action, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_and
void _cgo_69e7774c0d1c_Cfunc_h_and(void*);

void
·_Cfunc_h_and(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_and, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_attr_bool
void _cgo_69e7774c0d1c_Cfunc_h_attr_bool(void*);

void
·_Cfunc_h_attr_bool(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_attr_bool, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_bind_indirect
void _cgo_69e7774c0d1c_Cfunc_h_bind_indirect(void*);

void
·_Cfunc_h_bind_indirect(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_bind_indirect, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_bits
void _cgo_69e7774c0d1c_Cfunc_h_bits(void*);

void
·_Cfunc_h_bits(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_bits, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_butnot
void _cgo_69e7774c0d1c_Cfunc_h_butnot(void*);

void
·_Cfunc_h_butnot(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_butnot, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_ch
void _cgo_69e7774c0d1c_Cfunc_h_ch(void*);

void
·_Cfunc_h_ch(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_ch, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_ch_range
void _cgo_69e7774c0d1c_Cfunc_h_ch_range(void*);

void
·_Cfunc_h_ch_range(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_ch_range, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_choice
void _cgo_69e7774c0d1c_Cfunc_h_choice(void*);

void
·_Cfunc_h_choice(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_choice, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_compile
void _cgo_69e7774c0d1c_Cfunc_h_compile(void*);

void
·_Cfunc_h_compile(struct{void *y[4];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_compile, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_difference
void _cgo_69e7774c0d1c_Cfunc_h_difference(void*);

void
·_Cfunc_h_difference(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_difference, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_end_p
void _cgo_69e7774c0d1c_Cfunc_h_end_p(void*);

void
·_Cfunc_h_end_p(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_end_p, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_epsilon_p
void _cgo_69e7774c0d1c_Cfunc_h_epsilon_p(void*);

void
·_Cfunc_h_epsilon_p(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_epsilon_p, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_ignore
void _cgo_69e7774c0d1c_Cfunc_h_ignore(void*);

void
·_Cfunc_h_ignore(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_ignore, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_in
void _cgo_69e7774c0d1c_Cfunc_h_in(void*);

void
·_Cfunc_h_in(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_in, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_indirect
void _cgo_69e7774c0d1c_Cfunc_h_indirect(void*);

void
·_Cfunc_h_indirect(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_indirect, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_int16
void _cgo_69e7774c0d1c_Cfunc_h_int16(void*);

void
·_Cfunc_h_int16(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_int16, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_int32
void _cgo_69e7774c0d1c_Cfunc_h_int32(void*);

void
·_Cfunc_h_int32(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_int32, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_int64
void _cgo_69e7774c0d1c_Cfunc_h_int64(void*);

void
·_Cfunc_h_int64(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_int64, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_int8
void _cgo_69e7774c0d1c_Cfunc_h_int8(void*);

void
·_Cfunc_h_int8(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_int8, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_int_range
void _cgo_69e7774c0d1c_Cfunc_h_int_range(void*);

void
·_Cfunc_h_int_range(struct{void *y[4];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_int_range, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_left
void _cgo_69e7774c0d1c_Cfunc_h_left(void*);

void
·_Cfunc_h_left(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_left, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_length_value
void _cgo_69e7774c0d1c_Cfunc_h_length_value(void*);

void
·_Cfunc_h_length_value(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_length_value, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_many
void _cgo_69e7774c0d1c_Cfunc_h_many(void*);

void
·_Cfunc_h_many(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_many, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_many1
void _cgo_69e7774c0d1c_Cfunc_h_many1(void*);

void
·_Cfunc_h_many1(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_many1, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_middle
void _cgo_69e7774c0d1c_Cfunc_h_middle(void*);

void
·_Cfunc_h_middle(struct{void *y[4];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_middle, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_not
void _cgo_69e7774c0d1c_Cfunc_h_not(void*);

void
·_Cfunc_h_not(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_not, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_not_in
void _cgo_69e7774c0d1c_Cfunc_h_not_in(void*);

void
·_Cfunc_h_not_in(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_not_in, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_nothing_p
void _cgo_69e7774c0d1c_Cfunc_h_nothing_p(void*);

void
·_Cfunc_h_nothing_p(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_nothing_p, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_optional
void _cgo_69e7774c0d1c_Cfunc_h_optional(void*);

void
·_Cfunc_h_optional(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_optional, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_parse
void _cgo_69e7774c0d1c_Cfunc_h_parse(void*);

void
·_Cfunc_h_parse(struct{void *y[4];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_parse, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_parse_result_free
void _cgo_69e7774c0d1c_Cfunc_h_parse_result_free(void*);

void
·_Cfunc_h_parse_result_free(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_parse_result_free, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_pprint
void _cgo_69e7774c0d1c_Cfunc_h_pprint(void*);

void
·_Cfunc_h_pprint(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_pprint, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_repeat_n
void _cgo_69e7774c0d1c_Cfunc_h_repeat_n(void*);

void
·_Cfunc_h_repeat_n(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_repeat_n, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_right
void _cgo_69e7774c0d1c_Cfunc_h_right(void*);

void
·_Cfunc_h_right(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_right, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_sepBy
void _cgo_69e7774c0d1c_Cfunc_h_sepBy(void*);

void
·_Cfunc_h_sepBy(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_sepBy, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_sepBy1
void _cgo_69e7774c0d1c_Cfunc_h_sepBy1(void*);

void
·_Cfunc_h_sepBy1(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_sepBy1, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_sequence
void _cgo_69e7774c0d1c_Cfunc_h_sequence(void*);

void
·_Cfunc_h_sequence(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_sequence, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_token
void _cgo_69e7774c0d1c_Cfunc_h_token(void*);

void
·_Cfunc_h_token(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_token, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_uint16
void _cgo_69e7774c0d1c_Cfunc_h_uint16(void*);

void
·_Cfunc_h_uint16(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_uint16, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_uint32
void _cgo_69e7774c0d1c_Cfunc_h_uint32(void*);

void
·_Cfunc_h_uint32(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_uint32, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_uint64
void _cgo_69e7774c0d1c_Cfunc_h_uint64(void*);

void
·_Cfunc_h_uint64(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_uint64, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_uint8
void _cgo_69e7774c0d1c_Cfunc_h_uint8(void*);

void
·_Cfunc_h_uint8(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_uint8, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_whitespace
void _cgo_69e7774c0d1c_Cfunc_h_whitespace(void*);

void
·_Cfunc_h_whitespace(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_whitespace, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_write_result_unamb
void _cgo_69e7774c0d1c_Cfunc_h_write_result_unamb(void*);

void
·_Cfunc_h_write_result_unamb(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_write_result_unamb, &p);
}

#pragma cgo_import_static _cgo_69e7774c0d1c_Cfunc_h_xor
void _cgo_69e7774c0d1c_Cfunc_h_xor(void*);

void
·_Cfunc_h_xor(struct{void *y[3];}p)
{
	runtime·cgocall(_cgo_69e7774c0d1c_Cfunc_h_xor, &p);
}

