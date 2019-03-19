#!/bin/bash

echo "starting up db..."
docker-compose up -d db

echo "waiting for db to be ready..."
while ! curl localhost:3306 &> /dev/null; do
    sleep 2
done

CODECAMP_DBUSER="root"
CODECAMP_DBPASS="codecamp"
CODECAMP_DBADDR="localhost"
CODECAMP_DBNET="tcp"

go test -v -tags=integration ./pkg/api/mysql
