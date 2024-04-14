## Docker
Для запуска нужно сбилдить контейнер с golang-кодом, затем сбилдить на основаниее avito-app ещё контейнер, а затем поднять всё вместе с базой данных.
docker-compose down в данном контексте нужна, потому что при первом запуске база данных поднимается позже go-кода и к ней не удаётся подключится, поэтому требуется разовый перезапуск после билда.
```
docker build -t avito-app

docker-compose up --build avito-app

docker-compose down

docker-compose up
```
Эндпоинты:
<table>
    <thead>
        <tr>
            <th>Method</th>
            <th>Route</th>
            <th>Body</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="center">GET</td>
            <td align="center">/user_banner</td>
            <td align="center"></td>
        </tr>
        <tr>
            <td align="center">POST</td>
            <td align="center">/banner</td>
            <td align="center"></td>
        </tr>
       <tr>
            <td align="center">POST</td>
            <td align="center">/banner</td>
            <td align="center">`{"tags_ids": [1, 3, 4],"feature_id": 2,
                             "content":  {"title": "some_title", "text":
                             "some_text", "url":"some_url"},"is_active": true}`</td>
        </tr>
    </tbody>
</table>