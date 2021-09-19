package post

import "fmt"

type iPost interface {
	AddView() int
	CompareView()
	GetView() int
}

type post struct {
	view int
}

var instancePost *post

func GetInstancePost() iPost {
	if instancePost == nil {
		return &post{view: 1}
	}
	return instancePost
}

func (p *post) AddView() int {
	p.view++
	return p.view
}

func (p *post) CompareView() {
	if p.view == 10 {
		fmt.Println("Good")
	} else {
		fmt.Println("not Good")
	}
}

func (p *post) GetView() int {
	return p.view
}

func DemoPost() {
	post := GetInstancePost()
	fmt.Println(post)

	view := post.AddView()
	fmt.Println(view)

	view1 := post.AddView()
	fmt.Println(view1)

	post.CompareView()
}
