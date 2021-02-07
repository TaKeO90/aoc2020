INPUTS=sample.txt input

.PHONY: test
test: main
	./main $(INPUTS)

main: main.go
	go build -o main main.go
