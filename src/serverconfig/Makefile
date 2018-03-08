# 
BINARY=serverconfig

# 
VERSION=0.0.1
BUILD=`date +%FT%T%z`
HASH=`date +%FT%T%z`

# 
THIS_DIR=$(shell pwd)/../../
export GOPATH=$(THIS_DIR)

# 
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Hash=${HASH}"


default:
	cd $(GOPATH)
	go get gopkg.in/kyokomi/emoji.v1
	go build ${LDFLAGS} -o ${BINARY}
	

development:
	cd $(GOPATH)
	go get gopkg.in/kyokomi/emoji.v1


install:
	cp ./${BINARY} /usr/local/bin/${BINARY} 

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	
uninstall:
	if [ -f /usr/local/bin/${BINARY}  ] ; then rm /usr/local/bin/${BINARY}  ; fi 
	
