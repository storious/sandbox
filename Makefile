run:
	go run main.go

build:
	go build -o sandbox .

test:build
	./sandbox --Debug \
		--exe_path="./main"
		--input_path="./tests/app/std/in" \
		--output_path="./tests/app/std/out" \
		--error_path="./tests/app/std/out" 

clean:
	rm -f *.log