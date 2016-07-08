FROM apicht/rpi-golang:latest

RUN apt-get update && \
	apt-get install -y build-essential

ADD . /code
WORKDIR /code

RUN make build
CMD ["make"]
