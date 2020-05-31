FROM openjdk:11-jre
COPY play-scala-rest-api-example-1.0-SNAPSHOT play
ENV PLAY_VERSION 2.8.2
EXPOSE 9000

RUN chmod +x play/bin/play-scala-rest-api-example

CMD play/bin/play-scala-rest-api-example