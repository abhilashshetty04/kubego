FROM ubuntu

COPY ./kubego ./kubego

ENTRYPOINT ["./kubego"]
