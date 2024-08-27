ARG BASE_IMAGE=ubuntu

FROM ubuntu AS packages

RUN apt-get update && apt-get install -y \
	tini

FROM ${BASE_IMAGE}

COPY --from=packages /usr/bin/tini /usr/bin/tini
COPY --from=packages /usr/bin/tini-static /usr/bin/tini-static
COPY --from=packages /usr/share/doc/tini /usr/share/doc/tini

# COPY ./dist /go/src/app
# RUN chmod +x /go/src/app/*

# for non-root user
# RUN chown -R 1000:1000 /app
# USER 1000

WORKDIR /go/src/app

ENTRYPOINT ["tini", "--"]
CMD ["./dist/add-config"]
