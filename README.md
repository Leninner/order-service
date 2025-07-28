# Order Service

Microservicio para gestión de órdenes en TanEats.

## Estructura del Proyecto

```
order-service/
├── cmd/
│   └── main.go
├── internal/
│   ├── application/
│   ├── config/
│   ├── dataaccess/
│   ├── di/
│   └── domain/
├── __tests__/
├── go.mod
└── README.md
```

## Tecnologías

- **Go 1.24.2**
- **Arquitectura Hexagonal (Clean Architecture)**
- **Domain-Driven Design (DDD)**
- **Event-Driven Architecture**

## Desarrollo con GitHub Submódulos

### Configuración Inicial

Este servicio es parte de un proyecto modular usando Git submódulos. El repositorio principal está en: `https://github.com/leninner/taneats`

### Clonar el Proyecto Completo

```bash
git clone --recursive https://github.com/leninner/taneats.git
cd taneats/services/order-service
```

### Trabajar en el Servicio

1. **Navegar al directorio del servicio:**
   ```bash
   cd services/order-service
   ```

2. **Hacer cambios en el código**

3. **Commit y push de cambios:**
   ```bash
   git add .
   git commit -m "feat: add new order functionality"
   git push origin main
   ```

4. **Actualizar el submódulo en el repositorio principal:**
   ```bash
   cd ../..  # volver al root del proyecto
   git add services/order-service
   git commit -m "Update order-service submodule"
   git push origin main
   ```

### Actualizar Dependencias

```bash
go mod tidy
go mod download
```

### Ejecutar Tests

```bash
go test ./...
```

### Ejecutar el Servicio

```bash
go run cmd/main.go
```

## Estructura de Dominio

### Entidades
- `Order` - Entidad principal de orden
- `Customer` - Cliente que realiza la orden
- `Restaurant` - Restaurante que prepara la orden
- `Product` - Productos del restaurante

### Eventos de Dominio
- `OrderCreatedEvent`
- `OrderPaidEvent`
- `OrderCancelledEvent`

### Servicios de Aplicación
- `OrderApplicationService` - Orquestador de casos de uso

## Integración con Otros Servicios

- **Shared Module**: Utiliza tipos y utilidades compartidas
- **Restaurant Service**: Comunicación para aprobación de órdenes
- **Payment Service**: Integración para procesamiento de pagos

## Convenciones de Git

### Commits
- `feat:` nuevas características
- `fix:` correcciones de bugs
- `refactor:` refactorización de código
- `test:` agregar o modificar tests
- `docs:` documentación

### Branches
- `main` - código estable
- `develop` - desarrollo activo
- `feature/` - nuevas características
- `hotfix/` - correcciones urgentes

## Troubleshooting

### Problema: Submódulo no actualizado
```bash
git submodule update --remote services/order-service
```

### Problema: Dependencias no encontradas
```bash
go mod tidy
go mod download
```

### Problema: Tests fallando
```bash
go clean -testcache
go test ./...
``` 