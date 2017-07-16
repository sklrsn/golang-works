Image:=sklrsn/postgresql

all: build run

build:
			sudo docker build -t $(Image) .
run:
	 		docker run --rm -P --name pg_test $(Image)
