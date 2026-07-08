

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tweets (
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content VARCHAR(280) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


-- Добавляем пользователей
INSERT INTO users (username, password) VALUES 
('pavel', 'hashed_cheburek123'),
('vanya', 'secret_password');

-- Добавляем твиты (допустим, у Павла id = 1, а у Вани id = 2)
INSERT INTO tweets (user_id, content) VALUES 
(1, 'Мой первый твит в Blipw! Всем привет.'),
(1, 'Разрабатываю бэкенд на Go, язык просто пушка!'),
(2, 'А я тестирую базу данных PostgreSQL в Докере.');