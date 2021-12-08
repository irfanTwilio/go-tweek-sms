FROM centos:7.6.1810
RUN mkdir -p /opt/tweek
COPY go-tweek-sms /opt/tweek

WORKDIR /opt/tweek
ENTRYPOINT ["./go-tweek-sms"]

EXPOSE 8090