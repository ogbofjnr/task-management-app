Assumption I made based on task specification:

- "task can't overlap with other tasks" means user can't have multiple task at `in_progress` status.
- "task can occupy some certain amount of time" means that use should have ability to track time and get statistics
- "reminder period" user can set reminder for the task and target status. For example set reminder at +10 days and `todo` status 
would mean that if after 10 days the task would still be in `todo` status, user would receive the notification.
- "on-site" notifications implemented with websockets. In production, I would rather create standalone service and use message broker and socket.io since it has long-poll fallback. 


I used some boilerplate, but still it took about 12 hours to complete the task.
The task turned out to be quit big, I didn't implement some the features like tests, detailed validation, some endpoints etc

project init 
```
docker-compose up -d
./migrate -path migrations -database "postgresql://root:root@127.0.0.1:5432/pm?sslmode=disable" -verbose up
```






