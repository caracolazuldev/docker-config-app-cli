ARG BASE_IMAGE=ubuntu

FROM ubuntu AS packages

RUN apt-get update && apt-get install -y \
	tini

FROM ${BASE_IMAGE}

COPY --from=packages /usr/bin/tini /usr/bin/tini
COPY --from=packages /usr/bin/tini-static /usr/bin/tini-static
COPY --from=packages /usr/share/doc/tini /usr/share/doc/tini

COPY . /app
RUN chown -R 1000:1000 /app
RUN chmod +x /app/add-config.sh


WORKDIR /app
USER 1000

ENTRYPOINT ["tini", "--"]
CMD ["awk","-f", "/app/add-config.awk"]
