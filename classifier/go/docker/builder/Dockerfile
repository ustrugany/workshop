FROM 624165964329.dkr.ecr.ap-southeast-1.amazonaws.com/golang-builder-image:latest

ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/Propertyfinder".insteadOf "git@github.com:Propertyfinder"
RUN echo "machine github.com login ${GITHUB_TOKEN} password x-oauth-basic" >> /root/.netrc
