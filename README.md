# MicroGen 🚀

[![Go Version](https://img.shields.io/badge/Go-1.24.3+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/WhiCu/microgen)](https://goreportcard.com/report/github.com/WhiCu/microgen)

**MicroGen** - это мощный CLI инструмент для быстрого создания Go микросервисов с современной архитектурой и лучшими практиками.

## ✨ Особенности

- 🏗️ **Стандартная структура проекта** - следует [Go Project Layout](https://github.com/golang-standards/project-layout)
- 🚀 **Готовые шаблоны** - HTTP handlers, services, middleware, конфигурация
- 🔧 **Автоматическое управление зависимостями** - интеграция с `go mod tidy`
- 📁 **Гибкая настройка** - выбор директории назначения
- ⚡ **Taskfile** - автоматизация задач с помощью [Task](https://taskfile.dev/)

## 🚀 Быстрый старт

### Установка

```bash
go install github.com/WhiCu/microgen@latest
```

### Использование

```bash
# Создать новый микросервис в текущей директории
microgen gen

# Создать микросервис в указанной директории
microgen gen -d ./my-service

# Создать микросервис с автоматическим управлением зависимостями
microgen gen -d ./api-service --tidy

# Использовать алиасы
microgen generate -d ./user-service -t
mg gen -d ./payment-service --tidy
```

## 📁 Структура генерируемого проекта

```
my-service/
├── cmd/
│   └── app/
│       └── main.go          # Точка входа приложения
├── internal/
│   ├── app/
│   │   └── app.go           # Конфигурация приложения
│   ├── config/
│   │   ├── config.go        # Структуры конфигурации
│   │   └── load.go          # Загрузка конфигурации
│   ├── handler/
│   │   ├── handler.go       # Базовый handler
│   │   └── ping.go          # Health check endpoint
│   └── service/
│       └── service.go       # Бизнес-логика
├── test/
│   ├── reg.go              # Тестовые утилиты аутентификации
│   ├── client/
│   │   └── client.go        # HTTP клиент для тестов
│   ├── handler/
│   │   └── handler.go       # Тесты handlers
│   └── service/
│       └── service.go       # Тесты services
├── go.mod                   # Зависимости Go
├── go.sum                   # Хеши зависимостей
└── taskfile.yaml            # Автоматизация задач
```

## 🛠️ Команды

### `microgen gen`

Генерирует новый Go микросервис проект.

**Флаги:**
- `-d, --destation` - директория назначения для сгенерированного проекта (по умолчанию: текущая директория)
- `-t, --tidy` - автоматически выполнить `go mod tidy` после генерации для управления зависимостями

**Примеры:**
```bash
# Базовое использование
microgen gen

# Указать директорию
microgen gen -d ./my-service

# С автоматическим управлением зависимостями
microgen gen -d ./api-service --tidy
```

## 🔧 Зависимости

Проект использует следующие основные зависимости:

- [Gin](https://github.com/gin-gonic/gin) - HTTP веб-фреймворк
- [Cobra](https://github.com/spf13/cobra) - CLI фреймворк
- [Viper](https://github.com/spf13/viper) - управление конфигурацией
- [Cleanenv](https://github.com/ilyakaznacheev/cleanenv) - загрузка переменных окружения

## 📋 Требования

- Go 1.24.3 или выше
- Git
