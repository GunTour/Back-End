FROM golang

##buat folder APP
RUN mkdir /guntour

##set direktori utama
WORKDIR /guntour

##copy seluruh file ke completedep
ADD . .

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]
