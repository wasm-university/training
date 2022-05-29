#!/bin/bash
set -o allexport; source .env; set +o allexport
docker run -it ${IMAGE_NAME} bash

