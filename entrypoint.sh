#!/bin/bash

make createsuperuser

if [[ $1 = "debug" ]]; then
    tail -F app.log
elif [[ $1 = "stage" ]]; then
    make run
fi
