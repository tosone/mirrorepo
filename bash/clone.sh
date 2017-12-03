#!/usr/bin/env bash

git clone --progress $1 $2 2>&1 | tee -a $3