# Web layer layout (Laravel + Vue 3)

Этот каталог добавляет **веб-слой** поверх существующего `task-service`, не ломая текущие классы и требования к БД.

## Что за архитектура

Сейчас у вас получается **3-слойная архитектура**:

1. **Go (`task-service`) — Core API / домен**
   - источник бизнес-логики по задачам;
   - работа с БД (через текущие репозитории/сервисы);
   - стабильный доменный контракт (`Task`).

2. **Laravel (`web/server`) — Web API / BFF слой**
   - web-facing API для браузера;
   - валидация входящих запросов, агрегация данных, web-specific правила;
   - единая точка для фронта (auth/cors/rate-limit/формат ответов).

3. **Vue 3 (`web/client`) — UI слой**
   - страницы, компоненты, пользовательские сценарии;
   - вызов API, отображение данных, управление состоянием на клиенте.

## Кто за что отвечает

### Go (обязательная доменная часть)
- `internal/task/orm` — доменная модель задачи;
- `internal/task/repository` — доступ к данным;
- `internal/task/service` — бизнес-операции;
- `internal/transport/grpc` — транспортный слой сервиса.

Go лучше считать **главным владельцем доменной логики**.

### Laravel (адаптер для веба)
- `app/Domain/Task/DTO` — контракт данных на web-слое;
- `app/Domain/Task/Http/Requests` — валидация payload;
- `app/Domain/Task/Http/Controllers` — обработка web-запросов;
- `routes/api.php` — внешний HTTP-контракт для Vue.

Laravel — **не замена Go**, а адаптер между браузером и core-сервисом.

### Vue 3 (чисто клиент)
- `entities/task/model` — клиентские типы;
- `entities/task/api` — вызовы HTTP API;
- `widgets/pages` — композиция UI.

Vue не должен содержать бизнес-правила БД/домена — только UI + orchestration.

## Рекомендуемый поток запроса

```text
Browser (Vue 3)
   -> Laravel API (валидация, auth, формат)
      -> Go task-service (доменная логика)
         -> PostgreSQL
```

## Структура

```text
web/
  server/                         # Laravel API (BFF / web API)
    app/Domain/Task/
      DTO/TaskData.php            # Контракт данных Task
      Models/Task.php             # Eloquent-модель tasks (если Laravel ходит в БД напрямую)
      Actions/                    # Бизнес-действия (Create, Update, Delete)
      Http/
        Controllers/TaskController.php
        Requests/
          StoreTaskRequest.php
          UpdateTaskRequest.php
    database/migrations/
      2025_07_08_163209_create_tasks_table.php
      2025_07_15_162510_link_tasks_with_users.php
    routes/api.php

  client/                         # Vue 3 SPA
    src/
      app/                        # bootstrap приложения
      shared/api/                 # базовый http-клиент
      entities/task/
        model/types.ts            # единый тип Task
        api/taskApi.ts            # вызовы Laravel API
        ui/TaskItem.vue
      widgets/task-list/ui/TaskList.vue
      pages/tasks/ui/TasksPage.vue
```

## Принцип совместимости

1. **Имя таблицы и поля**: `tasks(id, task, is_done, created_at, updated_at, deleted_at, user_id)`.
2. **JSON-поля API**: `id`, `task`, `user_id` (при необходимости можно расширить `is_done`).
3. **Одинаковая ответственность слоёв**:
   - `DTO / model` = контракт данных.
   - `Action / service` = бизнес-логика.
   - `Controller / api` = транспорт.

## Важный выбор (чтобы не дублировать логику)

Есть 2 варианта:

1. **Laravel -> Go (рекомендовано)**
   - Go = single source of truth для бизнес-логики и БД-контракта;
   - Laravel только проксирует/адаптирует веб-запросы.

2. **Laravel напрямую в БД**
   - быстрее старт, но есть риск дублирования логики между Go и Laravel.

Если цель — "те же классы и требования к БД", лучше фиксировать правила в Go,
а Laravel держать тонким web-слоем.


## Можно ли уже посмотреть UI?

Да. Теперь клиент в `web/client` можно запустить как обычное Vite-приложение.

```bash
cd web/client
npm install
npm run dev
```

После запуска откройте URL из консоли Vite (обычно `http://localhost:5173`).



### Совсем без запуска сервера/сборщика

Да, можно. Откройте файл `web/client/preview.html` напрямую в браузере
(двойной клик или `file:///.../web/client/preview.html`).

> Это статическое офлайн-превью (mock), а не полноценный Vue runtime.
> Для настоящего приложения с API всё равно нужен `npm run dev`.

### Если backend еще не поднят

Чтобы просто посмотреть готовый интерфейс без Laravel/Go API:

```bash
cd web/client
VITE_USE_MOCK=true npm run dev
```

В этом режиме будут показаны моковые задачи.

### Если backend поднят

```bash
cd web/client
VITE_API_URL=http://localhost:8000/api npm run dev
```

Тогда UI будет ходить в ваш Laravel API.
