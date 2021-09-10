package main

type Member struct {
	id      string
	name    string
	email   string
	pass    string
	roles   []string
	age     int
	enabled bool
}

func PassStructAsValue(p Member) {
	p.id = "100"
	p.email = "demo@demo.vn"
	p.pass = "22324d"
	p.name = "data"
	p.roles = []string{"admin", "opeator"}
	p.age = 30
	p.enabled = true
}

func PassStructAsPointer(p *Member) {
	p.id = "100"
	p.email = "data@data.vn"
	p.pass = "3232dds"
	p.name = "vm"
	p.roles = []string{"user", "reader"}
	p.age = 50
	p.enabled = true
}
