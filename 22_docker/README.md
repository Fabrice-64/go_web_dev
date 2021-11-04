# DOCKER COMMANDS

## Docker Installation
docker version
docker run hello-world 
docker ps -a
docker ps

## Docker Images
docker search <search term>
docker run docker/whalesay cowsay boo
docker images
docker ps -a
docker ps

## Build Image
### Create Directory
mkdir mydockerbuild2
### Create File inside directory
nano Dockerfile
### Fill the File
FROM ubuntu:latest
RUN apt-get -y update && apt-get install -y curl
### Build the image
docker build -t <give it an image name> . 
[comment]:  "Penser au point en fin de ligne"
### Check available images
docker images
### Run the image with curl
curl --head www.google.com
### Exit 
exit









