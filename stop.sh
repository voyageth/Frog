#!/usr/bin/env bash

kill $(ps -ef | grep [v]oyageth/frog/server | awk '{print $2}')
