#!/usr/bin/env bash

curl http://localhost:8080/api/people \
    -X POST \
    -d '{ \"name\": \"Jan\", "title": "Developer"}'
