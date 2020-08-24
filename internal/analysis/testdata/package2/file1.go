package package2

func f1() {
	var m map[string]string
	m["key"]
}

type foo struct {
	foo      *foo
	mapField map[string]interface{}
}

func f2() {
	m := &foo{}
	m.mapField["key"]
	m.foo.foo.foo.mapField["key"]
}
