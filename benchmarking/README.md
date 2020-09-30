for package runecount

go test -bench=.


for package addnumbers

GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
GOGC=off go test -cpu 8 -run none -bench . -benchtime 6s //Parallelism will show better results


for package iodocs

GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s //Concurrency will show better results
GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s

