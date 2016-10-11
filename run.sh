#!/usr/bin/env bash

export GOPATH=/home/ec2-user/go
$GOPATH/bin/revel run github.com/voyageth/frog/server dev 8080
