CREATE TABLE
    IF NOT EXISTS students (
        id int check (id > 0) NOT NULL GENERATED ALWAYS AS IDENTITY,
        first_name text NOT NULL,
        last_name text NOT NULL,
        phone varchar(191) NOT NULL,
        email varchar(191) NOT NULL,
        birthday date DEFAULT NULL,
        gender smallint NOT NULL DEFAULT '0',
        university text,
        address text,
        created_at timestamp(3) DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp(3) DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        CONSTRAINT uni_students_email UNIQUE (email),
        CONSTRAINT uni_students_phone UNIQUE (phone)
    );