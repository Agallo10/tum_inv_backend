# Sistema de Inventario Tecnológico Municipal - Tumaco

## Descripción General

Este es un sistema backend desarrollado en **Go** usando el framework **Echo** para gestionar el inventario de equipos tecnológicos de la Alcaldía Municipal. La aplicación permite administrar equipos de cómputo, sus componentes, usuarios responsables, mantenimientos, reportes de servicio y toda la información asociada a la infraestructura tecnológica municipal.

## Tecnologías Principales

- **Lenguaje**: Go 1.23.3
- **Framework Web**: Echo v4
- **ORM**: GORM
- **Base de Datos**: PostgreSQL
- **Autenticación**: JWT (JSON Web Tokens)
- **Encriptación**: bcrypt para contraseñas

## Arquitectura del Proyecto

El proyecto sigue una **arquitectura en capas** (Clean Architecture) con separación clara de responsabilidades:

```
tum_inv_backend/
├── internal/
│   ├── api/                    # Capa de presentación
│   │   ├── controllers/        # Controladores HTTP
│   │   ├── middleware/         # Middleware de autenticación
│   │   └── routes/             # Definición de rutas
│   ├── domain/                 # Capa de dominio
│   │   ├── models/             # Modelos de datos
│   │   ├── repositories/       # Interfaces y implementaciones de repositorios
│   │   └── services/           # Lógica de negocio
│   └── infrastructure/         # Capa de infraestructura
│       ├── config/             # Configuración de la aplicación
│       ├── database/           # Conexión y migraciones de BD
│       └── seed/               # Datos iniciales (seeders)
├── docs/                       # Documentación
├── go.mod                      # Dependencias del proyecto
└── server.go                   # Punto de entrada de la aplicación
```

## Modelos de Datos Principales

### 1. **Estructura Organizacional**

#### Secretaria
Representa las secretarías municipales.
- Nombre, descripción, ubicación
- Secretario a cargo
- Relación: tiene múltiples dependencias

#### Dependencia
Representa las oficinas o dependencias dentro de cada secretaría.
- Nombre, ubicación, jefe de oficina
- Correo institucional y teléfono
- Relación: pertenece a una secretaría, tiene múltiples usuarios responsables

### 2. **Gestión de Usuarios**

#### Usuario (Sistema de Autenticación)
Usuario del sistema con capacidad de login.
- Nombre, apellido, email, username
- Password (hasheado con bcrypt)
- Rol: `admin`, `usuario`, `tecnico`
- Estado: activo/inactivo
- Último login

#### UsuarioResponsable
Funcionario responsable de un equipo.
- Nombres, apellidos, cédula
- Correo personal, celular
- Tipo de vinculación: Planta, Contratista, Otro
- Relación: pertenece a una dependencia, responsable de un equipo

### 3. **Inventario de Equipos**

#### Equipo
Dispositivo tecnológico principal.
- Tipo: Todo en Uno, Escritorio, Portátil, Impresora, Escáner, Otro
- Placa de inventario, marca, serial, modelo
- Estado del equipo
- Usuario responsable
- Fecha de diligenciamiento
- Observaciones generales

**Relaciones:**
- Periféricos (teclado, mouse, monitor)
- Hardware interno (disco, RAM, procesador)
- Software instalado
- Configuración de red
- Usuarios del sistema
- Accesos remotos
- Backups
- Reportes de servicio

#### EstadoEquipo
Estado actual del equipo.
- Opciones: Activo, Inactivo, En Mantenimiento, Dañado, Dado de Baja
- Descripción y estado activo/inactivo

### 4. **Componentes del Equipo**

#### Periferico
Dispositivos externos conectados.
- Tipos: Teclado, Mouse, Monitor, Otros
- Placa, marca, serial

#### HardwareInterno
Componentes internos.
- Componentes: Disco Duro, Memoria RAM, Procesador
- Tecnología (HDD, SSD, DDR4, etc.)
- Capacidad

