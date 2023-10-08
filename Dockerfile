FROM ubuntu:22.04
LABEL authors="yann"
WORKDIR /whisper-go
RUN sed -i 's#http://archive.ubuntu.com/#http://mirrors.tuna.tsinghua.edu.cn/#' /etc/apt/sources.list
# 更新软件包信息
RUN apt-get update && apt-get install  -y wget
RUN mkdir -p /whisper-go/whisper/models/
ADD dist/whisper-go /whisper-go/whisper-go
ADD whisper/main /whisper-go/whisper/
ADD whisper/models/ggml-base.bin /whisper-go/whisper/models/