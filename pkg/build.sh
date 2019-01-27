#! /bin/bash

# Build web UI
cd ~/gopath1/src/todogo/pkg/web
go install
cp ~/gopath1/bin/web ~/gopath1/bin/web_ui/web
cp -R ~/gopath1/src/todogo/pkg/template ~/gopath1/bin/web_ui