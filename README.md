
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

postdata: I need filter
{
   "id":"1684359534886646993",
   "phrase":"Hello, world!",
   "output":"main: build = 533 (fb62f92)\nmain: seed  = 42\nllama.cpp: loading model from ./WizardLM-7B-uncensored.ggml.q4_0.bin\nllama_model_load_internal: format     = ggjt v2 (latest)\nllama_model_load_internal: n_vocab    = 32001\nllama_model_load_internal: n_ctx      = 512\nllama_model_load_internal: n_embd     = 4096\nllama_model_load_internal: n_mult     = 256\nllama_model_load_internal: n_head     = 32\nllama_model_load_internal: n_layer    = 32\nllama_model_load_internal: n_rot      = 128\nllama_model_load_internal: ftype      = 2 (mostly Q4_0)\nllama_model_load_internal: n_ff       = 11008\nllama_model_load_internal: n_parts    = 1\nllama_model_load_internal: model size = 7B\nllama_model_load_internal: ggml ctx size =  68,20 KB\nllama_model_load_internal: mem required  = 5809,34 MB (+ 1026,00 MB per state)\nllama_init_from_file: kv self size  =  256,00 MB\n\nsystem_info: n_threads = 3 / 4 | AVX = 1 | AVX2 = 0 | AVX512 = 0 | AVX512_VBMI = 0 | AVX512_VNNI = 0 | FMA = 0 | NEON = 0 | ARM_FMA = 0 | F16C = 1 | FP16_VA = 0 | WASM_SIMD = 0 | BLAS = 0 | SSE3 = 1 | VSX = 0 | \nsampling: repeat_last_n = 64, repeat_penalty = 1,100000, presence_penalty = 0,000000, frequency_penalty = 0,000000, top_k = 40, tfs_z = 1,000000, top_p = 0,950000, typical_p = 1,000000, temp = 0,800000, mirostat = 0, mirostat_lr = 0,100000, mirostat_ent = 5,000000\ngenerate: n_ctx = 512, n_batch = 512, n_predict = 512, n_keep = 0\n\n\n Hello, world! I'm so excited to be here. This is my first post on my new blog. I will be sharing my thoughts, ideas, and experiences with you. Thank you for stopping by and supporting me in this new adventure. [end of text]\n\nllama_print_timings:        load time =  4493,72 ms\nllama_print_timings:      sample time =   102,87 ms /    48 runs   (    2,14 ms per token)\nllama_print_timings: prompt eval time =  4184,02 ms /     5 tokens (  836,80 ms per token)\nllama_print_timings:        eval time = 42784,89 ms /    47 runs   (  910,32 ms per token)\nllama_print_timings:       total time = 47398,80 ms\n",
   "created_at":"2023-05-17T23:38:54.886653598+02:00",
   "finished_at":"2023-05-17T23:39:42.401900595+02:00",
   "status":"completed"
}
