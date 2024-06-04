# Exoplanet Service

This is a microservice for managing exoplanets. It provides functionality to add, list, retrieve, update, and delete exoplanets. It also provides fuel estimation for a trip to a particular exoplanet.

## Endpoints

- `POST /exoplanets`: Add a new exoplanet
- `GET /exoplanets`: List all exoplanets
- `GET /exoplanets/{id}`: Get an exoplanet by ID
- `PUT /exoplanets/{id}`: Update an exoplanet
- `DELETE /exoplanets/{id}`: Delete an exoplanet
- `GET /exoplanets/{id}/fuel?crew_capacity={capacity}`: Get fuel estimation

## Running the Service

### Locally

```sh
go run main.go
```

### Using Docker

```
docker build -t exoplanet-service .
docker run -p 8080:8080 exoplanet-service
```

## Example Requests

### Add Exoplanet

```
curl -X POST http://localhost:8080/exoplanets -d '{
    "name": "Planet X",
    "description": "A mysterious planet",
    "distance": 50,
    "radius": 1.2,
    "mass": 2.5,
    "type": "Terrestrial"
}' -H "Content-Type: application/json"
```

### List Exoplanets

```
curl http://localhost:8080/exoplanets
```

### Get Exoplanet by ID

```
curl http://localhost:8080/exoplanets/{id}
```
Replace `{id}` with the actual ID of the exoplanet.

### Update Exoplanet

```
curl -X PUT http://localhost:8080/exoplanets/{id} -d '{
    "name": "Updated Planet",
    "description": "An updated description",
    "distance": 100,
    "radius": 1.5,
    "mass": 3.0,
    "type": "Terrestrial"
}' -H "Content-Type: application/json"
```
Replace `{id}` with the actual ID of the exoplanet.

### Delete Exoplanet
```
curl -X DELETE http://localhost:8080/exoplanets/{id}
```
Replace `{id}` with the actual ID of the exoplanet.
