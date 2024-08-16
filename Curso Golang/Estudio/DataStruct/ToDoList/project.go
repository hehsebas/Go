package main

import (
	"bufio"
	"fmt"
	"os"
)

type tarea struct {
	nombre      string
	description string
	completado  bool
}
type ListaToDo struct {
	tareas []tarea
}

// Metodo para agregar tarea a la lista ToDo
func (l *ListaToDo) AgregarTarea(t tarea) {
	l.tareas = append(l.tareas, t)
}

// Metodo para marcar tarea como "Completada"
func (l *ListaToDo) MarcarCompletado(index int) {
	l.tareas[index].completado = true
}

// Método para editar tarea
func (l *ListaToDo) EditarTarea(index int, t tarea) {
	l.tareas[index] = t
}

//Método para eliminar tarea

func (l *ListaToDo) EliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)

}
func main() {
	//Instancia de la lista de tareas
	lista := ListaToDo{}
	//Instancia bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for {
		var Opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir\n")
		fmt.Scanln(&Opcion)

		switch Opcion {
		case 1:
			var t tarea
			fmt.Println("Ingrese el nombre de la tarea\n")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea\n")
			t.description, _ = leer.ReadString('\n')
			lista.AgregarTarea(t)
			fmt.Println("Tarea agregada")
		case 2:
			var index int
			fmt.Println("Ingrese el nombre de la tarea que desea marcar como completada\n")
			fmt.Scanln(&index)
			lista.MarcarCompletado(index)
			fmt.Println("Tarea marcada como completada\n")
		case 3:
			var index int
			var t tarea
			fmt.Println("Ingrese el indice de la tarea que desea editar\n")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea que desea editar\n")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea que desea editar\n")
			t.description, _ = leer.ReadString('\n')
			lista.EditarTarea(index, t)
			fmt.Println("La tarea ha sido editada con exito\n")
		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea eliminar\n")
			fmt.Scanln(&index)
			lista.EliminarTarea(index)
			fmt.Println("La tarea ha sido eliminar con exito\n")
		case 5:
			fmt.Println("Saliendo del programa...\n")
			return
		default:
			fmt.Println("Opción invalida\n")
		}
		//Listar todas las tareas
		fmt.Println("Lista de Tareas")
		fmt.Println("------------------------------------------------")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i+1, t.nombre, t.description, t.completado)

		}
		fmt.Println("------------------------------------------------")
	}

}
