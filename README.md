To Do List тестовое задаие.

Цель:
Разработать REST API для системы управления задачами, которая позволяет пользователям создавать, просматривать, обновлять и удалять задачи.

Требования:
1. Создание задачи

Метод: POST /tasks
Описание: Создать новую задачу.

Запрос:
<pre><code>
Заголовки:
	Content-Type: application/json
Тело:
{
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)"
}
</code></pre>
Ответ:
<pre><code>
Успех (201 Created):
{
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
}
</code></pre>
Ошибка (400 Bad Request): Неправильный формат данных.<br>
Ошибка (500 Internal Server Error): Проблема на сервере.

2. Просмотр списка задач

Метод: GET /tasks
Описание: Получить список всех задач.

Запрос:

<pre><code>
Заголовки:
	-Type: application/json
Ответ:
Успех (200 OK):
[
    {
    "id": "int",
    "title": "string",
	"description": "string",
    "due_date": "string (RFC3339 format)",
	"created_at": "string (RFC3339 format)",
	"updated_at": "string (RFC3339 format)"
  	}
]
</code></pre>

Ошибка (500 Internal Server Error): Проблема на сервере.

3. Просмотр задачи

Метод: GET /tasks/{id}
Описание: Получить задачу по ID.
<pre><code>
Запрос:
Параметры пути:
    id: ID задачи (int)
Заголовки:
    Content-Type: application/json
Ответ:
Успех (200 OK):
{
  	"id": "int",
  	"title": "string",
  	"description": "string",
  	"due_date": "string (RFC3339 format)",
  	"created_at": "string (RFC3339 format)",
  	"updated_at": "string (RFC3339 format)"
}
</code></pre>
Ошибка (404 Not Found): Задача не найдена.<br>
Ошибка (500 Internal Server Error): Проблема на сервере.

4. Обновление задачи

Метод: PUT /tasks/{id}
Описание: Обновить задачу по ID.

<pre><code>
Метод: PUT /tasks/{id}
Описание: Обновить задачу по ID.
Запрос:
	Параметры пути:
		id: ID задачи (int)
	Заголовки:
		Content-Type: application/json
	Тело:
	{
  	"title": "string",
  	"description": "string",
  	"due_date": "string (RFC3339 format)"
	}
Ответ:
Успех (200 OK):
{
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
}
</code></pre>

Ошибка (400 Bad Request): Неправильный формат данных.<br>
Ошибка (404 Not Found): Задача не найдена.<br>
Ошибка (500 Internal Server Error): Проблема на сервере.

5. Удаление задачи

Метод: DELETE /tasks/{id}
Описание: Удалить задачу по ID.

<pre><code>
Запрос:
	Параметры пути:
		id: ID задачи (int)
	Заголовки:
		Content-Type: application/json
Ответ:
Успех (204 No Content): Задача удалена.
</code></pre>

Ошибка (404 Not Found): Задача не найдена.<br>
Ошибка (500 Internal Server Error): Проблема на сервере.
