package main

import (
	"bufio"
	"fmt"
	"os"
)

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isOpening(symbol rune) bool {
	return symbol == '(' || symbol == '[' || symbol == '{'
}

func isClosing(symbol rune) bool {
	return symbol == ')' || symbol == ']' || symbol == '}'
}

func checkBalance(expression string) bool {
	var stack []rune

	fmt.Printf("Expresión: %s\n", expression)
	fmt.Println("Pasos de la pila:")

	for _, char := range expression {
		if isOpening(char) {
			stack = append(stack, char)
			fmt.Printf("  Push: %c - pila: %q\n", char, stack)
		} else if isClosing(char) {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				fmt.Printf(" Error: se esperaba %c pero no se encontró\n", pairs[char])
				return false
			}
			fmt.Printf("  Pop: %c . pila antes: %q\n", char, stack)
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		fmt.Println("Se balanceo")
		return true
	} else {
		fmt.Printf("No se balanceo Pila final: %q\n\n", stack)
		return false
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Línea %d:\n", lineNum)
		checkBalance(scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
