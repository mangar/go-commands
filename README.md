# GO Command line tools

## Building

Inside the base dir, same level as ```src``` directory, type:

	export GOPATH=`pwd`



Go inside the command you want to build, for instance ```s3bucket``` and ```make````

	cd /src/s3bucket
	make





## Developing


Inside the base dir, same level as ```src``` directory, type:

	export GOPATH=`pwd``


Go inside the command you want to build, for instance ```s3bucket``` and open the editor

	cd /src/s3bucket
	



## Installing

__S3Upload__

	curl -sSL https://github.com/mangar/go-commands/archive/master.zip -o go-commands-master.zip; unzip go-commands-master.zip; cd go-commands-master/src/s3upload; make; make install



__S3Bucket__

	curl -sSL https://github.com/mangar/go-commands/archive/master.zip -o go-commands-master.zip; unzip go-commands-master.zip; cd go-commands-master/src/s3bucket; make; make install








