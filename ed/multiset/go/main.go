package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MultiSet[T comparable] struct {
	data     map[T]int
	size     int
	capacity int
}

func NewMultiSet[T comparable](cap int) *MultiSet[T] {
	return &MultiSet[T]{}
}

func Contais(m MultiSet[]) {

}

//+ search(value: int): (bool, int)         ' Se o elemento value estiver presente, retorna true e o índice da última ocorrência
//--                                        ' Se não estiver, retorna false e o índice onde ele deve ser inserido
//--
//+ Insert(value: int): void                ' Insere o valor na posição correta, mantendo a ordem e permitindo repetições
//- insert(value: int, index: int): error   ' Insere value no índice indicado, deslocando os elementos à direita
//--
//+ Erase(value: int): error                ' Remove uma ocorrência do valor, se existir; retorna erro caso não exista
//- erase(index: int): error                ' Remove o elemento na posição index, deslocando os demais
//--
//+ Contains(value: int): bool              ' Retorna true se o valor estiver presente no multiconjunto
//+ Count(value: int): int                  ' Retorna o número de ocorrências do valor no multiconjunto
//+ Unique(): int                           ' Retorna o número de valores distintos no multiconjunto
//+ Clear(): void                           ' Remove todos os elementos do multiconjunto
//+ String(): string                        ' Retorna uma representação textual dos elementos do multiconjunto
//}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	// ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			// value, _ := strconv.Atoi(args[1])
			// ms = NewMultiSet(value)
		case "insert":
			// for _, part := range args[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
