-- name: CreateUser :one
INSERT INTO users (name, email, permission_type, inactive_status, password)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListAllUsers :many
-- offset: $1
-- limit: $2
SELECT * FROM users ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateUserByID :exec
UPDATE users SET name = $2, email = $3, permission_type = $4, inactive_status = $5, password = $6
WHERE id = $1;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

-- name: CreateOrderStatus :one
INSERT INTO order_status (name) VALUES ($1) RETURNING *;

-- name: GetOrderStatusByID :one
SELECT * FROM order_status WHERE id = $1;

-- name: ListAllOrderStatus :many
-- offset: $1
-- limit: $2
SELECT * FROM order_status ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateOrderStatusByID :exec
UPDATE order_status SET name = $2 WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (client_id, equipment_id, order_number, reported_issue, diagnosis, solution, estimated_time, budget, status_id, assigned_to, priority)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: ListAllOrders :many
-- offset: $1
-- limit: $2
SELECT * FROM orders ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateOrderByID :exec
UPDATE orders SET client_id = $2, equipment_id = $3, order_number = $4, reported_issue = $5, diagnosis = $6, solution = $7, estimated_time = $8, budget = $9, status_id = $10, assigned_to = $11, priority = $12
WHERE id = $1;

-- name: DeleteOrderByID :exec
DELETE FROM orders WHERE id = $1;

-- name: CreateActivity :one
INSERT INTO activity (order_id, user_id, action, details) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetActivityByID :one
SELECT * FROM activity WHERE id = $1;

-- name: ListActivityByOrderID :many
-- offset: $1
-- limit: $2
SELECT * FROM activity WHERE order_id = $3 ORDER BY id OFFSET $1 LIMIT $2;

-- name: CreateClient :one
INSERT INTO clients (name, address, district, city, country, phone, ruc, contact_person, email, website, address_line_2, postal_code, fax, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;

-- name: GetClientByID :one
SELECT * FROM clients WHERE id = $1;

-- name: ListAllClients :many
-- offset: $1
-- limit: $2
SELECT * FROM clients ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateClientByID :exec
UPDATE clients SET name = $2, address = $3, district = $4, city = $5, country = $6, phone = $7, ruc = $8, contact_person = $9, email = $10, website = $11, address_line_2 = $12, postal_code = $13, fax = $14, notes = $15 WHERE id = $1;

-- name: DeleteClientByID :exec
DELETE FROM clients WHERE id = $1;

-- name: CreateComment :one
INSERT INTO comments (order_id, user_id, comment) VALUES ($1, $2, $3) RETURNING *;

-- name: GetCommentByID :one
SELECT * FROM comments WHERE id = $1;

-- name: ListCommentsByOrderID :many
-- offset: $1
-- limit: $2
SELECT * FROM comments WHERE order_id = $3 ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateCommentByID :exec
UPDATE comments SET comment = $2 WHERE id = $1;

-- name: DeleteCommentByID :exec
DELETE FROM comments WHERE id = $1;

-- name: CreateEquipmentType :one
INSERT INTO equipment_types (name) VALUES ($1) RETURNING *;

-- name: GetEquipmentTypeByID :one
SELECT * FROM equipment_types WHERE id = $1;

-- name: ListAllEquipmentTypes :many
-- offset: $1
-- limit: $2
SELECT * FROM equipment_types ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateEquipmentTypeByID :exec
UPDATE equipment_types SET name = $2 WHERE id = $1;

-- name: CreateEquipment :one
INSERT INTO equipments (type_id, name, serial_number, notes) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetEquipmentByID :one
SELECT * FROM equipments WHERE id = $1;

-- name: ListAllEquipments :many
-- offset: $1
-- limit: $2
SELECT * FROM equipments ORDER BY id OFFSET $1 LIMIT $2;

-- name: UpdateEquipmentByID :exec
UPDATE equipments SET type_id = $2, name = $3, serial_number = $4, notes = $5 WHERE id = $1;

-- name: DeleteEquipmentByID :exec
DELETE FROM equipments WHERE id = $1;
