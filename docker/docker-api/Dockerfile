FROM alpine:3.8

EXPOSE 8080

RUN adduser deploy -S

RUN mkdir /etc/crypto-wallet/

RUN chown deploy /etc/crypto-wallet/

USER deploy

COPY crypto-wallet /usr/local/bin/crypto-wallet

ENTRYPOINT ["/usr/local/bin/crypto-wallet"]
