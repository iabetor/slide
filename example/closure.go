package main

// START OMIT
func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	c := f(0)
	println(c())
	println(c())
	println(c())
	d := f(0)
	println(d())
	println(d())
}

// END OMIT
