#!/usr/bin/env bash

kubectl -n kube-system delete configmap nginx-template 
kubectl delete -f manifests/ingress/
