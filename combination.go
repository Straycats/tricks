package tricks

import "context"

func snoob(x uint64) uint64 {
	var smallest, ripple, ones uint64
	// x = xxx0 1111 0000
	smallest = x & -x             //     0000 0001 0000
	ripple = x + smallest         //     xxx1 0000 0000
	ones = x ^ ripple             //     0001 1111 0000
	ones = (ones >> 2) / smallest //     0000 0000 0111
	return ripple | ones          //     xxx1 0000 0111
}

func trans(n uint64, maxbits uint64) []uint64 {
	r := make([]uint64, Popcnt(n))

	index := 0
	for i := uint64(0); i < maxbits; i++ {
		if (uint64(1)<<i)&n != 0 {
			r[index] = i
			index++
		}
	}

	return r
}

func Combination(n uint64, r uint64) (<-chan []uint64, context.CancelFunc) {
	c := make(chan []uint64)

	cancelCtx, cancel := context.WithCancel(context.Background())

	go func() {
		defer close(c)

		if n > 64 || r > 64 || r == 0 {
			return
		}

		max := uint64(1) << n

		iter := uint64(0)
		for i := uint64(0); i < r; i++ {
			iter |= uint64(1) << i
		}

		for iter&max == 0 {
			select {
			case c <- trans(iter, n):
			case <-cancelCtx.Done():
				return
			}

			prev := iter
			iter = snoob(iter)
			if iter < prev {
				return
			}
		}
	}()

	return c, cancel
}
