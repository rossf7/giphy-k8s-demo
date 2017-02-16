FROM alpine:3.4

RUN apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

# Add binary and supporting files
COPY giphydemo index.html /

# Metadata params
ARG BUILD_DATE
ARG VCS_URL
ARG VCS_REF

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.description="Kubernetes demo image using Giphy API and k8s downward API" \
      org.label-schema.vcs-url=$VCS_URL \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.schema-version="1.0" 

ENTRYPOINT ["/giphydemo"]
