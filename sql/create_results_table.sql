CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    request_id INTEGER NOT NULL,          -- уникальный идентификатор запроса
    code TEXT NOT NULL,                   -- исполняемый код
    language TEXT NOT NUll,               -- язык программирования
    image TEXT NOT NULL,                  -- использованный образ
    output TEXT NOT NULL,                 -- результат выполнения кода
    created_at TIMESTAMPTZ DEFAULT now()
);