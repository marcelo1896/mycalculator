package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	valores := strings.Split(entrada, operador)
	operador1 := parsear(valores[0])
	operador2 := parsear(valores[1])
	switch operador {
	case "+":
		fmt.Println(operador1, " + ", operador2, " = ", operador1+operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1, " - ", operador2, " = ", operador1-operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1, " * ", operador2, " = ", operador1*operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1, " / ", operador2, " = ", operador1/operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, " no esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerpantalla() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
