#Executando

É necessário configurar três variáveis de ambientes: SERVER_PORT, GEOCODING_API_KEY e GEOCODING_API_URL para que
a aplicação funcione corretamente.

Na raiz do projeto já se encontra um binário funcional, necessitando somente das variáveis acima.

```bash
SERVER_PORT=8080 GEOCODING_API_KEY=SUA_API_KEY GEOCODING_API_URL=https://maps.googleapis.com/maps/api/geocode/json ./calindra 
```