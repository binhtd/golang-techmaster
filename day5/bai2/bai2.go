package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var employee []iEmployee
	firstPermanent := Permanent{
		empId:       0,
		basicSalary: 0,
		allowance:   0,
	}

	firstContract := Contract{
		empId:       0,
		basicSalary: 0,
		allowance:   0,
	}

	for i := 1; i <= 10; i++ {
		permanetClone := firstPermanent.Clone()
		permanetClone.SetEmpId(i)
		permanetClone.SetSalary(rand.Intn(100))
		permanetClone.SetAllowance(rand.Intn(100))
		employee = append(employee, permanetClone)
	}

	for i := 11; i <= 20; i++ {
		contractClone := firstContract.Clone()
		contractClone.SetEmpId(i)
		contractClone.SetSalary(rand.Intn(100))
		contractClone.SetAllowance(0)
		employee = append(employee, contractClone)
	}

	var totalSalaryCompanyMustPay int = 0
	for index := range employee {
		totalSalaryCompanyMustPay = totalSalaryCompanyMustPay + employee[index].GetSalary()
	}

	fmt.Printf("Total salary company must pay for all employee %d", totalSalaryCompanyMustPay)
	fmt.Println()
}
