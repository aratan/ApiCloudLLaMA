package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Result struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Job struct {
	ID        string
	Phrase    string
	CreatedAt time.Time
}

type JobManager struct {
	sync.RWMutex
	Jobs map[string]*Job
}

func NewJobManager() *JobManager {
	return &JobManager{
		Jobs: make(map[string]*Job),
	}
}

func (jm *JobManager) AddJob(phrase string) *Job {
	jm.Lock()
	defer jm.Unlock()

	jobID := generateJobID()
	job := &Job{
		ID:        jobID,
		Phrase:    phrase,
		CreatedAt: time.Now(),
	}
	jm.Jobs[jobID] = job

	return job
}

func (jm *JobManager) GetJob(jobID string) (*Job, bool) {
	jm.RLock()
	defer jm.RUnlock()

	job, ok := jm.Jobs[jobID]
	return job, ok
}

func (jm *JobManager) DeleteJob(jobID string) {
	jm.Lock()
	defer jm.Unlock()

	delete(jm.Jobs, jobID)
}

func generateJobID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func handleJobs(w http.ResponseWriter, r *http.Request) {
	// Obtener una lista de todos los trabajos en la cola
	jobManager := NewJobManager()
	jobs := make([]*Job, 0)
	jobManager.RLock()
	for _, job := range jobManager.Jobs {
		jobs = append(jobs, job)
	}
	jobManager.RUnlock()

	// Crear una estructura que contenga los datos a devolver en formato JSON
	response := struct {
		Jobs []*Job `json:"jobs"`
	}{
		Jobs: jobs,
	}

	// Codificar la estructura de respuesta en formato JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera de la respuesta HTTP para indicar que se devuelve JSON
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON en el cuerpo de la respuesta HTTP
	w.Write(jsonResponse)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Obtener la frase de la solicitud HTTP
	phrase := r.FormValue("phrase")

	// Agregar un trabajo a la cola de trabajos
	jobManager := NewJobManager()
	job := jobManager.AddJob(phrase)

	// Ejecutar la aplicación de terminal "llama" en segundo plano
	go func(job *Job) {
		cmd := exec.Command("./main", "-m", "./WizardLM-7B-uncensored.ggml.q4_0.bin", "-p", job.Phrase, "-n", "512")
		fmt.Println(cmd)
		// Obtener la salida de la aplicación
		output, err := cmd.Output()
		if err != nil {
			jobManager.DeleteJob(job.ID)
			return
		}

		// Almacenar la salida de la aplicación en un archivo
		err = ioutil.WriteFile(fmt.Sprintf("%s.txt", job.ID), output, 0644)
		if err != nil {
			jobManager.DeleteJob(job.ID)
			return
		}
	}(job)

	// Devolver una respuesta HTTP inmediata con la ID del trabajo
	result := &Result{Output: job.ID}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func handleResult(w http.ResponseWriter, r *http.Request) {
	// Obtener el valor del parámetro jobID de la solicitud HTTP
	jobID := r.URL.Query().Get("jobID")

	// Comprobar si el archivo jobID.txt existe en el directorio actual
	_, err := os.Stat(jobID + ".txt")
	exists := !os.IsNotExist(err)

	// Obtener la dirección IP del cliente que realiza la solicitud HTTP
	ip := r.RemoteAddr

	// Crear una estructura que contenga los datos a devolver en formato JSON
	response := struct {
		JobID  string    `json:"jobID"`
		Content string   `json:"content,omitempty"`
		Exists bool      `json:"exists"`
		Date   time.Time `json:"date"`
		IP     string    `json:"ip"`
	}{
		JobID:  jobID,
		Exists: exists,
		Date:   time.Now(),
		IP:     ip,
	}

	if exists {
		// Leer el contenido del archivo jobID.txt
		content, err := ioutil.ReadFile(fmt.Sprintf("%s.txt", jobID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Content = string(content)
	}

	// Codificar la estructura de respuesta en formato JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera de la respuesta HTTP para indicar que se devuelve JSON
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON en el cuerpo de la respuesta HTTP
	w.Write(jsonResponse)
}

func main() {
	fmt.Println("service enabled localhost:8080")
	fmt.Println("http://localhost:8080/llama?phrase=hola")
	fmt.Println("http://localhost:8080/result?jobID=1684222811777450241")
	fmt.Println("http://localhost:8080/jobs")

	// Asociar la función de manejo de solicitudes a la ruta /llama
	http.HandleFunc("/llama", handleRequest)

	// Asociar la función de manejo de solicitudes a la ruta /result
	http.HandleFunc("/result", handleResult)

	//ver las colas de trabajo
	http.HandleFunc("/jobs", handleJobs)

	// Iniciar el servidor HTTP en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
