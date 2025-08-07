package rapidhash

import (
	"hash/maphash"
	"math/rand/v2"
	"testing"
)

func fillWithRandomData(data []byte) {
	source := rand.NewPCG(69, 420)
	r := rand.New(source) //#nosec:G404
	for i := 0; i < len(data); {
		randVal := r.Uint64()
		for j := 0; j < 8 && i < len(data); j++ {
			data[i] = byte(randVal >> (j * 8))
			i++
		}
	}
}

func TestRapidhash(t *testing.T) {
	testCases := []struct {
		key      []byte
		expected uint64
	}{
		{[]byte(""), 0x338dc4be2cecdae},
		{[]byte("a"), 0x599f47df33a2e1eb},
		{[]byte("abc"), 0xcb475beafa9c0da2},
		{[]byte("hello"), 0x2e2d7651b45f7946},
		{[]byte("message digest"), 0x489e17c8eba5e6e7},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), 0x2e1abe6bd50a7a46},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0x7b36ad202067d2e5},
		{[]byte("#This is a sample string for testing purposes. It has been created to contain one hundred and twelve characters."), 0x463de64cbf157f56},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0x001c9413a0cb9203},
	}
	for _, tc := range testCases {
		got := Rapidhash(tc.key)
		if got != tc.expected {
			t.Errorf("Rapidhash(%q) = 0x%x; want 0x%x", tc.key, got, tc.expected)
		}
	}
}

func TestRapidhashMicro(t *testing.T) {
	testCases := []struct {
		key      []byte
		expected uint64
	}{
		{[]byte(""), 0x338dc4be2cecdae},
		{[]byte("a"), 0x599f47df33a2e1eb},
		{[]byte("abc"), 0xcb475beafa9c0da2},
		{[]byte("hello"), 0x2e2d7651b45f7946},
		{[]byte("message digest"), 0x489e17c8eba5e6e7},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), 0x2e1abe6bd50a7a46},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0x7b36ad202067d2e5},
		{[]byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890"), 0x73a21f3bca920f8a},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0x35e1b90fb7514912},
	}
	for _, tc := range testCases {
		got := RapidhashMicro(tc.key)
		if got != tc.expected {
			t.Errorf("RapidhashMicro(%q) = 0x%x; want 0x%x", tc.key, got, tc.expected)
		}
	}
}

func TestRapidhashNano(t *testing.T) {
	testCases := []struct {
		key      []byte
		expected uint64
	}{
		{[]byte(""), 0x338dc4be2cecdae},
		{[]byte("a"), 0x599f47df33a2e1eb},
		{[]byte("abc"), 0xcb475beafa9c0da2},
		{[]byte("hello"), 0x2e2d7651b45f7946},
		{[]byte("message digest"), 0x489e17c8eba5e6e7},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), 0x2e1abe6bd50a7a46},
		{[]byte("Test_string_with_various_chars_and_length_of_48!"), 0x3e9f5539d2db0a},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0xf248fe9d7e75acfe},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), 0x304fd3828c7e2c22},
	}
	for _, tc := range testCases {
		got := RapidhashNano(tc.key)
		if got != tc.expected {
			t.Errorf("RapidhashNano(%q) = 0x%x; want 0x%x", tc.key, got, tc.expected)
		}
	}
}

func TestRapidhashWithSeed(t *testing.T) {
	testCases := []struct {
		key      []byte
		seed     uint64
		expected uint64
	}{
		{[]byte(""), 42069, 0xf7911025ed422ddb},
		{[]byte("hello"), 42069, 0x989aab9c7f9994f},
	}
	for _, tc := range testCases {
		got := RapidhashWithSeed(tc.key, tc.seed)
		if got != tc.expected {
			t.Errorf("RapidhashWithSeed(%q, %d) = 0x%x; want 0x%x", tc.key, tc.seed, got, tc.expected)
		}
	}
}

