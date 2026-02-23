ifneq (,$(wildcard .env))
	include .env
	export
endif

SHELL := /bin/bash

.PHONY: help

help:
	@echo ""
	@echo "Доступные разделы:"
	@echo "  make help-migrate   — команды для работы с миграциями"
	@echo ""

include make/migrate.mk