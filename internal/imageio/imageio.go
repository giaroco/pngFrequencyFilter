// Package imageio handles loading and saving PNG files and converting
// between image.NRGBA and the per-channel float64 matrices used by the
// freqfilter package.
package imageio

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// Load reads a PNG file from path and returns it as an *image.NRGBA,
// converting it if necessary.
func Load(path string) (*image.NRGBA, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("imageio: open %s: %w", path, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("imageio: decode %s: %w", path, err)
	}

	if nrgba, ok := img.(*image.NRGBA); ok {
		return nrgba, nil
	}

	b := img.Bounds()
	nrgba := image.NewNRGBA(b)
	draw.Draw(nrgba, b, img, b.Min, draw.Src)
	return nrgba, nil
}

// SplitChannels extracts the red, green, and blue channels of img as
// H x W float64 matrices (values 0-255), and flattens the alpha channel
// into a row-major []uint8 of length W*H. It also returns the image's
// width and height.
func SplitChannels(img *image.NRGBA) (r, g, b [][]float64, a []uint8, w, h int) {
	bounds := img.Bounds()
	w, h = bounds.Dx(), bounds.Dy()

	r = make([][]float64, h)
	g = make([][]float64, h)
	b = make([][]float64, h)
	a = make([]uint8, w*h)

	for y := 0; y < h; y++ {
		r[y] = make([]float64, w)
		g[y] = make([]float64, w)
		b[y] = make([]float64, w)
		for x := 0; x < w; x++ {
			px := img.NRGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			r[y][x] = float64(px.R)
			g[y][x] = float64(px.G)
			b[y][x] = float64(px.B)
			a[y*w+x] = px.A
		}
	}
	return r, g, b, a, w, h
}

// MergeChannels reassembles R, G, B float64 matrices (clamped and rounded
// to [0,255]) and an unmodified row-major alpha slice back into an
// *image.NRGBA of size w x h.
func MergeChannels(r, g, b [][]float64, a []uint8, w, h int) *image.NRGBA {
	out := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			out.SetNRGBA(x, y, color.NRGBA{
				R: clampToUint8(r[y][x]),
				G: clampToUint8(g[y][x]),
				B: clampToUint8(b[y][x]),
				A: a[y*w+x],
			})
		}
	}
	return out
}

func clampToUint8(v float64) uint8 {
	v = math.Round(v)
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

// Save encodes img as a PNG and writes it to path.
func Save(path string, img *image.NRGBA) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("imageio: create %s: %w", path, err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("imageio: encode %s: %w", path, err)
	}
	return nil
}
