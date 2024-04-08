An extremely basic dummy login page and backend authentication endpoint for use in 
development OAuth code flows.

The login page will be located at `http://localhost:9000` \
The authentication endpoint will be located at `http://localhost:8080`

To use:
1. Update the redirect url in src/Update.elm (`http://localhost:5173/?code=123-123-123`) to point at whatever your frontend dev URL is
2. Build the frontend UI
```bash
cd frontend
npm run build
```
3. Build the backend auth endpoint 
4. Start up the docker container hosting the UI + auth endpoint
```bash
docker-compose up -d
```
