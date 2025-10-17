all: test

test:
	@go test

doc:
	@go tool doc -short
#	@go tool doc -C demos/demo01 -cmd -all

cov:
	@go test -coverprofile=output.cov
	@go tool cover -func=output.cov

clean:
	@make -C demos/demo01 clean
	@rm -f *~ output.*
