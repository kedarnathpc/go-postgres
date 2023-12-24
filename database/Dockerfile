FROM postgres:latest

# Environment variables for PostgreSQL
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres

# Copy the initialization script into the container
COPY ./init.sql /docker-entrypoint-initdb.d/