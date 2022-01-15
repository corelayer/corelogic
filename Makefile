run:
	go run main.go

testcompile:
	go run main.go > testcompile.txt && wc -l testcompile.txt

testdeploy:
	scp testcompile.txt nsroot@"$(DEVVPX)":/var/tmp/input.txt
	ssh -l nsroot "$(DEVVPX)" 'batch -f /var/tmp/input.txt -outfile /var/tmp/output.txt'

testresult:
	scp nsroot@"$(DEVVPX)":/var/tmp/output.txt .
	cat output.txt | grep ERROR

testrun:
	make testcompile
	make testdeploy
	make testresult

testclean:
	rm testcompile.txt
	rm output.txt
	