func TestRapidhashMicroWithSeed(t *testing.T) {
	testCases := []struct {
		key      []byte
		seed     uint64
		expected uint64
	}{
		{[]byte(""), 42069, 0xf7911025ed422ddb},
		{[]byte("hello"), 42069, 0x989aab9c7f9994f},
	}
	for _, tc := range testCases {
		got := RapidhashMicroWithSeed(tc.key, tc.seed)
		if got != tc.expected {
			t.Errorf("RapidhashMicroWithSeed(%q, %d) = 0x%x; want 0x%x", tc.key, tc.seed, got, tc.expected)
		}
	}
}

func TestRapidhashNanoWithSeed(t *testing.T) {
	testCases := []struct {
		key      []byte
		seed     uint64
		expected uint64
	}{
		{[]byte(""), 42069, 0xf7911025ed422ddb},
		{[]byte("hello"), 42069, 0x989aab9c7f9994f},
	}
	for _, tc := range testCases {
		got := RapidhashNanoWithSeed(tc.key, tc.seed)
		if got != tc.expected {
			t.Errorf("RapidhashNanoWithSeed(%q, %d) = 0x%x; want 0x%x", tc.key, tc.seed, got, tc.expected)
		}
	}
}

func FuzzRapidhashNanoMatchesRapidhash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte, seed uint64) {
		hashNano := RapidhashNanoWithSeed(data, seed)
		hashRapid := RapidhashWithSeed(data, seed)
		if len(data) > 48 {
			if hashNano == hashRapid {
				t.Errorf("RapidhashNano and Rapidhash match for input of length %d. Nano: 0x%x, Rapid: 0x%x", len(data), hashNano, hashRapid)
			}
		} else if hashNano != hashRapid {
			t.Errorf("RapidhashNano and Rapidhash mismatch for input of length %d. Nano: 0x%x, Rapid: 0x%x", len(data), hashNano, hashRapid)
		}
	})
}

func FuzzRapidhashMicroMatchesRapidhash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte, seed uint64) {
		hashMicro := RapidhashMicroWithSeed(data, seed)
		hashRapid := RapidhashWithSeed(data, seed)
		if len(data) > 80 {
			if hashMicro == hashRapid {
				t.Errorf("RapidhashMicro and Rapidhash match for input of length %d. Micro: 0x%x, Rapid: 0x%x", len(data), hashMicro, hashRapid)
			}
		} else if hashMicro != hashRapid {
			t.Errorf("RapidhashMicro and Rapidhash mismatch for input of length %d. Micro: 0x%x, Rapid: 0x%x", len(data), hashMicro, hashRapid)
		}
	})
}

func benchmarkRapidhash(b *testing.B, size int) {
	data := make([]byte, size)
	fillWithRandomData(data)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for b.Loop() {
		Rapidhash(data)
	}
}

func BenchmarkRapidhash4B(b *testing.B)   { benchmarkRapidhash(b, 4) }
func BenchmarkRapidhash8B(b *testing.B)   { benchmarkRapidhash(b, 8) }
func BenchmarkRapidhash12B(b *testing.B)  { benchmarkRapidhash(b, 12) }
func BenchmarkRapidhash16B(b *testing.B)  { benchmarkRapidhash(b, 16) }
func BenchmarkRapidhash32B(b *testing.B)  { benchmarkRapidhash(b, 32) }
func BenchmarkRapidhash64B(b *testing.B)  { benchmarkRapidhash(b, 64) }
func BenchmarkRapidhash128B(b *testing.B) { benchmarkRapidhash(b, 128) }
func BenchmarkRapidhash256B(b *testing.B) { benchmarkRapidhash(b, 256) }
func BenchmarkRapidhash1KB(b *testing.B)  { benchmarkRapidhash(b, 1024) }
func BenchmarkRapidhash10KB(b *testing.B) { benchmarkRapidhash(b, 10*1024) }

func benchmarkRapidhashMicro(b *testing.B, size int) {
	data := make([]byte, size)
	fillWithRandomData(data)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for b.Loop() {
		RapidhashMicro(data)
	}
}

