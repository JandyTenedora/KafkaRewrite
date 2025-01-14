# Kafka Rewrite in Go

This project is a simple Kafka Rewrite implemented in Go. It provides a foundational understanding of Kafka producers and how to handle basic message sending operations. It also implements a barebones Kafka broker to further simulate Kafka functionality.

---

## Project Structure

```
kafka-rewrite/
├── main.go
├── broker/
│   ├── broker.go
│   ├── storage.go
│   └── ack.go
├── producer/
│   ├── producer.go
│   ├── config.go
│   └── connection.go
├── utils/
│   ├── hash.go
│   └── logger.go
└── go.mod
```

---

## Kafka_Rewrite

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

### **2. `go.mod`**

**Purpose**: Manages dependencies and specifies the module name.

**Fields**:
- Module name (`module kafka-producer`).
- Go version (`go 1.20`).

**Responsibilities**:
- Ensure consistent dependency management.
- Simplify module initialization.

---

## Broker

### **1. `broker/broker.go`**

**Purpose**: Core logic for managing producer connections, message routing, and topic-based communication.

**Structs**:
- `Broker`:
  - Fields:
    - `Address` (string): Address where the broker listens for connections.
    - `Topics` (map[string][]byte): Stores messages organized by topic.
    - `Connections` ([]net.Conn): Tracks active producer connections.

**Methods**:
1. `NewBroker(address string) *Broker`:
   - Initializes a new broker instance.
   - Sets up storage for topics and prepares to accept connections.

2. `Start() error`:
   - Starts the broker by listening for incoming producer connections on the specified address.
   - Handles each connection in a separate goroutine.

3. `HandleConnection(conn net.Conn)`:
   - Reads messages from the producer connection.
   - Parses the message into a topic and message body.
   - Stores the message in the appropriate topic.

4. `Stop()`:
   - Gracefully shuts down the broker.
   - Closes all active connections.

---

### **2. `broker/storage.go`**

**Purpose**: Handles storage and retrieval of messages for topics.

**Structs**:
- `TopicStorage`:
  - Fields:
    - `Messages` (map[string][][]byte): A map of topic names to lists of messages.

**Methods**:
1. `NewTopicStorage() *TopicStorage`:
   - Initializes a new in-memory storage system for topics.

2. `AddMessage(topic string, message []byte)`:
   - Appends a message to the specified topic.

3. `GetMessages(topic string) [][]byte`:
   - Retrieves all messages for a specific topic.

4. `ListTopics() []string`:
   - Returns a list of all topic names in the storage.

---

### **3. `broker/ack.go`**

**Purpose**: Handles acknowledgment messages to confirm receipt of producer messages.

**Functions**:
1. `SendAck(conn net.Conn)`:
   - Sends an acknowledgment message back to the producer after successfully receiving and storing a message.

2. `ReceiveAck(conn net.Conn) error`:
   - Waits for an acknowledgment from the producer or broker.
   - Returns an error if no acknowledgment is received within a timeout.

--- 

## Producer

### **1. `producer/producer.go`**

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

### **2. `producer/config.go`**

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

### **3. `producer/connection.go`**

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

## Utils

### **1. `utils/hash.go`**

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


### **2. `utils/logger.go`**

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

### Future Enhancements:
1. **Message Batching**: Add a feature to batch messages before sending.
2. **Retry Logic**: Implement retries on connection or message-send failure.
3. **Custom Protocols**: Extend `WriteMessage` to mimic Kafka's protocol more closely.
4. **Testing**: Use Go’s testing framework (`testing` package) for unit tests.

