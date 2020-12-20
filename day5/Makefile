INPUT=sample.txt

.PHONY: solveI
solveI: main $(INPUT)
	./main -partI < $(INPUT)

.PHONY: solveII
solveII: main $(INPUT)
	./main -partII < $(INPUT)

main: main.go 
	go build -o main main.go
