FROM frolvlad/alpine-glibc:alpine-3.12_glibc-2.32
WORKDIR /arti
RUN apk add --upgrade git curl openssl libressl-dev alpine-sdk

## Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs > install_rust
RUN chmod +x install_rust
RUN ./install_rust -y

## Arti
RUN git clone https://gitlab.torproject.org/tpo/core/arti
WORKDIR /arti/arti
# NOTE: see https://gitlab.torproject.org/tpo/core/arti/-/issues/571 for
# rationale behind the next command
RUN apk update && \
  apk add alpine-sdk sqlite-dev openssl-dev perl && \
  ~/.cargo/bin/cargo install --features static arti

CMD ["/bin/sh"]
