# GenerarPDFReporte - Documentación

## Descripción

Los endpoints de PDF permiten generar y visualizar reportes de servicio técnico en formato PDF, siguiendo el formato oficial de la Alcaldía Distrital de Tumaco.

## Endpoints HTTP

### Descargar PDF

```
GET /api/reportes-servicio/:id/pdf?usuario_id=:usuario_id
```

### Visualizar PDF en navegador

```
GET /api/reportes-servicio/:id/pdf/view?usuario_id=:usuario_id
```

## Parámetros

| Parámetro | Tipo | Ubicación | Descripción |
|-----------|------|-----------|-------------|
| `id` | uint | Path | ID del reporte de servicio |
| `usuario_id` | uint | Query | ID del usuario del sistema que genera el PDF |

## Ejemplos CURL

### Descargar PDF

```bash
# Descargar el PDF del reporte ID 1, generado por el usuario ID 1
curl -o reporte.pdf "http://localhost:8080/api/reportes-servicio/1/pdf?usuario_id=1"

# Con autenticación (si está habilitada)
curl -H "Authorization: Bearer <token>" -o reporte.pdf "http://localhost:8080/api/reportes-servicio/1/pdf?usuario_id=1"
```

### Visualizar en navegador

```bash
# Abrir directamente en el navegador
open "http://localhost:8080/api/reportes-servicio/1/pdf/view?usuario_id=1"

# O con curl para guardar
curl -o reporte_view.pdf "http://localhost:8080/api/reportes-servicio/1/pdf/view?usuario_id=1"
```

## Respuesta de Éxito

### Status: 200 OK

- **Content-Type**: `application/pdf`
- **Content-Disposition**: 
  - Para `/pdf`: `attachment; filename=reporte_servicio_<id>.pdf`
  - Para `/pdf/view`: `inline; filename=reporte_servicio_<id>.pdf`

El cuerpo de la respuesta es el archivo PDF binario.

## Estructura del PDF

El PDF generado contiene dos páginas:

### Página 1 - Reporte Principal

1. **Encabezado**: Logo Alcaldía, Escudo de Colombia, información institucional
2. **Título**: "REPORTE DE SERVICIO TECNICO"
3. **Datos del Reporte**:
   - Fecha inicio / Fecha finalización
   - Dependencia / Ubicación
   - Equipo / Marca
   - Modelo / Serie
4. **Trabajo Realizado** (checkboxes):
   - Fila 1: Mantenimiento Preventivo | Revisión | Instalación | Configuración
   - Fila 2: Mantenimiento Correctivo | Ingreso | Salida | Concepto de Baja
   - Fila 3: Otro: [descripción]
5. **Secciones de Texto**:
   - Diagnóstico y/o Falla Reportada
   - Actividad Realizada
   - Observaciones
6. **Pie de página**: Información de contacto institucional

### Página 2 - Repuestos y Firmas

1. **Encabezado**: (mismo que página 1)
2. **Título**: "REPUESTOS EMPLEADOS Y/O REEMPLAZADO"
3. **Tabla de Repuestos**:
   - Cantidad
   - Serial o Número de Parte
   - Marca, Tecnología, Capacidad, Descripción
4. **Sección de Firmas** (dos columnas):
   - **Izquierda**: Funcionario y/o Contratista del Servicio (UsuarioResponsable)
     - Nombre
     - Cargo (TipoVinculacion)
     - Firma y C.C.
   - **Derecha**: Funcionario y/o Contratista de Sistemas (Usuario)
     - Nombre
     - Cargo (Rol)
     - Firma y C.C.

## Respuestas de Error

### 400 Bad Request

```json
{
  "error": "ID de reporte inválido"
}
```

```json
{
  "error": "usuario_id es requerido"
}
```

### 500 Internal Server Error

```json
{
  "error": "Error generando PDF: [detalle del error]"
}
```

## Datos Utilizados

El PDF combina datos de múltiples modelos:

| Sección | Modelo | Campos |
|---------|--------|--------|
| Datos Reporte | ReporteServicio | FechaInicio, FechaFinalizacion, Dependencia, Ubicacion, DiagnosticoFalla, ActividadRealizada, Observaciones |
| Equipo | Equipo | TipoDispositivo, Marca, Modelo, Serial |
| Trabajo Realizado | TipoMantenimiento | Tipo, Revision, Instalacion, Configuracion, Ingreso, Salida, ConceptoBaja, Otro, DescripcionOtro |
| Repuestos | Repuesto[] | Cantidad, SerialNumeroParte, Marca, Tecnologia, Capacidad, Descripcion |
| Firma Izquierda | UsuarioResponsable | NombresApellidos, Cedula, TipoVinculacion |
| Firma Derecha | Usuario | Nombre, Apellido, Cedula, Rol |

## Requisitos

- El reporte debe existir en la base de datos
- El usuario del sistema debe existir
- Los archivos de logos deben estar en `assets/logos/`:
  - `logo-004.png` - Logo Alcaldía
  - `logo-escudo.png` - Escudo de Colombia
  - `logo-marca-agua.png` - Marca de agua

## Notas Importantes

- El PDF tiene tamaño carta (216 x 279 mm)
- Incluye marca de agua del escudo de Colombia
- Las fechas se formatean como DD/MM/YYYY
- Si `Tipo` está vacío, se marca automáticamente la casilla "OTRO"
- El campo `ConceptoBaja` es un nuevo checkbox para reportar equipos dados de baja
- La cédula del usuario del sistema (`Usuario.Cedula`) aparece en la firma derecha
