# 
# Reference: 
# https://pt.slideshare.net/RaPz1/build-golang-projects-properly-with-makefiles
# 

# 
BINARY=cfdns

# 
VERSION=1.0.0
BUILD=`date +%FT%T%z`
HASH=`date +%FT%T%z`
# HASH=`git rev-parse HEAD`

# 
THIS_DIR=$(shell pwd)/../../
export GOPATH=$(THIS_DIR)

# 
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Hash=${HASH}"


default:
	cd $(GOPATH)
	go get gopkg.in/kyokomi/emoji.v1
	go build ${LDFLAGS} -o ${BINARY}
	

install:
	cp ./${BINARY} /usr/local/bin/${BINARY} 

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	
uninstall:
	if [ -f /usr/local/bin/${BINARY}  ] ; then rm /usr/local/bin/${BINARY}  ; fi 
	
