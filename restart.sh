#!/bin/bash
cd .deploy
docker-compose rm -f
docker-compose up --build -d
