run:
	go run main.go

output:
	go run main.go > output.txt

linecount:
	go run main.go > output.txt && wc -l output.txt