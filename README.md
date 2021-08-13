# CRON - GO

Para subir o projeto é nescessário ter o docker e [docker-compose](https://docs.docker.com/compose/install/).

### RUN
Execute:
```bash
docker-compose up --build
```
### FUNCIONAMENTO

O serviço de cron irá executar automáticamente carregando os bancos com os dados do subredit artificial.

Para utilizar os endpoints:

Exemplo para criação de post:

```javascript
curl --request POST \
  --url http://localhost:8080/posts \
  --header 'Content-Type: application/json' \
  --data '{
	"author":"nome do autor",
	"title":"titulo",
	"ups":250,
	"num_comments":100,
	"created_date": 1628721244 // unix 
}'
```
Buscar posts, 
```javascript
curl --request GET \
  --url 'http://localhost:8080/posts?orderby=ups&startdate=2021-01-01&enddate=2021-02-01'
```

Buscar User posts,
```javascript
curl --request GET \
  --url 'http://localhost:8080/user-posts?orderby=comments'
```
