.PHONY: tools
tools:
	$(MAKE) -C cmd/troptgen install

.PHONY: gen
gen: tools
	go generate ./...
	find ./cache -maxdepth 1 ! -path ./cache -type d | xargs -I % sh -c 'cd %; go generate ./...'

.PHONY: test
test: tools
	go test ./...
	find ./cache -maxdepth 1 ! -path ./cache -type d | xargs -I % sh -c 'cd %; go test ./...'

.PHONY: gittag
gittag:
	test -n "$(TAG)"  # $$TAG
	find cache -maxdepth 1 ! -path cache -type d | xargs -I % sh -c 'git tag %/$(TAG)'
	git tag $(TAG)
