FROM gcr.io/jenkinsxio-labs-private/jxl-base:0.0.55

ENTRYPOINT ["jx-jenkins"]

COPY ./build/linux/jx-jenkins /usr/bin/jx-jenkins