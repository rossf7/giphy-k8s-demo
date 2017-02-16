# giphy-k8s-demo

Demo image using the Kubernetes [Downward API](https://kubernetes.io/docs/user-guide/downward-api/) to get annotations and serve a random gif using the Giphy API.

[![](https://images.microbadger.com/badges/image/rossf7/giphy-k8s-demo.svg)](https://microbadger.com/images/rossf7/giphy-k8s-demo "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/commit/rossf7/giphy-k8s-demo.svg)](https://microbadger.com/images/rossf7/giphy-k8s-demo "Get your own image badge on microbadger.com")

# Running locally in Minikube

* Add a hosts entry for futurama.local to the IP address of the Minikube VM

```
$ minikube ip
```
* Create the k8s objects

```
$ kubectl apply -f examples/futurama-deployment.yaml
$ kubectl apply -f examples/futurama-svc.yaml
$ kubectl apply -f examples/futurama-ingress.yaml
```
