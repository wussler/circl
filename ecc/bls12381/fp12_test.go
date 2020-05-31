package bls12381

import (
	"testing"

	"github.com/cloudflare/circl/internal/test"
)

func randomFp12() *fp12 {
	return &fp12{
		*randomFp6(),
		*randomFp6(),
	}
}

func TestFp12(t *testing.T) {
	const testTimes = 1 << 6
	t.Run("mul_inv", func(t *testing.T) {
		var z fp12
		for i := 0; i < testTimes; i++ {
			x := randomFp12()
			y := randomFp12()

			// x*y*x^1 - y = 0
			z.Inv(x)
			z.Mul(&z, y)
			z.Mul(&z, x)
			z.Sub(&z, y)
			got := z.IsZero()
			want := true
			if got != want {
				test.ReportError(t, got, want, x, y)
			}
		}
	})
	t.Run("mul_sqr", func(t *testing.T) {
		var l0, l1, r0, r1 fp12
		for i := 0; i < testTimes; i++ {
			x := randomFp12()
			y := randomFp12()

			// (x+y)(x-y) = (x^2-y^2)
			l0.Add(x, y)
			l1.Sub(x, y)
			l0.Mul(&l0, &l1)
			r0.Sqr(x)
			r1.Sqr(y)
			r0.Sub(&r0, &r1)
			got := &l0
			want := &r0
			if !got.IsEqual(want) {
				test.ReportError(t, got, want, x, y)
			}
		}
	})
}

func BenchmarkFp12(b *testing.B) {
	x := randomFp12()
	y := randomFp12()
	z := randomFp12()
	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Add(x, y)
		}
	})
	b.Run("Mul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Mul(x, y)
		}
	})
	b.Run("Sqr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Sqr(x)
		}
	})
	b.Run("Inv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Inv(x)
		}
	})
}