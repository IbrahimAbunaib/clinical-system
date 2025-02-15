CREATE TABLE admin (
    adminID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('active', 'inactive')) NOT NULL
);
