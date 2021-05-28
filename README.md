# benchmarks-bls-libs

a set of benchmarks for bls libs

## Run bencmarks

```bash
# for herumi and blst
make blsbench

# for milagro
cd incubator-milagro-crypto-rust
cargo bench --features bench
```
## BenchMarks

The following collation of a detailed table, because we only focus on the running speed data, memory temporarily not concerned, because the content of the bounty at the time was mainly introduced blst than herumi 3 times faster：

| BLS library        |   Herumi  |  blst  |    milagro(BLS381, G2 signature) |     relic    |
| --------           | -----:   | ----: |  ----: |            -----:     |
| Single Sign        | 665471 ns/op      |  662231 ns/op    |    1574800 ns/op |     NULL   |
| Single Verify       | 1793354 ns/op      |   1574097 ns/op    |   5380100 ns/op |    NULL  |
| num 10 Aggregated Verify        | 9082633 ns/op      |   3833003 ns/op    |     22576000 ns/op |    NULL | 
| num 100 Aggregated Verify        | 71359633 ns/op      |   23694837 ns/op   |     194460000 ns/op |   NULL |
| num 1000 Aggregated Verify        | 725578500 ns/op     |   225515180 ns/op     |  NULL |      NULL |
| num 10 Fast Aggregated Verify        | 1681933 ns/op     |  1514945 ns/op     |     NULL  |   NULL |
| num 100 Fast Aggregated Verify        | 2018462 ns/op       |   1775963 ns/op    |   NULL |   NULL |
| num 1000 Fast Aggregated Verify        | 2627733 ns/op     |   2254356 ns/op      |   NULL  |    NULL

From the above table, we can see that blst's library is faster than herumi library in both generate of signature and verification of signature, but only the performance of ordinary aggregated verification of signature is 3 times faster than herumi library, if it is fast aggregated verification of signature, blst library is obviously not 3 times faster than herumi library, the advantage is not great.

Of course, all in all, the performance advantages of the blst library over the herumi library are comprehensive

But for the milagro library, the performance is really poor, an order of magnitude worse than the two previous libraries, and it may not only not support Fast Aggregated Verify, but also has some other unavailable issues

## Issues for other BLS libraries

the other librarys（Milagro, MIRACL, RELIC） mentioned by https://blog.quarkslab.com/technical-assessment-of-the-herumi-libraries.html is built by C, ASM, Rust etc. and there is no official bindings for Go language. And in terms of documentation, these libraries are very poorly documented compared to the blst and herumi libraries. So on the same premise, it is not convenient to do benchmark.

### Milagro

Just use Milagro official benchmarks and change some arguments

Milagro maybe does not support Fast Aggregated Verify like herumi and blst（https://github.com/apache/incubator-milagro-crypto-rust/blob/4e0d0c60a18c4b418932145002f12ad672f7f4de/src/bls381/core.rs#L756-L759）, So I create a issue to ask question: https://github.com/apache/incubator-milagro-crypto-rust/issues/49, and when Aggregated number is 1000, the benchmark program just paniced

```txt
Benchmarking aggregation/Verifying aggregate of 1000 signatures G2 on BLS381: Warming up for 3.0000 sthread 'main' panicked at 'assertion failed: basic::aggregate_verify_g2(&pks_g2_refs, &msgs_refs, &agg_sig_g2)', benches/bls381.rs:283:17
stack backtrace:
   0: rust_begin_unwind
             at /rustc/9bc8c42bb2f19e745a63f3445f1ac248fb015e53/library/std/src/panicking.rs:493:5
   1: core::panicking::panic_fmt
             at /rustc/9bc8c42bb2f19e745a63f3445f1ac248fb015e53/library/core/src/panicking.rs:92:14
   2: core::panicking::panic
             at /rustc/9bc8c42bb2f19e745a63f3445f1ac248fb015e53/library/core/src/panicking.rs:50:5
   3: criterion::Bencher<M>::iter
   4: <criterion::routine::Function<M,F,T> as criterion::routine::Routine<M,T>>::warm_up
   5: criterion::routine::Routine::sample
   6: criterion::analysis::common
   7: <criterion::benchmark::Benchmark<M> as criterion::benchmark::BenchmarkDefinition<M>>::run
   8: bls381::aggregate_verfication_n
   9: bls381::main
note: Some details are omitted, run with `RUST_BACKTRACE=full` for a verbose backtrace.
error: bench failed
```

This library performance and feature is not good

### RELIC

Just search bench and test folder of this project, There is no examples related to BLS12-381 for aggregated verification(It's BN Curve, another pairing-friendly curve), and I also found out that the maintainer of the project does not recommend this library for production environments and does not support aggregation(need to implement it by ourself according to paper), the issue is here https://github.com/relic-toolkit/relic/issues/65

so RELIC library is not good even compared to Milagro



