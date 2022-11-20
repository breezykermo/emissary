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
RUN ~/.cargo/bin/cargo build -F static-sqlite --release

CMD ["/bin/sh"]
