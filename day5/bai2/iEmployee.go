package main

type iEmployee interface {
	Clone() iEmployee
	GetSalary() int
	Print()
	SetEmpId(empId int)
	SetSalary(basicSalary int)
	SetAllowance(allowance int)
}
