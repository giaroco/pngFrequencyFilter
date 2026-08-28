package imageio

import (
	"image"
	"image/color"
	"path/filepath"
	"testing"
)

func TestSplitMergeChannelsRoundTrip(t *testing.T) {
	w, h := 4, 3
	src := image.NewNRGBA(image.Rect(0, 0, w, h))
	n := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			src.SetNRGBA(x, y, color.NRGBA{
				R: uint8(n % 256),
				G: uint8((n * 2) % 256),
				B: uint8((n * 3) % 256),
				A: uint8((n * 5) % 256),
			})
			n++
		}
	}

	r, g, b, a, gotW, gotH := SplitChannels(src)
	if gotW != w || gotH != h {
		t.Fatalf("got dims %dx%d, want %dx%d", gotW, gotH, w, h)
	}

	merged := MergeChannels(r, g, b, a, w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			want := src.NRGBAAt(x, y)
			got := merged.NRGBAAt(x, y)
			if got != want {
				t.Fatalf("pixel mismatch at (%d,%d): got %v, want %v", x, y, got, want)
			}
		}
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	w, h := 3, 2
	src := image.NewNRGBA(image.Rect(0, 0, w, h))
	src.SetNRGBA(0, 0, color.NRGBA{R: 10, G: 20, B: 30, A: 255})
	src.SetNRGBA(1, 0, color.NRGBA{R: 40, G: 50, B: 60, A: 128})
	src.SetNRGBA(2, 1, color.NRGBA{R: 70, G: 80, B: 90, A: 0})

	path := filepath.Join(t.TempDir(), "test.png")
	if err := Save(path, src); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	bounds := loaded.Bounds()
	if bounds.Dx() != w || bounds.Dy() != h {
		t.Fatalf("got dims %dx%d, want %dx%d", bounds.Dx(), bounds.Dy(), w, h)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if loaded.NRGBAAt(x, y) != src.NRGBAAt(x, y) {
				t.Fatalf("pixel mismatch at (%d,%d): got %v, want %v", x, y, loaded.NRGBAAt(x, y), src.NRGBAAt(x, y))
			}
		}
	}
}
