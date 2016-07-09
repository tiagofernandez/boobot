FROM tiagofernandez/rpi-golang:latest

ADD . /code
WORKDIR /code

RUN make build-all
CMD ["make", "run"]
