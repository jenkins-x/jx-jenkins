FROM gcr.io/jenkinsxio-labs-private/jx-cli-base:0.0.3

ENTRYPOINT ["jx-jenkins"]

COPY ./build/linux/jx-jenkins /usr/bin/jx-jenkins