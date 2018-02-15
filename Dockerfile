FROM alpine
COPY /code/code /
RUN apk add --no-cache ca-certificates &&\
    chmod +x code
EXPOSE 80/tcp
CMD [ "./code" ]