## Docker

### Running the image from DockerHub
There is a docker image published in [DockerHub](https://hub.docker.com/r/swaggerapi/swagger-editor/).

To use this, run the following:

```
docker pull swaggerapi/swagger-editor
docker run -d -p 80:8080 swaggerapi/swagger-editor
```

This will run Swagger Editor (in detached mode) on port 80 on your machine, so you can open it by navigating to `http://localhost` in your browser.

## Building and running an image locally

To build and run a docker image with the code checked out on your machine, run the following from the root directory of the project:

```
# Install npm packages (if needed)
npm install

# Build the app
npm run build

# Build an image
docker build -t swagger-editor .

# Run the container
docker run -d -p 80:8080 swagger-editor

```

You can then view the app by navigating to `http://localhost` in your browser.

## Test for Create App
1. Import related yaml into editor.

2. Click login, then click "try it out"

3. Click edit and type in the email and password.

Example:
{"email":"testuserankrfron5@mailinator.com", "password":"11111111nn"}

4. Click Execute and copy the object of "access_token"

5. Click Authorize on the top, right of the server.

6. type in "bearer " then past the access_token.
in the Value field it should be:
bearer ey....lelGN1KI

7. Click Authorize, then scroll down to /app/create and click "try it out"

8. Click edit and type in relative data. (Request here is for test, in formal case user will not be allowed to type in many field)

Example:
{"name": "wordpress_connectiontest", "type": "1", "chartname":"wordpress", "chartrepo":"stable", "chartversion": "5.7.1", "ns_name": "test_ns1", "ns_cpu_limit": "300", "ns_mem_limit": "500", "ns_storage_limit": "10", "ns_cluster_id": "f74536bb-9729-412a-b418-e45b4fc151fd", "ns_cluster_name": "mock_name"}

9. Click Execute and wait for its result.