#### Software
Programas instalados.
- Nombre, versión
- Tipo de licencia
- Categoría: Sistema Operativo, Paquete de Oficina, Navegador Web, Otro

#### ConfiguracionRed
Configuración de red del equipo.
- Dirección IP
- Asignación: Manual, Automática, Dinámica
- Nombre del dispositivo
- Tipo de conectividad

#### UsuarioSistema
Usuarios locales del sistema operativo.
- Nombre de usuario
- Contraseña
- Tipo: Administrador o no

#### AccesoRemoto
Credenciales de acceso remoto.
- Plataforma (por defecto: AnyDesk)
- Usuario, contraseña
- ID de conexión

#### Backup
Copias de seguridad realizadas.
- Fecha de backup
- Número de carpetas
- Peso total de archivos
- Ruta del backup
- Indicador si se realizó exitosamente

### 5. **Gestión de Mantenimiento**

#### ReporteServicio
Registro de intervenciones técnicas.
- Fecha de inicio y finalización
- Dependencia y ubicación
- Diagnóstico de la falla
- Actividad realizada
- Observaciones
- Equipo asociado (obligatorio)
- Usuario que creó el reporte

**Relaciones:**
- Tipo de mantenimiento
- Repuestos utilizados
- Creado por (Usuario del sistema)
- Equipo (con su UsuarioResponsable)

#### TipoMantenimiento
Clasificación del mantenimiento.
- Tipo: PREVENTIVO o CORRECTIVO
- Actividades: Revisión, Instalación, Configuración, Ingreso, Salida, Otro
- Descripción de "Otro"

#### Repuesto
Partes utilizadas en reparaciones.
- Cantidad
- Serial/Número de parte
- Marca, tecnología, capacidad
- Descripción
- Fecha de utilización

## Funcionalidades Principales

### 1. **Autenticación y Autorización**
- **Registro de usuarios** con roles (admin, usuario, técnico)
- **Login** con generación de token JWT
- **Refresh token** para renovar sesiones
- **Middleware de autenticación** para proteger rutas
- **Verificación de roles** para control de acceso

### 2. **Gestión de Estructura Organizacional**
- CRUD de Secretarías
- CRUD de Dependencias
- Consulta de dependencias por secretaría
- Consulta de usuarios por dependencia

### 3. **Gestión de Inventario**
- **Equipos**: CRUD completo
  - Listado general y con detalle completo
  - Filtrado por dependencia
  - "Hoja de vida" del equipo (toda la información relacionada)
- **Estados de equipo**: gestión de estados con activación/desactivación
- **Periféricos**: CRUD y consulta por equipo
- **Hardware interno**: CRUD y consulta por equipo
- **Software**: CRUD y consulta por equipo
- **Configuración de red**: CRUD y consulta por equipo
- **Usuarios del sistema**: CRUD y consulta por equipo
- **Accesos remotos**: CRUD y consulta por equipo
- **Backups**: CRUD y consulta por equipo

### 4. **Gestión de Usuarios Responsables**
- CRUD de usuarios responsables
- Búsqueda por cédula
- Consulta por dependencia

### 5. **Gestión de Mantenimiento**
- **Reportes de servicio**: 
  - CRUD completo
  - Creación de reporte completo (con tipo de mantenimiento incluido)
  - Consulta por equipo
  - Registro del usuario que crea el reporte
  - Usuario responsable obtenido automáticamente del equipo
- **Tipos de mantenimiento**: CRUD y consulta por reporte
- **Repuestos**: CRUD y consulta por reporte

## API REST Endpoints

### Autenticación (`/api/auth`)
- `POST /register` - Registrar nuevo usuario
- `POST /login` - Iniciar sesión
- `POST /refresh` - Renovar token
- `GET /profile` - Obtener perfil (protegido)

