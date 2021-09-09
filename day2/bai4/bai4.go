package bai4

import "sort"

/*
Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/

type Employee struct {
	Name        string
	SalaryRatio float64
	Subsidy     float64
}

const baseRatioSalary = 1500000

/*
sort employee by name
*/
func SortEmployeeByName(people []Employee) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
}

/*
sort emloyee by salary
*/
func SortEmployeeBySalary(people []Employee) {
	sort.Slice(people, func(i, j int) bool {
		return getSalary(people[i]) < getSalary(people[j])
	})
}

/*
get list employee has salary max 2nd in company
*/
func ListEmployeeIsMax2nd(people []Employee) []Employee {
	maxSalary := 0.0
	for i := range people {
		if maxSalary < getSalary(people[i]) {
			maxSalary = getSalary(people[i])
		}
	}
	max2ndSalary := 0.0
	for i := range people {
		if max2ndSalary < getSalary(people[i]) && getSalary(people[i]) != maxSalary {
			max2ndSalary = getSalary(people[i])
		}
	}

	max2ndSalaryPeople := []Employee{}
	for i := range people {
		if max2ndSalary == getSalary(people[i]) {
			max2ndSalaryPeople = append(max2ndSalaryPeople, people[i])
		}
	}
	return max2ndSalaryPeople
}

/*
get person salary
*/
func getSalary(person Employee) float64 {
	return person.SalaryRatio*baseRatioSalary + person.Subsidy
}
