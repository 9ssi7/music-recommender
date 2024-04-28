dev:
	docker compose -f ./.docker/docker-compose.yml up --build -d

stop:
	docker compose -f ./.docker/docker-compose.yml down	

build:
	docker build -t 9ssi7/music-recommender -f ./.docker/Dockerfile .