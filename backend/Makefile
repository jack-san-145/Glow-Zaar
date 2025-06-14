
DB_URL=mysql://root:root@tcp(127.0.0.1:3307)/GlowZaar

migrate_up:
	migrate -path ./database/migrations -database "$(DB_URL)" up
migrate_down:
	migrate -path ./database/migrations -database "$(DB_URL)" down