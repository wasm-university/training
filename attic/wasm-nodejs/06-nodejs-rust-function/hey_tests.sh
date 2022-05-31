#!/bin/bash

url_api=$(gp url 8080)
data='{"author":"Bob","text":"What is the meaning of life"}'

# This opens 100 connections, and sends 1000 requests. 
hey -n 1000 -c 100 -m POST -T "Content-Type: application/json" -d "${data}" ${url_api} 


