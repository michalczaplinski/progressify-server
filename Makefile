start:
	gin --path ./progressify -a 8080 -p 8081 run ./progressify/main.go

install_deps:
	go get github.com/gorilla/mux
	go get github.com/gorilla/context
	go get github.com/go-redis/redis
	go get github.com/asaskevich/govalidator
	go get github.com/graymeta/stow
	go get github.com/graymeta/stow/google
	go get github.com/graymeta/stow/s3

	go get github.com/codegangsta/gin

deploy:
	gcloud app deploy --project progressify-tool ./progressify/app.yaml

test:
	curl http://localhost:8081/
	echo
	curl http://localhost:8081/a
# vet:
# 	go vet $(go list ./... | grep -v /vendor/)