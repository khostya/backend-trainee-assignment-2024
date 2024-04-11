## Оформление решения
Необходимо предоставить публичный git-репозиторий на любом публичном хосте (GitHub / GitLab / etc), содержащий в master/main ветке:
1. Код сервиса
2. Makefile c командами сборки проекта / Описанная в README.md инструкция по запуску
3. Описанные в README.md вопросы/проблемы, с которыми столкнулись, и ваша логика их решений (если требуется)

## Запуск
```
    docker compose up postgres server
```

## CURL
```
curl -X 'GET' \
  'http://localhost:8080/user_banner?tag_id=0&feature_id=11&use_last_revision=true' \
  -H 'accept: application/json' \
  -H 'token: admin'
```

```
curl -X 'GET' \
  'http://localhost:8080/banner?feature_id=11&tag_id=0&limit=10&offset=0' \
  -H 'accept: application/json' \
  -H 'token: user'
```


```
curl -X 'PATCH' \
  'http://localhost:8080/banner/1' \
  -H 'accept: */*' \
  -H 'token: admin' \
  -H 'Content-Type: application/json' \
  -d '{
  "tag_ids": [
    0, 1, 2, 3
  ],
  "feature_id": 11,
  "content": {
    "ttle": "sometle",
    "text": "somet_ext",
    "url": "someurl"
  },
  "is_active": true
}'
```

```
curl -X 'POST' \
  'http://localhost:8080/banner' \
  -H 'accept: application/json' \
  -H 'token: admin' \
  -H 'Content-Type: application/json' \
  -d '{
  "tag_ids": [
    0, 1, 2, 3
  ],
  "feature_id": 11,
  "content": {
    "title": "some_title",
    "text": "some_text",
    "url": "some_url"
  },
  "is_active": true
}'
```

```
curl -X 'DELETE' \
  'http://localhost:8080/banner/1' \
  -H 'accept: */*' \
  -H 'token: admin'
```
