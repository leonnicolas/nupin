FROM gcr.io/distroless/static

ARG TARGETARCH
COPY ./bin/linux/${TARGETARCH}/nupin /bin/nupin

ENTRYPOINT ["/bin/nupin"]
