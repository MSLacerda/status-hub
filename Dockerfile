# Build Stage
FROM lacion/alpine-golang-buildimage:1.9.7 AS build-stage

LABEL app="build-status-hub"
LABEL REPO="https://github.com/MSLacerda/status-hub"

ENV PROJPATH=/go/src/github.com/MSLacerda/status-hub

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/MSLacerda/status-hub
WORKDIR /go/src/github.com/MSLacerda/status-hub

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/MSLacerda/status-hub"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/status-hub/bin

WORKDIR /opt/status-hub/bin

COPY --from=build-stage /go/src/github.com/MSLacerda/status-hub/bin/status-hub /opt/status-hub/bin/
RUN chmod +x /opt/status-hub/bin/status-hub

# Create appuser
RUN adduser -D -g '' status-hub
USER status-hub

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/status-hub/bin/status-hub"]
