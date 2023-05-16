# ApiCloudLLaMA
The idea is to make an api that everyone can consume in their GPT4-like applications but with free models, to democratize its use.

Thought in a modular system that reminds me of the shirts that are clicked in arduino for extra functions.

We use our application on the LLaMA executable, providing an Api with queue service.
Linux:
go build api.go
./api

Windows
GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe app.go

batou@kusanagi: ~/llama.cpp/./api

Browser:
curl host:8080/llama?phrase="hi"
curl http://localhost:8080/result?jobID=1684234007692666387


