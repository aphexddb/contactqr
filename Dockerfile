# Use phusion/baseimage as base image. To make your builds
# reproducible, make sure you lock down to a specific version, not
# to `latest`! See
# https://github.com/phusion/baseimage-docker/blob/master/Changelog.md
# for a list of version numbers.
FROM phusion/baseimage:0.11
LABEL maintainer="Gardiner Allen <aphexddb@gmail.com>"

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]

# install package
RUN apt-get update -y
RUN apt-get install -y jq curl wget

# add all docker files
ADD ./docker_files /

# contactqr service
ARG VERSION
ENV VERSION=${VERSION}
ADD ./release /opt/release
RUN chmod +x /opt/release/contactqr-${VERSION}-linux-amd64
RUN chmod +x /etc/service/contactqr/run
ARG PORT
ENV PORT=${PORT}
EXPOSE ${PORT}

# ui static files
ADD ./release/ui /opt/ui

# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*