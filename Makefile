Image:=sklrsn/robot

all: tests build run

build:
			sudo docker build -t $(Image) .
tests:
		go test

run:
	 sudo docker run $(Image)
