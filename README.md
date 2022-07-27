# tunity_test

## Tasks
1. Create a go program that given a number n calculates the number of steps for n to reach 1 based on the Collatz conjecture,
2. set up a kubernetes cluster (e.g. using minikube) and expose the code as a service.
3. add a readme describing how to set it up and a script you can use to query the service for all values from 1...K for a given number K.

Note: you should be able to send a request to your local cluster and get a valid response using a bash command or your browser.

The solution should include all code, build, and k8s related files in a dedicated github repository.

### Collatz Conjecture
The Collatz conjecture is one of the most famous unsolved problems in mathematics. The conjecture asks whether repeating two simple arithmetic operations will eventually transform every positive integer into 1. It concerns sequences of integers in which each term is obtained from the previous term as follows: if the previous term is even, the next term is one half of the previous term. If the previous term is odd, the next term is 3 times the previous term plus 1. The conjecture is that these sequences always reach 1, no matter which positive integer is chosen to start the sequence. (Wikipedia)

# USAGE
## Start Minikube
```
minikube start
```
#### Go to root project folder and type
```
kubectl create -f deployment.yaml
kubectl get pods
```
#### Wait to status is Running
```
NAME                           READY   STATUS    RESTARTS   AGE
tunity-test-586bc84dbb-4f4xp   1/1     Running   0          78s
```

&nbsp;
&nbsp;

### Running as WEB Service

#### Expose ports
```
kubectl expose deployment tunity-test --type=NodePort --name=tunity-test-svc --target-port=3000
```
#### Get your url
```
minikube service tunity-test-svc --url

Ex: http://127.0.0.1:53685

```

Then access the url above with the parameters:

```http://<minikube_ip>:<minikube_port>/collatz/<positive_number>```

For Example:

```http://127.0.0.1:53685/collatz/10```

&nbsp;
&nbsp;

### Running with your Terminal
Get your pod name
```
kubectl get pods

Example:
tunity-test-586bc84dbb-4f4xp   1/1     Running   0          16m
```
Enter in Kubernetes bash:
```
kubectl exec -it <pod_name> -- /bin/sh

Example:
kubectl exec -it tunity-test-586bc84dbb-4f4xp -- /bin/sh
```
Then execute the following command:
```
./main <number>
```
&nbsp;
&nbsp;


## Running Locally (without minikube)
### bash
Run this project with any positive number.
```
go run main.go <positive number>
```

### local service
Run this project without any parameter.
```
go run main.go
```

`Then access http://localhost:3000/collatz/<positive_number>`

# Results
The result of the current script is a JSON object with a list of the number of steps and the math result calculated in this step.

Example with given number `10`
```
[
	{"Step":1,"Result":10},
	{"Step":2,"Result":5},
	{"Step":3,"Result":16},
	{"Step":4,"Result":8},
	{"Step":5,"Result":4},
	{"Step":6,"Result":2},
	{"Step":7,"Result":1}
]
```

# Tests
Go to `/tests` folder and run

```
go test -v collatz_test.go
```
