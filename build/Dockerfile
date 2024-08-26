FROM golang:1.22 as build
WORKDIR /app

COPY . .

RUN ls -R

RUN go build -v -tags lambda.norpc -o main ./src/cmd/main.go

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]