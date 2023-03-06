## Interview Task: Concurrency and Signal Handling

### Overview
In this task, I was asked to modify an existing application to make it concurrent and handle interrupt signals. The application simulates an inventory system and adds inventory items to a total.

### Approach
To complete the task, I took the following approach:

1. Rewrite the app to be concurrent. I modified the AddToInventory function to use a goroutine to process the inventory items. I also used a sync.WaitGroup to ensure that all goroutines complete before the program exits.

2. Handle interrupt signals. I used the os/signal package to listen for the interrupt signal (ctrl+c) and cancel the processing of inventory items using a context. I also added a context.CancelFunc to ensure that all goroutines exit gracefully when the interrupt signal is received.

3. Test the application. I wrote two tests to verify that the application works correctly. The first test verifies that the application can add inventory items concurrently and calculate the total inventory correctly. The second test verifies that the application can handle the interrupt signal and cancel the processing of inventory items.

### Skills Demonstrated
In completing this task, I demonstrated the following skills:

- Concurrency: I used goroutines and the sync.WaitGroup to make the application concurrent and ensure that all goroutines complete before the program exits.
- Context Handling: I used the context package to handle the cancellation of inventory processing and ensure that all goroutines exit gracefully.
- Signal Handling: I used the os/signal package to listen for the interrupt signal and cancel the processing of inventory items.
- Testing: I wrote two tests to verify that the application works correctly.
