run:
	docker build -t tennis-client:latest -f Dockerfile.dev .
	docker run --network tennis-network -p 3000:3000 tennis-client:latest