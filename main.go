package main

import (
	"fmt"
	"./contenido"
)
func main() {
	var op int
	contenedor:=contenidos.ContenidoWeb{}
	cont := 1
	for cont != 0 {
		op = 0
		fmt.Println("...Opciones...\n1)Imagen\n2)Audio\n3)Video\n4)Mostrar\n0)Salir")
		fmt.Scan(&op)
		switch op{
		case 1: 
		var a string
		var b string
		var c string
			fmt.Println("...Imagen...")
			fmt.Println("Titulo: ")
			fmt.Scan(&a)
			fmt.Println("Formato: ")
			fmt.Scan(&b)
			fmt.Println("Canal: ")
			fmt.Scan(&c)
			img := contenidos.Imagen{Titulo: a, Formato: b, Canales: c}
			contenedor.Documentos = append(contenedor.Documentos,&img)
		case 2:
			var a string
			var b string
			var c string
			fmt.Println("...Audio...")
			fmt.Println("Titulo: ")
			fmt.Scan(&a)
			fmt.Println("Formato: ")
			fmt.Scan(&b)
			fmt.Println("Duracion: ")
			fmt.Scan(&c)
			aud := contenidos.Audio{Titulo: a, Formato: b, Duracion: c}
			contenedor.Documentos = append(contenedor.Documentos,&aud)
		case 3:
			var a string
			var b string
			var c string
			fmt.Println("...Video...")
			fmt.Println("Titulo: ")
			fmt.Scan(&a)
			fmt.Println("Formato: ")
			fmt.Scan(&b)
			fmt.Println("Frames: ")
			fmt.Scan(&c)
			video := contenidos.Video{Titulo: a, Formato: b, Frames: c}
			contenedor.Documentos = append(contenedor.Documentos,&video)
		case 4:
			contenedor.MostrarT()
		default:
			fmt.Println("Fin")
			cont = 0
		}
	}
}