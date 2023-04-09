curl -i -X POST -H 'Content-Type: application/json' \
    -d '{
            "login": "test_login",
            "password": "test_password"
        }' \
     http://localhost:1323/signin