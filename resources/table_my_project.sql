CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS experience_type (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    skill_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE IF NOT EXISTS experiences (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL, -- FK ke tabel USERS (Diperlukan berdasarkan ERD, relasi 1:N)
    experience_name VARCHAR(255) NOT NULL,
    experience_image VARCHAR(255),
    start_date DATE NOT NULL,
    end_date DATE,
    experience_type_id INTEGER NOT NULL, -- FK ke experience_type (nama kolom diubah)
    role_id INTEGER, -- FK ke roles
    description TEXT,
    source VARCHAR(255),
    media VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,

    -- Definisi Foreign Key
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users (id),
    CONSTRAINT fk_experience_type
        FOREIGN KEY (experience_type_id)
        REFERENCES experience_type (id),
    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES roles (id)
);
CREATE TABLE IF NOT EXISTS skillsets (
    id SERIAL PRIMARY KEY,
    skill_id INTEGER NOT NULL, -- FK ke skills (Diubah dari id_skill)
    experience_id INTEGER NOT NULL, -- FK ke experiences (Diubah dari id_experience)
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,

    -- Mencegah entri pengalaman-skill yang sama ganda
    UNIQUE (skill_id, experience_id), 

    -- Definisi Foreign Key
    CONSTRAINT fk_skill
        FOREIGN KEY (skill_id)
        REFERENCES skills (id),
    CONSTRAINT fk_experience
        FOREIGN KEY (experience_id)
        REFERENCES experiences (id)
);
CREATE TABLE IF NOT EXISTS detail_experiences (
    id SERIAL PRIMARY KEY,
    experience_id INTEGER NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT fk_experience_detail
        FOREIGN KEY(experience_id)
        REFERENCES experiences(id)
);