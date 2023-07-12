build:
	go build .

run:
	go build . && export PORT=3000 && export MODE=DEBUG && ./mspinventory
