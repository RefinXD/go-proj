-- name: GetEmployee :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1;

-- name: GetEmployeeByName :one
SELECT * FROM employees
WHERE name = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY name;

-- name: CreateEmployee :one
INSERT INTO Employees (
  name, dob, department, job_title,address,join_date, created_date
) VALUES (
  $1, $2 , $3, $4 , $5 , $6 , NOW()
)
RETURNING *;

-- name: UpdateEmployee :one
UPDATE employees
  set name = coalesce($2,name),
  dob = coalesce($3,dob),
  department = coalesce($4,department),
  job_title = coalesce($5,job_title),
  address = coalesce($6,address),
  join_date = coalesce($7,join_date),
  updated_date = NOW()  
WHERE id = $1
RETURNING *;
-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1;