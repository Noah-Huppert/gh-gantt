# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:v0.12.6 as builder

RUN mkdir -p $GOPATH/src/github.com/Noah-Huppert/gh-gantt/server
WORKDIR $GOPATH/src/github.com/Noah-Huppert/gh-gantt/server

# this will cache the npm install step, unless package.json changes
ADD package.json .
RUN npm install --no-progress
ADD . .
RUN dep ensure
RUN buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD exec /bin/app
