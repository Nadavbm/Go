run-pg-compose:
	docker-compose -f docker-compose.yml down && \
	docker-compose -f docker-compose.yml up --build -d && \
	sleep 5 && \
	echo "Postgres is running!"

run-pg-container:
	docker volume rm data -f && \
	docker run --rm \
	-p 5432:5432 \
	-v data:/var/lib/postgresql/data \
	--env POSTGRES_USER=postgres \
	--env POSTGRES_PASSWORD=pass1234 \
	--env POSTGRES_DB=address \
	postgres:11