package L1_1

//Declaration of structure Human
type Human struct {
	name string
}

//Initialization fields
func (h *Human) NewHuman() {

}

//incapsulating of struture Human in Action struct
type Action struct {
	h Human
}

//Action from parent structure
func (a *Action) NewAction() {
	a.h.NewHuman()
}
