package sieve

import (
	"fmt"
	"math"
)

type Sieve2357 struct {
	Sieve
}

var d_pattern357 = [48]int64{
	2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2, 6, 4, 2, 6, 4, 6, 8, 4, 2, 4, 2, 4, 8, 6, 4, 6, 2, 4, 6, 2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10, 2, 10,
}

var m_pattern357 = [210]int64{
	-1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, -1, 1, -1, -1, -1, 1, -1, 2, -1, -1, -1, 2, -1, -1, -1, -1, -1, 1, -1, 2, -1, -1, -1, -1, -1, 2, -1, -1, -1, 2, -1, 2, -1, -1, -1, 2, -1, -1, -1, -1, -1, 2, -1, -1, -1, -1, -1, 2, -1, 2, -1, -1, -1, -1, -1, 2, -1, -1, -1, 2, -1, 2, -1, -1, -1, -1, -1, 2, -1, -1, -1, 2, -1, -1, -1, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, 1, -1, -1, -1, 2, -1, 2, -1, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, -1, 2, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, 2, -1, -1, -1, 2, -1, 3, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1,
}

func init() {
}

func (sv *Sieve2357) Init(N int64) {
	sv.N = N
	required_size := N/140 + 1 // 48 / 2*3*5*7*32 = 24 / 3 * 5 * 7 * 32 = 3 / 3 * 5 * 7 * 4 = 1 / 5 * 7 * 4 = 140
	sv.bits = make([]uint32, required_size)
	len_bits := len(sv.bits)

	fmt.Println("Sieve size (bytes): ", len_bits*4)
	for i := 0; i < len_bits; i++ {
		sv.bits[i] = 0b11111111111111111111111111111111
	}
}

func (sv *Sieve2357) GetNextJump(idx int) int64 {
	return d_pattern357[idx]
}

func (sv *Sieve2357) Mark(n int64) {

	var t int64 = (n - 11) * 8
	var ii int64 = t / 35
	var m byte = byte(n % 210)

	// magic seq of m :  0, 1, 9, 10, 3, 4, 13, 6
	var j int = int(m_pattern357[m])

	if j == -1 {
		return
	}

	ii += int64(j)

	b := int(ii / 32)
	bi := int(ii % 32)

	sv.bits[b] = sv.bits[b] & mask_mark[bi]
}

// func (sv *Sieve2357) IsPrime(n int64) bool {

// 	var t int64 = (n - 11) * 24
// 	var ii int64 = t / 105
// 	var m byte = byte(t % 105)

// 	var j int = int(m_pattern357[m])
// 	if j == -1 {
// 		return false
// 	}

// 	ii += int64(j)

// 	b := int(ii / 32)
// 	bi := int(ii % 32)

// 	return sv.bits[b]&mask_get[bi] != 0
// }

func IsPrime(n int64) bool {
	for i := int64(2); i < int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (sv *Sieve2357) Count() int64 {
	var prime int64 = 0
	var count int64 = 4

	d_idx := -1

	sv.Begin()
	for prime = 11; prime <= sv.N; prime += sv.GetNextJump(d_idx) {
		//var x string = "x"
		if sv.Get() != 0 {
			count++
			// if !IsPrime(prime) {
			// 	fmt.Printf("%3d \n", prime)
			// }
			//x = " "
		}

		//fmt.Printf("%3d %3d %3s \n", count, prime, x)

		sv.Next()

		d_idx++
		if d_idx == 48 {
			d_idx = 0
		}
	}
	return count
}
