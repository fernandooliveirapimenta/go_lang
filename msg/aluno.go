package msg

import "fmt"

type AlunoF struct {
	Nome string
	Idade int8
}

func (a AlunoF) Estudar(){
	fmt.Printf("Aluno %s estudando \n", a.Nome)
}
