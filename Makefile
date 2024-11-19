.PHONY: build deploy

build:
	sam build

deploy:
	sam deploy --profile miyabiii0310
