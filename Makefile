all-run:
	cd client/ && npm run dev &
	docker-compose up &
	cd firebase-front/ && firebase use yuc && firebase serve &
	wait