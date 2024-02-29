## Service Name : Subscription Management System

## Features

- **Customer Management**: Manage customer information and their associated contracts.
- **Contract Management**: Handle various subscription contracts, including start and end dates, billing intervals, and statuses.
- **Subscription Management**: Define and manage different subscription services, including pricing and duration.
- **Invoice Generation**: Automatically generate and manage invoices based on contract billing intervals.

## Database Structure

### Tables

### Customers

- `ID`: Primary Key, Unique Identifier
- `Name`: String
- `Email`: String, Unique
- `CreatedAt`: DateTime
- `UpatedAt`: DateTime

### Contracts

- `ID`: Primary Key, Unique Identifier
- `CustomerID`: Foreign Key, References `Customers`
- `SubscriptionID`: Foreign Key, References `Subscriptions`
- `BillingInterval`: Integer 
- `InstallmentAmount`: Float64
- `Duration`: Integer
- `DurationUnit`: Integer
- `Status`: Integer [Enum (e.g., Active, Expired, Cancelled)]
- `ContractStartDate`: DateTime
- `ContractEndDate`: DateTime
- `CreatedAt`: DateTime
- `UpdatedAt`: DateTime


### Subscriptions

- `ID`: Primary Key, Unique Identifier
- `Name`: String
- `Code`: String, Unique
- `Description`: String
- `Price`: Decimal
- `IsActive`: Bool(true/false)
- `CreatedAt`: DateTime
- `UpdatedAt`: DateTime

### Invoices

- `ID`: Primary Key, Unique Identifier
- `CustomerID`: Foreign Key, References `Customer`
- `SubscriptionID`: Foreign Key, References `Subscription`
- `IssueDate`: DateTime
- `DueDate`: DateTime
- `Amount`: Decimal
- `IsSendToCustomer`: Bool(true/false)

## Relationships

- A `Customer` can have multiple `Contracts`.
- Each `Contract` is associated with one `Subscription`.
- `Invoices` are generated for each `Contract` based on the `BillingInterval`.

## Installation

Clone the repository:
   ```bash
   git clone https://github.com/shadhin-int/Subscription-Management-System.git
   ```
Navigate to directory
```
cd Subscription-Management-System
```

Install Dependency
```
go mod tidy
```
Configure DB
```
Change the db configuration from .env file
```

Database Migration
```
go run migrations/migrate.go
```
Load Test Data
```
go run migrations/load_data.go
```

Run Project
```
go run main.go
```

### Note:
```
By Default the cron job running in 0.0.0.0 time
```

# Further Improvements

## Introduction of Kafka

Introduce Kafka to the system to facilitate event-driven architecture for invoice generation and notification processing. Kafka will serve as the messaging backbone to produce data for invoices whose bills are due and to consume notifications for invoice generation.

## Produce Data for Invoice Generation

Implement Kafka producers to produce data for invoices whose bills are due. These producers will publish messages containing contract information to Kafka topics whenever a bill is due.

## Consume Data for Invoice Generation

Develop Kafka consumers to consume data from Kafka topics containing contract information for invoice generation. These consumers will retrieve contract details, generate invoices, and send them to users.

## Notification Server

Implement a notification server to handle the processing of invoices. The notification server will consume invoice data from Kafka topics, generate invoices, and send them to users via email or other communication channels.

## Retry Mechanism for Failed Invoices

Introduce a retry mechanism for failed invoices to ensure robustness and reliability in invoice generation and delivery. Failed invoices should be retried automatically with backoff and retry strategies until successfully processed.

## Monitoring and Error Handling

Implement monitoring and error handling mechanisms to track the status of invoice generation and delivery processes. Use logging and monitoring tools to identify and troubleshoot errors, ensuring smooth operation of the system.


## Conclusion

Introducing Kafka to produce data for invoice generation and implementing a notification server will enhance the system's scalability, reliability, and performance. By implementing retry mechanisms, monitoring tools, and security measures, the system will be well-equipped to handle invoice generation and delivery efficiently and reliably.
