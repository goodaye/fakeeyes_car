GOCC = go
TARGET = bin/fakeeyes_car
M_SRC = main.go

all : ${TARGET}


${TARGET} : ${M_SRC}
	@echo "build fakeeyes_car"
	GOARCH=arm GOARM=5 ${GOCC} build -o ${TARGET} $^

clean:
	/bin/rm -rf ${TARGET}