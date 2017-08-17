# GO Command line tools


## S3Upload


__Building__

The command bellow will make the binary and install it at: ```/usr/local/bin```.
After the instalation you can test it by running: ```/usr/local/bin/s3upload --version```


	make
	make install


__Developing__

The command bellow will install the depencies needed to run the project


	make development




---



__TODO__

- comando para baixar o código todo e fazer o make na maquina...



\curl -sSL https://codeload.github.com/mangar/docker/zip/master | bash -s stable


curl -sSL https://github.com/mangar/go-commands/archive/feature/make.zip -o go-commands-master.zip; unzip go-commands-master.zip; cd go-commands-feature-make/src/s3upload; make; make install