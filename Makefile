SHELL:=/bin/bash -O extglob
BINARY=devbook  # Cambiar el nombre del binario si es necesario
VERSION=0.1
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Directorio raíz del proyecto
ROOT_DIR := $(shell pwd)

# Ruta del archivo .env en el directorio raíz
ENV_FILE := $(ROOT_DIR)/.env

up:
	docker-compose up

down:
	docker-compose down --remove-orphans