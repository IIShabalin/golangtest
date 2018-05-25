Step 1: Pull the image from the Repository and create a Container on the Cluster

$ kubectl run my-app --image=gcr.io/some-repo/my-app:v1 --port=3000
deployment "my-app" created

$ kubectl get pods

NAME              READY     STATUS    RESTARTS   AGE
my-app            1/1       Running   0          10m

Step 2: Expose the Kubernetes Deployment through a Load Balancer

$ kubectl expose deployment my-app --type=LoadBalancer --port=8080 --target-port=3000
service "my-app" exposed

Step 3: Find the external IP of your Container

$ kubectl get svc

NAME     TYPE           CLUSTER-IP      EXTERNAL-IP
my-app   LoadBalancer   10.11.452.237   56.170.30.123

(Extra) Step 4: Use Kubernetes Rolling Updates

$ kubectl set image deployment/my-app  my-app=gcr.io/some-repo/my-app:v2

(Extra) Step 4: Clean Up

$ kubectl delete deployment my-app
$ kubectl delete svc my-app