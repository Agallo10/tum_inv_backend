# Autenticación - Documentación

## Descripción

El sistema de autenticación permite gestionar usuarios que acceden al sistema de inventario. Solo los usuarios con rol `admin` pueden crear nuevos usuarios.

## Endpoints HTTP

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| POST | `/api/auth/register` | Registrar nuevo usuario | No (pero solo admin debería usarlo) |
| POST | `/api/auth/login` | Iniciar sesión | No |
| POST | `/api/auth/refresh` | Renovar token | No |
| GET | `/api/auth/profile` | Obtener perfil del usuario autenticado | Sí (JWT) |
| GET | `/api/auth/users` | Listar todos los usuarios registrados | Sí (JWT, solo admin) |

## Roles del Sistema

| Rol | Descripción |
|-----|-------------|
| `admin` | Administrador del sistema. Puede crear usuarios y tiene acceso completo. |
| `tecnico` | Técnico de soporte. Puede crear reportes de servicio y gestionar equipos. |
| `usuario` | Usuario básico. Acceso limitado de solo lectura. |

---

## 1. Iniciar Sesión (Login)

Autentica un usuario y devuelve tokens JWT.

### Request

```bash
curl -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

### Respuesta Exitosa (200 OK)

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2026-01-20T12:00:00Z",
  "usuario": {
    "ID": 1,
    "Nombre": "Administrador",
    "Apellido": "Sistema",
    "Cedula": "1107090505",
    "Email": "admin@municipio.gov.co",
    "Username": "admin",
    "Rol": "admin",
    "Activo": true
  }
}
```

### Errores Comunes

```json
// 400 Bad Request - Datos faltantes
{
  "error": "Nombre de usuario y contraseña son obligatorios"
}

// 401 Unauthorized - Credenciales incorrectas
{
  "error": "Credenciales inválidas"
}
```

---

## 2. Registrar Usuario (Solo Admin)

Crea un nuevo usuario en el sistema. **Solo debe ser ejecutado por administradores.**

### Request

```bash
# Primero obtener el token de admin
TOKEN=$(curl -s -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# Crear nuevo usuario técnico
curl -X POST "http://localhost:8080/api/auth/register" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "nombre": "Juan Carlos",
    "apellido": "Pérez García",
    "cedula": "1234567890",
    "email": "juan.perez@municipio.gov.co",
    "username": "jperez",
    "password": "miContraseña123",
    "rol": "tecnico"
  }'
```

### Campos del Request

| Campo | Tipo | Obligatorio | Descripción |
|-------|------|-------------|-------------|
| `nombre` | string | Sí | Nombre del usuario |
| `apellido` | string | Sí | Apellido del usuario |
| `cedula` | string | Sí | Número de cédula de identidad (requerido para firmas) |
| `email` | string | Sí | Correo electrónico (único) |
| `username` | string | Sí | Nombre de usuario para login (único) |
| `password` | string | Sí | Contraseña (mínimo 6 caracteres) |
| `rol` | string | No | Rol del usuario: `admin`, `tecnico`, `usuario` (default: `usuario`) |

### Respuesta Exitosa (201 Created)

```json
{
  "ID": 3,
  "CreatedAt": "2026-01-20T10:30:00Z",
  "UpdatedAt": "2026-01-20T10:30:00Z",
  "DeletedAt": null,
  "Nombre": "Juan Carlos",
  "Apellido": "Pérez García",
  "Cedula": "",
  "Email": "juan.perez@municipio.gov.co",
  "Username": "jperez",
  "Password": "",
  "Rol": "tecnico",
  "Activo": true,
  "UltimoLogin": null
}
```

### Errores Comunes

```json
// 400 Bad Request - Datos faltantes
{
  "error": "Todos los campos son obligatorios"
}

// 400 Bad Request - Usuario ya existe
{
  "error": "el nombre de usuario ya está en uso"
}

// 400 Bad Request - Email ya existe
{
  "error": "el email ya está registrado"
}
```

---

## 3. Renovar Token (Refresh)

Obtiene un nuevo token JWT usando el refresh token.

### Request

```bash
curl -X POST "http://localhost:8080/api/auth/refresh" \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

### Respuesta Exitosa (200 OK)

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2026-01-20T14:00:00Z",
  "usuario": {
    "ID": 1,
    "Nombre": "Administrador",
    "Apellido": "Sistema",
    "Username": "admin",
    "Rol": "admin"
  }
}
```

### Errores Comunes

```json
// 400 Bad Request
{
  "error": "Token de actualización es obligatorio"
}

// 401 Unauthorized
{
  "error": "Token de actualización inválido o expirado"
}
```

