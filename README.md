# Drone Delivery Management Backend

This project is a simple **Drone Delivery Management Backend** written in **Go**, designed as a take-home / interview-style service.  
It follows a **clean MVC-like structure**, uses **no external dependencies**, and keeps the logic intentionally readable and explicit.

The system supports **end users**, **drones**, and **admins**, each authenticated via a lightweight self-signed JWT.

---

## Features

### Authentication
- JWT-based authentication
- Token issued via `/auth/token`
- Roles supported:
    - `admin`
    - `enduser`
    - `drone`
- Tokens are passed as `Authorization: Bearer <token>`


---
### End Users
End users can:
- Submit delivery orders (origin → destination)
- Withdraw orders that have not yet been picked up
- View their submitted orders and status

---

### Drones
Drones can:
- Reserve available delivery jobs
- Pick up orders
- Mark orders as delivered or failed
- Update their current GPS location (heartbeat)
- Mark themselves as broken
- Automatically create a handoff job when broken

---

### Admins
Admins can:
- View all orders in bulk
- Update order origin or destination
- View all registered drones
- Mark drones as broken or fixed

---

## How to Run

This section describes how to run the Drone Delivery Management Backend locally.



### Prerequisites

- Go **1.20** or newer installed  
  Verify with:
  `  go version
  `
- `git clone github.com/Omaroma/drone-backend`
- `cd drone-backend`
- `go mod tidy`
- `go run main.go`

also you can run tests by

`go test ./tests
`

---
## Authentication
All protected endpoints require a Bearer token:
Authorization: Bearer <token>
Tokens are issued by the authentication endpoint and contain:
- name
- role (admin, enduser, drone)
- expiration timestamp
---


POST /auth/token


**Request**
```json
{
  "name": "alice",
  "role": "enduser"
}
```
Response
```json
{
  "token": "base64payload.signature"
}
```
---
## End User APIs
Submit Order

Creates a new delivery order.

Endpoint

```bash
POST /orders
```

Auth role: enduser

**Request Body**

```json
{
  "origin": {
    "lat": 52.52,
    "lon": 13.40
  },
  "destination": {
    "lat": 52.50,
    "lon": 13.45
  }
}
```

**Response**

```json
{
  "id": "20240201123045",
  "owner": "alice",
  "origin": { "lat": 52.52, "lon": 13.40 },
  "destination": { "lat": 52.50, "lon": 13.45 },
  "status": "pending",
  "created_at": "2024-02-01T12:30:45Z"
}

```

**Withdraw Order**

Withdraws an order that has not yet been picked up.
```bash
DELETE /orders/withdraw?id=<order_id>
```
Auth role: enduser

Rules

* Only pending orders can be withdrawn
* Only the owner can withdraw their order

**Response**
```json
204
```

**Drone APIs**

**Reserve Job**

Reserves the first available pending order.
```bash
POST /drone/reserve
```
auth role: drone

**Response**
```json
{
  "id": "20240201123045",
  "status": "reserved",
  "assigned_to": "drone-1"
}
```

**Update Drone Status (Heartbeat)**

auth role: drone

```bash
POST /drone/status
```

**Request Body**
```json
{
  "location": {
    "lat": 52.51,
    "lon": 13.41
  },
  "broken": false
}
```
**Behavior**

1. Updates drone heartbeat timestamp
2. If marked as broken:
* Drone stops
* A new handoff job is created automatically

**Response**
```json
{
  "id": "drone-1",
  "location": { "lat": 52.51, "lon": 13.41 },
  "broken": false,
  "updated_at": "2024-02-01T12:40:00Z"
}
```
## Drone APIs

### Complete Order

Marks an assigned order as delivered or failed.

**Endpoint**
```bash
POST /drone/complete
```

**Auth Role** 

drone


**Request Body**
```json
{
  "order_id": "20240201123045",
  "status": "delivered"
}
```
Allowed Status Values
* delivered
* failed

**Response**
```json
{
  "id": "20240201123045",
  "status": "delivered",
  "assigned_to": "drone-1"
}
```

## Admin APIs
List All Orders

Returns all orders in the system, regardless of owner.

Endpoint
```bash
GET /admin/orders
```

auth role: admin

**Response**
```json
[
  {
    "id": "20240201123045",
    "owner": "alice",
    "status": "pending",
    "assigned_to": "",
    "origin": { "lat": 52.52, "lon": 13.40 },
    "destination": { "lat": 52.50, "lon": 13.45 }
  }
]
```

**Update Order Origin or Destination**

Allows an admin to update the origin and/or destination of an order.
```bash
PUT /admin/orders/update
```
**Request Body**
```json
{
  "id": "20240201123045",
  "origin": {
    "lat": 52.53,
    "lon": 13.42
  },
  "destination": {
    "lat": 52.51,
    "lon": 13.46
  }
}
```

**Notes**

* Either `origin`, `destination`, or both may be provided
* Fields not provided remain unchanged

**Response**
```json
{
  "id": "20240201123045",
  "origin": { "lat": 52.53, "lon": 13.42 },
  "destination": { "lat": 52.51, "lon": 13.46 }
}
```

**List Drones**

Returns all drones currently known to the system.
```bash 
GET /admin/drones
```

**Response**
```json
[
  {
    "id": "drone-1",
    "location": { "lat": 52.51, "lon": 13.41 },
    "broken": false,
    "updated_at": "2024-02-01T12:40:00Z"
  }
]
```

**Update Drone Status**

Marks a drone as broken or fixed.

**Request Body**
```json
{
  "id": "drone-1",
  "broken": true
}
```

Behavior

If a drone is marked as broken:

The drone stops operating

A handoff job is created for its current cargo

Marking a drone as fixed does not remove previously created handoff jobs

**Response**
```json
{
  "id": "drone-1",
  "broken": true
}
```

