# Ejemplos de CURL - Sistema de Inventario

## 1. Autenticación

### Login como Técnico
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "tecnico",
    "password": "tecnico123"
  }'
```

**Respuesta esperada:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2025-12-08T15:30:00Z",
  "usuario": {
    "ID": 2,
    "nombre": "Técnico",
    "apellido": "Soporte",
    "email": "tecnico@municipio.gov.co",
    "username": "tecnico",
    "rol": "tecnico",
    "activo": true
  }
}
```

**Guardar el token:** Copia el valor del campo `token` para usarlo en las siguientes peticiones.

---

### Login como Admin
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

---

## 2. Consultar Usuarios Responsables

### Listar todos los usuarios responsables
```bash
curl -X GET http://localhost:8080/api/usuarios-responsables \
  -H "Content-Type: application/json"
```

**Respuesta esperada:**
```json
[
  {
    "ID": 1,
    "nombres": "Juan Carlos",
    "apellidos": "Pérez García",
    "cedula": "1000000001",
    "correo_personal": "juan.perez@email.com",
    "celular": "3001234567",
    "tipo_vinculacion": "Planta",
    "dependencia_id": 1
  },
  {
    "ID": 2,
    "nombres": "María Elena",
    "apellidos": "Rodríguez López",
    "cedula": "1000000002",
    "correo_personal": "maria.rodriguez@email.com",
    "celular": "3009876543",
    "tipo_vinculacion": "Contratista",
    "dependencia_id": 2
  }
]
```

---

### Buscar usuario responsable por cédula
```bash
curl -X GET "http://localhost:8080/api/usuarios-responsables/buscar?cedula=1000000002" \
  -H "Content-Type: application/json"
```

---

## 3. Crear Reporte de Servicio Completo

### Ejemplo 1: Reporte creado por usuario técnico (ID: 2)

```bash
curl -X POST http://localhost:8080/api/reportes-servicio/completo \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TU_TOKEN_AQUI" \
  -d '{
    "creado_por_id": 2,
    "equipo_id": 1,
    "fecha_inicio": "2025-12-08T09:00:00Z",
    "fecha_finalizacion": "2025-12-08T11:30:00Z",
    "dependencia": "Secretaría de Tecnología e Innovación",
    "ubicacion": "Edificio Central - Piso 2",
    "diagnostico_falla": "Equipo no enciende, posible falla en fuente de poder",
    "actividad_realizada": "Reemplazo de fuente de poder ATX 500W",
    "observaciones": "Se realizó limpieza interna del equipo",
    "tipo_mantenimiento": {
      "tipo": "CORRECTIVO",
      "revision": true,
      "instalacion": false,
      "configuracion": false,
      "ingreso": false,
      "salida": true,
      "otro": false,
      "descripcion_otro": ""
    },
    "repuestos": [
      {
        "cantidad": 1,
        "serial_numero_parte": "PSU-500W-2025",
        "marca": "Thermaltake",
        "tecnologia": "ATX",
        "capacidad": "500W",
        "descripcion": "Fuente de poder ATX 500W 80+ Bronze",
        "fecha_utilizacion": "2025-12-08T10:00:00Z"
      }
    ]
  }'
```

**Explicación:**
- `creado_por_id: 2` → El usuario "tecnico" crea el reporte
- `equipo_id: 1` → Equipo asociado al reporte (obligatorio)
- El usuario responsable se obtiene automáticamente del equipo

---

### Ejemplo 2: Mantenimiento Preventivo

```bash
curl -X POST http://localhost:8080/api/reportes-servicio/completo \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TU_TOKEN_AQUI" \
  -d '{
    "creado_por_id": 1,
    "equipo_id": 2,
    "fecha_inicio": "2025-12-08T14:00:00Z",
    "dependencia": "Secretaría de Salud",
    "ubicacion": "Centro Administrativo - Piso 2",
    "diagnostico_falla": "",
    "actividad_realizada": "Mantenimiento preventivo programado: limpieza de hardware, actualización de antivirus, optimización de disco",
    "observaciones": "Equipo en buen estado general",
    "tipo_mantenimiento": {
      "tipo": "PREVENTIVO",
      "revision": true,
      "instalacion": false,
      "configuracion": true,
      "ingreso": false,
      "salida": false,
      "otro": false
    },
    "repuestos": []
  }'
```

**Explicación:**
- `creado_por_id: 1` → El usuario "admin" crea el reporte
- `equipo_id: 2` → Equipo al que se le realiza el mantenimiento
- No se usaron repuestos (mantenimiento preventivo)

---

### Ejemplo 3: Con equipo específico

Primero, obtén el ID de un equipo:

```bash
curl -X GET http://localhost:8080/api/equipos \
  -H "Content-Type: application/json"
```

Luego crea el reporte:

