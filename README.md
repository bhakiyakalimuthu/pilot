Pilot project

****Create profile****
`curl -X POST "http://localhost:8080/profile/create" -d '{"name":"bhaki","phoneNumber":"0708753884","address":"test","age":35,"dob":"06061986","sex":"male","height":176,"weight":68}'
`

****Get profile****

`curl -X GET "http://localhost:8080/profile/1a740acc-8970-4634-82f6-df82751736d8"`

****Create medical data****

`curl -X POST "http://localhost:8080/profile/med-data/create" -d '{"patientId":"jk8Sm327PKAdZjv3evmyy6","demographicFactor":{"year":1986,"sex":"male"}}'`


****Docker run****

`docker run --rm   --name pg-docker-testing -e POSTGRES_DB=users -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=dbpassword -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  -d -t kartoza/postgis`