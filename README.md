# MSP Inventory

## Supported Tools
- Azure AD
- IT Glue
- Sophos
- Liongard
- CyberCNS
- Tactical RMM

## Description

Pick your source of truth then sync and alert on inventory changes.

## Stack

Go
- gofiber
- gorm
HTMX
TailwindCSS
PostgreSQL (sqlite for debug)

```
docker run --name mspinventory -e POSTGRES_USER=mspinventory -e POSTGRES_PASSWORD=mspinventory -d postgres
```
