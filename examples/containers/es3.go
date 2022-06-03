//in questo esempio possiamo eseguire qualsiasi comando passato a go run es2.go run <comando>
//peró abbiamo implementato un livello di sicurezza rispetto ad es1 perché ora siamo in grado di preteggere l'hostname
//ed isolare i pid dall'host
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
	case "child":
		child()
	default:
		panic("ay yo, excuse me?!?")
	}
}

//per isolare i pid della macchina host rispetto a quelli del container
func run() {
	//il comando run ora non esegue solamente il processo ma richiama /proc/self/exe
	//che é possibile definire come il fare una fork exec del processo che stiamo eseguendo
	
	//passiamo il comando a run, questo fa una fork di se stesso e setta come primo argomento child
	//ed assieme aggiunge tutti gli altri argomenti del comando.
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	//settiamo gli std
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	//fare il fork exec é necessario perché newpid deve essere "instanziato" prima di 
	//eseguire i comandi richiesti.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	//vediamo che tutto funziona ma continuiamo a vedere gli stessi processi della macchina host
	//ma questa volta il problema non sono i namespace ma il modo in cui guardiamo i processi in esecuzione
	//il comando `ps` per trovare i processi guarda nella cartella /proc quindi per far si che tutto funzioni
	//dobbiamo dare una cartella /proc proprietaria (la cartella con tutte le informazioni di tutti i processi) 

	//esegui il comando
	must(cmd.Run())
}

func child() {
	//stampiamo tutti gli argomenti ed il pid del processo
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	//creiamo una variabile con il comando da eseguire (args[2])
	//e tutti gli argomenti da passare al comando (args[3:]...)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	//settiamo gli std
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	//i namespace non servono piú in questo comando perché
	//sono stati instanziati dal comando "run" che ha eseguto "child"

	//esegui il comando
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
