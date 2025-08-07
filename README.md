# rapidhash

Based on the original C implementation https://github.com/Nicoshev/rapidhash

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/oittaa/rapidhash
cpu: AMD Ryzen 7 9800X3D 8-Core Processor
BenchmarkRapidhash4B-16                 515146736                2.321 ns/op    1723.67 MB/s
BenchmarkRapidhash8B-16                 516139838                2.328 ns/op    3436.85 MB/s
BenchmarkRapidhash12B-16                518024334                2.316 ns/op    5181.98 MB/s
BenchmarkRapidhash16B-16                499713609                2.334 ns/op    6855.78 MB/s
BenchmarkRapidhash32B-16                379044699                3.190 ns/op    10032.46 MB/s
BenchmarkRapidhash64B-16                248604324                4.859 ns/op    13171.08 MB/s
BenchmarkRapidhash128B-16               162710464                7.348 ns/op    17419.92 MB/s
BenchmarkRapidhash256B-16               89311137                12.30 ns/op     20805.47 MB/s
BenchmarkRapidhash1KB-16                27597392                40.89 ns/op     25041.02 MB/s
BenchmarkRapidhash10KB-16                3093076               390.4 ns/op      26229.07 MB/s
BenchmarkRapidhashMicro4B-16            503888270                2.369 ns/op    1688.74 MB/s
BenchmarkRapidhashMicro8B-16            507778536                2.380 ns/op    3361.70 MB/s
BenchmarkRapidhashMicro12B-16           498228735                2.377 ns/op    5049.33 MB/s
BenchmarkRapidhashMicro16B-16           505840923                2.379 ns/op    6726.59 MB/s
BenchmarkRapidhashMicro32B-16           365915658                3.249 ns/op    9848.42 MB/s
BenchmarkRapidhashMicro64B-16           243734956                4.944 ns/op    12945.96 MB/s
BenchmarkRapidhashMicro128B-16          139107153                8.602 ns/op    14881.11 MB/s
BenchmarkRapidhashMicro256B-16          71132540                15.39 ns/op     16631.47 MB/s
BenchmarkRapidhashMicro1KB-16           21272820                55.52 ns/op     18442.83 MB/s
BenchmarkRapidhashMicro10KB-16           2198941               539.6 ns/op      18978.58 MB/s
BenchmarkRapidhashNano4B-16             502298546                2.324 ns/op    1721.43 MB/s
BenchmarkRapidhashNano8B-16             520876132                2.319 ns/op    3450.45 MB/s
BenchmarkRapidhashNano12B-16            520572804                2.306 ns/op    5202.95 MB/s
BenchmarkRapidhashNano16B-16            517591276                2.315 ns/op    6911.62 MB/s
BenchmarkRapidhashNano32B-16            357987770                3.283 ns/op    9746.65 MB/s
BenchmarkRapidhashNano64B-16            230143780                5.191 ns/op    12327.99 MB/s
BenchmarkRapidhashNano128B-16           142455987                8.485 ns/op    15086.24 MB/s
BenchmarkRapidhashNano256B-16           73219266                15.29 ns/op     16744.36 MB/s
BenchmarkRapidhashNano1KB-16            21143647                55.41 ns/op     18481.90 MB/s
BenchmarkRapidhashNano10KB-16            2196829               543.5 ns/op      18841.16 MB/s
BenchmarkMaphash4B-16                   562809742                2.124 ns/op    1883.38 MB/s
BenchmarkMaphash8B-16                   539521122                2.154 ns/op    3714.53 MB/s
BenchmarkMaphash12B-16                  566849648                2.120 ns/op    5659.20 MB/s
BenchmarkMaphash16B-16                  543771522                2.067 ns/op    7739.37 MB/s
BenchmarkMaphash32B-16                  564645726                2.119 ns/op    15100.71 MB/s
BenchmarkMaphash64B-16                  518455281                2.325 ns/op    27526.10 MB/s
BenchmarkMaphash128B-16                 316494427                3.760 ns/op    34040.17 MB/s
BenchmarkMaphash256B-16                 99602784                10.30 ns/op     24844.88 MB/s
BenchmarkMaphash1KB-16                  19653373                60.55 ns/op     16910.71 MB/s
BenchmarkMaphash10KB-16                  1817319               660.6 ns/op      15500.36 MB/s
PASS
ok      github.com/oittaa/rapidhash     47.197s
```
