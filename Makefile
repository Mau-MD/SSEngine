build:
	(cd libs/proto && make build)
	docker build -t sse-master -f services/master/Dockerfile .
	docker build -t sse-streamer -f services/streamer/Dockerfile .