.PHONY: all up_main up_external down_main down_external build_main build_external

all: up_main up_external

up_main:
	cd main && docker-compose up -d

up_external:
	cd external && docker-compose up -d

down_main:
	cd main && docker-compose down

down_external:
	cd external && docker-compose down

build_main:
	cd main && docker-compose build

build_external:
	cd external && docker-compose build

