FROM golang:1.14.3-alpine AS build

WORKDIR /src
COPY  . .

RUN go build -o main .

CMD [ "/src/out" ]

FROM scratch as bin

COPY --from=build /src /
CMD [ "/main" ]