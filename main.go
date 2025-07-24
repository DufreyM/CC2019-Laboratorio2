package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var precedence = map[rune]int{
	'*': 3,
	'.': 2,
	'|': 1,
}

func isOperator(c rune) bool {
	return strings.ContainsRune("*.|", c)
}

func isLeftAssoc(op rune) bool {
	return op != '*'
}

func expandEscapes(input string) string {
	return strings.NewReplacer(
		`\n`, "\n",
		`\t`, "\t",
		`\{`, "{",
		`\}`, "}",
		`\\`, `\`,
	).Replace(input)
}

func insertConcatOperators(regex string) string {
	var result []rune
	runes := []rune(regex)
	escaped := false

	isLiteral := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r) || r == 'ε'
	}

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if escaped {
			result = append(result, '\\', c)
			escaped = false
			continue
		}

		if c == '\\' {
			escaped = true
			continue
		}

		result = append(result, c)

		if i+1 < len(runes) {
			next := runes[i+1]
			if (isLiteral(c) || c == '*' || c == ')' || c == '+' || c == '?') &&
				(isLiteral(next) || next == '(' || next == '\\') {
				result = append(result, '.')
			}
		}
	}
	return string(result)
}

func handleExtensions(expr string) string {
	var output []rune
	runes := []rune(expr)
	escaped := false

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if escaped {
			output = append(output, '\\', c)
			escaped = false
			continue
		}

		if c == '\\' {
			escaped = true
			continue
		}

		if (c == '+' || c == '?') && len(output) > 0 {
			var group []rune
			if output[len(output)-1] == ')' {
				count := 0
				for j := len(output) - 1; j >= 0; j-- {
					if output[j] == ')' {
						count++
					} else if output[j] == '(' {
						count--
					}
					group = append([]rune{output[j]}, group...)
					if count == 0 {
						output = output[:j]
						break
					}
				}
			} else {
				group = []rune{output[len(output)-1]}
				output = output[:len(output)-1]
			}

			if c == '+' {
				output = append(output, '(')
				output = append(output, group...)
				output = append(output, ')', '.')
				output = append(output, '(')
				output = append(output, group...)
				output = append(output, ')', '*', '.')
			} else if c == '?' {
				output = append(output, '(')
				output = append(output, group...)
				output = append(output, '|', 'ε', ')')
			}
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}

func shuntingYard(expr string) (string, []string) {
	var output []rune
	var stack []rune
	var steps []string

	runes := []rune(expr)
	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if c == '\\' && i+1 < len(runes) {
			output = append(output, '\\', runes[i+1])
			steps = append(steps, fmt.Sprintf("Escaped char: \\%c -> Output: %s", runes[i+1], string(output)))
			i++
			continue
		}

		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == 'ε' {
			output = append(output, c)
			steps = append(steps, fmt.Sprintf("Token %c -> Output: %s", c, string(output)))
		} else if c == '(' {
			stack = append(stack, c)
			steps = append(steps, fmt.Sprintf("Push '(' -> Stack: %v", string(stack)))
		} else if c == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				output = append(output, op)
				steps = append(steps, fmt.Sprintf("Pop %c -> Output: %s", op, string(output)))
			}
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			}
		} else if isOperator(c) {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if isOperator(top) &&
					((isLeftAssoc(c) && precedence[c] <= precedence[top]) || (!isLeftAssoc(c) && precedence[c] < precedence[top])) {
					stack = stack[:len(stack)-1]
					output = append(output, top)
					steps = append(steps, fmt.Sprintf("Pop %c -> Output: %s", top, string(output)))
				} else {
					break
				}
			}
			stack = append(stack, c)
			steps = append(steps, fmt.Sprintf("Push %c -> Stack: %v", c, string(stack)))
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		steps = append(steps, fmt.Sprintf("Flush Stack -> Output: %s", string(output)))
	}

	return string(output), steps
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error abriendo archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		raw := scanner.Text()
		fmt.Printf("Expresión original (%d): %q\n", lineNumber, raw)

		expanded := expandEscapes(raw)
		preprocessed := insertConcatOperators(handleExtensions(expanded))

		fmt.Printf("Preprocesada: %s\n", preprocessed)

		postfix, steps := shuntingYard(preprocessed)
		fmt.Println("Resultado postfix:", postfix)
		fmt.Println("Pasos:")
		for _, s := range steps {
			fmt.Println(" -", s)
		}
		fmt.Println(strings.Repeat("-", 40))
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error leyendo archivo: %v", err)
	}
}
