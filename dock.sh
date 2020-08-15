#!/usr/bin/env bash
# You probably already have a container named fuzzbuzz running
docker build  -t fuzzbuzz-interview . && docker run -p 8080:8080 fuzzbuzz-interview
