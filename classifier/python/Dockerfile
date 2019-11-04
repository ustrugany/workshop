FROM python:3.7.4

WORKDIR /app

COPY app .

RUN apt-get update && \
    apt-get install libsasl2-dev

RUN pip install --default-timeout=100 --no-cache-dir -r requirements.txt

ENV TORCH_HOME /app/.cache
