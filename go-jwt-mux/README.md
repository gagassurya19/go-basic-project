## postman
```
{
	"info": {
		"_postman_id": "56b86457-7c4b-43df-aff5-7483bd28c557",
		"name": "go-jwt-mux",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "26065462"
	},
	"item": [
		{
			"name": "register",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"nama_lengkap\": \"Gagas Surya Laksana\",\n    \"username\":\"qweqdfsca\",\n    \"password\": \"testtasdasdest\"\n}"
				},
				"url": {
					"raw": "localhost:8080/register",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"register"
					]
				}
			},
			"response": []
		},
		{
			"name": "login",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "localhost:8080/login",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "logout",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/logout",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"logout"
					]
				}
			},
			"response": []
		},
		{
			"name": "product",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/logout",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"logout"
					]
				}
			},
			"response": []
		}
	]
}
```