---

## 4. Obtener Perfil (Autenticado)

Obtiene los datos del usuario actualmente autenticado.

### Request

```bash
curl -X GET "http://localhost:8080/api/auth/profile" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Respuesta Exitosa (200 OK)

```json
{
  "ID": 1,
  "CreatedAt": "2026-01-13T21:07:28Z",
  "UpdatedAt": "2026-01-20T10:00:00Z",
  "DeletedAt": null,
  "Nombre": "Administrador",
  "Apellido": "Sistema",
  "Cedula": "1107090505",
  "Email": "admin@municipio.gov.co",
  "Username": "admin",
  "Password": "",
  "Rol": "admin",
  "Activo": true,
  "UltimoLogin": "2026-01-20T10:00:00Z"
}
```

### Errores Comunes

```json
// 401 Unauthorized - Token faltante o inválido
{
  "error": "Token de autorización requerido"
}

// 404 Not Found
{
  "error": "Usuario no encontrado"
}
```

---

## 5. Listar Usuarios (Solo Admin)

Obtiene la lista completa de todos los usuarios registrados en el sistema.

### Request

```bash
# Obtener token de admin
TOKEN=$(curl -s -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# Listar todos los usuarios
curl -X GET "http://localhost:8080/api/auth/users" \
  -H "Authorization: Bearer $TOKEN"
```

### Respuesta Exitosa (200 OK)

```json
{
  "usuarios": [
    {
      "ID": 1,
      "CreatedAt": "2026-01-13T21:07:28Z",
      "UpdatedAt": "2026-01-20T10:00:00Z",
      "DeletedAt": null,
      "Nombre": "Administrador",
      "Apellido": "Sistema",
      "Cedula": "1107090505",
      "Email": "admin@municipio.gov.co",
      "Username": "admin",
      "Password": "",
      "Rol": "admin",
      "Activo": true,
      "UltimoLogin": "2026-01-20T10:00:00Z"
    },
    {
      "ID": 2,
      "CreatedAt": "2026-01-13T21:07:28Z",
      "UpdatedAt": "2026-01-20T10:00:00Z",
      "DeletedAt": null,
      "Nombre": "Técnico",
      "Apellido": "Soporte",
      "Cedula": "1107090506",
      "Email": "tecnico@municipio.gov.co",
      "Username": "tecnico",
      "Password": "",
      "Rol": "tecnico",
      "Activo": true,
      "UltimoLogin": "2026-01-19T15:30:00Z"
    }
  ],
  "total": 2
}
```

### Errores Comunes

```json
// 401 Unauthorized - Token faltante o inválido
{
  "error": "Token de autorización requerido"
}

// 500 Internal Server Error
{
  "error": "Error obteniendo usuarios"
}
```

---

## Ejemplos de Flujo Completo

### Flujo: Admin crea un nuevo técnico

```bash
#!/bin/bash

# 1. Login como admin
echo "=== Iniciando sesión como admin ==="
RESPONSE=$(curl -s -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}')

TOKEN=$(echo $RESPONSE | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")
echo "Token obtenido: ${TOKEN:0:50}..."

# 2. Crear nuevo técnico
echo -e "\n=== Creando nuevo técnico ==="
curl -s -X POST "http://localhost:8080/api/auth/register" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "nombre": "María",
    "apellido": "González",
    "cedula": "0987654321",
    "email": "maria.gonzalez@municipio.gov.co",
    "username": "mgonzalez",
    "password": "tecnico2026",
    "rol": "tecnico"
  }' | python3 -m json.tool

# 3. Verificar login del nuevo usuario
echo -e "\n=== Verificando login del nuevo usuario ==="
curl -s -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "mgonzalez", "password": "tecnico2026"}' | python3 -m json.tool
```

---

## Usuarios por Defecto (Seed)

El sistema viene con los siguientes usuarios pre-configurados:

| Username | Password | Rol | Email |
|----------|----------|-----|-------|
| `admin` | `admin123` | admin | admin@municipio.gov.co |
| `tecnico` | `tecnico123` | tecnico | tecnico@municipio.gov.co |

> ⚠️ **Importante:** Cambiar las contraseñas por defecto en producción.

---

## Notas de Seguridad

1. **Tokens JWT**: Los tokens tienen una duración limitada. Usar el endpoint `/refresh` para renovarlos.
2. **Contraseñas**: Se almacenan hasheadas con bcrypt.
3. **Roles**: Solo el admin puede crear usuarios. Implementar middleware de autorización en la interfaz.
4. **HTTPS**: En producción, siempre usar HTTPS para proteger las credenciales en tránsito.
