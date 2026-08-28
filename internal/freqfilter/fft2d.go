// Package freqfilter implements a 2D FFT-based low-pass filter for image
// channels represented as real-valued matrices.
package freqfilter

import "gonum.org/v1/gonum/dsp/fourier"

// FFT2D computes the unnormalized 2D discrete Fourier transform of a real
// H x W matrix, producing an H x W complex matrix of frequency-domain
// (u,v) coefficients. It is computed as a separable transform: a 1D FFT
// across each row, followed by a 1D FFT across each column.
func FFT2D(real2d [][]float64) [][]complex128 {
	h := len(real2d)
	w := len(real2d[0])

	grid := make([][]complex128, h)
	for i := range real2d {
		grid[i] = make([]complex128, w)
		for j, v := range real2d[i] {
			grid[i][j] = complex(v, 0)
		}
	}

	rowFFT := fourier.NewCmplxFFT(w)
	for i := range grid {
		rowFFT.Coefficients(grid[i], grid[i])
	}

	colFFT := fourier.NewCmplxFFT(h)
	col := make([]complex128, h)
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			col[i] = grid[i][j]
		}
		colFFT.Coefficients(col, col)
		for i := 0; i < h; i++ {
			grid[i][j] = col[i]
		}
	}

	return grid
}

// IFFT2D computes the normalized inverse 2D discrete Fourier transform of
// an H x W complex frequency-domain matrix, producing an H x W complex
// spatial-domain matrix. For real-valued input images the imaginary part
// of the result is negligible (floating-point noise) and can be dropped
// by the caller. The transform is the inverse of FFT2D: a 1D inverse FFT
// across each column, followed by a 1D inverse FFT across each row, with
// the result normalized by H*W.
func IFFT2D(freq [][]complex128) [][]complex128 {
	h := len(freq)
	w := len(freq[0])

	grid := make([][]complex128, h)
	for i := range freq {
		grid[i] = make([]complex128, w)
		copy(grid[i], freq[i])
	}

	colFFT := fourier.NewCmplxFFT(h)
	col := make([]complex128, h)
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			col[i] = grid[i][j]
		}
		colFFT.Sequence(col, col)
		for i := 0; i < h; i++ {
			grid[i][j] = col[i]
		}
	}

	rowFFT := fourier.NewCmplxFFT(w)
	norm := complex(float64(h*w), 0)
	for i := range grid {
		rowFFT.Sequence(grid[i], grid[i])
		for j := range grid[i] {
			grid[i][j] /= norm
		}
	}

	return grid
}

// shiftIdx maps index i (0 <= i < n) in an unshifted sequence of length n
// to its position in the zero-frequency-centered (shifted) sequence.
// It mirrors gonum's (*CmplxFFT).ShiftIdx, which depends only on the
// sequence length.
func shiftIdx(i, n int) int {
	h := n / 2
	if i < h {
		return i + (n+1)/2
	}
	return i - h
}

// unshiftIdx is the inverse of shiftIdx.
func unshiftIdx(i, n int) int {
	h := (n + 1) / 2
	if i < h {
		return i + n/2
	}
	return i - h
}

// FFTShift rearranges an H x W frequency-domain grid so that the
// zero-frequency (DC) component moves from grid[0][0] to the center of
// the grid, at index (H/2, W/2). This matches the convention used by
// numpy.fft.fftshift.
func FFTShift(grid [][]complex128) [][]complex128 {
	return remapGrid(grid, shiftIdx)
}

// IFFTShift is the inverse of FFTShift: it moves the DC component from
// the center of the grid back to (0,0).
func IFFTShift(grid [][]complex128) [][]complex128 {
	return remapGrid(grid, unshiftIdx)
}

func remapGrid(grid [][]complex128, idx func(i, n int) int) [][]complex128 {
	h := len(grid)
	w := len(grid[0])
	out := make([][]complex128, h)
	for i := 0; i < h; i++ {
		si := idx(i, h)
		out[i] = make([]complex128, w)
		for j := 0; j < w; j++ {
			out[i][j] = grid[si][idx(j, w)]
		}
	}
	return out
}