### Equipos (`/api/equipos`)
- `POST /` - Crear equipo
- `GET /` - Listar todos los equipos
- `GET /AllDetalle` - Listar con todos los detalles
- `GET /:id` - Obtener equipo por ID
- `PUT /:id` - Actualizar equipo
- `DELETE /:id` - Eliminar equipo
- `GET /:dependenciaId/dependencia` - Equipos por dependencia
- `GET /:equipoId/hv` - Hoja de vida del equipo
- `GET /:equipoId/perifericos` - Periféricos del equipo
- `GET /:equipoId/software` - Software del equipo
- `GET /:equipoId/hardware-interno` - Hardware interno del equipo
- `GET /:equipoId/configuracion-red` - Configuración de red
- `GET /:equipoId/usuarios-sistema` - Usuarios del sistema
- `GET /:equipoId/accesos-remotos` - Accesos remotos
- `GET /:equipoId/backups` - Backups del equipo
- `GET /:equipoId/reportes-servicio` - Reportes del equipo

### Secretarías (`/api/secretarias`)
- `POST /` - Crear secretaría
- `GET /` - Listar secretarías
- `GET /:id` - Obtener por ID
- `PUT /:id` - Actualizar
- `DELETE /:id` - Eliminar
- `GET /:id/dependencias` - Dependencias de la secretaría

### Dependencias (`/api/dependencias`)
- `POST /` - Crear dependencia
- `GET /` - Listar dependencias
- `GET /:id` - Obtener por ID
- `PUT /:id` - Actualizar
- `DELETE /:id` - Eliminar
- `GET /:id/usuarios` - Usuarios de la dependencia
- `GET /:secretariaId/dependencias` - Por secretaría

### Estados de Equipo (`/api/estados-equipo`)
- `POST /` - Crear estado
- `GET /` - Listar todos
- `GET /activos` - Listar activos
- `GET /:id` - Obtener por ID
- `PUT /:id` - Actualizar
- `DELETE /:id` - Eliminar
- `PATCH /:id/toggle-activo` - Activar/desactivar
- `GET /:id/equipos` - Equipos por estado

### Usuarios Responsables (`/api/usuarios-responsables`)
- CRUD completo
- `GET /buscar` - Buscar por cédula
- `GET /:dependenciaId/dependencia` - Por dependencia

### Reportes de Servicio (`/api/reportes-servicio`)
- CRUD completo
- `POST /completo` - Crear reporte con tipo de mantenimiento
- `GET /:reporteId/tipos-mantenimiento` - Tipos de mantenimiento
- `GET /:reporteId/repuestos` - Repuestos utilizados

### Otros Endpoints
Similar estructura CRUD para:
- Periféricos (`/api/perifericos`)
- Software (`/api/software`)
- Hardware Interno (`/api/hardware-interno`)
- Configuración de Red (`/api/configuraciones-red`)
- Usuarios del Sistema (`/api/usuarios-sistema`)
- Accesos Remotos (`/api/accesos-remotos`)
- Backups (`/api/backups`)
- Tipos de Mantenimiento (`/api/tipos-mantenimiento`)
- Repuestos (`/api/repuestos`)

## Configuración

La aplicación usa variables de entorno (archivo `.env`):

```env
# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=inventario
DB_SSLMODE=disable
DB_TIMEOUT=10

# Aplicación
APP_PORT=8080
APP_ENV=development

# Seguridad
JWT_SECRET=tu_clave_secreta_jwt_super_segura
```

## Características de Seguridad

1. **Contraseñas hasheadas** con bcrypt
2. **Autenticación JWT** con tokens de acceso y refresh
3. **Middleware de autenticación** para proteger endpoints
4. **Verificación de roles** (admin, usuario, técnico)
5. **CORS** habilitado
6. **Secure middleware** de Echo
7. **Timeout** de 30 segundos en las peticiones

## Datos Iniciales (Seeders)

Al iniciar la aplicación, se ejecutan seeders que crean:

