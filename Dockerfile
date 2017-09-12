FROM tomwhiston/micro-golang:test
# If we also want to test the c code we will need to install some extra stuff here
RUN apk add --no-cache -U cmake make clang && \
    go get github.com/schrej/godacov && \
    go get -u github.com/haya14busa/goverage
ENTRYPOINT ["/go/src/app/tests/script.sh"]