start:
	gin --path ./progressify -a 8080 -p 8081 run ./progressify/main.go

install_deps:
	go get github.com/gorilla/mux
	# go get github.com/gorilla/context
	go get github.com/go-redis/redis
	# go get github.com/asaskevich/govalidator
	# go get github.com/graymeta/stow
	# go get github.com/graymeta/stow/google
	# go get github.com/graymeta/stow/s3
	go get github.com/smartystreets/goconvey
	go get github.com/disintegration/imaging
	go get github.com/pkg/errors

	# DEPENDENCIES FOR LOCAL DEVELOPMENT
	go get github.com/codegangsta/gin

deploy:
	gcloud app deploy --project progressify-tool ./progressify/app.yaml

test:

	# try to get a key that does not exist
	http localhost:8081/https://www.w3schools.com/css/trolltunga.jpg
	#
	#
	#
	# try to get a non-existing URL
	http localhost:8081/https://www.w3schools.com/css/this_image_does_not_exist.jpg
	#
	#
	#
	# try to fetch a resource that is not an image
	curl -L localhost:8081/https://czaplinski.io/
	#
	#
	#
	# get a key that does exist
	redis-cli set "https://www.w3schools.com/css/trolltunga.jpg" "https://www.w3schools.com/css/trolltunga.jpg"
	http localhost:8081/https://www.w3schools.com/css/trolltunga.jpg
	redis-cli del "https://www.w3schools.com/css/trolltunga.jpg"

# vet:
# 	go vet $(go list ./... | grep -v /vendor/)