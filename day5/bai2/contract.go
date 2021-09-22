package main

import "fmt"

type Contract struct {
	empId       int
	basicSalary int
	allowance   int
}

func (p *Contract) Clone() iEmployee {
	return &Contract{
		empId:       p.empId,
		basicSalary: p.basicSalary,
		allowance:   0,
	}
}

func (p *Contract) GetSalary() int {
	return p.basicSalary
}

func (p *Contract) Print() {
	fmt.Printf("empID: %d, basicSalary: %d", p.empId, p.basicSalary)
}

func (p *Contract) SetEmpId(empId int) {
	p.empId = empId
}

func (p *Contract) SetSalary(basicSalary int) {
	p.basicSalary = basicSalary
}

func (p *Contract) SetAllowance(allowance int) {
	p.allowance = allowance
}
