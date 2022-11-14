FROM rust:1-buster

WORKDIR /arti
RUN apt-get update && apt-get install git
RUN git clone https://gitlab.torproject.org/tpo/core/arti

WORKDIR /arti/arti
RUN cargo build --release

CMD ["/bin/sh"]
