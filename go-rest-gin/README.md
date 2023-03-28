## POSTMAN
```
{
	"info": {
		"_postman_id": "8ffef545-4093-4546-88c0-7c45c416d0ac",
		"name": "hello-golang",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "26065462"
	},
	"item": [
		{
			"name": "products",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/products",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"products"
					]
				}
			},
			"response": []
		},
		{
			"name": "product/id",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/product/3",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"product",
						"3"
					]
				}
			},
			"response": []
		},
		{
			"name": "product",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"nama_produk\": \"tas A\",\n    \"deskripsi\": \"tas cewe\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/api/product",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"product"
					]
				}
			},
			"response": []
		},
		{
			"name": "product Copy",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"nama_produk\": \"tas A\",\n    \"deskripsi\": \"tas cewe\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/api/product/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"product",
						"1"
					]
				}
			},
			"response": []
		},
		{
			"name": "product/id",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/product/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"product",
						"1"
					]
				}
			},
			"response": []
		}
	]
}
```
