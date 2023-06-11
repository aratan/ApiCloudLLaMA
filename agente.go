package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Función para imprimir el mensaje en color rojo
func printError(message string) {
	fmt.Printf("\033[31m%s\033[0m\n", message)
}
func printResult(message string){
	fmt.Printf("\033[94m%s\033[0m\n", message)
}

func limitarTexto(texto string, limite int) string {
	if len(texto) > limite {
		return texto[:limite]
	}
	return texto
}

func main() {
	// Parsear los argumentos desde la línea de comandos
	urlPtr := flag.String("url", "", "URL del buscador")
	startTagPtr := flag.String("startTag", "", "resultados")
	endTagPtr := flag.String("endTag", "", "wikipedia")
	flag.Parse()

	// Validar los argumentos
	if *urlPtr == "" || *startTagPtr == "" || *endTagPtr == "" {
		fmt.Println("Usage: go run main.go -url <URL> -startTag <startTag> -endTag <endTag>")
		os.Exit(1)
	}
	// todo a minusculas
	
	url := strings.ToLower(*urlPtr)
	startTag := strings.ToLower(*startTagPtr)
	endTag := strings.ToLower(*endTagPtr)

	// Realizar la petición HTTP a la página web
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Leer el contenido de la respuesta HTTP
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyBytes)

	// Extraer el texto de los elementos HTML y mostrarlo
	//fmt.Println("Texto completo:")

	// Filtros

	// Eliminar código HTML
	re := regexp.MustCompile(`<(.*?)>`)
	body = re.ReplaceAllString(body, "")

	// Eliminar código CSS
	re = regexp.MustCompile(`<style\b[^>]*>(.*?)</style>`)
	body = re.ReplaceAllString(body, "")

	// Eliminar código JavaScript
	re = regexp.MustCompile(`<script\b[^>]*>(.*?)</script>`)
	body = re.ReplaceAllString(body, "")

	//fmt.Println(texto)
	//fmt.Println(body)

	// Buscar la posición del inicio y fin de la cadena
	startIndex := strings.Index(body, startTag)
	endIndex := strings.Index(body, endTag)

	if startIndex == -1 || endIndex == -1 {
		fmt.Println(body)
		printError("No se encontraron las etiquetas correspondientes.")
		
		return
	}

	// Obtener el contenido entre las etiquetas
	content := body[startIndex+len(startTag) : endIndex]
	// limitar la respuesta para el prompts
	textoLimitado := limitarTexto(content, 200)
	printResult(textoLimitado)
	//fmt.Println("Contenido entre las etiquetas:")
	//printResult(content)
	return
	//fmt.Println(content)
}
//   ./links -url "https://www.bing.com/search?q=Warcraft" -startTag "resultados" -endTag "wikipedia"
