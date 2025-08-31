all: test

test:
	@go test

clean:
	@make -C demos/demo01 clean
	@rm -f *~
