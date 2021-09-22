package main

import "fmt"

type Permanent struct {
	empId       int
	basicSalary int
	allowance   int
}

func (p *Permanent) Clone() iEmployee {
	return &Permanent{
		empId:       p.empId,
		basicSalary: p.basicSalary,
		allowance:   p.allowance,
	}
}

func (p *Permanent) GetSalary() int {
	return p.basicSalary + p.allowance
}

func (p *Permanent) Print() {
	fmt.Printf("empID: %d, basicSalary: %d, allowance: %d", p.empId, p.basicSalary, p.allowance)
}

func (p *Permanent) SetEmpId(empId int) {
	p.empId = empId
}

func (p *Permanent) SetSalary(basicSalary int) {
	p.basicSalary = basicSalary
}

func (p *Permanent) SetAllowance(allowance int) {
	p.allowance = allowance
}
