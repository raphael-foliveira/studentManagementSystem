

build:
	go build .

run:
	./studentManagementSystem

watch:
	while inotifywait -r -e modify,move,create,delete ./ ; do \
		make build && make run ; \
	done