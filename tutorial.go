package main

import "fmt"

// INSTALL https://golang.org/doc/install?download=go1.12.6.linux-amd64.tar.gz

/*
go build tutorial.go
./tutorial

go build -o executableName.exe tutorial.go
./executableName.exe

go run tutorial.go
*/

func main() {
    ejemplo_print()
    ejemplo_for()
    cuentas()
}

func ejemplo_print(){
    fmt.Printf("hello, world\n")
}

func ejemplo_for(){
    for x:=0; x<3; x++{
        fmt.Printf("hello, %d\t",x)
    }
    fmt.Printf("\n")
    // en GO SOLO hay ciclo for
    x := 0
    for x<3{
        x++
        fmt.Printf("hello, %d\t",x)
    }
    fmt.Printf("\n")
}

























func cuentas(){
    salario := 22000
    renta := 3000
    comidas := 6000
    extras := 2000
    apartados := 2000
    fmt.Printf("%d \n",salario - renta - comidas - extras -apartados)
    fmt.Printf("%d \n",salario/2 - renta - 2000 - 2000)
}
