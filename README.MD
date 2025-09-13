# Spy Cat Agency (SCA) — Test Assignment

This is a test assignment: a small REST CRUD service to manage spy cats, their missions, and mission targets.

## How to run
1) Copy environment variables:
```bash
cp .env.example .env
```
- `.env.example` lists all variables required in `.env`. Copy and adjust as needed.
- `CAT_CLIENT_API_KEY`: obtain a key at https://thecatapi.com/ and put it into your `.env`.

2) Start the app and database with Docker:
```bash
docker compose up --build
```
The service listens at the host and port configured in `.env` (`HTTP_SERVER_HOST` and `HTTP_SERVER_PORT`). For example: http://localhost:4040.

## API Routes
Note: List endpoints support pagination with optional query parameters: `limit` (default: 20) and `offset` (default: 0).

### Cats
- POST /v1/cats
  - Description: Create a cat.
  - Example body:
    ```json
    {
      "name": "Motya",
      "yearsOfExperience": 3,
      "breed": "American Shorthair",
      "salary": 1800
    }
    ```

- GET /v1/cats/{catId}
  - Description: Get a cat by ID.

- GET /v1/cats?limit=10&offset=0
  - Description: List cats (pagination).
  - Query params (optional): `limit`, `offset` (defaults: 20, 0).

- PATCH /v1/cats/{catId}
  - Description: Partially update a cat.
  - Example body:
    ```json
    { "salary": 1900 }
    ```

- DELETE /v1/cats/{catId}
  - Description: Delete a cat.

### Missions
- POST /v1/missions
  - Description: Create a mission with 1–3 targets.
  - Example body:
    ```json
    {
      "targets": [
        { "name": "Target #1", "country": "Greece" },
        { "name": "Target #2", "country": "China" }
      ]
    }
    ```

- GET /v1/missions/{missionId}
  - Description: Get a mission by ID.

- GET /v1/missions?limit=1&offset=0
  - Description: List missions (pagination).
  - Query params (optional): `limit`, `offset` (defaults: 20, 0).

- PATCH /v1/missions/{missionId}
  - Description: Assign a cat to a mission.
  - Example body:
    ```json
    { "catId": "<catId>" }
    ```

- DELETE /v1/missions/{missionId}
  - Description: Delete a mission.

- POST /v1/missions/{missionId}/targets
  - Description: Add targets to a mission (up to 3 per mission).
  - Example body:
    ```json
    {
      "targets": [
        { "name": "Target #3", "country": "Poland" }
      ]
    }
    ```

### Targets
- PATCH /v1/targets/{targetId}
  - Description: Update a target (e.g., mark as completed).
  - Example body:
    ```json
    { "completed": true, "notes": "" }
    ```

- DELETE /v1/targets/{targetId}
  - Description: Delete a target.

## Postman
- Import `spy-cat-agency.postman_collection.json`. The collection uses `ADDRESS = http://localhost:4040` by default.
