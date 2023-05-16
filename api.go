package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	// Obtener la ID del trabajo de la solicitud HTTP
	jobID := r.FormValue("jobID")

	// Obtener el trabajo correspondiente de la cola de trabajos
	jobManager := NewJobManager()
	job, ok := jobManager.GetJob(jobID)
	if !ok {
		http.Error(w, "Trabajo no encontrado", http.StatusNotFound)
		return
	}

	// Leer la salida de la aplicación desde el archivo
	output, err := ioutil.ReadFile(fmt.Sprintf("%s.txt", job.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Eliminar el trabajo de la cola de trabajos
	jobManager.DeleteJob(job.ID)

	// Crear una estructura de resultado con la salida de la aplicación
	result := &Result{Output: string(output)}

	// Codificar la estructura de resultado en JSON y enviarla como respuesta HTTP
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func main() {
    fmt.Println("service enabled localhost:8080")
    fmt.Println("http://localhost:8080/llama?phrase=hola")
    fmt.Println("http://localhost:8080/result?jobID=1684222811777450241")
	// Asociar la función de manejo de solicitudes a la ruta /llama
	http.HandleFunc("/llama", handleRequest)

	// Asociar la función de manejo de solicitudes a la ruta /result
	http.HandleFunc("/result", handleResult)

	// Iniciar el servidor HTTP en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
