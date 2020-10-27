package contenidos

import "fmt"

type Imagen struct {
	Titulo string 
	Formato string 
	Canales string
}
type Audio struct {
	Titulo string
	Formato string
	Duracion string
}
type Video struct {
	Titulo string
	Formato string
	Frames string
}

func (i *Imagen) Mostrar() string {
	return "Titulo: " + i.Titulo + "\nFormato: " + i.Formato + "\nCanales: " + i.Canales+"\n"
}
func (a *Audio) Mostrar() string {
	return "Titulo: " + a.Titulo + "\nFormato: " + a.Formato + "\nDuracion: " + a.Duracion+"\n"
}
func (v *Video) Mostrar() string {
	return "Titulo: " + v.Titulo + "\nFormato: " + v.Formato + "\nFrames: " + v.Frames+"\n"
}
func (c *ContenidoWeb) MostrarT() {
	for _, i := range c.Documentos {
		fmt.Println(i.Mostrar())
	}
}
type Multimedia interface {
	Mostrar() string
}
type ContenidoWeb struct {
	Documentos []Multimedia
}

