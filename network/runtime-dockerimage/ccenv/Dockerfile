FROM golang:1.17.7-bullseye
RUN apt update && apt-get install -y \
	binutils-gold \
	g++ \
	gcc \
	git \
	musl-dev\
	curl

RUN curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.7.0.tar.gz" | tar -C /usr/local -xz
RUN ldconfig

RUN mkdir -p /chaincode/output /chaincode/input
RUN addgroup --gid 500 chaincode && adduser --disabled-login --uid 500 --home /home/chaincode --ingroup chaincode chaincode
RUN chown -R chaincode:chaincode /chaincode
USER chaincode
