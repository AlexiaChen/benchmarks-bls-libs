package main

import (
	"crypto/rand"
	"runtime"
	"testing"

	blst "github.com/supranational/blst/bindings/go"
)

// Min-PK
type blstPublicKey = blst.P1Affine
type blstSignature = blst.P2Affine
type blstAggregateSignature = blst.P2Aggregate
type blstAggregatePublicKey = blst.P1Aggregate
type blstSecretKey = blst.SecretKey

var dstMinPk = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

func init() {
	// Use all cores when testing and benchmarking
	blst.SetMaxProcs(runtime.GOMAXPROCS(0))
}

func randBLSTSecretKey() *blstSecretKey {
	var t [32]byte
	_, _ = rand.Read(t[:])
	secretKey := blst.KeyGen(t[:])
	return secretKey
}

func blstPubkey(sk *blstSecretKey) *blstPublicKey {
	return new(blstPublicKey).From(sk)
}

func BenchmarkBlstSign(b *testing.B) {
	sk := randBLSTSecretKey()
	msg := []byte("hello foo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		new(blstSignature).Sign(sk, msg, dstMinPk)
	}
}

func BenchmarkBlstVerify(b *testing.B) {

	sk := randBLSTSecretKey()
	pk := blstPubkey(sk)

	msg := []byte("hello foo")
	sig := new(blstSignature).Sign(sk, msg, dstMinPk)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sig.Verify(true, pk, true, msg, dstMinPk)
	}
}

func benchmarkBlstAggregateVerify(n int, b *testing.B) {
	messages := make([][]byte, n)
	publicKeys := make([]*blstPublicKey, n)
	signatures := make([]*blstSignature, n)

	for i := 0; i < n; i++ {
		message := make([]byte, 32)
		rand.Read(message)
		secretKey := randBLSTSecretKey()
		publicKey := blstPubkey(secretKey)
		signature := new(blstSignature).Sign(secretKey, message, dstMinPk)
		signatures[i] = signature
		publicKeys[i] = publicKey
		messages[i] = message
	}

	aggregatedSignature := new(blst.P2Aggregate)
	aggregatedSignature.Aggregate(signatures, false)

	aggreSig := aggregatedSignature.ToAffine()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		aggreSig.AggregateVerify(false, publicKeys, false, messages, dstMinPk)
	}

}

func benchmarkBlstFastAggregateVerify(n int, b *testing.B) {
	message := []byte("hello foo")
	publicKeys := make([]*blstPublicKey, n)
	signatures := make([]*blstSignature, n)

	for i := 0; i < n; i++ {
		message := make([]byte, 32)
		rand.Read(message)
		secretKey := randBLSTSecretKey()
		publicKey := blstPubkey(secretKey)
		signature := new(blstSignature).Sign(secretKey, message, dstMinPk)
		signatures[i] = signature
		publicKeys[i] = publicKey
	}

	aggregatedSignature := new(blst.P2Aggregate)
	aggregatedSignature.Aggregate(signatures, false)

	aggreSig := aggregatedSignature.ToAffine()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		aggreSig.FastAggregateVerify(false, publicKeys, message, dstMinPk)
	}

}

func BenchmarkBlstAggregateVerify10(b *testing.B) {
	benchmarkBlstAggregateVerify(10, b)
}

func BenchmarkBlstAggregateVerify100(b *testing.B) {
	benchmarkBlstAggregateVerify(100, b)
}

func BenchmarkBlstAggregateVerify1000(b *testing.B) {
	benchmarkBlstAggregateVerify(1000, b)
}

func BenchmarkBlstFastAggregateVerify10(b *testing.B) {
	benchmarkBlstFastAggregateVerify(10, b)
}

func BenchmarkBlstFastAggregateVerify100(b *testing.B) {
	benchmarkBlstFastAggregateVerify(100, b)
}

func BenchmarkBlstFastAggregateVerify1000(b *testing.B) {
	benchmarkBlstFastAggregateVerify(1000, b)
}
