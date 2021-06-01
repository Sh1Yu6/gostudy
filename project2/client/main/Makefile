src = $(wildcard *.go)

client: $(src)
	go build -o $@ $^

.PHONY:
	clean

clean:
	rm client