func BenchmarkRapidhashMicro4B(b *testing.B)   { benchmarkRapidhashMicro(b, 4) }
func BenchmarkRapidhashMicro8B(b *testing.B)   { benchmarkRapidhashMicro(b, 8) }
func BenchmarkRapidhashMicro12B(b *testing.B)  { benchmarkRapidhashMicro(b, 12) }
func BenchmarkRapidhashMicro16B(b *testing.B)  { benchmarkRapidhashMicro(b, 16) }
func BenchmarkRapidhashMicro32B(b *testing.B)  { benchmarkRapidhashMicro(b, 32) }
func BenchmarkRapidhashMicro64B(b *testing.B)  { benchmarkRapidhashMicro(b, 64) }
func BenchmarkRapidhashMicro128B(b *testing.B) { benchmarkRapidhashMicro(b, 128) }
func BenchmarkRapidhashMicro256B(b *testing.B) { benchmarkRapidhashMicro(b, 256) }
func BenchmarkRapidhashMicro1KB(b *testing.B)  { benchmarkRapidhashMicro(b, 1024) }
func BenchmarkRapidhashMicro10KB(b *testing.B) { benchmarkRapidhashMicro(b, 10*1024) }

func benchmarkRapidhashNano(b *testing.B, size int) {
	data := make([]byte, size)
	fillWithRandomData(data)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for b.Loop() {
		RapidhashNano(data)
	}
}

func BenchmarkRapidhashNano4B(b *testing.B)   { benchmarkRapidhashNano(b, 4) }
func BenchmarkRapidhashNano8B(b *testing.B)   { benchmarkRapidhashNano(b, 8) }
func BenchmarkRapidhashNano12B(b *testing.B)  { benchmarkRapidhashNano(b, 12) }
func BenchmarkRapidhashNano16B(b *testing.B)  { benchmarkRapidhashNano(b, 16) }
func BenchmarkRapidhashNano32B(b *testing.B)  { benchmarkRapidhashNano(b, 32) }
func BenchmarkRapidhashNano64B(b *testing.B)  { benchmarkRapidhashNano(b, 64) }
func BenchmarkRapidhashNano128B(b *testing.B) { benchmarkRapidhashNano(b, 128) }
func BenchmarkRapidhashNano256B(b *testing.B) { benchmarkRapidhashNano(b, 256) }
func BenchmarkRapidhashNano1KB(b *testing.B)  { benchmarkRapidhashNano(b, 1024) }
func BenchmarkRapidhashNano10KB(b *testing.B) { benchmarkRapidhashNano(b, 10*1024) }

func benchmarkMaphash(b *testing.B, size int) {
	seed := maphash.MakeSeed()
	data := make([]byte, size)
	fillWithRandomData(data)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for b.Loop() {
		maphash.Bytes(seed, data)
	}
}

func BenchmarkMaphash4B(b *testing.B)   { benchmarkMaphash(b, 4) }
func BenchmarkMaphash8B(b *testing.B)   { benchmarkMaphash(b, 8) }
func BenchmarkMaphash12B(b *testing.B)  { benchmarkMaphash(b, 12) }
func BenchmarkMaphash16B(b *testing.B)  { benchmarkMaphash(b, 16) }
func BenchmarkMaphash32B(b *testing.B)  { benchmarkMaphash(b, 32) }
func BenchmarkMaphash64B(b *testing.B)  { benchmarkMaphash(b, 64) }
func BenchmarkMaphash128B(b *testing.B) { benchmarkMaphash(b, 128) }
func BenchmarkMaphash256B(b *testing.B) { benchmarkMaphash(b, 256) }
func BenchmarkMaphash1KB(b *testing.B)  { benchmarkMaphash(b, 1024) }
func BenchmarkMaphash10KB(b *testing.B) { benchmarkMaphash(b, 10*1024) }
