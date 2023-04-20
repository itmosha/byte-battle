CREATE TABLE profiles (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    CONSTRAINT fk_user FOREIGN KEY(id) REFERENCES users(id) ON DELETE CASCADE,
    
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,

    status VARCHAR(255) NOT NULL UNIQUE,
    image_url VARCHAR(255) NOT NULL UNIQUE,
    tasks_completed INTEGER DEFAULT 0,

    active_courses INTEGER ARRAY,
    completed_courses INTEGER ARRAY
);