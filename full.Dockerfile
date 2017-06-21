FROM golang:latest 
COPY ./dist/varna .

CMD ["./varna"]