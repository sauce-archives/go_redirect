all:
	go get github.com/neelance/godockerize
	godockerize build -t quay.io/saucelabs/go_redirect github.com/saucelabs/go_redirect
push:
	docker push quay.io/saucelabs/go_redirect
