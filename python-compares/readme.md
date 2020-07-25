# Notes:
- simple apps that are creating 100 users in memory and returning the user by ID
- This was meant to do a simple speed comparison w/ Locust


## docker notes:
- docker build 
    - `docker build -t flask-app .`
- create initial run of docker container from build image `flask-app` naming container as `flask`:
    - `docker run --name flask -p 5000:5000 -d flask-app`
        - `it` instead of `-d` will run in foreground
            -IE: `docker run -p 5000:5000 -it flask-app` 
                - ctrl+c will stop it and you wont need `docker stop` to stop the app

    - stop running:
        - `docker stop flask`
    - start again:
        - `docker start flask`
    - check running containers:
        - `docker container ps`
    - remove app:
        - `docker rm flask` 
        - if the image is rebuilt and you want to re-run the app in the background again with the same name


    - Docker Mimic GCP `n1-standard-1` GKE env
        - 1 vCPU/ 3.75 GBs RAM
        - `docker run -p 127.0.0.1:5000:5000 -it --memory="3.75g" --cpus="1" flask-app`
    
