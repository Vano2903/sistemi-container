//in questo esempio possiamo eseguire qualsiasi comando passato a go run es1.go run <comando>
//peró non abbiamo alcun tipo di sicurezza/isolamento dalla macchina host, infatti se eseguiamo una shell
//possiamo vedere lo stesso filesystem dell'host, lo stesso hostname e gli stessi processi con gli stessi pid che vede l'host
package main

import (
	"fmt"
	"os"
	"os/exec"
)

//vogliamo ricreare il comando docker run <container> [args]
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("ay yo, excuse me?!?")
	}
}

func run() {
	//stampiamo tutti gli argomenti
	fmt.Printf("running %v\n", os.Args[2:])

	//creiamo una variabile con il comando da eseguire (args[2])
	//e tutti gli argomenti da passare al comando (args[3:]...)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	//settiamo gli std
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	//esegui il comando
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
