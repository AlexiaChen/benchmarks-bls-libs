package main

import (
	"crypto/rand"
	"runtime"
	"testing"

	herumi "github.com/herumi/bls-eth-go-binary/bls"
	blst "github.com/supranational/blst/bindings/go"
)

// Min-PK
type blstPublicKey = blst.P1Affine
type blstSignature = blst.P2Affine
type blstAggregateSignature = blst.P2Aggregate
type blstAggregatePublicKey = blst.P1Aggregate
type blstSecretKey = blst.SecretKey

type herumiPublicKey = herumi.PublicKey
type _herumiPublickeyAggregator = herumi.PublicKey
type herumiSecretKey = herumi.SecretKey
type herumiSignature = herumi.Sign

var dstMinPk = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

func init() {
	// Use all cores when testing and benchmarking
	blst.SetMaxProcs(runtime.GOMAXPROCS(0))

	if err := herumi.Init(herumi.BLS12_381); err != nil {
		panic(err)
	}
	if err := herumi.SetETHmode(herumi.EthModeDraft07); err != nil {
		panic(err)
	}
	herumi.VerifyPublicKeyOrder(true)
	herumi.VerifySignatureOrder(true)
}

func randBLSTSecretKey() *blstSecretKey {
	var t [32]byte
	_, _ = rand.Read(t[:])
	secretKey := blst.KeyGen(t[:])
	return secretKey
}

func randHerumiSecretKey() *herumiSecretKey {
	secretKey := new(herumi.SecretKey)
	secretKey.SetByCSPRNG()
	return secretKey
}

func blstPubkey(sk *blstSecretKey) *blstPublicKey {
	return new(blstPublicKey).From(sk)
}

func herumiPubkey(sk *herumiSecretKey) *herumiPublicKey {
	return sk.GetPublicKey()
}

func BenchmarkBlstSign(b *testing.B) {
	sk := randBLSTSecretKey()
	msg := []byte("hello foo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		new(blstSignature).Sign(sk, msg, dstMinPk)
	}
}

func BenchmarkHerumiSign(b *testing.B) {
	sk := randHerumiSecretKey()
	msg := []byte("hello foo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sk.Sign(string(msg))
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

func BenchmarkHerumiVerify(b *testing.B) {
	sk := randHerumiSecretKey()
	pk := herumiPubkey(sk)
	msg := []byte("hello foo")
	sig := sk.Sign(string(msg))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sig.Verify(
			pk,
			string(msg),
		)
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

func benchmarkHerumiAggregateVerify(n int, b *testing.B) {
	messages := make([][]byte, n)
	publicKeys := make([]herumiPublicKey, n)
	signatures := make([]herumiSignature, n)

	for i := 0; i < n; i++ {
		message := make([]byte, 32)
		rand.Read(message)
		secretKey := randHerumiSecretKey()
		publicKey := herumiPubkey(secretKey)
		signature := secretKey.Sign(string(message))
		signatures[i] = *signature
		publicKeys[i] = *publicKey
		messages[i] = message
	}

	aggregatedSignature := new(herumiSignature)
	aggregatedSignature.Aggregate(signatures)

	_messages := []byte{}
	size := len(publicKeys)
	for i := 0; i < size; i++ {
		_messages = append(_messages, messages[i][:]...)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		aggregatedSignature.AggregateVerify(publicKeys, _messages)
	}
}

func benchmarkBlstFastAggregateVerify(n int, b *testing.B) {
	message := []byte("hello foo")
	publicKeys := make([]*blstPublicKey, n)
	signatures := make([]*blstSignature, n)

	for i := 0; i < n; i++ {
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

func benchmarkHerumiFastAggregateVerify(n int, b *testing.B) {
	message := []byte("hello foo")
	publicKeys := make([]herumiPublicKey, n)
	signatures := make([]herumiSignature, n)

	for i := 0; i < n; i++ {
		secretKey := randHerumiSecretKey()
		publicKey := herumiPubkey(secretKey)
		signature := secretKey.Sign(string(message))
		signatures[i] = *signature
		publicKeys[i] = *publicKey

	}

	aggregatedSignature := new(herumiSignature)
	aggregatedSignature.Aggregate(signatures)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		aggregatedSignature.FastAggregateVerify(publicKeys, message)
	}
}

func BenchmarkBlstAggregateVerify10(b *testing.B) {
	benchmarkBlstAggregateVerify(10, b)
}

func BenchmarkHerumiAggregateVerify10(b *testing.B) {
	benchmarkHerumiAggregateVerify(10, b)
}

func BenchmarkBlstAggregateVerify100(b *testing.B) {
	benchmarkBlstAggregateVerify(100, b)
}

func BenchmarkHerumiAggregateVerify100(b *testing.B) {
	benchmarkHerumiAggregateVerify(100, b)
}

func BenchmarkBlstAggregateVerify1000(b *testing.B) {
	benchmarkBlstAggregateVerify(1000, b)
}

func BenchmarkHerumiAggregateVerify1000(b *testing.B) {
	benchmarkHerumiAggregateVerify(1000, b)
}

//////////   Fast ////////////////////////

func BenchmarkBlstFastAggregateVerify10(b *testing.B) {
	benchmarkBlstFastAggregateVerify(10, b)
}

func BenchmarkHerumiFastAggregateVerify10(b *testing.B) {
	benchmarkHerumiFastAggregateVerify(10, b)
}

func BenchmarkBlstFastAggregateVerify100(b *testing.B) {
	benchmarkBlstFastAggregateVerify(100, b)
}

func BenchmarkHerumiFastAggregateVerify100(b *testing.B) {
	benchmarkHerumiFastAggregateVerify(100, b)
}

func BenchmarkBlstFastAggregateVerify1000(b *testing.B) {
	benchmarkBlstFastAggregateVerify(1000, b)
}

func BenchmarkHerumiFastAggregateVerify1000(b *testing.B) {
	benchmarkHerumiFastAggregateVerify(1000, b)
}
