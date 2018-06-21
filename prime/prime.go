package prime

func PrintPrimeUntil(n int) (primes []int) {
	in := make(chan int)
	wait := make(chan struct{})
	go filterPrime(in, wait, &primes)
	for i := 2; i <= n; i++ {
		in <- i
	}
	close(in)
	<-wait
	return primes
}

func filterPrime(in chan int, wait chan struct{}, primes *[]int) {
	prime, ok := <-in
	if !ok {
		close(wait)
		return
	}
	//fmt.Printf("%d ", prime)
	*primes = append(*primes, prime)
	out := make(chan int)
	go filterPrime(out, wait, primes)
	for num := range in {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)
}

func PrintPrimeUntil2(n int) (primes []int) {
	for i := 2; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				goto next
			}
		}
		//fmt.Printf("%d ", i)
		primes = append(primes, i)
	next:
	}
	return primes
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(n int, ch chan<- int) {
	for i := 2; i <= n; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
	close(ch)
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i, ok := <-in // Receive value from 'in'.
		if !ok {
			close(out)
			break
		}
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func PrintPrimeUntil3(n int) (primes []int) {
	ch := make(chan int) // Create a new channel.
	go generate(n, ch)   // Launch Generate goroutine.

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		//fmt.Printf("%d ", prime)
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}
