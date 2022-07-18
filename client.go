package main

import (
	"fmt"
	"net/rpc"

	"./args"
)

func client() {
	menu := "1) Agregar calificación\n" +
			"2) Obtener promedio\n" +
			"3) Obtener promedio global\n" +
			"4) Obtener promedio por materia\n" +
			"5) Salir\n"
	var op int64
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for {
		fmt.Print(menu)
		fmt.Scanln(&op)
		
		switch op {
			case 1:
				var name, subject,reply string
				var note float64
				fmt.Print("Nombre: ")
				fmt.Scanln(&name)
				fmt.Print("Materia: ")
				fmt.Scanln(&subject)
				fmt.Print("Calificación: ")
				fmt.Scanln(&note)
				args := args.Args{Name:name,Subject:subject,Note:note}
				// invocar agregar calificación por rpc
				err = c.Call("Server.AddNoteBySubject", args, &reply)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(reply)
				}
			case 2:
				var name string
				var result float64
				fmt.Print("Nombre: ")
				fmt.Scanln(&name)
				// invocar obtener promedio por rpc
				err := c.Call("Server.GetStudentAVG", name, &result)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Resultado: ", result)
				}
			case 3:
				var result float64
				// invocar obtener promedio general por rpc
				err := c.Call("Server.AVGsByStudents", 0.0, &result)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Resultado: ", result)
				}
			case 4:
				var subject string
				var result float64
				fmt.Print("Materia: ")
				fmt.Scanln(&subject)
				// invocar obtener promedio de materia por rpc
				err := c.Call("Server.AVGsBySubjects", subject, &result)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Resultado: ", result)
				}
			case 5:
				c.Close()
				fmt.Println("Saliendo...")
				return
			default:
				fmt.Println("Opción incorrecta")
		}
	}
}


func main() {
	client()
}