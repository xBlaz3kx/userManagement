# Simple user and group management API

This is a demo User and Group management API, used to demonstrate my programming skills for the company **3fs**.

Known "issues" and possible improvements:

1. When adding a user to the group, transactions should be used, but omitted due to the simplicity (it would require a
   replica set with MongoDB, which needs more configuration to auto-deploy using docker-compose)
2. Using environment variables instead of configuration files - just my preference to use configuration files because
   they're easier to modify
3. Validation layer: simplified due to lack of concrete requirements.
4. Access management - no requirements specified

## Tech stack

| Component | Technology used |
| :----: | :----: |
| Language | Go |
| REST API router/server | Gin framework & endless|
| Database | MongoDB |
| Container Engine | Docker |
| Configuration file library | [Fig](https://github.com/kkyr/fig) |
| Mongo ODM driver | [Mgm](https://github.com/Kamva/mgm) |
| OpenAPI | V2.0 |

## Running the demo

To run the demo API, simply run:

```bash
docker-compose up -d 
```

