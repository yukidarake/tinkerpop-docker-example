FROM openjdk:8-jre-alpine

RUN apk --no-cache --update add tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    wget https://archive.apache.org/dist/tinkerpop/3.4.0/apache-tinkerpop-gremlin-console-3.4.0-bin.zip && \
    unzip apache-tinkerpop-gremlin-console-3.4.0-bin.zip && \
    mv apache-tinkerpop-gremlin-console-3.4.0 gremlin-console
