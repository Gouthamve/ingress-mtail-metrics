# Deploying the Nginx Ingress controller

This example aims to demonstrate the deployment of an nginx ingress controller and
scrape metrics off using mtail. By using `prometheus/scrape: true`,
Prometheus will scrape it automatically and you start using the generated metrics right away.

## Default Backend

The default backend is a Service capable of handling all url paths and hosts the
nginx controller doesn't understand. This most basic implementation just returns
a 404 page:

```console
$ kubectl apply -f default-backend.yaml
deployment "default-http-backend" created
service "default-http-backend" created

$ kubectl -n kube-system get po
NAME                                    READY     STATUS    RESTARTS   AGE
default-http-backend-2657704409-qgwdd   1/1       Running   0          28s
```

## Custom Template and Other Configs

```console
$ kubectl -n kube-system create configmap nginx-template --from-file=nginx.tmpl                 
configmap "nginx-template" created
$ kubectl apply -f nginx-mtail-config-gen.yaml 
configmap "nginx-mtail-template" created
```

## Controller

You can deploy the controller as follows:

```console
$ kc apply -f nginx-ingress-controller.yaml 
deployment "nginx-ingress-controller" created
```
