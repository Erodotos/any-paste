# Any Paste: A simple web application build with Fiber

<p align="center">
     <img src="https://gofiber.io/assets/images/logo.svg" width="250">
     <img src="https://github.com/kubernetes/kubernetes/raw/master/logo/logo.png" width="100">
</p>


## Table of contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Project architecture exmplation and Demo](#project-architecture-explanation-and-demo)

## Introduction
This project is about creating and deploying a web application comprising a backend and frontend microservices on a [Kubernetes](https://kubernetes.io/) cluster. When looking over the internet for similar projects endeavouring to understand the whole development and deployment picture, I could not find a particular example expanding on networking between microservices while using a frontend, backend and a database to store data. Thus, after assembling information from different sources, I created a simple web application using [Fiber](https://gofiber.io/) (a Go web framework). 

The application is called **Any Paste**, and it is a cheap clone of Pastebin.com. In other words, a user can submit any text, and Any Paste will store it in a database. Afterwards, the users can retrieve a specific piece of posted text, providing its identification number. 

## Prerequisites
You need access to a local or remote Kubernetes cluster, such as [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [AKS](https://azure.microsoft.com/en-us/products/kubernetes-service). You also need the [`kubectl`](https://kubernetes.io/docs/reference/kubectl/kubectl/) cli tool, which will allow you to connect to the desired Kubernetes cluster.

## Installation
The current installation does not include automation tools such as [`helm`](https://helm.sh/) or [`kustomize`](https://kustomize.io/) to automatically package and install our Kubernetes manifest files.

1. Connect to your Kubernetes cluster via kubectl - for each Kubernetes provider there is a different way to get on your cluster. Please ask Google or ChatGPT to show you the way :)
   
2. Create a new namespace for the application

   `kubectl create namespace any-paste`

3. Install the mysql instance using the following commands (note that in general it is a bad idea to deploy databases on Kubernetes. Make sure you watch my full video for more details):
   
     `kubectl apply -f .k8s/database/volume.yaml --namespace any-paste`

     `kubectl apply -f .k8s/database/secret.yaml --namespace any-paste`

     `kubectl apply -f .k8s/database/config-map.yaml --namespace any-paste`

     `kubectl apply -f .k8s/database/deployment.yaml --namespace any-paste`

     `kubectl apply -f .k8s/database/service.yaml --namespace any-paste`

4. Install the backend of our application using the following commands:
   
   `kubectl apply -f .k8s/api/deployment.yaml --namespace any-paste`

   `kubectl apply -f .k8s/api/service.yaml --namespace any-paste`

5. Install the frontend of our application using the following commands:
   
   `kubectl apply -f .k8s/website/deployment.yaml --namespace any-paste`
   
   `kubectl apply -f .k8s/website/service.yaml --namespace any-paste`

6. Finally, we need to deploy an ingress resource to manage the incoming traffic towards the cluster. Note that all deployed services (database, frontend and backend) are all deployed on a private virtual network and are not accessible from the outside world. Hence, we need to deploy ingress, which will route incoming traffic to our application. For this, use the following command (make sure you have enabled the Nginx ingress controller):
   
   `kubectl apply -f .k8s/ingress.yaml --namespace any-paste`



## Project architecture explanation and Demo
