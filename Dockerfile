FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.17-bullseye
WORKDIR /testing-aws-backend
ENV ALLOWED_CROSS_SITE_ORIGIN="https://frontend.test.mysite.fabianjrivasportfolio.com"
COPY . /testing-aws-backend
RUN go mod tidy
RUN go build main.go
CMD ["./main"]