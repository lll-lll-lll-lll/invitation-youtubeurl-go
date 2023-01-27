# ⚠︎ already db table created
all-run:
	cd client/ && npm run dev &
	docker-compose up &
	wait