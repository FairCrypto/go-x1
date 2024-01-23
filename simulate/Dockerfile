FROM ubuntu:22.04

RUN apt update -y \
      && apt install -y golang wget git make curl

COPY .. /src/
RUN cd /src \
      && make x1 \
      && mv ./build/x1 /usr/local/bin/ \
      && rm -rf /src

entrypoint ["/usr/local/bin/x1"]
