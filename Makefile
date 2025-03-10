run:
	go run main.go


test:clean
	go run main.go --Debug --version

clean:
	rm -f *.log