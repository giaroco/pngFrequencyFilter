// Command pngfreqfilter loads a colored PNG, band-pass filters it in the
// frequency domain (via a 2D FFT per color channel) — optionally cutting
// off high frequencies and/or blocking low frequencies with a circular
// obstruction at the center of the u/v plane — and writes the filtered
// result back out as a PNG.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/giaroco/pngFrequencyFilter/internal/freqfilter"
	"github.com/giaroco/pngFrequencyFilter/internal/imageio"
)

func main() {
	in := flag.String("in", "", "input PNG path (required)")
	out := flag.String("out", "", "output PNG path (required)")
	cutoff := flag.Float64("cutoff", 1.0, "normalized frequency cutoff in [0,1]; 0 keeps only the DC component (flat average color), 1 keeps almost everything")
	obstruction := flag.Float64("obstruction", 0.0, "normalized radius in [0,1] of a circular obstruction blocking low frequencies at the center of the u/v plane; 0 disables the obstruction")
	flag.Parse()

	if *in == "" || *out == "" {
		flag.Usage()
		log.Fatal("both -in and -out are required")
	}

	img, err := imageio.Load(*in)
	if err != nil {
		log.Fatalf("loading image: %v", err)
	}

	r, g, b, a, w, h := imageio.SplitChannels(img)

	fr, err := freqfilter.ApplyBandPass(r, *obstruction, *cutoff)
	if err != nil {
		log.Fatalf("filtering red channel: %v", err)
	}
	fg, err := freqfilter.ApplyBandPass(g, *obstruction, *cutoff)
	if err != nil {
		log.Fatalf("filtering green channel: %v", err)
	}
	fb, err := freqfilter.ApplyBandPass(b, *obstruction, *cutoff)
	if err != nil {
		log.Fatalf("filtering blue channel: %v", err)
	}

	result := imageio.MergeChannels(fr, fg, fb, a, w, h)

	if err := imageio.Save(*out, result); err != nil {
		log.Fatalf("saving image: %v", err)
	}

	fmt.Printf("wrote %s (%dx%d, obstruction=%.3f, cutoff=%.3f)\n", *out, w, h, *obstruction, *cutoff)
}
