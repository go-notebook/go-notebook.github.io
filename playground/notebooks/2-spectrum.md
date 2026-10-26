# Finding a tone in noise

The example of go-notebook's README: a signal of two sines, at 50 Hz and 120 Hz, buried in Gaussian noise. Its discrete Fourier transform shows the two frequencies as peaks. The arrays, the transform and the plot are [go-ndarray](https://github.com/go-ndarray/ndarray) and [go-plot](https://github.com/go-plot/plot), compiled: only the cells are interpreted.

```go
import (
	"math"

	"github.com/go-ndarray/ndarray"
	"github.com/go-ndarray/ndarray/fft"
	"github.com/go-ndarray/ndarray/random"
	"github.com/go-plot/plot"
)

const (
	n    = 4096   // samples
	rate = 1000.0 // samples per second
	f1   = 50.0   // Hz
	f2   = 120.0  // Hz
)
```

The time axis, then the clean signal.

```go
t, _ := ndarray.Arange(0, n, 1)
t = t.DivScalar(rate)

s1 := t.MulScalar(2 * math.Pi * f1).Sin()
s2 := t.MulScalar(2 * math.Pi * f2).Sin().MulScalar(0.5)
clean, _ := s1.Add(s2)
clean.Slice(ndarray.To(5))
```

Noise from a seeded generator, so the notebook gives the same numbers on every run (and NumPy's numbers for `default_rng(7)`).

```go
noise, _ := random.DefaultRNG(7).Normal(0, 2, n)
x, _ := clean.Add(noise)
x.Slice(ndarray.To(5))
```

The spectrum. The noise is four times the weaker tone's amplitude, yet both peaks stand out.

```go
spectrum, _ := fft.RFFT(x)
power := spectrum.Abs()
freqs, _ := fft.RFFTFreq(n, 1/rate)

peak, _ := power.ArgMax()
freqs.At(peak)
```

The second tone: the largest bin away from the first one.

```go
masked, _ := freqs.Sub(ndarray.Scalar(freqs.At(peak)))
far, _ := masked.Abs().Greater(ndarray.Scalar(5.0))
rest, _ := ndarray.Where(far, power, ndarray.Scalar(0.0))
second, _ := rest.ArgMax()
freqs.At(second)
```

The signal and its spectrum.

```go
fig, axs := plot.Subplots(2, 1, plot.FigSize(7, 5))
top, bottom := axs[0][0], axs[1][0]

t250, _ := t.Slice(ndarray.To(250))
x250, _ := x.Slice(ndarray.To(250))
clean250, _ := clean.Slice(ndarray.To(250))
top.Plot(t250, x250, plot.Label("noisy"))
top.Plot(t250, clean250, plot.Label("clean"))
top.SetXLabel("time (s)")
top.Legend()

bottom.Plot(freqs, power)
bottom.SetXLim(0, 200)
bottom.SetXLabel("frequency (Hz)")
fig.TightLayout()
fig
```
