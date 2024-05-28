# Тестовое задания компании Ozon

Сервер постов и комментариев с использованием GraphQL
Позволяет создавать посты и комментарии

## Примечания
Этот сервер предназначен для тествого задания компании Ozon.
Он не сохраняет данные в бд!

### Запуск сервера

1. **Запуск Докера**:

`make start`

Вывод:

`Server is running on port 8080`

2. **Проверить работу сервера**:

Откройте Postman:

- EndPoint: `http://localhost:8080/graphql`
- Method:   `POST`
- Headers:  `Content-Type: application/json`
- X-Access-Key `admin`

### Примеры создания постов и комментов

#### Создание поста

- **Запрос**:

```json
{
 "query": "mutation { createPost(title: \"Название поста\", content: \"Содержание поста\") { id title content } }"
}
```

- **Ответ**:

```json
{
  "data": {
    "createPost": {
      "id": 1,
      "title": "Название поста",
      "content": "Содержание поста"
    }
  }
}
```
#### Создание комментария
- **Запрос**:

```json
{
  "query": "mutation { createComment(postId: 1, content: \"Комментарий к посту 1\") { id postId content } }"
}
```

- **Ответ**:

```json
{
  "data": {
    "createComment": {
      "id": 1,
      "postId": 1,
      "content": "Комментарий к посту 1"
    }
  }
}
```
#### Получение всех постов с комментариями
- **Запрос**:

```json
{
  "query": "{ posts { id title content comments { id content } } }"
}
```

- **Ответ**:

```json
{
  "data": {
    "posts": [
      {
        "id": 1,
        "title": "Название поста",
        "content": "Содержание поста",
        "comments": [
          {
            "id": 1,
            "content": "Комментарий к посту 1"
          }
        ]
      }
    ]
  }
}
```
