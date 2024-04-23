package main

import (
    "fmt"
    "os"
    "strconv"
)

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func divide(a, b int) float64 {
    return float64(a) / float64(b)
}

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("Введите два числа для выполнения операции")
        return
    }

    num1, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    num2, err := strconv.Atoi(args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    switch args[0] {
    case "+":
        result := add(num1, num2)
        fmt.Printf("%d + %d = %d\n", num1, num2, result)
    case "-":
        result := subtract(num1, num2)
        fmt.Printf("%d - %d = %d\n", num1, num2, result)
    case "*":
        result := multiply(num1, num2)
        fmt.Printf("%d * %d = %d\n", num1, num2, result)
    case "/":
        result := divide(num1, num2)
        fmt.Printf("%f / %d = %f\n", float64(num1), num2, result)
    default:
        fmt.Println("Неизвестная операция")
    }
}