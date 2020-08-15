#!/usr/bin/env bash
# You probably already have a container named fuzzbuzz running
docker build  -t fuzzbuzz-interview . && docker run fuzzbuzz-interview
