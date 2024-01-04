package sieve

import "fmt"

type Sieve235 struct {
	Sieve
}

var d_pattern [8]int64
var m_pattern [15]int64

func init() {
	d_pattern = [8]int64{4, 2, 4, 2, 4, 6, 2, 6}
	m_pattern = [15]int64{0, 0, -1, 1, 1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1}

	// m "magic" pattern: 0, 1,  _, 3, 4,  _, 6,  _,  _, 9, 10, _, _, 13, _
}

func (sv *Sieve235) Init(N int64) {
	sv.N = N
	required_size := N*4/(15*32) + 1
	sv.bits = make([]uint32, required_size)
	len_bits := len(sv.bits)

	fmt.Println("Sieve size (bytes): ", len_bits*4)
	for i := 0; i < len_bits; i++ {
		sv.bits[i] = 0b11111111111111111111111111111111
	}
}

func (sv *Sieve235) GetNextJump(idx int) int64 {
	return d_pattern[idx]
}

func (sv *Sieve235) Mark(n int64) {

	var t int64 = (n - 7) * 4
	var ii int64 = t / 15
	var m byte = byte(t % 15)

	// magic seq of m :  0, 1, 9, 10, 3, 4, 13, 6
	var j int = int(m_pattern[m])
	if j == -1 {
		return
	}

	ii += int64(j)

	b := int(ii / 32)
	bi := int(ii % 32)

	sv.bits[b] = sv.bits[b] & mask_mark[bi]
}

func (sv *Sieve235) IsPrime(n int64) bool {

	var t int64 = (n - 7) * 4
	var ii int64 = t / 15
	var m byte = byte(t % 15)

	var j int = int(m_pattern[m])
	if j == -1 {
		return false
	}

	ii += int64(j)

	b := int(ii / 32)
	bi := int(ii % 32)

	return sv.bits[b]&mask_get[bi] != 0
}

func (sv *Sieve235) Count() int64 {
	var prime int64 = 0
	var count int64 = 3

	d_idx := -1

	sv.Begin()
	for prime = 7; prime <= sv.N; prime += sv.GetNextJump(d_idx) {
		if sv.Get() != 0 {
			count++
		}
		sv.Next()

		d_idx++
		if d_idx == 8 {
			d_idx = 0
		}
	}
	return count
}
