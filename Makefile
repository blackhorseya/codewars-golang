.PHONY: test-with-coverage
test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out