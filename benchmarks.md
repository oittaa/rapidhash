# Benchmarks

This isn't super optimized implementation, but here are some of the things that have been tried.

```
go test -bench=BenchmarkRapidhash -count=10 ./... > before.txt
# Make edits
go test -bench=BenchmarkRapidhash -count=10 ./... > after.txt
~/go/bin/benchstat before.txt after.txt
```

## Using pointers in mum()

before
```go
func mum(a, b uint64) (uint64, uint64) {
	hi, lo := bits.Mul64(a, b)
	return lo, hi
}
func mix(a, b uint64) uint64 {
	lo, hi := mum(a, b)
	return lo ^ hi
}
```

after
```go
func mum(a, b *uint64) {
	hi, lo := bits.Mul64(*a, *b)
	*a = lo
	*b = hi
}
func mix(a, b uint64) uint64 {
	mum(&a, &b)
	return a ^ b
}
```


```
goos: linux
goarch: amd64
pkg: rapidhash
cpu: AMD Ryzen 7 9800X3D 8-Core Processor
                      │  before.txt  │              after.txt              │
                      │    sec/op    │   sec/op     vs base                │
Rapidhash4B-16           2.311n ± 0%   2.768n ± 0%  +19.75% (p=0.000 n=10)
Rapidhash8B-16           2.361n ± 0%   2.742n ± 0%  +16.11% (p=0.000 n=10)
Rapidhash12B-16          2.363n ± 1%   2.742n ± 0%  +16.06% (p=0.000 n=10)
Rapidhash16B-16          2.359n ± 0%   2.739n ± 0%  +16.11% (p=0.000 n=10)
Rapidhash32B-16          3.205n ± 0%   3.484n ± 0%   +8.72% (p=0.000 n=10)
Rapidhash64B-16          4.878n ± 0%   5.140n ± 0%   +5.37% (p=0.000 n=10)
Rapidhash128B-16        15.390n ± 3%   7.708n ± 0%  -49.91% (p=0.000 n=10)
Rapidhash256B-16         12.53n ± 0%   12.49n ± 1%        ~ (p=0.230 n=10)
Rapidhash1KB-16          42.78n ± 0%   42.95n ± 1%   +0.40% (p=0.000 n=10)
Rapidhash10KB-16         412.8n ± 0%   410.4n ± 1%   -0.57% (p=0.009 n=10)
RapidhashMicro4B-16      2.309n ± 1%   2.717n ± 0%  +17.67% (p=0.000 n=10)
RapidhashMicro8B-16      2.305n ± 0%   2.716n ± 0%  +17.85% (p=0.000 n=10)
RapidhashMicro12B-16     2.304n ± 1%   2.715n ± 0%  +17.79% (p=0.000 n=10)
RapidhashMicro16B-16     2.304n ± 1%   2.717n ± 0%  +17.90% (p=0.000 n=10)
RapidhashMicro32B-16     3.177n ± 0%   3.466n ± 0%   +9.08% (p=0.000 n=10)
RapidhashMicro64B-16     4.924n ± 0%   5.178n ± 0%   +5.17% (p=0.000 n=10)
RapidhashMicro128B-16    8.553n ± 0%   8.571n ± 0%   +0.20% (p=0.027 n=10)
RapidhashMicro256B-16    15.29n ± 1%   15.27n ± 0%        ~ (p=0.483 n=10)
RapidhashMicro1KB-16     54.30n ± 1%   53.93n ± 0%   -0.67% (p=0.000 n=10)
RapidhashMicro10KB-16    524.7n ± 1%   524.1n ± 0%        ~ (p=0.159 n=10)
RapidhashNano4B-16       2.314n ± 1%   2.722n ± 0%  +17.61% (p=0.000 n=10)
RapidhashNano8B-16       2.296n ± 1%   2.719n ± 0%  +18.40% (p=0.000 n=10)
RapidhashNano12B-16      2.305n ± 0%   2.723n ± 1%  +18.13% (p=0.000 n=10)
RapidhashNano16B-16      2.303n ± 1%   2.717n ± 0%  +18.00% (p=0.000 n=10)
RapidhashNano32B-16      3.265n ± 1%   3.496n ± 0%   +7.06% (p=0.000 n=10)
RapidhashNano64B-16      5.193n ± 0%   5.489n ± 1%   +5.69% (p=0.000 n=10)
RapidhashNano128B-16     8.361n ± 2%   8.587n ± 0%   +2.70% (p=0.001 n=10)
RapidhashNano256B-16     15.31n ± 1%   15.64n ± 0%   +2.16% (p=0.001 n=10)
RapidhashNano1KB-16      55.02n ± 1%   56.26n ± 1%   +2.26% (p=0.000 n=10)
RapidhashNano10KB-16     537.3n ± 0%   553.5n ± 1%   +3.01% (p=0.000 n=10)
geomean                  8.371n        8.871n        +5.97%

                      │  before.txt  │               after.txt               │
                      │     B/s      │      B/s       vs base                │
Rapidhash4B-16          1.612Gi ± 0%    1.346Gi ± 0%  -16.50% (p=0.000 n=10)
Rapidhash8B-16          3.155Gi ± 0%    2.718Gi ± 0%  -13.87% (p=0.000 n=10)
Rapidhash12B-16         4.731Gi ± 1%    4.076Gi ± 0%  -13.84% (p=0.000 n=10)
Rapidhash16B-16         6.317Gi ± 0%    5.440Gi ± 0%  -13.87% (p=0.000 n=10)
Rapidhash32B-16         9.301Gi ± 0%    8.554Gi ± 0%   -8.04% (p=0.000 n=10)
Rapidhash64B-16         12.22Gi ± 0%    11.60Gi ± 0%   -5.10% (p=0.000 n=10)
Rapidhash128B-16        7.745Gi ± 3%   15.465Gi ± 0%  +99.67% (p=0.000 n=10)
Rapidhash256B-16        19.03Gi ± 0%    19.10Gi ± 1%        ~ (p=0.218 n=10)
Rapidhash1KB-16         22.29Gi ± 0%    22.20Gi ± 1%   -0.40% (p=0.000 n=10)
Rapidhash10KB-16        23.10Gi ± 0%    23.24Gi ± 1%   +0.57% (p=0.009 n=10)
RapidhashMicro4B-16     1.613Gi ± 1%    1.371Gi ± 0%  -15.01% (p=0.000 n=10)
RapidhashMicro8B-16     3.232Gi ± 0%    2.743Gi ± 0%  -15.15% (p=0.000 n=10)
RapidhashMicro12B-16    4.850Gi ± 1%    4.116Gi ± 0%  -15.12% (p=0.000 n=10)
RapidhashMicro16B-16    6.466Gi ± 1%    5.484Gi ± 0%  -15.18% (p=0.000 n=10)
RapidhashMicro32B-16    9.380Gi ± 0%    8.598Gi ± 0%   -8.33% (p=0.000 n=10)
RapidhashMicro64B-16    12.11Gi ± 0%    11.51Gi ± 0%   -4.91% (p=0.000 n=10)
RapidhashMicro128B-16   13.94Gi ± 0%    13.91Gi ± 0%   -0.21% (p=0.029 n=10)
RapidhashMicro256B-16   15.60Gi ± 1%    15.61Gi ± 0%        ~ (p=0.529 n=10)
RapidhashMicro1KB-16    17.56Gi ± 1%    17.68Gi ± 0%   +0.68% (p=0.000 n=10)
RapidhashMicro10KB-16   18.18Gi ± 1%    18.20Gi ± 0%        ~ (p=0.165 n=10)
RapidhashNano4B-16      1.610Gi ± 1%    1.369Gi ± 0%  -14.97% (p=0.000 n=10)
RapidhashNano8B-16      3.245Gi ± 1%    2.740Gi ± 0%  -15.56% (p=0.000 n=10)
RapidhashNano12B-16     4.849Gi ± 0%    4.104Gi ± 1%  -15.36% (p=0.000 n=10)
RapidhashNano16B-16     6.471Gi ± 1%    5.485Gi ± 0%  -15.24% (p=0.000 n=10)
RapidhashNano32B-16     9.127Gi ± 1%    8.525Gi ± 0%   -6.59% (p=0.000 n=10)
RapidhashNano64B-16     11.48Gi ± 0%    10.86Gi ± 1%   -5.38% (p=0.000 n=10)
RapidhashNano128B-16    14.26Gi ± 2%    13.88Gi ± 0%   -2.63% (p=0.001 n=10)
RapidhashNano256B-16    15.57Gi ± 1%    15.24Gi ± 0%   -2.10% (p=0.002 n=10)
RapidhashNano1KB-16     17.33Gi ± 1%    16.95Gi ± 1%   -2.21% (p=0.000 n=10)
RapidhashNano10KB-16    17.75Gi ± 0%    17.23Gi ± 1%   -2.91% (p=0.000 n=10)
geomean                 8.126Gi         7.668Gi        -5.63%

```

