package main

import "fmt"
import (
    "os"
    "io"
    "log"
    "regexp"
    "strings"
    "io/ioutil"
)

// go run tutorial.go argumento_1 argumento_2 argumento_3

func main() {
    ejemplo_print()
    ejemplo_for()
    ejemplo_array()
    ejemplo_logs()
    ejemplo_files()
    argumentos(os.Args)
    fmt.Println(func_arguments([]float64{1.123, 4.333, 27.0/5}))
    concatenar()
    regex_tests()

}

func ejemplo_print(){
    fmt.Printf("\t\t\tPRINT Y LOG\n\n")
    fmt.Printf("hello, world\n")
    log.Println("Hello World\n")
}

func ejemplo_for(){
    fmt.Printf("\t\t\tFORs\n\n")
    for x:=0; x<3; x++{
        fmt.Printf("hello, %d\t",x)
    }
    fmt.Printf("\n")
    // en GO SOLO hay ciclo for
    x := 0
    for x<3{
        x++
        fmt.Printf("hello, %d\n",x)
    }

    a := [5]float64{1.5, 2.0, 3.5, 4.5, 4.155}
    sum := float64(0)
    for index := range a {
        sum = sum + a[index]
        fmt.Printf("Suma: %0.2f\n", sum) //con dos decimales
    }
    daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

    for index, value := range daysOfWeek {
        fmt.Printf("Day %d of week = %s\n", index, value)
    }

}

func ejemplo_array(){
    fmt.Printf("\t\t\tARRAYS\n\n")
    // crea array de longitud 5 inicializado con 0s
    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    // crea array de longitud 5 inicializado con numeros 1-5
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // crea matriz de 2x3
    var twoD [2][3]int
    fmt.Println("len twoD:", len(twoD))
    fmt.Println("len twoD[0]:", len(twoD[0]))
    for i := 0; i < len(twoD); i++ {
        for j := 0; j < len(twoD[0]); j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("Matriz: ", twoD)

    // crea array de strings
    names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larry Page"}
    fmt.Println("Array de strings: ", names)
    // crea array de floats
    fls := [4]float64{3.5, 7.2, 4.8, 9.5}
    fmt.Println("Array de floats: ", fls)

    c := [3][4]float64{
        {0,1, 3.0},
        {4.5, -3, 7.4, 2},
        {6.1234567632, 2, 11},
    }
    fmt.Println(c)
}

func ejemplo_logs(){
    fmt.Printf("\t\t\tLOGS FILE\n\n")
    daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    f, err := os.OpenFile("/tmp/testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()

    mw := io.MultiWriter(os.Stdout, f)
    log.SetOutput(mw)
    log.Println("This is a test log entry:")
    log.Println(daysOfWeek)
}

func ejemplo_files(){
    fmt.Printf("\t\t\tFILES\n\n")
    fileName := "/tmp/testlogfile"
    fmt.Println("READ FILE")
    dat, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    fmt.Print(string(dat))

    fmt.Println("WRITE FILE")
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString("new line added there originally\n"); err != nil {
        panic(err)
    }
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    fmt.Print(string(data))

    fmt.Println("DELETE FILE")
    err = os.Remove(fileName)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println("File Deleted")
}

func argumentos(args []string){
    fmt.Printf("\t\t\tARGUMENTOS DE SCRIPT\n\n")
    // go run tutorial.go argumento_1 argumento_2 argumento_3 123
    argsWithFilePath := args
    argsWithoutPath  := args[1:]
    arg              := args[1]

    fmt.Println(argsWithFilePath)
    fmt.Println(argsWithoutPath)
    fmt.Println(arg)
}

func func_arguments(arrayFloat []float64) float64{
    fmt.Printf("\t\t\tARGUMENTOS DE FUNCION\n\n")
    total := 0.0
    for _, v := range arrayFloat {
        total += v
    }
    promedio := total / float64(len(arrayFloat))
    fmt.Printf("%0.5f\n",promedio)
    return total / float64(len(arrayFloat))
}

func concatenar(){
    fmt.Printf("\t\t\tCONCATENAR ARRAR DE STRINGS (JOIN)\n\n")
    frase := []string{"Una","vaca","vestida","de","uniforme"}
    fmt.Println(strings.Join(frase, " "))
}


func regex_tests(){
    fmt.Printf("\t\t\tREGEX\n\n")
    // ejemplo 1
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)
    // ejemplo 2
    match, _ = regexp.MatchString("p([a-z]+)ch", "NOMATCH")
    fmt.Println(match)

    // ejemplo 3
    html := `<div id="img-preload" style="display: none;">
        <img src="https://i.ytimg.com/vi/nNBaEkjtupA/hqdefault.jpg?sqp=-oaymwEjCNACELwBSFryq4qpAxUIARUAAAAAGAElAADIQj0AgKJDeAE=&amp;rs=AOn4CLB4yn_D3J7M2eeWzw_OMEw96OgAbg">
        <img src="https://i.ytimg.com/vi/c5G9g7qXUns/hqdefault.jpg?sqp=-oaymwEjCNACELwBSFryq4qpAxUIARUAAAAAGAElAADIQj0AgKJDeAE=&amp;rs=AOn4CLC7htsUDo3iFp-j8ONttBAvJZkd9w">
        <img src="https://i.ytimg.com/vi/_7j5Tcbo9Zc/hqdefault.jpg?sqp=-oaymwEjCNACELwBSFryq4qpAxUIARUAAAAAGAElAADIQj0AgKJDeAE=&amp;rs=AOn4CLDYJBicKVQAGn0yabfJWg3G3pRWSw">
        <img src="https://i.ytimg.com/vi/TAAXwrgd1U8/hqdefault.jpg?sqp=-oaymwEjCNACELwBSFryq4qpAxUIARUAAAAAGAElAADIQj0AgKJDeAE=&amp;rs=AOn4CLBMesiqQ1dWWPdy7wyTOcBqog3XBA">
        <img src="https://i.ytimg.com/vi/YfRy9j1E0Qk/hqdefault.jpg?sqp=-oaymwEjCNACELwBSFryq4qpAxUIARUAAAAAGAElAADIQj0AgKJDeAE=&amp;rs=AOn4CLCg0tkKTndW8nmkUB5Xg-u_S2845w">
    </div>`
    getSrc := "(https.+jpg)"
    r, _ := regexp.Compile(getSrc)
    fmt.Println(r.FindAllString(html, 3)) // segundo argumendo es el numero de matchs a retornar

    // ejemplo 4
    r1, _ := regexp.Compile("NOMATCH")
    fmt.Println(r1.FindAllString(html, -1)) // -1 regresa todos los matches

}
