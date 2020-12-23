package meander

// Facade with Public method which will return the public view of a struct
type Facade interface {
	Public() interface{}
}

// Public return the public view of a struct
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
