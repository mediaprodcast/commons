## Commons

This repository houses shared code, libraries, and configurations for the Media Processing Platform. This promotes code reusability, consistency, and maintainability across all platform services.

**Contents:**

* **Logging:**
    * `logger.go`: Provides a standardized logging interface with configurable log levels, structured logging formats (e.g., JSON), and integration with logging frameworks (e.g., Zap, logrus).
* **Configuration:**
    * `config.go`: Defines interfaces and structs for application configurations (e.g., database connection strings, API keys, service endpoints).
    * `config_loader.go`: Provides functions for loading configuration from environment variables, configuration files (e.g., YAML, JSON), or a configuration server (e.g., Consul).
* **Error Handling:**
    * `errors.go`: Defines custom error types and helper functions for handling and propagating errors throughout the system.
* **Data Structures:**
    * `models.go`: Defines common data structures used across services (e.g., Job struct, MediaMetadata struct).
* **Utils:**
    * Helper functions for common tasks (e.g., file system operations, string manipulation, time utilities).
* **Testing:**
    * Shared test utilities (e.g., test helpers for setting up mock objects, asserting expected behavior).

**Usage:**

* Each service within the platform will import and utilize the shared code from this "common" repository.
* This allows for consistent error handling, logging, and configuration management across all services.
* Changes made to the "common" repository will be reflected in all dependent services upon updating their dependencies.

---

## Commands

The `Makefile` provided in this repository facilitates protocol buffer compilation and output directory management. Below are the key commands:

### Build Protocol Buffers
To compile the protocol buffer (`.proto`) files and generate the necessary Go files for gRPC:

```bash
make all
