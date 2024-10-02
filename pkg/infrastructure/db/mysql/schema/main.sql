-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(1000),
    permission_type VARCHAR(30) CHECK (permission_type IN ('admin', 'supervisor', 'employee')) NOT NULL,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    inactive_status VARCHAR(30) CHECK (inactive_status IN ('active', 'inactive')) DEFAULT 'active',
    UNIQUE (email)
);

-- Tabla de estado de órdenes
CREATE TABLE order_status (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Datos iniciales para el estado de órdenes
INSERT INTO order_status (name) VALUES ('Pending'), ('In progress'), ('Completed'), ('Cancelled');

-- Tabla de clientes
CREATE TABLE clients (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255),
    district VARCHAR(100),
    city VARCHAR(100),
    country VARCHAR(100),
    phone VARCHAR(15),
    ruc VARCHAR(11) UNIQUE,
    contact_person VARCHAR(100),
    email VARCHAR(100),
    website VARCHAR(255),
    address_line_2 VARCHAR(255),
    postal_code VARCHAR(20),
    fax VARCHAR(15),
    notes TEXT
);

-- Tabla de tipo de equipos
CREATE TABLE equipment_types (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Tabla de equipos
CREATE TABLE equipments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    type_id INT,
    name VARCHAR(100) NOT NULL,
    serial_number VARCHAR(100) NOT NULL,
    notes TEXT,
    FOREIGN KEY (type_id) REFERENCES equipment_types(id)
);

-- Tabla de órdenes
CREATE TABLE orders (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    client_id INT,
    equipment_id INT,
    order_number VARCHAR(100) NOT NULL,
    reported_issue TEXT,
    diagnosis TEXT,
    solution TEXT,
    estimated_time INTERVAL,
    budget DECIMAL(10, 2),
    status_id INT,
    assigned_to INT,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    priority VARCHAR(10) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high')),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (equipment_id) REFERENCES equipments(id),
    FOREIGN KEY (status_id) REFERENCES order_status(id),
    FOREIGN KEY (assigned_to) REFERENCES users(id)
);

-- Tabla de comentarios
CREATE TABLE comments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id INT,
    user_id INT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    comment TEXT,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Tabla de actividades
CREATE TABLE activity (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id INT,
    user_id INT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    action VARCHAR(100),
    details TEXT,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Tabla de asociación de órdenes a usuarios
CREATE TABLE user_orders (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT,
    order_id INT,
    assignment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

-- Indices para búsquedas eficientes
CREATE INDEX idx_order_number ON orders(order_number);
CREATE INDEX idx_client_name ON clients(name);
CREATE INDEX idx_equipment_name ON equipments(name);
CREATE INDEX idx_equipment_type ON equipment_types(name);

-- Tipos de equipos para servicios técnicos
INSERT INTO equipment_types (name) VALUES 
('PC de Escritorio'),
('Portátil'),
('Servidor'),
('Impresora Láser'),
('Impresora de Inyección de Tinta'),
('Scanner'),
('Router'),
('Switch'),
('Módem'),
('Monitor'),
('Proyector'),
('Teléfono IP'),
('Tablet'),
('Smartphone'),
('Unidad de Almacenamiento Externa'),
('Disco Duro Interno'),
('Disco Duro Externo'),
('Unidad de Estado Sólido (SSD)'),
('Memoria RAM'),
('Fuente de Alimentación'),
('Placa Base (Motherboard)'),
('Tarjeta Gráfica'),
('Tarjeta de Sonido'),
('Teclado'),
('Ratón'),
('Webcam'),
('Altavoces'),
('Micrófono'),
('Cámara de Seguridad'),
('Sistema de Alimentación Ininterrumpida (UPS)'),
('Equipo de Red'),
('Controlador de Dominio'),
('Servidor de Correo Electrónico'),
('Servidor de Archivos'),
('Servidor Web'),
('Dispositivo IoT'),
('Impresora Multifunción'),
('Fax'),
('Repetidor WiFi'),
('Sistema de Videoconferencia'),
('Servidor NAS'),
('Conmutador'),
('Reproductor de Medios Digitales'),
('Consola de Juegos'),
('Dispositivo de Realidad Virtual'),
('Plotter'),
('Impresora 3D'),
('Cortafuegos (Firewall)'),
('Servidor de Bases de Datos'),
('Controlador de Dominio Secundario'),
('Servidor Proxy'),
('Antivirus Corporativo'),
('Sistema de Respaldo y Recuperación de Datos');


-- Trigger para enviar correo cuando una orden es registrada
CREATE OR REPLACE FUNCTION send_order_email() RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_notify('order_registered', NEW.order_number);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER order_after_insert
AFTER INSERT ON orders
FOR EACH ROW
EXECUTE FUNCTION send_order_email();