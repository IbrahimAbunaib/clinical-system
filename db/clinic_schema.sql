CREATE TABLE admin (
    admin_id SERIAL PRIMARY KEY,        -- Auto-incrementing unique admin ID
    full_name VARCHAR(150) NOT NULL,    -- Full name of the admin
    email VARCHAR(150) UNIQUE NOT NULL, -- Unique email
    password TEXT NOT NULL,             -- Hashed password
    role VARCHAR(50) NOT NULL,          -- Role (e.g., "Super Admin", "Manager")
    status VARCHAR(20) DEFAULT 'active',-- Status (e.g., "active", "inactive", "suspended")
    created_at TIMESTAMP DEFAULT NOW(), -- Auto timestamp when created
    updated_at TIMESTAMP DEFAULT NOW()  -- Auto timestamp when updated
);
