## How to use
- Install docker and docker-compose
- Clone this repository:
```
$ git clone https://github.com/kevinabdul/cake-project
```
- Move to cake-project directory
```
$ cd cake-project
```

- run make mysql-service. This could take a while depending on your computer's capacity
```
$ make mysql-service
```

- run make cake-service
```
$ make cake-service
```

- run the database migration
```
$ make migrate-up
```

- if you want to rollback every migration to initial state, run this command:
```
$ make migrate-reset
```