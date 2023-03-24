#!/bin/bash
docker run -td --env repo=$2 $1 > container.txt
