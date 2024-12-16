-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) DEFAULT NULL,
    last_name VARCHAR(255) DEFAULT NULL,
    is_premium INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    referal_code VARCHAR(255) UNIQUE
);


CREATE TABLE IF NOT EXISTS portfolios (
    id BIGSERIAL PRIMARY KEY,
    owner_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    profit DECIMAL(10, 2) DEFAULT 0.0,

    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Таблица криптовалют
CREATE TABLE cryptocurrencies (
     id SERIAL PRIMARY KEY,
     symbol VARCHAR(10) UNIQUE NOT NULL, -- Символ (например, BTC, ETH)
     name VARCHAR(255) NOT NULL         -- Полное название криптовалюты
);

-- Таблица активов портфелей
CREATE TABLE portfolio_assets (
    id SERIAL PRIMARY KEY,
    portfolio_id INT NOT NULL ,
    cryptocurrency_id INT NOT NULL,
    balance NUMERIC(20, 8) NOT NULL DEFAULT 0,  -- Текущий баланс криптовалюты
    average_price NUMERIC(20, 8) DEFAULT 0,     -- Средняя цена покупки
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (portfolio_id, cryptocurrency_id), -- Уникальная пара портфель-криптовалюта
    FOREIGN KEY (portfolio_id) REFERENCES portfolios (id) ON DELETE CASCADE,
    FOREIGN KEY (cryptocurrency_id) REFERENCES cryptocurrencies (id) ON DELETE CASCADE
);


CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    portfolio_id INT NOT NULL,
    transaction_type VARCHAR(50) CHECK (transaction_type IN ('buy', 'sell', 'transfer')),  -- Тип сделки
    crypto_id INT NOT NULL,
    amount NUMERIC(20, 8) NOT NULL,  -- Количество криптовалюты
    buy_price NUMERIC(20, 8) NOT NULL,  -- Последняя цена покупки
    price NUMERIC(20, 8) NOT NULL,  -- Цена на момент сделки
    total NUMERIC(20, 8) NOT NULL,  -- Сумма сделки
    timestamp TIMESTAMP DEFAULT current_timestamp,
    profit DECIMAL(10, 2) DEFAULT 0.0,
    profit_usd DECIMAL(10, 2) DEFAULT 0.0,

    FOREIGN KEY (portfolio_id) REFERENCES portfolios (id) ON DELETE CASCADE,
    FOREIGN KEY (crypto_id) REFERENCES cryptocurrencies (id)
);

CREATE TABLE points_transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    amount INT NOT NULL,          -- Сколько начислено/списано
    type VARCHAR(50) NOT NULL,   -- Тип транзакции (например, "tournament", "referral", "mining")
    created_at TIMESTAMP DEFAULT current_timestamp
);


CREATE TABLE referrals (
     id SERIAL PRIMARY KEY,
     referrer_id INT REFERENCES users(id) ON DELETE CASCADE, -- Пригласивший пользователь
     referred_id INT REFERENCES users(id) ON DELETE CASCADE, -- Приглашённый пользователь
     created_at TIMESTAMP DEFAULT current_timestamp          -- Дата создания связи
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS portfolios;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cryptocurrencies;
DROP TABLE IF EXISTS portfolio_assets;
DROP TABLE IF EXISTS points_transactions;
DROP TABLE IF EXISTS referrals;

