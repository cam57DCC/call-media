FROM alpine

RUN apk add --update supervisor && rm  -rf /tmp/* /var/cache/apk/*

WORKDIR /root/

ADD conf.d/supervisord.conf /etc/

ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]