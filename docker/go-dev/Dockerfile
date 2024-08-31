ARG BASE_IMAGE=golang:latest

FROM ${BASE_IMAGE}

ARG USER_ID=1000
ARG GROUP_ID=1000

RUN addgroup --gid ${GROUP_ID} go \
	&& adduser --disabled-password --gecos '' --uid ${USER_ID} --gid ${GROUP_ID} go

USER go

