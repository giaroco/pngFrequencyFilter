# pngFrequencyFilter

A Go command-line tool that low-pass filters a PNG image in the frequency domain.

It loads a colored PNG, runs a 2D FFT on each of the red, green, and blue
channels to get their frequency-domain (u,v) planes, zeroes out every
frequency component beyond a configurable cutoff, runs a 2D inverse FFT
back to the spatial domain, and writes the result out as a new PNG. The
alpha channel is passed through unmodified.

## Build

```sh
go build -o pngfreqfilter .
```

## Usage

```sh
./pngfreqfilter -in <input.png> -out <output.png> -cutoff <0.0-1.0>
```

| Flag       | Required | Default | Description                                  |
|------------|----------|---------|-----------------------------------------------|
| `-in`      | yes      | —       | Input PNG path                                |
| `-out`     | yes      | —       | Output PNG path                               |
| `-cutoff`  | no       | `1.0`   | Normalized frequency cutoff, in `[0.0, 1.0]`  |

### Cutoff

After FFT-shifting each channel's frequency plane so DC sits at its
center, every frequency bin `(u,v)` gets a normalized radius:

```
r = sqrt((u / (H/2))^2 + (v / (W/2))^2)
```

which is `1.0` at the axis-aligned Nyquist edge and up to `~1.414` at the
corners. Any bin with `r > cutoff` is zeroed before the inverse transform.

- `-cutoff 1.0` (default) keeps nearly everything — only the extreme
  corner diagonal frequencies are cut.
- `-cutoff 0.0` keeps only the DC component, flattening each channel to
  its average color.
- Values in between progressively blur the image by removing
  higher-frequency detail.

### Example

```sh
go build -o pngfreqfilter .
./pngfreqfilter -in photo.png -out photo_blurred.png -cutoff 0.2
```

## Project layout

```
main.go                        # CLI entry point
internal/imageio/              # PNG load/save, channel split/merge
internal/freqfilter/           # 2D FFT/IFFT, fftshift, and low-pass mask
```

## Testing

```sh
go test ./...
```
