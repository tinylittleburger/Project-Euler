package utilities

type Sieve struct {
	primes []bool
}

func NewSieve() Sieve {
	return Sieve{}
}

func (s *Sieve) IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if len(s.primes) <= n {
		s.primes = sieve(int64(n))
	}

	return !s.primes[n]
}

func (s *Sieve) IsPrime64(n int64) bool {
	if n <= 1 {
		return false
	}

	if int64(len(s.primes)) <= n {
		s.primes = sieve(n)
	}

	return !s.primes[n]
}

func sieve(l int64) []bool {
	a := make([]bool, l+1)

	for i := 2; i < len(a); i++ {
		if !a[i] {
			for j := 2; i*j < len(a); j++ {
				a[i*j] = true
			}
		}
	}

	return a
}
