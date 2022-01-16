#! /bin/zsh

fswatch -or ./ | xargs -n1 -I{} make docker-build-run