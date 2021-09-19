package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Print

func demoStringFunc() {
	p("Contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("HasPrfix:", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "st"))

	p("Index: ", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:", s.Repeat("x", 10))
	p("REplace:", s.Replace("test", "est", "o", -1))
	p("Split:", s.Split("Hello World", " "))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
