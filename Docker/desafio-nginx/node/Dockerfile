FROM node:latest

WORKDIR /usr/app
COPY . .

RUN npm install

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
