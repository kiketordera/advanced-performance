# Basic project for rendegin Golang unsing gin

In this project you can find an easy, ready and documented implementation for a simple landing page written in Golang and plain SCSS & HTML with best practices.

The project is equiped with all functionallity you need to create your own project from here.

## This is an <h2> tag

###### This is an <h6> tag

## For making Hotreloading for first time

export GOPATH=$HOME/go
go get -u github.com/cosmtrek/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
alias air='$(go env GOPATH)/bin/air'
air

## For making Hotreloading AFTER for first time

export GOPATH=$HOME/go
alias air='$(go env GOPATH)/bin/air'
air

# If you have a permission denied error, run the following:

chmod u+x air

# If we want to use the hot reload with the make file, write:

make watch
