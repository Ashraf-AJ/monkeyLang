package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Set(k string, v Object) Object {
	e.store[k] = v
	return v
}

func (e *Environment) Get(k string) (Object, bool) {
	obj, ok := e.store[k]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(k)
	}
	return obj, ok
}
