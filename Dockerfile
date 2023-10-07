FROM alpine:latest

# Define the project name | 定义项目名称
ARG PROJECT=mms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mms.yaml
# Define the author | 定义作者
ARG AUTHOR="yuansu.china.work@gmail.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /build/${PROJECT}_api ./
COPY --from=builder /build/etc/${CONFIG_FILE} ./etc/

EXPOSE 9104

ENTRYPOINT ./${PROJECT}_api -f etc/${CONFIG_FILE}