Big speedup in Rapidhash128B, but every other case would be slower.
Not implemented at the moment.

## Breaking out of the loop at the end

before
```go
			for i > 112 {
				// ....
				p = p[112:]
				i -= 112
			}
```

after
```go
			for {
				// ....
				p = p[112:]
				i -= 112
				if i <= 112 {
					break
				}
			}
```

```
goos: linux
goarch: amd64
pkg: rapidhash
cpu: AMD Ryzen 7 9800X3D 8-Core Processor
                      │  before.txt  │              after.txt              │
                      │    sec/op    │   sec/op     vs base                │
Rapidhash4B-16           2.311n ± 0%   2.316n ± 1%        ~ (p=0.236 n=10)
Rapidhash8B-16           2.361n ± 0%   2.317n ± 1%   -1.91% (p=0.000 n=10)
Rapidhash12B-16          2.363n ± 1%   2.315n ± 1%   -2.01% (p=0.000 n=10)
Rapidhash16B-16          2.359n ± 0%   2.327n ± 2%        ~ (p=0.224 n=10)
Rapidhash32B-16          3.205n ± 0%   3.185n ± 1%   -0.61% (p=0.005 n=10)
Rapidhash64B-16          4.878n ± 0%   4.877n ± 0%        ~ (p=0.616 n=10)
Rapidhash128B-16        15.390n ± 3%   7.330n ± 2%  -52.37% (p=0.000 n=10)
Rapidhash256B-16         12.53n ± 0%   12.20n ± 0%   -2.63% (p=0.000 n=10)
Rapidhash1KB-16          42.78n ± 0%   40.61n ± 0%   -5.07% (p=0.000 n=10)
Rapidhash10KB-16         412.8n ± 0%   384.6n ± 2%   -6.83% (p=0.000 n=10)
RapidhashMicro4B-16      2.309n ± 1%   2.366n ± 1%   +2.47% (p=0.001 n=10)
RapidhashMicro8B-16      2.305n ± 0%   2.380n ± 0%   +3.25% (p=0.000 n=10)
RapidhashMicro12B-16     2.304n ± 1%   2.372n ± 1%   +2.93% (p=0.001 n=10)
RapidhashMicro16B-16     2.304n ± 1%   2.374n ± 0%   +2.99% (p=0.001 n=10)
RapidhashMicro32B-16     3.177n ± 0%   3.238n ± 0%   +1.90% (p=0.000 n=10)
RapidhashMicro64B-16     4.924n ± 0%   4.942n ± 0%   +0.38% (p=0.010 n=10)
RapidhashMicro128B-16    8.553n ± 0%   8.579n ± 0%   +0.30% (p=0.019 n=10)
RapidhashMicro256B-16    15.29n ± 1%   15.39n ± 0%   +0.72% (p=0.001 n=10)
RapidhashMicro1KB-16     54.30n ± 1%   55.46n ± 1%   +2.14% (p=0.000 n=10)
RapidhashMicro10KB-16    524.7n ± 1%   542.6n ± 0%   +3.41% (p=0.000 n=10)
RapidhashNano4B-16       2.314n ± 1%   2.321n ± 0%        ~ (p=0.084 n=10)
RapidhashNano8B-16       2.296n ± 1%   2.314n ± 0%   +0.74% (p=0.003 n=10)
RapidhashNano12B-16      2.305n ± 0%   2.312n ± 1%   +0.33% (p=0.011 n=10)
RapidhashNano16B-16      2.303n ± 1%   2.317n ± 1%   +0.61% (p=0.011 n=10)
RapidhashNano32B-16      3.265n ± 1%   3.285n ± 1%   +0.63% (p=0.015 n=10)
RapidhashNano64B-16      5.193n ± 0%   5.184n ± 1%        ~ (p=0.060 n=10)
RapidhashNano128B-16     8.361n ± 2%   8.471n ± 0%   +1.32% (p=0.014 n=10)
RapidhashNano256B-16     15.31n ± 1%   15.35n ± 0%        ~ (p=0.402 n=10)
RapidhashNano1KB-16      55.02n ± 1%   55.44n ± 0%   +0.77% (p=0.000 n=10)
RapidhashNano10KB-16     537.3n ± 0%   545.0n ± 0%   +1.44% (p=0.000 n=10)
geomean                  8.371n        8.182n        -2.26%

                      │  before.txt  │               after.txt                │
                      │     B/s      │      B/s       vs base                 │
Rapidhash4B-16          1.612Gi ± 0%    1.609Gi ± 1%         ~ (p=0.218 n=10)
Rapidhash8B-16          3.155Gi ± 0%    3.217Gi ± 1%    +1.95% (p=0.000 n=10)
Rapidhash12B-16         4.731Gi ± 1%    4.828Gi ± 1%    +2.06% (p=0.000 n=10)
Rapidhash16B-16         6.317Gi ± 0%    6.402Gi ± 2%         ~ (p=0.218 n=10)
Rapidhash32B-16         9.301Gi ± 0%    9.358Gi ± 1%    +0.61% (p=0.007 n=10)
Rapidhash64B-16         12.22Gi ± 0%    12.22Gi ± 0%         ~ (p=0.631 n=10)
Rapidhash128B-16        7.745Gi ± 3%   16.262Gi ± 2%  +109.96% (p=0.000 n=10)
Rapidhash256B-16        19.03Gi ± 0%    19.55Gi ± 0%    +2.71% (p=0.000 n=10)
Rapidhash1KB-16         22.29Gi ± 0%    23.48Gi ± 0%    +5.35% (p=0.000 n=10)
Rapidhash10KB-16        23.10Gi ± 0%    24.80Gi ± 2%    +7.34% (p=0.000 n=10)
RapidhashMicro4B-16     1.613Gi ± 1%    1.574Gi ± 1%    -2.43% (p=0.002 n=10)
RapidhashMicro8B-16     3.232Gi ± 0%    3.130Gi ± 0%    -3.16% (p=0.000 n=10)
RapidhashMicro12B-16    4.850Gi ± 1%    4.711Gi ± 0%    -2.86% (p=0.001 n=10)
RapidhashMicro16B-16    6.466Gi ± 1%    6.278Gi ± 0%    -2.89% (p=0.000 n=10)
RapidhashMicro32B-16    9.380Gi ± 0%    9.203Gi ± 0%    -1.89% (p=0.000 n=10)
RapidhashMicro64B-16    12.11Gi ± 0%    12.06Gi ± 0%    -0.38% (p=0.011 n=10)
RapidhashMicro128B-16   13.94Gi ± 0%    13.90Gi ± 0%    -0.30% (p=0.015 n=10)
RapidhashMicro256B-16   15.60Gi ± 1%    15.49Gi ± 1%    -0.71% (p=0.002 n=10)
RapidhashMicro1KB-16    17.56Gi ± 1%    17.20Gi ± 1%    -2.08% (p=0.000 n=10)
RapidhashMicro10KB-16   18.18Gi ± 1%    17.58Gi ± 0%    -3.29% (p=0.000 n=10)
RapidhashNano4B-16      1.610Gi ± 1%    1.605Gi ± 0%         ~ (p=0.089 n=10)
RapidhashNano8B-16      3.245Gi ± 1%    3.220Gi ± 0%    -0.76% (p=0.003 n=10)
RapidhashNano12B-16     4.849Gi ± 0%    4.833Gi ± 1%    -0.32% (p=0.011 n=10)
RapidhashNano16B-16     6.471Gi ± 1%    6.433Gi ± 1%    -0.60% (p=0.011 n=10)
RapidhashNano32B-16     9.127Gi ± 1%    9.070Gi ± 1%    -0.63% (p=0.015 n=10)
RapidhashNano64B-16     11.48Gi ± 0%    11.50Gi ± 1%         ~ (p=0.052 n=10)
RapidhashNano128B-16    14.26Gi ± 2%    14.07Gi ± 0%    -1.31% (p=0.015 n=10)
RapidhashNano256B-16    15.57Gi ± 1%    15.53Gi ± 0%         ~ (p=0.481 n=10)
RapidhashNano1KB-16     17.33Gi ± 1%    17.20Gi ± 0%    -0.76% (p=0.000 n=10)
RapidhashNano10KB-16    17.75Gi ± 0%    17.50Gi ± 0%    -1.43% (p=0.000 n=10)
geomean                 8.126Gi         8.314Gi         +2.31%
```

Big speedup in Rapidhash128B, and doesn't completely destroy performance with the smaller inputs.
Implemented at the moment.
