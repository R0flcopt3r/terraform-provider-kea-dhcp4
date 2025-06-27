FROM golang:alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk add git
ARG CGO_ENABLED=0
ARG GOARCH=amd64 
ARG GOOS=linux 
RUN go build -a -o terraform-provider-kea-dhcp4_v1.0.0

FROM scratch AS exporter
COPY --from=builder /build/terraform-provider-kea* /

FROM hashicorp/terraform:latest AS runner
RUN mkdir -p /root/.terraform.d/plugins/terraform.local/r0flcopt3r/kea-dhcp4/1.0.0/linux_amd64
COPY --from=builder /build/terraform-provider-kea-dhcp4_v1.0.0 /root/.terraform.d/plugins/terraform.local/r0flcopt3r/kea-dhcp4/1.0.0/linux_amd64/
RUN mkdir /tffiles
WORKDIR /tffiles
COPY ./test-data/terraform/*.tf /tffiles/
COPY ./entrypoint.terraform.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
