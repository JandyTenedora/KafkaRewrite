# Kafka Producer in Go

This project is a simple Kafka producer implemented in Go. It provides a foundational understanding of Kafka producers and how to handle basic message sending operations. The implementation is simplified for learning purposes and doesn't include all Kafka protocol features.

---

## Project Structure

```
kafka-producer/
├── main.go
├── producer/
│   ├── producer.go
│   ├── config.go
│   └── connection.go
├── utils/
│   └── hash.go
│   └── logger.go
└── go.mod
```

---

### **1. `main.go`**

**Purpose**: Entry point of the application. 

- Initializes the producer configuration.
- Creates an instance of the producer.
- Sends an example message to a Kafka topic.

**Instructions**:
- Import necessary packages for logging and producer usage.
- Define a configuration with Kafka broker address and topic.
- Call the `NewProducer` function to initialize a producer.
- Use the `SendMessage` method to send a sample message.
- Handle errors appropriately and log the results.

**Libraries Used**: 
- `log`: For logging messages.
- Custom `producer` package for Kafka producer functionality.

---

### **2. `producer/producer.go`**

**Purpose**: Handles high-level producer logic, including message sending and producer setup.

**Structs**:
- `Producer`:
  - Contains a configuration object.
  - Manages the connection to the Kafka broker.

**Methods**:
1. `NewProducer(config Config) (*Producer, error)`:
   - Initializes a new Kafka producer instance.
   - Establishes a connection to the Kafka broker using the configuration.

2. `SendMessage(message string) error`:
   - Accepts a message string.
   - Serializes the message if necessary.
   - Sends the message to the broker using the connection.

**Responsibilities**:
- Encapsulate producer setup logic.
- Abstract message serialization and sending.

**Libraries Used**:
- `errors`: For error handling.

---

### **3. `producer/config.go`**

**Purpose**: Manages configuration settings for the producer.

**Structs**:
- `Config`:
  - Fields include:
    - `BrokerAddress` (string): Address of the Kafka broker (e.g., `localhost:9092`).
    - `Topic` (string): Kafka topic to send messages to.

**Responsibilities**:
- Provide a centralized way to manage and access producer configurations.

**Libraries Used**: None.

---

### **4. `producer/connection.go`**

**Purpose**: Handles low-level network communication with the Kafka broker.

**Structs**:
- `Connection`:
  - Contains the TCP connection to the Kafka broker.

**Methods**:
1. `NewConnection(brokerAddress string) (*Connection, error)`:
   - Establishes a TCP connection to the Kafka broker.
   - Handles errors if the connection cannot be established.

2. `WriteMessage(topic string, message []byte) error`:
   - Sends a message to the broker.
   - Constructs a simple format combining topic and message.
   - Writes the message to the broker connection.

**Responsibilities**:
- Abstract low-level TCP communication.
- Simplify sending messages to the broker.

**Libraries Used**:
- `net`: For TCP connection handling.
- `errors`: For error handling.

---

### **5. `utils/hash.go`**

**Purpose**: Provides reusable logging functionality.

**Functions**:
1. `HashStringToInt32(input string)`:
   - Logs informational messages with a specific prefix.

**Responsibilities**:
- Provide hash functions required across project 

**Libraries Used**:
- `fmt`: For printing.
- `hash/fnv`: For introducing hash libraries.

---


### **6. `utils/logger.go`**

**Purpose**: Provides reusable logging functionality.

**Functions**:
1. `Info(message string)`:
   - Logs informational messages with a specific prefix.

2. `Error(message string)`:
   - Logs error messages with a specific prefix.

**Responsibilities**:
- Standardize logging across the application.

**Libraries Used**:
- `log`: For logging messages.

---

### **7. `go.mod`**

**Purpose**: Manages dependencies and specifies the module name.

**Fields**:
- Module name (`module kafka-producer`).
- Go version (`go 1.20`).

**Responsibilities**:
- Ensure consistent dependency management.
- Simplify module initialization.

### Suggested Enhancements for Learning:
1. **Message Batching**: Add a feature to batch messages before sending.
2. **Retry Logic**: Implement retries on connection or message-send failure.
3. **Custom Protocols**: Extend `WriteMessage` to mimic Kafka's protocol more closely.
4. **Testing**: Use Go’s testing framework (`testing` package) for unit tests.



