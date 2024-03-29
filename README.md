## How to run the server (listens to / and /hello on port 8081)
```$ go run .```

## Deploy this app to a minikube cluster using argocd
[Tutorial](https://medium.com/@mehmetodabashi/installing-argocd-on-minikube-and-deploying-a-test-application-caa68ec55fbf "Installing ArgoCD on Minikube and deploying a test application, Mehmet Odabasi, PhD")

### Start the minikube cluster
- ```$ minikube start --driver=docker```

### Setup argocd
1. Create argocd namespace
   ```$ kubectl create namespace argocd```

1. Apply argocd installation manifest
   ```$ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml```

1. Verify argocd resources created
   ```$ kubectl get all -n argocd```

1. Once all pods show ready in previous command, startup the argocd web UI at localhost:8080
   ```$ kubectl port-forward svc/argocd-server -n argocd 8080:443```

1. Navigate to https://localhost:8080, login as admin, get the randomly generated password by running
   ```$ kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo```


Install Argo Workflows
> ```$ kubectl create namespace argo```
> ```$ kubectl apply -n argo -f https://github.com/argoproj/argo-workflows/releases/download/v3.5.2/install.yaml```