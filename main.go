package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Ready to run tasks!")

	tasksToRun := os.Args[1:]
	fmt.Println("Tasks to run:", tasksToRun)

	b, err := ioutil.ReadFile("gtask.toml")
	if err != nil {
		log.Fatal("Can't open `gtask.toml` => ", err)
	}
	gtaskToml := string(b)

	fmt.Println("================================================================================")
	fmt.Println("gtask.toml contents")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf(gtaskToml)
	fmt.Println("================================================================================")

	// task := "my_task"
	// arg0 := "python"
	// arg1 := "my_task.py"
	//
	// fmt.Println("================================================================================")
	// fmt.Printf("Running task %q\n", task)
	// fmt.Println("--------------------------------------------------------------------------------")
	//
	// cmd := exec.Command(arg0, arg1)
	// stdout, err := cmd.Output()
	// if err != nil {
	//     fmt.Println(err.Error())
	//     return
	// }
	// fmt.Print(string(stdout))
	// fmt.Println("================================================================================")
}