```bash
curl -X POST http://localhost:8080/api/reportes-servicio/completo \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TU_TOKEN_AQUI" \
  -d '{
    "creado_por_id": 2,
    "equipo_id": 1,
    "fecha_inicio": "2025-12-08T08:00:00Z",
    "fecha_finalizacion": "2025-12-08T09:00:00Z",
    "dependencia": "Secretaría de Educación",
    "ubicacion": "Oficina 301",
    "diagnostico_falla": "Lentitud en el sistema operativo",
    "actividad_realizada": "Ampliación de memoria RAM de 4GB a 8GB",
    "observaciones": "Mejoró significativamente el rendimiento",
    "tipo_mantenimiento": {
      "tipo": "CORRECTIVO",
      "revision": false,
      "instalacion": true,
      "configuracion": false,
      "ingreso": false,
      "salida": false,
      "otro": false
    },
    "repuestos": [
      {
        "cantidad": 1,
        "serial_numero_parte": "RAM-DDR4-4GB-2025",
        "marca": "Kingston",
        "tecnologia": "DDR4",
        "capacidad": "4GB",
        "descripcion": "Memoria RAM DDR4 4GB 2666MHz",
        "fecha_utilizacion": "2025-12-08T08:30:00Z"
      }
    ]
  }'
```

---

## 4. Consultar Reportes

### Listar todos los reportes
```bash
curl -X GET http://localhost:8080/api/reportes-servicio \
  -H "Content-Type: application/json"
```

### Obtener un reporte específico con todas sus relaciones
```bash
curl -X GET http://localhost:8080/api/reportes-servicio/1 \
  -H "Content-Type: application/json"
```

**Respuesta esperada:**
```json
{
  "ID": 1,
  "creado_por_id": 2,
  "equipo_id": 1,
  "fecha_inicio": "2025-12-08T09:00:00Z",
  "fecha_finalizacion": "2025-12-08T11:30:00Z",
  "dependencia": "Secretaría de Tecnología e Innovación",
  "ubicacion": "Edificio Central - Piso 2",
  "diagnostico_falla": "Equipo no enciende, posible falla en fuente de poder",
  "actividad_realizada": "Reemplazo de fuente de poder ATX 500W",
  "observaciones": "Se realizó limpieza interna del equipo",
  "creado_por": {
    "ID": 2,
    "nombre": "Técnico",
    "apellido": "Soporte",
    "username": "tecnico",
    "rol": "tecnico"
  },
  "equipo": {
    "ID": 1,
    "placa_inventario": "PC-001",
    "marca": "Dell",
    "serial": "ABC123",
    "usuario_responsable": {
      "ID": 1,
      "nombres_apellidos": "Juan Carlos Pérez García",
      "cedula": "1000000001",
      "tipo_vinculacion": "Planta"
    }
  },
  "tipo_mantenimiento": {
    "ID": 1,
    "tipo": "CORRECTIVO",
    "revision": true,
    "salida": true
  },
  "repuestos": [
    {
      "ID": 1,
      "cantidad": 1,
      "serial_numero_parte": "PSU-500W-2025",
      "marca": "Thermaltake",
      "descripcion": "Fuente de poder ATX 500W 80+ Bronze"
    }
  ]
}
```

---

## 5. Verificar Equipo con Usuario Responsable

### Obtener equipo con su usuario responsable
El usuario responsable se incluye automáticamente al consultar el reporte a través del equipo:

```bash
curl -X GET http://localhost:8080/api/reportes-servicio/1 \
  -H "Content-Type: application/json"
```

En la respuesta verás el campo `equipo` con los datos del equipo y su `usuario_responsable`.

---

## Flujo Completo de Prueba

### 1. Login
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "tecnico", "password": "tecnico123"}' \
  | grep -o '"token":"[^"]*' | cut -d'"' -f4)

echo "Token: $TOKEN"
```

### 2. Obtener ID del usuario logueado
```bash
curl -X GET http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer $TOKEN"
```

### 3. Listar equipos disponibles
```bash
curl -X GET http://localhost:8080/api/equipos
```

### 4. Crear reporte con el usuario como creador
```bash
curl -X POST http://localhost:8080/api/reportes-servicio/completo \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "creado_por_id": 2,
    "equipo_id": 1,
    "fecha_inicio": "2025-12-08T10:00:00Z",
    "dependencia": "Secretaría de Tecnología",
    "ubicacion": "Oficina Principal",
    "actividad_realizada": "Configuración de red y actualización de sistema",
    "tipo_mantenimiento": {
      "tipo": "PREVENTIVO",
      "revision": true,
      "configuracion": true
    }
  }'
```

### 5. Verificar el reporte creado
```bash
curl -X GET http://localhost:8080/api/reportes-servicio/1
```

---

## Notas Importantes

1. **Token JWT**: Reemplaza `TU_TOKEN_AQUI` con el token obtenido en el login
2. **IDs**: Los IDs de usuarios y equipos pueden variar según tu base de datos
3. **Fechas**: Usa formato ISO 8601 (RFC3339): `2025-12-08T10:00:00Z`
4. **Validaciones**:
   - `creado_por_id` es obligatorio
   - `equipo_id` es obligatorio
   - `tipo` debe ser "PREVENTIVO" o "CORRECTIVO"
   - El usuario responsable se obtiene automáticamente del equipo

---

## Verificar Seeders

Para confirmar que los seeders se ejecutaron correctamente:

```bash
# Ver usuarios
curl -X GET http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'

# Ver usuarios responsables
curl -X GET http://localhost:8080/api/usuarios-responsables

# Ver secretarías
curl -X GET http://localhost:8080/api/secretarias

# Ver estados de equipo
curl -X GET http://localhost:8080/api/estados-equipo
```
