# File Upload to very large JSON files

# Containers commands
   #### Please, for more details, take a look at the Makefile.

   ## To Build containers
      - make build-containers

   ## To Start containers with a daemon
      - make daemon-start-containers

   ## To Start containers with debug mode
      - make start-containers

   ## To Stop containers
      - make stop-containers

# Run tests
   #### Please, for more details, take a look at the Makefile.
	   - make run-tests

# Port Domain Application ( GRPC  )
   #### Please, for more details, take a look at the Makefile.
   ## To Generate Proto File
      - make generate-port-proto

   ## To Generate Mock Port Repository
      - make generate-mock-portDomain-repository

# Client API Application ( Rest )
   ## Default URL to upload JSON file (method POST)
      - localhost:3000/port/fileupload

      Attach the JSON File to body of the request
      Ex: (on the postman application, use the binary mode to do it)

   ## Default URL to find by key (method GET)
      - localhost:3000/port/{key}


# To change environment variables on the services
   #### Please, for more details, take a look at the Makefile.

   ## API Server
      - API_PORT --> used to change the default port (default port: 3000)
      - SERVER_GRPC_URL --> used to change address of the GRPC Server (Ex: localhost:9090)

   ## GRPC Server
      - GRPC-PORT --> used to change default port of the GRPC Server (default port: 9090)
      - MONGO_PORT --> used to change port to connect on mongo server (default port: 27017)
      - MONGO_HOST --> used to change address to connect on mongo server (Ex: localhost)


