curl -i -X POST  \
    -H 'Content-Type: application/json' \
    -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Vl16d9RIxtWDeGXgh3cdK-KRvesGhjr96qcYqDncj8k' \
    http://localhost:1323/api/users/signout

curl -i -X POST -H 'Content-Type: application/json' \
    -d '{
            "login": "test_login",
            "password": "test_password"
        }' \
     http://localhost:1323/api/users/signup


curl -i -X POST -H 'Content-Type: application/json' \
    -d '{
            "login": "test_login",
            "password": "test_password"
        }' \
     http://localhost:1323/api/users/signin



curl -i -X POST  \
    -H 'Content-Type: application/json' \
    -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Vl16d9RIxtWDeGXgh3cdK-KRvesGhjr96qcYqDncj8k' \
     http://localhost:1323/api/users/signout

