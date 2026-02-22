# myLinter (loglint)

`loglint` — кастомный линтер для проверки текстов лог‑сообщений.

Он находит проблемы в строковых литералах, переданных в вызовы логгеров `log.*`, `slog.*`, `zap.*`:

- сообщение начинается с заглавной буквы (нужно со строчной)
- встречаются не‑ASCII/не‑английские символы
- встречаются запрещённые спецсимволы/эмодзи/повторяющаяся пунктуация
- встречаются потенциально чувствительные данные (password/token/api_key/…)

Проект интегрируется в `golangci-lint` **автоматическим способом (The Automatic Way)** через `golangci-lint custom` и **module plugins**.

---

## Требования

- Go (версия должна подходить вашему окружению; проект использует `go 1.26` в `go.mod`)
- `git`
- `golangci-lint` v2 (нужна команда `golangci-lint custom`)

---

## Установка и сборка кастомного golangci-lint

1) Перейдите в корень проекта

2) В корне уже лежит конфигурация сборки `.custom-gcl.yml`.

Сейчас она подключает локальный модуль как plugin:

- `module: github.com/aleksejevelkin/myLinter`
- `path: .`
- `import: github.com/aleksejevelkin/myLinter` (важно: именно этот импорт запускает `init()` в `plugin.go` и регистрирует линтер)

3) Соберите кастомный бинарник:

```bash
golangci-lint custom
```

По умолчанию появится бинарник `./custom-gcl`.

Если хотите увидеть логи сборки:

```bash
golangci-lint custom -v
```

---

## Настройка golangci-lint

В корне проекта лежит пример конфигурации `.golangci.yml` в формате v2, где:

- линтер включается через `linters.enable: [loglint]`
- плагин объявлен в `linters.settings.custom.loglint` с `type: "module"`

Фрагмент:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglint

  settings:
    custom:
      loglint:
        type: "module"
        description: "Проверка log-сообщений: строчные буквы, только английский, спецсимволы/эмодзи, чувствительные данные"
        settings: {}
```

---

## Конфигурация правил

Настройка делается в `.golangci.yml` в секции:

- `linters.settings.custom.loglint.settings`

Доступные параметры:

- `lowercase` (bool) — проверка, что сообщение не начинается с заглавной буквы
- `englishOnly` (bool) — проверка на не‑ASCII/не‑английские символы
- `specialChars` (bool) — проверка на запрещённые спецсимволы/эмодзи/повторы
- `sensitive` (bool) — проверка на чувствительные данные
- `sensitiveKeywords` ([]string) — кастомный список ключевых слов для sensitive‑правила (опционально)

По умолчанию все правила включены.

Пример: отключить проверку спецсимволов и переопределить список sensitive‑ключей:

```yaml
linters:
  settings:
    custom:
      loglint:
        type: "module"
        settings:
          specialChars: false
          sensitiveKeywords:
            - password
            - token
```

---

## Использование

Запускайте линт **кастомным бинарником**, который вы собрали на шаге выше:

Проверить весь проект:

```bash
./custom-gcl run ./...
```

Проверить только примеры:

```bash
./custom-gcl run ./example/...
```

Посмотреть список доступных линтеров:

```bash
./custom-gcl linters
```

---

## Использование через custom-gcl (пошагово)

`custom-gcl` — это кастомный бинарник `golangci-lint`, собранный командой `golangci-lint custom` на основе `.custom-gcl.yml`. Он уже «вшивает» в себя наш module‑плагин `loglint`.

### 1) Собрать/обновить custom-gcl

Собирать нужно после изменений в `plugin.go`, `analyzer/` или `checkers/`:

```bash
rm -f custom-gcl
golangci-lint custom -v
```

> Подсказка: `-v` полезен, чтобы увидеть, что во время сборки добавился импорт `_ "github.com/aleksejevelkin/myLinter"`.

### 2) Запуск линтера

Проверить весь репозиторий:

```bash
./custom-gcl run ./...
```

Проверить только примеры:

```bash
./custom-gcl run ./example/...
```

Запустить с подробным выводом:

```bash
./custom-gcl run -v ./...
```

### 3) Посмотреть доступные линтеры

```bash
./custom-gcl linters
```

### 4) Как custom-gcl находит конфиг

`./custom-gcl run` ищет `.golangci.yml` стандартным способом (в текущей директории и выше по дереву).

Если нужно явно указать конфиг:

```bash
./custom-gcl run -c .golangci.yml ./...
```

---

## Авто‑исправления (SuggestedFixes)

Для части проблем `loglint` добавляет `SuggestedFixes`:

- `lowercase`: делает первую букву строчной (например, `"Hello"` → `"hello"`)
- `specialChars`: удаляет запрещённые символы (`@ # $ % ^ & * ~`) и сжимает повторы `! ? .`

`golangci-lint` **показывает** эти подсказки, но **не применяет их автоматически**. Обычно они применляются через редактор/IDE (например, через `gopls` — Code Action / Quick Fix).

---

## Пример срабатываний

В папке `example/` лежат небольшие файлы для ручной проверки.

- `example/good.go` — корректные сообщения
- `example/bad_lowercase.go` — начинается с заглавной буквы
- `example/bad_english.go` — не‑английские символы
- `example/bad_special.go` — спецсимволы/повторяющаяся пунктуация
- `example/bad_sensitive.go` — чувствительные данные
- `example/ignored.go` — пример, который **не должен** сработать (строка не литерал в вызове)

Например, в `example/bad_sensitive.go` есть:

```go
log.Println("password: 12345")
```

А в `example/bad_english.go` есть:

```go
log.Println("привет мир")
```

Ожидаемые сообщения будут в формате:

```
log message issue: ...
```

Быстрый прогон по примерам:

```bash
./custom-gcl run ./example/...
```

---

## Архитектура проекта (кратко)

- `checkers/` — набор правил (проверки строк)
- `analyzer/` — анализатор на базе `go/analysis`, ищет строковые литералы в вызовах логгеров и применяет `checkers`
- `plugin.go` — module‑плагин для `golangci-lint`: регистрирует линтер `loglint` через `plugin-module-register`