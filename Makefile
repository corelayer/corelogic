run:
	go run main.go

testcompile:
	go run main.go > testcompile.txt

testdeploy:
	scp testcompile.txt nsroot@"$(DEVVPX)":/var/tmp/input.txt

testclean:
	rm testcompile.txt
	rm output.txt
	