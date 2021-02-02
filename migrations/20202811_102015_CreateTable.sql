CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100),
)

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    app_name VARCHAR(100) NOT NULL ,
    app_email VARCHAR(100) NOT NULL UNIQUE,
    app_password VARCHAR(100),

    FOREIGN KEY(user_id) REFERENCES users(id)
)
