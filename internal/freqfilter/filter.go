package freqfilter

import (
	"fmt"
	"math"
)

// LowPass zeroes out every bin in an already FFTShift-ed H x W
// frequency-domain grid whose normalized radial distance from the center
// (DC) exceeds cutoff. The radius for bin (i,j) is normalized per axis
// against that axis' Nyquist frequency, so it is 1.0 at the edge midpoints
// and up to sqrt(2) at the corners:
//
//	u := (i - H/2) / (H/2)
//	v := (j - W/2) / (W/2)
//	r := sqrt(u^2 + v^2)
//
// cutoff must be in [0,1]. A cutoff of 0 keeps only the DC bin; a cutoff
// of 1 keeps everything except the extreme corner diagonal frequencies.
func LowPass(freq [][]complex128, cutoff float64) ([][]complex128, error) {
	if cutoff < 0 || cutoff > 1 {
		return nil, fmt.Errorf("freqfilter: cutoff must be in [0,1], got %v", cutoff)
	}

	h := len(freq)
	w := len(freq[0])
	halfH := float64(h) / 2
	halfW := float64(w) / 2
	centerRow := h / 2
	centerCol := w / 2

	out := make([][]complex128, h)
	for i := range freq {
		out[i] = make([]complex128, w)
		u := float64(i-centerRow) / halfH
		for j := range freq[i] {
			v := float64(j-centerCol) / halfW
			r := math.Sqrt(u*u + v*v)
			if r <= cutoff {
				out[i][j] = freq[i][j]
			}
		}
	}
	return out, nil
}

// ApplyLowPass runs the full FFT -> shift -> low-pass -> unshift -> IFFT
// pipeline on a single real-valued H x W image channel, returning the
// filtered channel as a real H x W matrix.
func ApplyLowPass(channel [][]float64, cutoff float64) ([][]float64, error) {
	h := len(channel)
	w := len(channel[0])

	freq := FFT2D(channel)
	shifted := FFTShift(freq)

	masked, err := LowPass(shifted, cutoff)
	if err != nil {
		return nil, err
	}

	unshifted := IFFTShift(masked)
	spatial := IFFT2D(unshifted)

	out := make([][]float64, h)
	for i := range spatial {
		out[i] = make([]float64, w)
		for j := range spatial[i] {
			out[i][j] = real(spatial[i][j])
		}
	}
	return out, nil
}
