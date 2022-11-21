## How to use
- Install docker and make
- Clone this repository:
```
$ git clone https://github.com/kevinabdul/cake-project
```
- Move to cake-project directory
```
$ cd cake-project
```
- Run make mysql-service. This could take a while depending on your computer's capacity
```
$ make mysql-service
```
- Run make cake-service. Although its possible to make it wait forever, the cake service will only wait for database initialization for 10 minutes.
```
$ make cake-service
```
- If the 10 minutes deadline is excedeed, simply start the docker container again.
```
$ sudo docker start cake-service
```
- Run the database migration. This will create cakes table and seed it with two initial values.
```
$ make migrate-up
```


### Extras
- If you want to rollback every migration to initial state, run this command:
```
$ make migrate-reset
```