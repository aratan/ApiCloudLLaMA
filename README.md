Project ACL
============

Description
-----------

This project, ApiCloudLLaMA, aims to create an API that can be easily consumed by developers in their applications. Similar to the modular concept of click-in shields in Arduino, our goal is to provide an API with open-source models, like GPT4, to democratize their usage.

Installation
------------

1.  Clone the repository to your local machine.
    
    ```shell
    git clone https://github.com/username/repository.git
    ```
    
2.  Navigate to the project directory.
    
    ```shell
    cd project-directory
    ```
    
3.  Compile and run the API by executing the following command:
    *   For Linux:
        
        ```shell
        go build -o bin/app app.go
        ./bin/app
        ```
        
    *   For Windows:
        
        ```shell
        set GOOS=windows
        set GOARCH=amd64
        go build -o bin/app-amd64.exe app.go
        bin\app-amd64.exe
        ```
        

Usage
-----

1.  Once the API is running, you can interact with it using an HTTP client like cURL.
    *   For example, to make a GET request with the phrase "hola" as a parameter, use the following command:
        
        ```shell
        curl host:8080/llama?phrase="hola"
        
        {
           "output":"1684253956720900396"
        } 
        ```
        
2.  You can check the result of a specific job using its jobID. For example:
    
    ```shell
    curl http://localhost:8080/result?jobID=1684253956720900396
    
    {
      "jobID":"1684253956720900396",
      "content":" hola, amigo!\n¿Cómo estás? ¿Has estado muy cansado últimamente?",
      "exists":true,
      "date":"2023-05-16T18:22:19.179872455+02:00",
      "ip":"192.168.1.42:33638"
    }
    ```
    

Contributing
------------

Contributions are welcome! To contribute to this project, please follow these steps:

1.  Fork the repository.
2.  Create a new branch.
3.  Make your changes and test them.
4.  Commit your changes.
5.  Push the branch to your forked repository.
6.  Open a pull request in the main repository.

License
-------
