# benchmarks-bls-libs

a set of benchmarks for bls libs

## Run bencmarks

```bash
make bench
```
## BenchMarks

```txt
BenchmarkBlstSign-8                                 1722            662231 ns/op             480 B/op          2 allocs/op
BenchmarkHerumiSign-8                               1567            665471 ns/op             304 B/op          2 allocs/op
BenchmarkBlstVerify-8                                724           1574097 ns/op            4177 B/op         14 allocs/op
BenchmarkHerumiVerify-8                              628           1793354 ns/op              16 B/op          1 allocs/op
BenchmarkBlstAggregateVerify10-8                     322           3833003 ns/op           27459 B/op         25 allocs/op
BenchmarkHerumiAggregateVerify10-8                   141           9082633 ns/op             602 B/op          2 allocs/op
BenchmarkBlstAggregateVerify100-8                     63          23694837 ns/op           27329 B/op         25 allocs/op
BenchmarkHerumiAggregateVerify100-8                   15          71359633 ns/op            5578 B/op          8 allocs/op
BenchmarkBlstAggregateVerify1000-8                     5         225515180 ns/op           27347 B/op         25 allocs/op
BenchmarkHerumiAggregateVerify1000-8                   2         725578500 ns/op           82032 B/op          6 allocs/op
BenchmarkBlstFastAggregateVerify10-8                 667           1514945 ns/op            6434 B/op         36 allocs/op
BenchmarkHerumiFastAggregateVerify10-8               800           1681933 ns/op               0 B/op          0 allocs/op
BenchmarkBlstFastAggregateVerify100-8                798           1775963 ns/op            6435 B/op         36 allocs/op
BenchmarkHerumiFastAggregateVerify100-8              667           2018462 ns/op               0 B/op          0 allocs/op
BenchmarkBlstFastAggregateVerify1000-8               536           2254356 ns/op            6424 B/op         36 allocs/op
BenchmarkHerumiFastAggregateVerify1000-8             393           2627733 ns/op               0 B/op          0 allocs/op
```

The following collation of a detailed table, because we only focus on the running speed data, memory temporarily not concerned, because the content of the bounty at the time was mainly introduced blst than herumi 3 times faster：

| BLS library        |   Herumi  |  blst  |
| --------           | -----:   | ----: |
| Single Sign        | 665471 ns/op      |  662231 ns/op    |
| Single Verify       | 1793354 ns/op      |   1574097 ns/op    |
| num 10 Aggregated Verify        | 9082633 ns/op      |   3833003 ns/op    |
| num 100 Aggregated Verify        | 71359633 ns/op      |   23694837 ns/op   |
| num 1000 Aggregated Verify        | 725578500 ns/op     |   225515180 ns/op     |
| num 10 Fast Aggregated Verify        | 1681933 ns/op     |  1514945 ns/op     |
| num 100 Fast Aggregated Verify        | 2018462 ns/op       |   1775963 ns/op    |
| num 1000 Fast Aggregated Verify        | 2627733 ns/op     |   2254356 ns/op      |

From the above table, we can see that blst's library is faster than herumi library in both generate of signature and verification of signature, but only the performance of ordinary aggregated verification of signature is 3 times faster than herumi library, if it is fast aggregated verification of signature, blst library is obviously not 3 times faster than herumi library, the advantage is not great.

Of course, all in all, the performance advantages of the blst library over the herumi library are comprehensive

## Other BLS libraries

the other librarys（Milagro, MIRACL, RELIC） mentioned by https://blog.quarkslab.com/technical-assessment-of-the-herumi-libraries.html is built by C, ASM, Rust etc. and there is no official bindings for Go language. And in terms of documentation, these libraries are very poorly documented compared to the blst and herumi libraries. So on the same premise, it is not convenient to do benchmark.