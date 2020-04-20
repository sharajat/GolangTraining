package main

import (
	"testing"
)

func TestContactDetails(t *testing.T)  {
	d := newdetails()
	if len(d.ContactDetails.mobile) != 10{
		t.Errorf("Incorrect Number, Number is of 10 digits but got %d digits.",len(d.ContactDetails.mobile))
	}
	if d.ContactDetails.address != "Pune"{
		t.Errorf("Student doest not belong from pune. He/she belongs from %s .",d.ContactDetails.address)
	}

}

func Test_age(t *testing.T){
	d := newdetails()
	if d.class == "Engineering"{
		if d.age < 17{
			t.Errorf("Student is small to be an Engineer. Required age of Engineer must be greater than 17 but %s's age is %d.",d.firstName,d.age)
		}
	}
}
