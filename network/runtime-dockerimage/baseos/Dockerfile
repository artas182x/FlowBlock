FROM debian:bullseye

RUN addgroup --gid 500 chaincode && adduser --disabled-login --uid 500 --home /home/chaincode --ingroup chaincode chaincode
RUN apt-get update && apt-get install -y tzdata curl
RUN curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.7.0.tar.gz" | tar -C /usr/local -xz
RUN ldconfig
USER chaincode
