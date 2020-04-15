package main

// All assignments of day1

import "fmt"

func assign1()  {
	// Print hello world
	fmt.Println("Hello World of assignments")
}

func assign2()  {
	// Print the triangle of *
	for i:=1; i<=5; i++ {
		for j:=i;  j>0; j-- {
			fmt.Print("*"," ")
		}
		fmt.Println()
	}
}

func assign3()  {
	// Print the triangle of numbers
	val := 1
	for i:=1; i<=5; i++ {
		for j:=i;  j>0; j-- {
			fmt.Print(val," ")
			val ++
		}
		fmt.Println()
	}

}

func assign4()  {
	// Factorial of number
	num := 4
	fact := 1
	for i:=1; i<=num; i++ {
		fact = fact*i
	}
	fmt.Printf("Factorial of %d = %d",num ,fact)

}

func assign5()  {
	// Is number a Prime number
	num := 2
	if num > 1 {
		isprime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isprime = false
				fmt.Printf("%d is not a prime number", num)
				break
			}
		}
		if isprime == true {
			fmt.Printf("%d is a prime number", num)
		}
	} else {
		fmt.Println("Number must be greater then 1")

	}

}


func main()  {
	assign1()
	fmt.Println("\n")
	assign2()
	fmt.Println("\n")
	assign3()
	fmt.Println("\n")
	assign4()
	fmt.Println("\n")
	assign5()

}

