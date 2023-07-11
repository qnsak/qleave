CREATE TABLE employee
(
    id VARCHAR(10) NOT NULL,
    name VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password text NOT NULL,
    department_id INTEGER NOT NULL,
    salary INTEGER NOT NULL,
    director_id character VARCHAR(10),
    entry_at TIMESTAMP,
    leaved_at TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE department
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(30) NOT NULL
);

CREATE TABLE paid_leave
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_id VARCHAR(10) NOT NULL,
    total INTEGER NOT NULL,
    use INTEGER NOT NULL,
    year CHARACTER(4) NOT NULL,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("employee_id") REFERENCES "employee" ("id")
);

CREATE TABLE leave_requisition
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type_id INTEGER NOT NULL,
    status INTEGER  DEFAULT 0 NOT NULL,
    reason TEXT NOT NULL,
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    dock_rate INTEGER,
    employee_id VARCHAR(10),
    director_id VARCHAR(10),
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("type_id") REFERENCES "leave_type" ("id"),
    FOREIGN KEY ("employee_id") REFERENCES "employee" ("id"),
    FOREIGN KEY ("director_id") REFERENCES "employee" ("id")
);

CREATE TABLE leave
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    requisition_id INTEGER NOT NULL,
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    FOREIGN KEY ("requisition_id") REFERENCES "leave_requisition" ("id")
);

CREATE TABLE leave_type
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(30) NOT NULL,
    status INTEGER NOT NULL,
    dock_rate INTEGER NOT NULL,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
