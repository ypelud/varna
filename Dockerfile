FROM scratch
COPY script/ca-certificates.crt /etc/ssl/certs/
COPY dist/varna /

EXPOSE 8080

ENTRYPOINT ["/varna"]
