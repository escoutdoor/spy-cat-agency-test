# Spy Cat Agency (SCA) — Test Assignment

This is a test assignment: a small REST CRUD service to manage spy cats, their missions, and mission targets.

## How to run
1) Copy environment variables:
```bash
cp .env.example .env
```
- .env.example lists all variables required in .env. Copy and adjust as needed.
- CAT_CLIENT_API_KEY: obtain a key at https://api.thecatapi.com and put it into your .env.

2) Start the app and database with Docker:
```bash
docker compose up --build
```
The service listens at the host and port configured in `.env` (`HTTP_SERVER_HOST` and `HTTP_SERVER_PORT`). For example: http://localhost:4040.

## API Routes

### Cats
- POST /v1/cats
  - Description: Create a cat.
  - Example body:
    ```json
    {
      "name": "Murchik",
      "yearsOfExperience": 1,
      "breed": "American Shorthair",
      "salary": 1800
    }
    ```

- GET /v1/cats/f01c4ac0-40d8-4202-9b14-0ec9bb7b4780
  - Description: Get a cat by ID.

- GET /v1/cats?limit=10&offset=0
  - Description: List cats (pagination).

- PATCH /v1/cats/f01c4ac0-40d8-4202-9b14-0ec9bb7b4780
  - Description: Partially update a cat.
  - Example body:
    ```json
    { "salary": 1900 }
    ```

- DELETE /v1/cats/37d7834a-4919-471f-95d1-1b96d336fab4
  - Description: Delete a cat.

### Missions
- POST /v1/missions
  - Description: Create a mission with 1–3 targets.
  - Example body:
    ```json
    {
      "targets": [
        { "name": "Target #1", "country": "Portugal" },
        { "name": "Target #2", "country": "Spain" }
      ]
    }
    ```

- GET /v1/missions/4524deec-aaa9-420e-aa78-0007b194c189
  - Description: Get a mission by ID.

- GET /v1/missions?limit=1&offset=0
  - Description: List missions (pagination).

- PATCH /v1/missions/4524deec-aaa9-420e-aa78-0007b194c189
  - Description: Assign a cat to a mission.
  - Example body:
    ```json
    { "catId": "37d7834a-4919-471f-95d1-1b96d336fab3" }
    ```

- DELETE /v1/missions/c62e33f2-6e54-4534-850b-2e535e819f24
  - Description: Delete a mission.

### Targets
- POST /v1/targets/4524deec-aaa9-420e-aa78-0007b194c189
  - Description: Add targets to a mission (up to 3 per mission).
  - Example body:
    ```json
    {
      "targets": [
        { "name": "Target #3", "country": "Poland" }
      ]
    }
    ```

- PATCH /v1/targets/a489efa4-2109-48da-bb50-d86a01917156
  - Description: Update a target (e.g., mark as completed).
  - Example body:
    ```json
    { "completed": true }
    ```

- DELETE /v1/targets/1413dcd9-cafe-41fd-aa78-ee2356a07d96
  - Description: Delete a target.

## Postman
- Import `spy-cat-agency.postman_collection.json`. The collection uses `ADDRESS = http://localhost:4040` by default.
