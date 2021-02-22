FROM jeanblanchard/alpine-glibc:latest

COPY bin/app /usr/bin/app
COPY configs/config.json /var/data/config.json 

EXPOSE 8081 
EXPOSE 8082

CMD ["app", "-configs", "/var/data/config.json"]
