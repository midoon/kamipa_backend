CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,     
    student_nisn VARCHAR(50) NOT NULL UNIQUE,         
    email VARCHAR(100) NOT NULL UNIQUE,            
    password VARCHAR(255) NOT NULL,                               
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,            
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP       
); 