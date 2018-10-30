# dpl ?= deploy.env
# include $(dpl)
# export $(shell sed 's/=.*//' $(dpl))

# build:
# 	docker build -t $(APP_NAME) .

# run: 
# 	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

# stop:
# 	docker stop $(APP_NAME); docker rm $(APP_NAME)

build:
	go build
run:
	./coconut
test: 
	go test 