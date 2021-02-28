# This reload the app whenever a go file is changed.
build:
	go build -o server main.go

run: build
	./server

watch:
	ulimit -n 10000 
	reflex -s -r '\.go$$' make run