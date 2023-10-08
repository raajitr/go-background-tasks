FROM python:alpine3.8 AS build
WORKDIR /app
COPY ./postgres/migration/ .

RUN pip install faker
RUN python script.py

FROM postgres:alpine

COPY --from=build /app/2_data.sql /docker-entrypoint-initdb.d
ADD /postgres/migration/1_schema.sql /docker-entrypoint-initdb.d

RUN chmod a+r /docker-entrypoint-initdb.d/*