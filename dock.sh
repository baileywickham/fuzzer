#!/usr/bin/env bash
docker build  -t fuzzer . && docker run -p 8080:8080 fuzzer
