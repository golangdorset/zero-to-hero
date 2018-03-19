package main

import "fmt"

func simpleTypes() {
	// Here we have an integer variable. What do you think the zero value
	// of an integer is going to be? What would happen if we used a different
	// type of integer?
	var i int
	fmt.Printf("integer: %v\n", i)

	// We can also define an integer with variable like so:
	j := 0
	fmt.Printf("integer: %v\n", j)

	// Here we have a 64 bit float variable. What do you think the zero value
	// of a float is going to be? What would happen if we used a different
	// type of float?
	var f float64
	fmt.Printf("float: %v\n", f)

	// We can also define an float64 with variable like so:
	// Note that the syntax is the same as the integer defined on line 13. The
	// type is inferred by the initialise the variable with.
	g := 0.0
	fmt.Printf("float: %v\n", g)

	// Here we have a boolean variable. What do you think the zero value
	// of a boolean is going to be?
	var b bool
	fmt.Printf("boolean: %v\n", b)

	// Here we have a string variable. What do you think the zero value
	// of a string is going to be?
	var s string
	fmt.Printf("string: %q\n", s)
}

func moarSimpleTypes() {
	// Here we have an integer variable. What do you think will be displayed
	// when we use the & syntax?
	var i int
	fmt.Printf("integer: %v\n", &i)

	// Here we have a 64 bit float variable. What do you think will be displayed
	// when we use the & syntax?
	var f float64
	fmt.Printf("float: %v\n", &f)

	// Here we have a boolean variable. What do you think will be displayed
	// when we use the & syntax?
	var b bool
	fmt.Printf("boolean: %v\n", &b)

	// Here we have a string variable. What do you think will be displayed
	// when we use the & syntax?
	var s string
	fmt.Printf("string: %v\n", &s)
}

func pointers() {
	// Here we have pointers to the same types we used above. Can you guess what
	// the zero values will be in this case?
	var i *int
	var f *float64
	var b *bool
	var s *string

	fmt.Printf("integer: %v, float: %v, boolean: %v, string: %v\n", i, f, b, s)
}

func moarPointers() {
	// Here we've created a closure which takes a pointer to an integer and sets
	// the pointed value to the value of a zero valued integer variable.
	reset := func(i *int) {
		var j int
		*i = j
	}

	// We've defined an integer variable with an initial value of 1. We now pass
	// the integer to the zero pointer function.
	i := 1
	reset(&i)

	// What will this output?
	fmt.Printf("integer: %v\n", i)
}

func main() {
	simpleTypes()

	fmt.Println()

	moarSimpleTypes()

	fmt.Println()

	pointers()

	fmt.Println()

	moarPointers()
}
