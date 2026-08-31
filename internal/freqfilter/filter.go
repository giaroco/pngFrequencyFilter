package freqfilter

import (
	"fmt"
	"math"
)

// BandPass zeroes out every bin in an already FFTShift-ed H x W
// frequency-domain grid whose normalized radial distance from the center
// (DC) falls outside [obstruction, cutoff]. The radius for bin (i,j) is
// normalized per axis against that axis' Nyquist frequency, so it is 1.0
// at the edge midpoints and up to sqrt(2) at the corners:
//
//	u := (i - H/2) / (H/2)
//	v := (j - W/2) / (W/2)
//	r := sqrt(u^2 + v^2)
//
// cutoff is the existing low-pass radius: bins with r > cutoff are
// zeroed. obstruction simulates a circular obstruction blocking the
// low frequencies at the center of the u/v plane: bins with r <
// obstruction (including DC itself when obstruction > 0) are zeroed.
//
// Both obstruction and cutoff must be in [0,1]. obstruction=0 disables
// the obstruction (the DC bin and its immediate neighborhood are kept),
// giving the original low-pass-only behavior.
func BandPass(freq [][]complex128, obstruction, cutoff float64) ([][]complex128, error) {
	if obstruction < 0 || obstruction > 1 {
		return nil, fmt.Errorf("freqfilter: obstruction must be in [0,1], got %v", obstruction)
	}
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
			if r >= obstruction && r <= cutoff {
				out[i][j] = freq[i][j]
			}
		}
	}
	return out, nil
}

// ApplyBandPass runs the full FFT -> shift -> band-pass -> unshift -> IFFT
// pipeline on a single real-valued H x W image channel, returning the
// filtered channel as a real H x W matrix. See BandPass for the meaning
// of obstruction and cutoff.
func ApplyBandPass(channel [][]float64, obstruction, cutoff float64) ([][]float64, error) {
	h := len(channel)
	w := len(channel[0])

	freq := FFT2D(channel)
	shifted := FFTShift(freq)

	masked, err := BandPass(shifted, obstruction, cutoff)
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
