build:
	docker compose build kafka-api kafka-consumer playwright-test
run:
	docker compose up -d
push:
	docker tag kafka-api:1.0.0 ghcr.io/sittijedw/kafka-api:1.0.0
	docker push ghcr.io/sittijedw/kafka-api:1.0.0
	docker tag kafka-consumer:1.0.0 ghcr.io/sittijedw/kafka-consumer:1.0.0
	docker push ghcr.io/sittijedw/kafka-consumer:1.0.0
	docker tag playwright-test:1.0.0 ghcr.io/sittijedw/playwright-test:1.0.0
	docker push ghcr.io/sittijedw/playwright-test:1.0.0