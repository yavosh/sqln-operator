# SQLN Operator

A k8s operator for educational purposes, driving the experimental and a bit 
silly [sqln](https://github.com/yavosh/sqln) database.
The operator is using `kubebuilder` for the operator skeleton.

## Goals 

Create a stateful set of sqln (sqlite).
Create a backup schedule and export the database to cold storage.

## Using 

### Building

    make docker-build docker-push IMG=yavosh/sqln-operator

### Deploy to your configured cluster

    make deploy IMG=yavosh/sqln-operator:latest

### Deploy from docker hub

TODO ...

## Future goals 


## Intro

A dummy k8s operator which manages the sqln container. Created for testing the
operator patter in k8s the sqln container is just an sqlite data with a network (grpc)
interface.

## Development

To run the operator locally without deploying to k8s, after making changes:

    make manifests
    make install
    make run


## K8S Operator resources

- https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial.html
- https://betterprogramming.pub/building-a-highly-available-kubernetes-operator-using-golang-fe4a44c395c2
- https://iximiuz.com/en/posts/kubernetes-operator-pattern/
