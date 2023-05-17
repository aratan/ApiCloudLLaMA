package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Result struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Job struct {
	ID         string    `json:"id"`
	Phrase     string    `json:"phrase"`
	Output     string    `json:"output,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at,omitempty"`
	Status     string    `json:"status,omitempty"`
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
		Status:    "in_progress",
	}
	jm.Jobs[jobID] = job

	return job
}

func (jm *JobManager) GetJob(jobID string) (*Job, error) {
	jm.RLock()
	defer jm.RUnlock()

	job, ok := jm.Jobs[jobID]
	if !ok {
		return nil, fmt.Errorf("job %s not found", jobID)
	}

	return job, nil
}

func generateJobID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func main() {
	jobManager := NewJobManager()

	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var reqBody struct {
				Phrase string `json:"phrase"`
			}
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}

			// Verify JWT token
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// Create job
			job := jobManager.AddJob(reqBody.Phrase)

			go func() {
				//cmd := exec.Command("espeak", "-v", "en-us", "-s", "130", "-p", "50", "-g", "10", job.Phrase)
				cmd := exec.Command("./main", "-m", "./WizardLM-7B-uncensored.ggml.q4_0.bin", "-p", job.Phrase, "-n", "512", "-s", "42", "-t", "3")
                out, err := cmd.CombinedOutput()

				job.FinishedAt = time.Now()
				if err != nil {
					job.Status = "failed"
					job.Output = ""
					log.Printf("error while processing job '%s': %v", job.ID, err)
				} else {
					job.Status = "completed"
					job.Output = string(out)
					log.Printf("job '%s' completed successfully", job.ID)
				}
			}()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"job_id": job.ID})

		case http.MethodGet:
			jobID := r.URL.Query().Get("job_id")
			if jobID == "" {
				http.Error(w, "missing job_id", http.StatusBadRequest)
				return
			}

			// Verify JWT token
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			job, err := jobManager.GetJob(jobID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(job)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
// curl http://localhost:8080/token
// curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQzNTMyMDR9.DD63YjCdt2upWJkMCZR2OcbPJEwnHDuhDaxEg-v5IPk" -H "Content-Type: application/json" -d '{"phrase": "Hello, world!"}' http://localhost:8080/job
// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQzNTMyMDR9.DD63YjCdt2upWJkMCZR2OcbPJEwnHDuhDaxEg-v5IPk" http://localhost:8080/job?job_id=1684349700082163228
