// +build amd64

package common

import (
	"golang.org/x/sys/cpu"
)

// ZetasAVX2 contains all ζ used in NTT (like the Zetas array), but also
// the values int16(zeta * 62209) for each zeta, which is used in
// Montgomery reduction.  There is some duplication and reordering as
// compared to Zetas to make it more covenient for use with AVX2.
var ZetasAVX2 = [...]int16{
	// level 1: int16(Zetas[1]*62209) and Zetas[1]
	31499, 2571,

	// level 2
	//
	// int16(Zetas[2]*62209), Zetas[2], int16(Zetas[3]*62209), Zetas[3]
	14746, 2970, 788, 1812,

	// level 3, like level 2.
	13525, 1493, -12402, 1422, 28191, 287, -16694, 202,

	0, 0, // padding

	// layer 4
	//
	// The precomputed multiplication and zetas are grouped by 16 at a
	// time as used in the set of butterflies, etc.
	-20906, -20906, -20906, -20906, -20906, -20906, -20906, -20906,
	27758, 27758, 27758, 27758, 27758, 27758, 27758, 27758,
	3158, 3158, 3158, 3158, 3158, 3158, 3158, 3158,
	622, 622, 622, 622, 622, 622, 622, 622,
	-3799, -3799, -3799, -3799, -3799, -3799, -3799, -3799,
	-15690, -15690, -15690, -15690, -15690, -15690, -15690, -15690,
	1577, 1577, 1577, 1577, 1577, 1577, 1577, 1577,
	182, 182, 182, 182, 182, 182, 182, 182,
	10690, 10690, 10690, 10690, 10690, 10690, 10690, 10690,
	1359, 1359, 1359, 1359, 1359, 1359, 1359, 1359,
	962, 962, 962, 962, 962, 962, 962, 962,
	2127, 2127, 2127, 2127, 2127, 2127, 2127, 2127,
	-11201, -11201, -11201, -11201, -11201, -11201, -11201, -11201,
	31164, 31164, 31164, 31164, 31164, 31164, 31164, 31164,
	1855, 1855, 1855, 1855, 1855, 1855, 1855, 1855,
	1468, 1468, 1468, 1468, 1468, 1468, 1468, 1468,

	// layer 5
	-5827, -5827, -5827, -5827, 17364, 17364, 17364, 17364,
	-26360, -26360, -26360, -26360, -29057, -29057, -29057, -29057,
	573, 573, 573, 573, 2004, 2004, 2004, 2004,
	264, 264, 264, 264, 383, 383, 383, 383,
	5572, 5572, 5572, 5572, -1102, -1102, -1102, -1102,
	21439, 21439, 21439, 21439, -26241, -26241, -26241, -26241,
	2500, 2500, 2500, 2500, 1458, 1458, 1458, 1458,
	1727, 1727, 1727, 1727, 3199, 3199, 3199, 3199,
	-28072, -28072, -28072, -28072, 24313, 24313, 24313, 24313,
	-10532, -10532, -10532, -10532, 8800, 8800, 8800, 8800,
	2648, 2648, 2648, 2648, 1017, 1017, 1017, 1017,
	732, 732, 732, 732, 608, 608, 608, 608,
	18427, 18427, 18427, 18427, 8859, 8859, 8859, 8859,
	26676, 26676, 26676, 26676, -16162, -16162, -16162, -16162,
	1787, 1787, 1787, 1787, 411, 411, 411, 411,
	3124, 3124, 3124, 3124, 1758, 1758, 1758, 1758,

	// layer 6
	-5689, -5689, -6516, -6516, 1497, 1497, 30967, 30967,
	-23564, -23564, 20179, 20179, 20711, 20711, 25081, 25081,
	1223, 1223, 652, 652, 2777, 2777, 1015, 1015,
	2036, 2036, 1491, 1491, 3047, 3047, 1785, 1785,
	-12796, -12796, 26617, 26617, 16065, 16065, -12441, -12441,
	9135, 9135, -649, -649, -25986, -25986, 27837, 27837,
	516, 516, 3321, 3321, 3009, 3009, 2663, 2663,
	1711, 1711, 2167, 2167, 126, 126, 1469, 1469,
	19884, 19884, -28249, -28249, -15886, -15886, -8898, -8898,
	-28309, -28309, 9076, 9076, -30198, -30198, 18250, 18250,
	2476, 2476, 3239, 3239, 3058, 3058, 830, 830,
	107, 107, 1908, 1908, 3082, 3082, 2378, 2378,
	13427, 13427, 14017, 14017, -29155, -29155, -12756, -12756,
	16832, 16832, 4312, 4312, -24155, -24155, -17914, -17914,
	2931, 2931, 961, 961, 1821, 1821, 2604, 2604,
	448, 448, 2264, 2264, 677, 677, 2054, 2054,

	// layer 7
	-334, 11182, -11477, 13387, -32226, -14233, 20494, -21655,
	-27738, 13131, 945, -4586, -14882, 23093, 6182, 5493,
	2226, 430, 555, 843, 2078, 871, 1550, 105,
	422, 587, 177, 3094, 3038, 2869, 1574, 1653,
	32011, -32502, 10631, 30318, 29176, -18741, -28761, 12639,
	-18485, 20100, 17561, 18525, -14430, 19529, -5275, -12618,
	3083, 778, 1159, 3182, 2552, 1483, 2727, 1119,
	1739, 644, 2457, 349, 418, 329, 3173, 3254,
	-31183, 20297, 25435, 2146, -7382, 15356, 24392, -32384,
	-20926, -6279, 10946, -14902, 24215, -11044, 16990, 14470,
	817, 1097, 603, 610, 1322, 2044, 1864, 384,
	2114, 3193, 1218, 1994, 2455, 220, 2142, 1670,
	10336, -21497, -7933, -20198, -22501, 23211, 10907, -17442,
	31637, -23859, 28644, -20257, 23998, 7757, -17422, 23132,
	2144, 1799, 2051, 794, 1819, 2475, 2459, 478,
	3221, 3021, 996, 991, 958, 1869, 1522, 1628,
}

// Sets p to a + b.  Does not normalize coefficients.
func (p *Poly) Add(a, b *Poly) {
	if cpu.X86.HasAVX2 {
		addAVX2(
			(*[N]int16)(p),
			(*[N]int16)(a),
			(*[N]int16)(b),
		)
	} else {
		p.addGeneric(a, b)
	}
}

// Sets p to a - b.  Does not normalize coefficients.
func (p *Poly) Sub(a, b *Poly) {
	if cpu.X86.HasAVX2 {
		subAVX2(
			(*[N]int16)(p),
			(*[N]int16)(a),
			(*[N]int16)(b),
		)
	} else {
		p.subGeneric(a, b)
	}
}

// Executes an in-place forward "NTT" on p.
//
// Assumes the coefficients are in absolute value ≤q.  The resulting
// coefficients are in absolute value ≤7q.  If the input is in Montgomery
// form, then the result is in Montgomery form and so (by linearity of the NTT)
// if the input is in regular form, then the result is also in regular form.
func (p *Poly) NTT() {
	if cpu.X86.HasAVX2 {
		nttAVX2((*[N]int16)(p))
	} else {
		p.nttGeneric()
	}
}
