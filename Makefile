run:
	go run main.go

gomock:
	go generate ./...

dockerize:
	docker build -t todo-app-backend .

docker-run:
	docker run -p 8000:8000 todo-app-backend

test:
	go test -v ./...