#!/bin/bash

sudo protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/*.proto

sudo docker build -t local/gcd -f DockerfileGCD .
sudo docker build -t local/api -f DockerfileAPI .

sudo kubectl apply -f api.yaml
sudo kubectl apply -f gcd.yaml