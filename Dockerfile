# Grab the latest alpine image
FROM alpine:latest
LABEL maintainer="Gardiner Allen <aphexddb@gmail.com>"

# install packages
RUN apk add --no-cache --update jq curl wget bash

# add all docker files
ADD ./docker_files /

# service version configured at build time
ARG VERSION
ENV VERSION=${VERSION}
ADD ./release /opt/contactqr
RUN chmod +x /opt/contactqr/contactqr-${VERSION}-linux-amd64
RUN chmod +x /etc/service/contactqr/run

# service port default configured at build time, can be overriden wuth ENV value
ARG PORT
ENV PORT=${PORT}
EXPOSE ${PORT}

# Run the image as a non-root user
RUN adduser -D serviceuser
USER serviceuser

# Run the app.  CMD is required to run on Heroku
CMD ["/etc/service/contactqr/run"]