1. **5 Secretarías** con sus datos completos
2. **10 Dependencias** distribuidas entre las secretarías
3. **5 Estados de equipo** (Activo, Inactivo, En Mantenimiento, Dañado, Dado de Baja)
4. **2 Usuarios del sistema**:
   - **Admin**: `admin` / `admin123`
   - **Técnico**: `tecnico` / `tecnico123`

Los seeders verifican si los datos ya existen antes de crearlos, evitando duplicados.

## Migraciones Automáticas

Al conectarse a la base de datos, GORM ejecuta `AutoMigrate` para crear/actualizar todas las tablas necesarias basándose en los modelos definidos.

## Flujo de Inicio de la Aplicación

1. **Carga de configuración** desde variables de entorno
2. **Inicialización de Echo** con middlewares:
   - Logger
   - Recover (recuperación de panics)
   - CORS
   - Secure
   - Timeout
3. **Conexión a PostgreSQL** con reintentos (3 intentos)
4. **Configuración del pool de conexiones**
5. **Ejecución de migraciones** automáticas
6. **Ejecución de seeders** para datos iniciales
7. **Configuración de rutas** con inyección de dependencias
8. **Inicio del servidor** en puerto 8080

## Inyección de Dependencias

El proyecto usa un patrón de inyección de dependencias manual:

```
Repositorios → Servicios → Controladores → Rutas
```

Cada capa depende de interfaces, facilitando el testing y el mantenimiento.

## Middleware

### JWTMiddleware
- **Authenticate**: Valida el token JWT en el header Authorization
- **RequireRole**: Verifica que el usuario tenga un rol específico

Los tokens deben enviarse en el formato:
```
Authorization: Bearer <token>
```

## Características de Rendimiento

- **Pool de conexiones** a la base de datos (10 idle, 100 max open)
- **Conexiones reutilizables** con tiempo de vida de 1 hora
- **Logging condicional** (detallado en desarrollo, silencioso en producción)
- **Timeout de requests** de 30 segundos

## Casos de Uso Principales

### 1. Registro de Nuevo Equipo
1. Crear/buscar usuario responsable
2. Crear equipo con estado inicial
3. Agregar periféricos
4. Registrar hardware interno
5. Instalar y registrar software
6. Configurar red
7. Crear usuarios del sistema
8. Configurar acceso remoto

### 2. Mantenimiento Correctivo
1. Crear reporte de servicio (con usuario creador)
2. Asociar equipo afectado (obligatorio)
3. Definir tipo de mantenimiento (CORRECTIVO)
4. Registrar repuestos utilizados
5. Completar diagnóstico y actividades realizadas
6. El usuario responsable se obtiene automáticamente del equipo

### 3. Consulta de Hoja de Vida
- Obtener toda la información de un equipo:
  - Datos básicos
  - Usuario responsable y dependencia
  - Componentes (periféricos, hardware, software)
  - Configuración de red
  - Historial de mantenimientos
  - Backups realizados

## Ventajas del Sistema

1. **Trazabilidad completa** de equipos y mantenimientos
2. **Organización estructurada** por secretarías y dependencias
3. **Control de acceso** basado en roles
4. **Historial detallado** de intervenciones técnicas
5. **Gestión de inventario** integral
6. **API RESTful** fácil de integrar con frontends
7. **Escalable** gracias a su arquitectura en capas
8. **Mantenible** con separación clara de responsabilidades

## Posibles Mejoras Futuras

- [ ] Implementar paginación en listados
- [ ] Agregar filtros y búsquedas avanzadas
- [ ] Sistema de notificaciones
- [ ] Reportes en PDF
- [ ] Dashboard con estadísticas
- [ ] Historial de cambios (auditoría)
- [ ] Integración con Active Directory
- [ ] Alertas de mantenimiento preventivo programado
- [ ] API de exportación de datos
- [ ] Documentación con Swagger/OpenAPI
