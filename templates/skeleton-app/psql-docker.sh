#!/bin/bash
docker run --rm --name pg-skeleton-app -e POSTGRES_DB=skeleton-app -e POSTGRES_USER=skeleton-app -e POSTGRES_PASSWORD=OeWOpyZv -d -p 35432:5432 -v "$(pwd)/.postgres:/var/lib/postgresql/data" postgres
