#!/usr/bin/env bash

kubectl -n kube-system create configmap nginx-template --from-file=manifests/ingress/nginx.tmpl
kubectl apply -f manifests/ingress/
