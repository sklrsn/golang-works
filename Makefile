Image:=sklrsn/robot

all: tests build 

build:
			sudo docker build -t $(Image) .
tests:
		go test

up:
	 sudo docker run $(Image)
