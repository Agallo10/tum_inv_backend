# CrearReporteConTipo - Documentación

## Descripción

La función `CrearReporteConTipo` permite crear un reporte de servicio completo con todas sus relaciones asociadas (tipos de mantenimiento y repuestos) en una sola operación transaccional.

## Endpoint HTTP

```
POST /api/reportes-servicio/completo
```

## Estructura de Datos de Entrada

### Ejemplo de JSON Request Body:

```json
{
  "creado_por_id": 2,
  "equipo_id": 1,
  "fecha_inicio": "2024-10-08T09:00:00Z",
  "fecha_finalizacion": "2024-10-08T17:00:00Z",
  "dependencia": "Dirección de Sistemas de Información",
  "ubicacion": "Centro de Innovación - Oficina 401",
  "diagnostico_falla": "Computador presenta lentitud y errores de aplicación",
  "actividad_realizada": "Mantenimiento preventivo completo del equipo",
  "observaciones": "Se realizó limpieza física y actualización de software",
  "tipo_mantenimiento": {
    "tipo": "PREVENTIVO",
    "revision": true,
    "instalacion": false,
    "configuracion": true,
    "ingreso": true,
    "salida": true,
    "concepto_baja": false,
    "otro": false,
    "descripcion_otro": ""
  },
  "repuestos": [
    {
      "cantidad": 1,
      "serial_numero_parte": "MEM-8GB-DDR4-001",
      "marca": "Kingston",
      "tecnologia": "DDR4",
      "capacidad": "8GB",
      "descripcion": "Memoria RAM DDR4 8GB 2666MHz",
      "fecha_utilizacion": "2024-10-08T10:30:00Z"
    }
  ]
}
```

## Campos Obligatorios

### Reporte Principal:

- `creado_por_id` (uint): ID del usuario del sistema que crea el reporte
- `equipo_id` (uint): ID del equipo asociado al reporte (obligatorio)
- `fecha_inicio` (datetime): Fecha y hora de inicio del servicio
- `dependencia` (string): Nombre de la dependencia
- `ubicacion` (string): Ubicación donde se realizó el servicio
- `actividad_realizada` (string): Descripción de la actividad realizada

### Tipo de Mantenimiento:

- `tipo` (string): Puede ser "PREVENTIVO", "CORRECTIVO" o vacío (se mostrará como OTRO)
- `concepto_baja` (bool): Indica si es un concepto de baja de equipo

### Repuestos (opcionales):

Si se especifican repuestos, cada uno debe tener:

- `cantidad` (int): Cantidad utilizada (> 0)
- `serial_numero_parte` (string): Serial o número de parte
- `descripcion` (string): Descripción del repuesto

## Respuesta de Éxito

### Status: 201 Created

```json
{
  "message": "Reporte creado exitosamente",
  "reporte": {
    "ID": 1,
    "CreatedAt": "2024-10-08T15:30:00Z",
    "UpdatedAt": "2024-10-08T15:30:00Z",
    "DeletedAt": null,
    "creado_por_id": 2,
    "equipo_id": 1,
    "fecha_inicio": "2024-10-08T09:00:00Z",
    "fecha_finalizacion": "2024-10-08T17:00:00Z",
    "dependencia": "Dirección de Sistemas de Información",
    "ubicacion": "Centro de Innovación - Oficina 401",
    "diagnostico_falla": "Computador presenta lentitud y errores de aplicación",
    "actividad_realizada": "Mantenimiento preventivo completo del equipo",
    "observaciones": "Se realizó limpieza física y actualización de software",
    "CreadoPor": {
      "ID": 2,
      "nombre": "Técnico",
      "apellido": "Soporte",
      "username": "tecnico",
      "rol": "tecnico"
    },
    "Equipo": {
      "ID": 1,
      "placa_inventario": "PC-001",
      "marca": "Dell",
      "serial": "ABC123",
      "UsuarioResponsable": {
        "ID": 1,
        "nombres_apellidos": "Juan Carlos Pérez García",
        "cedula": "12345678",
        "tipo_vinculacion": "Planta"
      }
    },
    "TipoMantenimiento": {
      "ID": 1,
      "ReporteID": 1,
      "tipo": "PREVENTIVO",
      "revision": true,
      "instalacion": false,
      "configuracion": true,
      "ingreso": true,
      "salida": true,
      "otro": false,
      "descripcion_otro": ""
    },
    "Repuestos": [
      {
        "ID": 1,
        "ReporteID": 1,
        "cantidad": 1,
        "serial_numero_parte": "MEM-8GB-DDR4-001",
        "marca": "Kingston",
        "tecnologia": "DDR4",
        "capacidad": "8GB",
        "descripcion": "Memoria RAM DDR4 8GB 2666MHz",
        "fecha_utilizacion": "2024-10-08T10:30:00Z"
      }
    ]
  }
}
```

## Respuestas de Error

### 400 Bad Request

```json
{
  "error": "Debe especificar el tipo de mantenimiento"
}
```

### 500 Internal Server Error

```json
{
  "error": "Error al crear el reporte: [detalle del error]"
}
```

## Características de la Función

1. **Transaccional**: Toda la operación se ejecuta en una transacción. Si alguna parte falla, se hace rollback completo.

2. **Validaciones**: Se valida el tipo de mantenimiento y el equipo antes de crear el reporte.

3. **Relaciones**: Maneja automáticamente las relaciones:
   - one-to-one: tipo de mantenimiento
   - one-to-many: repuestos
   - many-to-one: creado por (Usuario), equipo (con su usuario responsable)

4. **Atomicidad**: O se crea todo correctamente o no se crea nada.

5. **Separación de Responsabilidades**: La lógica de base de datos está en el repositorio, las validaciones de negocio en el servicio.

6. **Respuesta Completa**: Devuelve el reporte completo con todas sus relaciones cargadas, incluyendo el equipo y su usuario responsable.

## Uso desde el Código

```go
// Ejemplo de uso en un servicio
reporteData := &dto.CrearReporteCompletoDTO{
    // ... llenar datos
}

reporte, err := reporteService.CrearReporteConTipo(reporteData)
if err != nil {
    // Manejar error
    return err
}

// El reporte se creó exitosamente con todas sus relaciones
fmt.Printf("Reporte ID: %d creado exitosamente\n", reporte.ID)
```

## Notas Importantes

- El usuario que crea el reporte (`creado_por_id`) debe existir en la base de datos
- El equipo (`equipo_id`) es obligatorio y debe existir
- El usuario responsable se obtiene automáticamente del equipo asociado
- Si no se especifican repuestos, el array puede estar vacío
- El tipo de mantenimiento es obligatorio
- Todas las fechas deben estar en formato RFC3339 (ISO 8601)
- La función es thread-safe gracias al uso de transacciones de base de datos
