package interfaces

type Animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	EstaVivo() bool
}
