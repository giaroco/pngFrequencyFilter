// Command pngfreqfilter loads a colored PNG, low-pass filters it in the
// frequency domain (via a 2D FFT per color channel), and writes the
// filtered result back out as a PNG.
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

	fr, err := freqfilter.ApplyLowPass(r, *cutoff)
	if err != nil {
		log.Fatalf("filtering red channel: %v", err)
	}
	fg, err := freqfilter.ApplyLowPass(g, *cutoff)
	if err != nil {
		log.Fatalf("filtering green channel: %v", err)
	}
	fb, err := freqfilter.ApplyLowPass(b, *cutoff)
	if err != nil {
		log.Fatalf("filtering blue channel: %v", err)
	}

	result := imageio.MergeChannels(fr, fg, fb, a, w, h)

	if err := imageio.Save(*out, result); err != nil {
		log.Fatalf("saving image: %v", err)
	}

	fmt.Printf("wrote %s (%dx%d, cutoff=%.3f)\n", *out, w, h, *cutoff)
}
