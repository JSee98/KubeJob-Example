This repo hold the code for the blog on Kubernetes Jobs. The code is simple example for spawning real-time jobs from a Go server. 

## RUNNING THE CODE

Start the server

`go run main.go`

Make the POST request with a concurrency value of 10

`curl -X POST -d '{"concurrency": 10}' http://localhost:8080/kubejob`


NOTE: This a crude example and obviously not production ready system. Use it as inspiration only. For more details please read up on kube clients for go. 
