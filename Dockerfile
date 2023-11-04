FROM alpine

COPY dist/go-bank /bin/

EXPOSE 5001

ENTRYPOINT [ "/bin/go-bank" ]