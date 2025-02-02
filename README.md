# Mock AWS SES API (Gin + SQLite)

## Setup & Run
1. Clone the repository:
```bash
git https://github.com/saiguptha2003/Mock-AWS-SES-Api.git
cd Mock-AWS-SES-Api
```
2. Install dependencies
```bash
go mod tidy
```

3. Run the Server
```bash
go run main.go
```

4. The API will be available at:
http://localhost:8080

## API Endpoints & Usage
##### 1. POST Send Single Email
###### Url http://localhost:8080/v1/email/send

###### Discription: Send a single email
###### Request
```json
{
  "from": "admin@example.com",
  "to": "user@example.com",
  "subject": "Test Email",
  "body": "Hello, this is a test email."
}
```
###### Response
```json
{
    "email": {
        "ID": 3,
        "CreatedAt": "2025-02-02T20:41:59.6947773+05:30",
        "UpdatedAt": "2025-02-02T20:41:59.6947773+05:30",
        "DeletedAt": null,
        "to": "user@example.com",
        "from": "admin@example.com",
        "subject": "Test Email",
        "body": "Hello, this is a test email.",
        "status": "success",
        "email_type": "",
        "sent_at": "2025-02-02T20:41:59.6942862+05:30",
        "retry_count": 0,
        "latency": 0
    },
    "message": "Email processed"
}
```


##### 2. GET Email Statistics
###### Url http://localhost:8080/v1/email/statistics

###### Discription: Get total sent, success, failed, and bounced emails.
###### Response
```json
{
    "average_latency": 0,
    "bounced": 0,
    "failed": 0,
    "failure_rate": 0,
    "successful": 3,
    "top_recipients": null,
    "total_emails": 3
}
```

##### 3. GET Search Emails
###### Url http://localhost:8080/v1/email/search?from=admin@example.com&status=success"

###### Discription: Search emails based on filters.
###### Response
```json
{
    "emails": [
        {
            "ID": 1,
            "CreatedAt": "2025-02-02T20:19:20.4437197+05:30",
            "UpdatedAt": "2025-02-02T20:19:20.4437197+05:30",
            "DeletedAt": null,
            "to": "user@example.com",
            "from": "admin@mock.com",
            "subject": "Test",
            "body": "Hello",
            "status": "success",
            "email_type": "",
            "sent_at": "0001-01-01T00:00:00Z",
            "retry_count": 0,
            "latency": 0
        }
    ]
}
```

##### 4. POST Retry Failed Emails
###### Url http://localhost:8080/v1/email/retry

###### Discription: Retry sending failed emails.
###### Request
```json
{
  "from": "admin@example.com"
}

```
###### Response
```json
{
    "emails_retried": 0,
    "message": "Retries processed"
}
```


##### 1. POST Send Bulk Emails
###### Url http://localhost:8080/v1/email/send-bulk

###### Discription: Send multiple emails at once.
###### Request
```json
{
  "from": "admin@example.com",
  "to": ["user1@example.com", "user2@example.com", "user3@example.com"],
  "subject": "Bulk Email Test",
  "body": "Hello, this is a bulk test email."
}
```
###### Response
```json
{
    "failed": 0,
    "message": "Bulk email processing complete",
    "successful": 3,
    "total_latency": 0.0172353,
    "total_sent": 3
}
```
