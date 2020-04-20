package main

import (
	"fmt"
	"io/ioutil"
)

type Student struct {
	firstName string
	lastName string
	age int
	class string
	ContactDetails
}

type ContactDetails struct {
	mobile string
	email string
	address string
}

func newdetails() Student {
	details := Student{
		firstName:      "Rajat",
		lastName:       "Sharma",
		age:            15,
		class:          "Engineering",
		ContactDetails: ContactDetails{
			mobile:  "8628043496",
			email:   "rajatcu143@gmail.com",
			address: "Pune",
		},
	}
	return details
}

func (s Student) print()  {
	fmt.Println(s)
}

func (s *Student) update(newage int) {
	s.age = newage
}

func (s Student) savetofile(filename string)  {
	_ = ioutil.WriteFile(filename, []byte(fmt.Sprintf("%+v", s)), 0666)
}

func main()  {
	d := newdetails()
	fmt.Println("Details Before update")
	d.print()
	d.update(16)
	fmt.Println("Details After update")
	d.print()
	fmt.Println("Saving Latest Student Details to file")
	d.savetofile("StudentDetails")

}