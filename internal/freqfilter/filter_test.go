package freqfilter

import "testing"

func TestBandPassZeroCutoffKeepsOnlyDC(t *testing.T) {
	h, w := 4, 6
	freq := make([][]complex128, h)
	for i := range freq {
		freq[i] = make([]complex128, w)
		for j := range freq[i] {
			freq[i][j] = complex(1, 1)
		}
	}

	masked, err := BandPass(freq, 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	centerRow, centerCol := h/2, w/2
	for i := range masked {
		for j := range masked[i] {
			if i == centerRow && j == centerCol {
				if masked[i][j] != complex(1, 1) {
					t.Fatalf("expected DC bin to be preserved, got %v", masked[i][j])
				}
				continue
			}
			if masked[i][j] != 0 {
				t.Fatalf("expected bin (%d,%d) to be zeroed, got %v", i, j, masked[i][j])
			}
		}
	}
}

func TestBandPassFullCutoffKeepsAxisAlignedEdges(t *testing.T) {
	h, w := 4, 4
	freq := make([][]complex128, h)
	for i := range freq {
		freq[i] = make([]complex128, w)
		for j := range freq[i] {
			freq[i][j] = complex(1, 0)
		}
	}

	masked, err := BandPass(freq, 0, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// The bin directly above center (axis-aligned, radius exactly 1.0) must survive.
	if masked[0][w/2] == 0 {
		t.Fatalf("expected axis-aligned edge bin (0,%d) to survive cutoff=1.0", w/2)
	}
	// The corner bin (radius > 1.0) must be zeroed.
	if masked[0][0] != 0 {
		t.Fatalf("expected corner bin (0,0) to be zeroed at cutoff=1.0, got %v", masked[0][0])
	}
}

func TestBandPassObstructionBlocksCenter(t *testing.T) {
	h, w := 8, 8
	freq := make([][]complex128, h)
	for i := range freq {
		freq[i] = make([]complex128, w)
		for j := range freq[i] {
			freq[i][j] = complex(1, 0)
		}
	}

	const obstruction = 0.5
	masked, err := BandPass(freq, obstruction, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	centerRow, centerCol := h/2, w/2

	// DC itself must be blocked by any positive obstruction.
	if masked[centerRow][centerCol] != 0 {
		t.Fatalf("expected DC bin to be zeroed by obstruction=%v, got %v", obstruction, masked[centerRow][centerCol])
	}
	// A bin just outside the obstruction radius but well within cutoff must survive.
	if masked[0][centerCol] == 0 {
		t.Fatalf("expected bin (0,%d) outside the obstruction to survive", centerCol)
	}
}

func TestBandPassObstructionDisabledByDefault(t *testing.T) {
	h, w := 4, 6
	freq := make([][]complex128, h)
	for i := range freq {
		freq[i] = make([]complex128, w)
		for j := range freq[i] {
			freq[i][j] = complex(1, 1)
		}
	}

	masked, err := BandPass(freq, 0, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	centerRow, centerCol := h/2, w/2
	if masked[centerRow][centerCol] != complex(1, 1) {
		t.Fatalf("expected DC bin to be preserved when obstruction=0, got %v", masked[centerRow][centerCol])
	}
}

func TestBandPassInvalidParams(t *testing.T) {
	freq := [][]complex128{{1, 1}, {1, 1}}
	if _, err := BandPass(freq, -0.1, 1.0); err == nil {
		t.Fatal("expected error for obstruction below 0")
	}
	if _, err := BandPass(freq, 1.1, 1.0); err == nil {
		t.Fatal("expected error for obstruction above 1")
	}
	if _, err := BandPass(freq, 0, -0.1); err == nil {
		t.Fatal("expected error for cutoff below 0")
	}
	if _, err := BandPass(freq, 0, 1.1); err == nil {
		t.Fatal("expected error for cutoff above 1")
	}
}

func TestApplyBandPassNoOpAtFullCutoffCornersOnly(t *testing.T) {
	// A smooth (low-frequency-dominant, pure-DC) channel should survive a
	// generous cutoff with no obstruction almost exactly.
	h, w := 8, 8
	channel := make([][]float64, h)
	for i := range channel {
		channel[i] = make([]float64, w)
		for j := range channel[i] {
			channel[i][j] = 128
		}
	}

	out, err := ApplyBandPass(channel, 0, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := range channel {
		for j := range channel[i] {
			diff := out[i][j] - channel[i][j]
			if diff < -1e-6 || diff > 1e-6 {
				t.Fatalf("mismatch at (%d,%d): got %v, want %v", i, j, out[i][j], channel[i][j])
			}
		}
	}
}

func TestApplyBandPassObstructionFlattensPureDCChannel(t *testing.T) {
	// A flat (pure-DC) channel has all its energy at r=0, so any positive
	// obstruction should blank it out to (approximately) zero.
	h, w := 8, 8
	channel := make([][]float64, h)
	for i := range channel {
		channel[i] = make([]float64, w)
		for j := range channel[i] {
			channel[i][j] = 128
		}
	}

	out, err := ApplyBandPass(channel, 0.1, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := range channel {
		for j := range channel[i] {
			if out[i][j] < -1e-6 || out[i][j] > 1e-6 {
				t.Fatalf("mismatch at (%d,%d): got %v, want ~0", i, j, out[i][j])
			}
		}
	}
}
