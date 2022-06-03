//in questo esempio possiamo eseguire qualsiasi comando passato a go run es2.go run <comando>
//peró abbiamo implementato un livello di sicurezza rispetto ad es1 perché ora siamo in grado di preteggere l'hostname
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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

	//aggiungiamo dei namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, //come abbiamo detto prima CLONE_NEWUTS é il namespace che gestisce gli hostname e domini
	}

	//in questo caso quando eseguiamo la shell vediamo che, se cambiamo l'hostname effettivamente cambia
	//ma se usciamo dalla shell l'hostname torna quello originale quindi: abbiamo protetto l'hostname!!

	//comqune i pid della macchina host e il filesystem sono comqune accessibili al nostr container quindi
	//la sicurezza non é ancora cosí impeccabile...

	//esegui il comando
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
