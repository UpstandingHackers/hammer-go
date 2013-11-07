#line 3 "/Users/cvincent/go/src/hammer/hammer.go"
	#include "hammer/src/hammer.h"



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_69e7774c0d1c_Cfunc_h_action(void *v)
{
	struct {
		HParser* p0;
		HAction p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_action((void*)a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_and(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_and((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_attr_bool(void *v)
{
	struct {
		HParser* p0;
		HPredicate p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_attr_bool((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_bind_indirect(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
	} __attribute__((__packed__)) *a = v;
	h_bind_indirect((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_bits(void *v)
{
	struct {
		size_t p0;
		bool p1;
		char __pad12[4];
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_bits(a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_butnot(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_butnot((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_ch(void *v)
{
	struct {
		uint8_t p0;
		char __pad1[7];
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_ch(a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_ch_range(void *v)
{
	struct {
		uint8_t p0;
		uint8_t p1;
		char __pad2[6];
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_ch_range(a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_choice(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_choice((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_compile(void *v)
{
	struct {
		HParser* p0;
		HParserBackend p1;
		char __pad12[4];
		void* p2;
		int r;
		char __pad28[4];
	} __attribute__((__packed__)) *a = v;
	a->r = h_compile((void*)a->p0, a->p1, (void*)a->p2);
}

void
_cgo_69e7774c0d1c_Cfunc_h_difference(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_difference((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_end_p(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_end_p();
}

void
_cgo_69e7774c0d1c_Cfunc_h_epsilon_p(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_epsilon_p();
}

void
_cgo_69e7774c0d1c_Cfunc_h_ignore(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_ignore((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_in(void *v)
{
	struct {
		uint8_t* p0;
		size_t p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_in((void*)a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_indirect(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_indirect();
}

void
_cgo_69e7774c0d1c_Cfunc_h_int16(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_int16();
}

void
_cgo_69e7774c0d1c_Cfunc_h_int32(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_int32();
}

void
_cgo_69e7774c0d1c_Cfunc_h_int64(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_int64();
}

void
_cgo_69e7774c0d1c_Cfunc_h_int8(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_int8();
}

void
_cgo_69e7774c0d1c_Cfunc_h_int_range(void *v)
{
	struct {
		HParser* p0;
		int64_t p1;
		int64_t p2;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_int_range((void*)a->p0, a->p1, a->p2);
}

void
_cgo_69e7774c0d1c_Cfunc_h_left(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_left((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_length_value(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_length_value((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_many(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_many((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_many1(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_many1((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_middle(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		HParser* p2;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_middle((void*)a->p0, (void*)a->p1, (void*)a->p2);
}

void
_cgo_69e7774c0d1c_Cfunc_h_not(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_not((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_not_in(void *v)
{
	struct {
		uint8_t* p0;
		size_t p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_not_in((void*)a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_nothing_p(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_nothing_p();
}

void
_cgo_69e7774c0d1c_Cfunc_h_optional(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_optional((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_parse(void *v)
{
	struct {
		HParser* p0;
		uint8_t* p1;
		size_t p2;
		const HParseResult* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_parse((void*)a->p0, (void*)a->p1, a->p2);
}

void
_cgo_69e7774c0d1c_Cfunc_h_parse_result_free(void *v)
{
	struct {
		HParseResult* p0;
	} __attribute__((__packed__)) *a = v;
	h_parse_result_free((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_pprint(void *v)
{
	struct {
		FILE* p0;
		HParsedToken* p1;
		int p2;
		int p3;
	} __attribute__((__packed__)) *a = v;
	h_pprint((void*)a->p0, (void*)a->p1, a->p2, a->p3);
}

void
_cgo_69e7774c0d1c_Cfunc_h_repeat_n(void *v)
{
	struct {
		HParser* p0;
		size_t p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_repeat_n((void*)a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_right(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_right((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_sepBy(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_sepBy((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_sepBy1(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_sepBy1((void*)a->p0, (void*)a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_sequence(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_sequence((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_token(void *v)
{
	struct {
		uint8_t* p0;
		size_t p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_token((void*)a->p0, a->p1);
}

void
_cgo_69e7774c0d1c_Cfunc_h_uint16(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_uint16();
}

void
_cgo_69e7774c0d1c_Cfunc_h_uint32(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_uint32();
}

void
_cgo_69e7774c0d1c_Cfunc_h_uint64(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_uint64();
}

void
_cgo_69e7774c0d1c_Cfunc_h_uint8(void *v)
{
	struct {
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_uint8();
}

void
_cgo_69e7774c0d1c_Cfunc_h_whitespace(void *v)
{
	struct {
		HParser* p0;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_whitespace((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_write_result_unamb(void *v)
{
	struct {
		HParsedToken* p0;
		const char* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_write_result_unamb((void*)a->p0);
}

void
_cgo_69e7774c0d1c_Cfunc_h_xor(void *v)
{
	struct {
		HParser* p0;
		HParser* p1;
		const HParser* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (__typeof__(a->r)) h_xor((void*)a->p0, (void*)a->p1);
}

