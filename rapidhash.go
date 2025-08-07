package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

// Default secret parameters.
var rapidSecret [8]uint64 = [8]uint64{
	0x2d358dccaa6c78a5,
	0x8bb84b93962eacc9,
	0x4b33a62ed433d4a3,
	0x4d5a2da51de1aa47,
	0xa0761d6478bd642f,
	0xe7037ed1a0b428db,
	0x90ed1765281c388c,
	0xaaaaaaaaaaaaaaaa,
}

// 64*64 -> 128bit multiply function.
func mum(a, b uint64) (uint64, uint64) {
	hi, lo := bits.Mul64(a, b)
	return lo, hi
}

// Multiply and xor mix function.
func mix(a, b uint64) uint64 {
	lo, hi := mum(a, b)
	return lo ^ hi
}

func read64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)

}

func read32(data []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(data))
}

func RapidhashWithSeed(data []byte, seed uint64) uint64 {
	return rapidhashInternal(data, seed, rapidSecret)
}

func Rapidhash(data []byte) uint64 {
	return rapidhashInternal(data, 0, rapidSecret)
}

func rapidhashInternal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	seed ^= mix(seed^secret[2], secret[1])
	var a, b uint64
	p := data
	i := bufferlen

	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
			}
		} else if bufferlen > 0 {
			a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
			b = (uint64)(data[bufferlen>>1])
		}

	} else {
		var (
			see1, see2 = seed, seed
			see3, see4 = seed, seed
			see5, see6 = seed, seed
		)

		if i > 112 {
			for {
				seed = mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = mix(read64(p[16:])^secret[1], read64(data[24:])^see1)
				see2 = mix(read64(p[32:])^secret[2], read64(data[40:])^see2)
				see3 = mix(read64(p[48:])^secret[3], read64(data[56:])^see3)
				see4 = mix(read64(p[64:])^secret[4], read64(data[72:])^see4)
				see5 = mix(read64(p[80:])^secret[5], read64(data[88:])^see5)
				see6 = mix(read64(p[96:])^secret[6], read64(data[104:])^see6)
				p = p[112:]
				i -= 112
				if i <= 112 {
					break
				}
			}

			seed ^= see1
			see2 ^= see3
			see4 ^= see5
			seed ^= see6
			see2 ^= see4
			seed ^= see2
		}

		if i > 16 {
			seed = mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
				if i > 48 {
					seed = mix(read64(p[32:])^secret[1], read64(p[40:])^seed)
					if i > 64 {
						seed = mix(read64(p[48:])^secret[1], read64(p[56:])^seed)
						if i > 80 {
							seed = mix(read64(p[64:])^secret[2], read64(p[72:])^seed)
							if i > 96 {
								seed = mix(read64(p[80:])^secret[1], read64(p[88:])^seed)
							}
						}
					}
				}
			}
		}
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	a, b = mum(a, b)
	return mix(a^secret[7], b^secret[1]^i)
}

func RapidhashMicroWithSeed(data []byte, seed uint64) uint64 {
	return rapidhashMicroInternal(data, seed, rapidSecret)
}

func RapidhashMicro(data []byte) uint64 {
	return rapidhashMicroInternal(data, 0, rapidSecret)
}

func rapidhashMicroInternal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	p := data
	i := bufferlen
	var a, b uint64

	seed ^= mix(seed^secret[2], secret[1])
	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
			}
		} else if bufferlen > 0 {
			a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
			b = (uint64)(data[bufferlen>>1])
		}

	} else {
		if i > 80 {
			var see1, see2, see3, see4 = seed, seed, seed, seed
			for {
				seed = mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
				see2 = mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
				see3 = mix(read64(p[48:])^secret[3], read64(p[56:])^see3)
				see4 = mix(read64(p[64:])^secret[4], read64(p[72:])^see4)
				p = p[80:]
				i -= 80
				if i <= 80 {
					break
				}
			}
			seed ^= see1
			see2 ^= see3
			seed ^= see4
			seed ^= see2
		}
		if i > 16 {
			seed = mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
				if i > 48 {
					seed = mix(read64(p[32:])^secret[1], read64(p[40:])^seed)
					if i > 64 {
						seed = mix(read64(p[48:])^secret[1], read64(p[56:])^seed)
					}
				}
			}
		}
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	a, b = mum(a, b)
	return mix(a^secret[7], b^secret[1]^i)
}

func RapidhashNanoWithSeed(data []byte, seed uint64) uint64 {
	return rapidhashNanoInternal(data, seed, rapidSecret)
}

func RapidhashNano(data []byte) uint64 {
	return rapidhashNanoInternal(data, 0, rapidSecret)
}

func rapidhashNanoInternal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	p := data
	i := bufferlen
	var a, b uint64

	seed ^= mix(seed^secret[2], secret[1])
	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
			}
		} else if bufferlen > 0 {
			a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
			b = (uint64)(data[bufferlen>>1])
		}

	} else {
		if i > 48 {
			var see1, see2 = seed, seed
			for {
				seed = mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
				see2 = mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
				p = p[48:]
				i -= 48
				if i <= 48 {
					break
				}
			}
			seed ^= see1
			seed ^= see2
		}
		if i > 16 {
			seed = mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
			}
		}
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	a, b = mum(a, b)
	return mix(a^secret[7], b^secret[1]^i)
}
