CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NOT NULL,
    address VARCHAR(100),
    passport_serie INT NOT NULL,
    passport_number INT NOT NULL
);

CREATE INDEX idx_surname ON users(surname);
CREATE INDEX idx_name ON users(name);
CREATE INDEX idx_passport_serie ON users(passport_serie);
CREATE INDEX idx_passport_number ON users(passport_number);
CREATE INDEX idx_patronymic ON users(patronymic);
CREATE INDEX idx_address ON users(address);

INSERT INTO users (passport_number, passport_serie, surname, name, patronymic, address)
SELECT
    (FLOOR(RANDOM() * 900000) + 100000)::INT AS passport_number, 
    (FLOOR(RANDOM() * 9000) + 1000)::INT AS passport_serie, 
    CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) AS surname,
    CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) AS name,
    CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) || CHR((65 + FLOOR(RANDOM() * 26))::INT) AS patronymic,
    'г. ' || CHR((65 + FLOOR(RANDOM() * 26))::INT) || ' ул. ' || CHR((65 + FLOOR(RANDOM() * 26))::INT) || ' д. ' || (FLOOR(RANDOM() * 100) + 1)::INT AS address
FROM generate_series(1, 100);

INSERT INTO users (passport_serie, passport_number, surname, name, patronymic, address) VALUES (1234,567890, 'Иванов', 'Иван','Иванович','г. Москва, ул. Ленина, д. 5, кв. 1');