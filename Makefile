BUILD_RELEASE_TAG = \
	buildReleaseTag() { \
		id=$$(docker image ls arnaudmorisset/serv-e:prod_latest --format '{{.ID}}'); \
		echo prod_$$id; \
	}; buildReleaseTag

build-image:
	docker build --tag arnaudmorisset/serv-e:prod_latest .
	docker tag arnaudmorisset/serv-e:prod_latest arnaudmorisset/serv-e:$$($(BUILD_RELEASE_TAG))

build-and-push-image: build-image
	docker push arnaudmorisset/serv-e:$$($(BUILD_RELEASE_TAG))
