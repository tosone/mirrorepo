#!/usr/bin/env bash

git clone --progress $1 $2 2>&1 | tee -a $3
git ls-files | xargs -n1 git blame --line-porcelain | sed -n 's/^author //p' | sort -f | uniq -ic | sort -nr
git log --all --format='%an <%ae>' | sort | uniq
git log --reverse --pretty=oneline --format="%ar" | head -n 1 # repo age
git ls-files | wc -l | tr -d ' ' # file count

git shortlog -n -s -e| awk '{args[NR]=$0;sum+=$0}END{for(i=1;i<=NR;++i){printf "%s♪%2.1f%%\n",args[i],100*args[i]/sum}}' | column -t -s♪ | wc -l # git summary