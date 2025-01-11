#!/bin/bash
docker image build -f Dockerfile -t forum-docker .
docker container run -p 8001:8001 --detach --name forum forum-docker