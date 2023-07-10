build:
	go build .

run:
	go build . && export PORT=3000 && ./mspinventory
