GOCC = go
TARGET_X86 = bin/fakeeyes_car_x86
TARGET_ARM = bin/fakeeyes_car_arm
M_SRC = ./*.go

all : ${TARGET_ARM} ${TARGET_X86}


${TARGET_ARM} : ${M_SRC}
	# /bin/rm ${TARGET}
	@echo "build fakeeyes_car_arm"
	GOARCH=arm GOARM=5 ${GOCC} build -o ${TARGET_ARM} $^
${TARGET_X86} : ${M_SRC}
	# /bin/rm ${TARGET}
	@echo "build fakeeyes_car_x86"
	${GOCC} build -o ${TARGET_X86} $^
	# ${GOCC} build -o ${TARGET} ${M_SRC}

clean:
	/bin/rm -rf ${TARGET_ARM}
	/bin/rm -rf ${TARGET_X86}