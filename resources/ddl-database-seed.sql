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

INSERT INTO
    students (
        first_name,
        last_name,
        phone,
        email,
        birthday,
        gender,
        university,
        address
    )
VALUES
    (
        'Perry',
        'Platypus',
        '84987654320',
        'perry_platypus@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Tech University',
        '1 Elm Street'
    ),
    (
        'Phineas',
        'Flynn',
        '84987654321',
        'phineas_flynn@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Abc University',
        '2 Elm Street'
    ),
    (
        'Ferb',
        'Fletcher',
        '84987654322',
        'ferb_fletcher@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Xyz University',
        '3 Elm Street'
    ),
    (
        'Candace',
        'Flynn',
        '84987654323',
        'candace_flynn@gmail.com',
        '2000-01-01 07:00:00',
        0,
        'Tech University',
        '4 Elm Street'
    ),
    (
        'Isabella',
        'Garcia-Shapiro',
        '84987654324',
        'isabella_gracia_shapiro@gmail.com',
        '2000-01-01 07:00:00',
        0,
        'Abc University',
        '5 Elm Street'
    ),
    (
        'Buford',
        'Van Stomm',
        '84987654325',
        'buford_van_stomm@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Xyz University',
        '6 Elm Street'
    ),
    (
        'Baljeet',
        'Tjinder',
        '84987654326',
        'baljeet_tjinder@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Tech University',
        '7 Elm Street'
    ),
    (
        'Jeremy',
        'Johnson',
        '84987654327',
        'jeremy_johnson@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Abc University',
        '8 Elm Street'
    ),
    (
        'Stacy',
        'Hirano',
        '84987654328',
        'stacy_hirano@gmail.com',
        '2000-01-01 07:00:00',
        0,
        'Xyz University',
        '9 Elm Street'
    ),
    (
        'Vanessa',
        'Doofenshmirtz',
        '84987654329',
        'vanessa_doffenshmirtz@gmail.com',
        '2000-01-01 07:00:00',
        0,
        'Tech University',
        '10 Elm Street'
    ),
    (
        'Dr. Heinz',
        'Doofenshmirtz',
        '84987654330',
        'heinz_doffenshmirtz@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Abc University',
        '11 Elm Street'
    ),
    (
        'Major',
        'Monogram',
        '84987654331',
        'major_monogram@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Xyz University',
        '12 Elm Street'
    ),
    (
        'Carl',
        'Karl',
        '84987654332',
        'carl_karl@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Tech University',
        '13 Elm Street'
    ),
    (
        'Norm',
        'Norm',
        '84987654333',
        'norm_norm@gmail.com',
        '2000-01-01 07:00:00',
        1,
        'Abc University',
        '14 Elm Street'
    );