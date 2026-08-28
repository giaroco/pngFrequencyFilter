package freqfilter

import (
	"math"
	"math/rand"
	"testing"
)

func TestFFT2DRoundTrip(t *testing.T) {
	h, w := 7, 5 // deliberately non-power-of-two and non-square
	rng := rand.New(rand.NewSource(1))
	orig := make([][]float64, h)
	for i := range orig {
		orig[i] = make([]float64, w)
		for j := range orig[i] {
			orig[i][j] = rng.Float64() * 255
		}
	}

	freq := FFT2D(orig)
	back := IFFT2D(freq)

	const tol = 1e-9
	for i := range orig {
		for j := range orig[i] {
			got := real(back[i][j])
			if math.Abs(got-orig[i][j]) > tol {
				t.Fatalf("round trip mismatch at (%d,%d): got %v, want %v", i, j, got, orig[i][j])
			}
			if math.Abs(imag(back[i][j])) > tol {
				t.Fatalf("unexpected imaginary residue at (%d,%d): %v", i, j, imag(back[i][j]))
			}
		}
	}
}

func TestFFTShiftRoundTrip(t *testing.T) {
	for _, dims := range [][2]int{{6, 5}, {5, 6}, {5, 5}, {6, 6}} {
		h, w := dims[0], dims[1]
		grid := make([][]complex128, h)
		n := 0
		for i := range grid {
			grid[i] = make([]complex128, w)
			for j := range grid[i] {
				grid[i][j] = complex(float64(n), float64(-n))
				n++
			}
		}

		shifted := FFTShift(grid)
		back := IFFTShift(shifted)

		for i := range grid {
			for j := range grid[i] {
				if back[i][j] != grid[i][j] {
					t.Fatalf("dims %dx%d: shift round trip mismatch at (%d,%d): got %v, want %v", h, w, i, j, back[i][j], grid[i][j])
				}
			}
		}
	}
}

func TestFFTShiftCentersDC(t *testing.T) {
	h, w := 4, 6
	real2d := make([][]float64, h)
	for i := range real2d {
		real2d[i] = make([]float64, w)
	}
	real2d[0][0] = 1 // constant-ish impulse keeps DC dominant in a simple, checkable way

	freq := FFT2D(real2d)
	shifted := FFTShift(freq)

	// DC (the sum of all input samples) must land at (h/2, w/2) after shift.
	if shifted[h/2][w/2] != freq[0][0] {
		t.Fatalf("DC bin not centered: got %v, want %v", shifted[h/2][w/2], freq[0][0])
	}
}
