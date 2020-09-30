for package runecount

go test -bench=.


for package addnumbers

GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
//Parallelism will show better results
GOGC=off go test -cpu 8 -run none -bench . -benchtime 6s


for package iodocs

//Concurrency will show better results on single core itself
GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s

GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s

