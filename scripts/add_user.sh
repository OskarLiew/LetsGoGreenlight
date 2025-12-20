BODY='{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
curl -i -d "$BODY" localhost:4000/v1/users
