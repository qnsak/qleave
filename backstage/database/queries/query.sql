-- name: GetUserByEmail :one
SELECT  e.id,
        e.name,
        e.email,
        e.password,
        e.department_id,
        e.director_id,
        d.name as department_title
FROM    employee as e
JOIN    department as d ON e.department_id = d.id
WHERE   email = ?
LIMIT 1;

-- name: ListUserBydDepartmentId :many
SELECT * FROM employee
WHERE department_id = ?
ORDER BY id;

-- name: CreateUser :one
INSERT INTO employee (name, email, password, department_id, salary)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateName :exec
UPDATE employee
SET name = ?
WHERE email = ?;

-- name: DeleteUser :exec
UPDATE employee
SET status = 0
WHERE email = ?;

-- name: GetUserInfo :one
SELECT  e.name,
        e.email,
        d.name as department_title,
        pl.total as paid_leave_total,
        pl.use as paid_leave_use
FROM    employee as e
JOIN    department as d ON e.department_id = d.id
JOIN    paid_leave as pl ON e.id = pl.employee_id
WHERE   e.id = ?
AND     pl.year = ?
LIMIT 1;