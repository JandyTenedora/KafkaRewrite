# **Kafka-Go Reimplementation**

This project is a simplified reimplementation of Apache Kafka in Go. This project serves as an opportunity to practice Go while familiarising myself with Kafka internals. The implementation includes core components such as brokers, producers, consumers, and networking.

---

## **Project Structure**

The repository is organized into directories and files for each major Kafka component. Below is the directory structure with detailed explanations of each file and its purpose.

```
kafka-go-reimplementation/
├── cmd/
│   ├── broker/
│   │   └── main.go
│   └── producer/
│       └── main.go
├── pkg/
│   ├── broker/
│   │   ├── broker.go
│   │   ├── partition.go
│   │   └── replication.go
│   ├── producer/
│   │   ├── producer.go
│   │   ├── encoding.go
│   │   └── error.go
│   ├── consumer/
│   │   ├── consumer.go
│   │   ├── decoding.go
│   │   └── offset.go
│   ├── network/
│   │   ├── broker_server.go
│   │   ├── produce_handler.go
│   │   └── fetch_handler.go
│   └── utils/
│       ├── log_compaction.go
│       └── compression.go
├── internal/
│   ├── message/
│   │   └── message.go
│   ├── offset_manager/
│   │   └── offset_manager.go
│   ├── errors/
│   │   └── errors.go
└── go.mod
```

---

## **Detailed File and Function Overview**

### **1. `cmd/` Directory**
This directory contains the entry points for the broker and producer applications.

- **`cmd/broker/main.go`**
  - Entry point for the broker.
  - Starts the broker server, listens for connections, and routes requests.

- **`cmd/producer/main.go`**
  - Entry point for the producer.
  - Initializes the producer, prepares messages, and sends them to brokers.

---

### **2. `pkg/` Directory**
This directory contains the core components of the Kafka-like system.

#### **`pkg/broker/`**
- **`broker.go`**
  - Defines the `Broker` struct and its core methods:
    - `func HandleProduceRequest(msg Message) error`: Handles produce requests by adding messages to partitions.
    - `func HandleFetchRequest(topic string, partition int, offset int) ([]Message, error)`: Fetches messages from partitions.
  
- **`partition.go`**
  - Manages partitions and leader election:
    - `func AssignPartition(topic string) *Partition`: Assigns a partition to a topic.
    - `func ElectPartitionLeader(topic string, partition int) error`: Elects a leader for a partition.

- **`replication.go`**
  - Handles message replication and failover:
    - `func ReplicateToFollowers(topic string, partition int, msg Message) error`: Replicates messages to follower brokers.

#### **`pkg/producer/`**
- **`producer.go`**
  - Defines the `Producer` struct and its methods:
    - `func Produce(topic string, msg Message) error`: Sends a message to a broker.

- **`encoding.go`**
  - Encodes messages for transmission:
    - `func EncodeMessage(msg Message) []byte`: Serializes a message into binary format.

- **`error.go`**
  - Handles producer-specific errors:
    - `func HandleProducerError(err error)`: Logs or handles production errors.

#### **`pkg/consumer/`**
- **`consumer.go`**
  - Defines the `Consumer` struct and its methods:
    - `func Consume() ([]Message, error)`: Fetches messages from the broker.

- **`decoding.go`**
  - Decodes messages received from brokers:
    - `func DecodeMessage(data []byte) (Message, error)`: Deserializes binary data into a message.

- **`offset.go`**
  - Manages offsets for consumers:
    - `func UpdateOffset(offset int)`: Updates the consumer’s offset to track consumed messages.

#### **`pkg/network/`**
- **`broker_server.go`**
  - Starts the broker server and handles connections:
    - `func StartBrokerServer(address string) error`: Sets up a TCP/HTTP server for the broker.

- **`produce_handler.go`**
  - Handles incoming produce requests from producers:
    - `func HandleProduceConnection(conn net.Conn)`: Processes produce requests.

- **`fetch_handler.go`**
  - Handles incoming fetch requests from consumers:
    - `func HandleFetchConnection(conn net.Conn)`: Processes fetch requests.

#### **`pkg/utils/`**
- **`log_compaction.go`**
  - Implements log compaction for topics:
    - `func CompactLog(topic string, partition int)`: Removes older messages.

- **`compression.go`**
  - Handles message compression:
    - `func CompressMessage(msg Message) ([]byte, error)`: Compresses a message.

---

### **3. `internal/` Directory**
Contains internal utilities and definitions that are not exposed externally.

- **`internal/message/message.go`**
  - Defines the `Message` struct:
    - Fields: `Key`, `Value`, `Headers`, etc.

- **`internal/offset_manager/offset_manager.go`**
  - Implements offset management for consumers:
    - `func ManageOffset(consumerID string, offset int)`: Tracks offsets.

- **`internal/errors/errors.go`**
  - Provides centralized error handling utilities:
    - `func NewError(message string) error`: Creates a custom error.

---

### **How to Run the Project**

#### **Run the Broker**
To start the broker server, navigate to the project root and run:
```bash
go run cmd/broker/main.go
```

#### **Run the Producer**
To start the producer, navigate to the project root and run:
```bash
go run cmd/producer/main.go
```

---

### **Future Enhancements**
1. **Log Compaction**: Optimize storage by removing old data.
2. **Consumer Groups**: Implement group coordination for load balancing.
3. **Fault Tolerance**: Simulate broker crashes and recoveries with leader failover.
4. **Message Compression**: Add GZIP or Snappy for improved performance.

