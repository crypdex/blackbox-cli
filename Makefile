TAG?=

release: require-tag
	git tag ${TAG}
	git push origin ${TAG}
	goreleaser --rm-dist --debug

release-test:
	goreleaser --snapshot --skip-publish --rm-dist --debug


require-tag:
ifndef TAG
	$(error 'TAG' is undefined)
else
	@echo "configured for ${TAG}"
endif