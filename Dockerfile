FROM alpine:3.10
COPY server ./
CMD [ "./server" ]

# jainpayal-macbookpro:bluegreen jainpayal$ export GOOS=linux
# jainpayal-macbookpro:bluegreen jainpayal$ export CGO_ENABLED=0
# jainpayal-macbookpro:bluegreen jainpayal$ go build -o server