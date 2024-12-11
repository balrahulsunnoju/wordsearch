#!/usr/bin/env bash
set -e

var="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
for _ in $(seq "$1"); do
    for _ in $(seq "$2"); do
         # pick a 1 char substring starting at a random position
        echo -n "${var:$(( RANDOM % ${#var} )):1}"
    done
    echo # newline
done
