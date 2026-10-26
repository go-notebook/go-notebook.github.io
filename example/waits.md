# Waiting times

```go
import "fmt"
import "github.com/go-ndarray/ndarray/random"

waits, _ := random.DefaultRNG(7).Gamma(2.5, 1.2, 2000)
```

Ruby fits a gamma law; Go reads the fit back.

```ruby
fit = Stats::Gamma.fit(waits, fixed: {loc: 0})
shape, scale = fit.a.round(3), fit.scale.round(3)
format("shape %.3f, scale %.3f", shape, scale)
```

```go
mean, _ := waits.Mean()
fmt.Sprintf("model %.3f, data %.3f", shape*scale, mean)
```
