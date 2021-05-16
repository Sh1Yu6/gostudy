src = $(wildcard *.go)

server: $(src)
	go build -o $@ $^

.PHONY:
	clean

clean:
	rm server
