-- name: GetLeaveListByUserId :many
SELECT  id,
        type_id,
        status,
        employee_id,
        director_id,
        reason,
        start_at,
        end_at,
        dock_rate
FROM    leave_requisition 
WHERE   employee_id = ?
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: GetTotalLeaveListByEmployeeId :one
SELECT COUNT(lr.id) AS total
FROM leave_requisition AS lr
WHERE lr.employee_id = ?;

-- name: GetLeaveListByDirectorId :many
SELECT  lr.id,
        lr.type_id,
        lr.status,
        lr.employee_id,
        lr.director_id,
        lr.reason,
        lr.start_at,
        lr.end_at,
        employee.name,
        leave_type.name AS LeaveTitle
FROM    leave_requisition AS lr
JOIN employee ON lr.employee_id = employee.id
JOIN leave_type ON lr.type_id = leave_type.id
WHERE lr.director_id = ?
AND lr.status = 0
LIMIT ? OFFSET ?;


-- name: GetLeaveRequisitionById :one
SELECT  lr.start_at,
        lr.end_at
FROM leave_requisition AS lr
WHERE lr.id = ?;

-- name: CreateLeaveSuccesss :one
INSERT INTO leave (requisition_id, start_at, end_at)
VALUES (?, ?, ?)
RETURNING *;


-- name: GetTotalLeaveListByDirectorId :one
SELECT COUNT(lr.id) AS total
FROM leave_requisition AS lr
WHERE lr.director_id = ?
AND lr.status = 0;

-- name: CreateLeave :one
INSERT INTO leave_requisition (type_id, status, employee_id, director_id, reason, start_at, end_at, dock_rate)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateLeave :exec
UPDATE leave_requisition
SET status = ?
WHERE id = ?;

-- name: DeleteLeave :exec
UPDATE leave_requisition
SET status = 2
WHERE id = ?;

-- name: SreachLeaveListByStartAt :many
SELECT  lr.id,
        lr.status,
        lr.employee_id,
        lr.director_id,
        lr.reason,
        lr.start_at,
        lr.end_at,
        employee.name,
        leave_type.name
FROM    leave_requisition AS lr 
JOIN employee ON lr.employee_id = employee.id
JOIN leave_type ON lr.type_id = leave_type.id
WHERE lr.start_at BETWEEN ? AND ?;

-- name: GetAttendanceRecord :many
SELECT  employee.name,
        lr.reason
FROM    employee 
JOIN leave_requisition AS lr ON employee.id = lr.employee_id
WHERE employee.department_id = ?
AND lr.start_at > ?
AND lr.end_at < ?;

-- name: GetEmployeeListByDepartment :many
SELECT  id,
        name,
        email,
        password,
        department_id
FROM    employee 
WHERE   department_id = ?;

-- name: GetLeavesIsSuccess :many
SELECT  lr.reason,
        lr.start_at,
        lr.end_at,
        leave_type.name AS leaveTitle
FROM    leave_requisition AS lr
JOIN leave_type ON lr.type_id = leave_type.id
WHERE lr.employee_id = ?
AND lr.status = 1
AND lr.start_at > ?;

-- name: GetLeaveTypes :many
SELECT  id,
        name,
        dock_rate
FROM    leave_type
WHERE status = 1;

-- name: GetNationalHoliday :many
SELECT  reason,
        start_at,
        end_at
FROM    leave_requisition 
WHERE   type_id = 5;

-- name: GetLeaveByStartDate :many
SELECT  l.start_at,
        l.end_at,
        e.name,
        lr.id
FROM leave as l
JOIN leave_requisition as lr on lr.id = l.requisition_id
JOIN employee as e on e.id = lr.employee_id
WHERE lr.status = 1
AND e.department_id = ?
AND l.start_at >= ?
AND l.start_at <= ?;
