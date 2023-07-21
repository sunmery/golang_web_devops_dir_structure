#!/bin/bash

kubectl apply -f golang-app-deployment.yaml
kubectl apply -f web-app-deployment.yaml
kubectl apply -f postgres-deployment.yaml
