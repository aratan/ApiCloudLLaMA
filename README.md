
Project (Api Cloud LLaMA)
============

Description
-----------

This project, ApiCloudLLaMA, aims to create an API that can be easily consumed by developers in their applications. Similar to the modular concept of click-in shields in Arduino, our goal is to provide an API with open-source models, like GPT4, to democratize their usage.

<img src="https://th.bing.com/th/id/OIG.mdR6q5sRWVj2WDWk7THM?pid=ImgGn" width="200" height="200" >


Installation
------------

1.  Clone the repository to your local machine.
    
    ```shell
    git clone https://github.com/aratan/ApiCloudLLaMA.git
    ```
    
2.  Navigate to the project directory.
    
    ```shell
    cd ApiCloudLLaMA
    ```
    
3.  Compile and run the API by executing the following command:
    *   For Linux:
        
        ```shell
        go build -o api api.go
        ./api
        ```
        
    *   For Windows:
        
        ```shell
        set GOOS=windows
        set GOARCH=amd64
        go build -o app-amd64.exe api.go
        app-amd64.exe
        ```
        

Usage
-----

1.  Once the API is running, you can interact with it using an HTTP client like cURL.
    *   For example, to make a GET request with the phrase "hola" as a parameter, use the following command:
        

    
     ```shell
     
     curl http://localhost:8080/token
     
     curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '''{"phrase": "Hola amigo"}''' http://localhost:8080/job
     
     curl -H "Authorization: Bearer <token>" http://localhost:8080/job?job_id=<job_id>
     
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
3. Docker:
    https://hub.docker.com/repository/docker/systemdeveloper868/apicloudllama/    
    docker run apicloudllama


Contributing
------------

Contributions are welcome! To contribute to this project, please follow these steps:

1.  Fork the repository.
2.  Create a new branch.
3.  Make your changes and test them.
4.  Commit your changes.
5.  Push the branch to your forked repository.
6.  Open a pull request in the main repository.
7.  You can also donate Account: eth 0x3317eba7cF6a56a9b81A6B7148e4F5B2d67027bc

Here are some ideas for features or enhancements you might consider adding in the future to your ApiCloudLLaMA project:

Authentication and authorization: implement an authentication and authorization system to protect access to the API. This will allow you to control who can consume services and perform operations.


Complete API documentation: Create detailed API documentation that includes information on available endpoints, accepted parameters, response codes and usage examples. This will make it easier for developers to understand and correctly use the API.


Additional endpoints: Extend the API functionality by adding new endpoints to perform different types of operations related to language models. For example, you could offer endpoints for translation, text generation based on a specific context, text summarization, among others.


Support for different language models: Add the ability to use different language models in the API, allowing users to select the model that best suits their needs.


Scalability and performance: Optimize API performance and make it scalable to handle a larger load of requests. You can consider using techniques such as caching, load sharing and implementing a queuing system to process requests efficiently.


Activity monitoring and logging: Implement a logging and monitoring system to track API usage, log errors and perform performance analysis. This will allow you to identify and fix problems quickly, as well as gain valuable insight into how the API is being used.


Integration with cloud services: Explore the possibility of integrating your API with popular cloud services such as file storage, databases or authentication services. This will extend the capabilities of your API and provide users with more options for interacting with it.


Unit testing and test automation: Develop a comprehensive set of unit tests to ensure API stability and quality. Also consider test automation to facilitate bug detection and streamline the development process.

License
-------
Licencia Pública General GNU v3.0
Víctor Arbiol Martínez

