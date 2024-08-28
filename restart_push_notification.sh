#!/bin/bash

# Define the port to check
PORT=9090

# Find the process ID (PID) of the service running on the specified port
PID=$(lsof -t -i :$PORT)

# Check if a PID was found
if [ -z "$PID" ]; then
    echo "No service running on port $PORT."
else
    # Kill the process
    kill -9 $PID
    echo "Service running on port $PORT has been killed."
    # Add a small delay to ensure the port is fully released
    sleep 3
fi

# Start the application in the background with nohup
echo "Starting pushNotifications application..."
nohup ./pushNotifications > nohup.out 2>&1 &

# Check if the application started successfully
sleep 3
if pgrep -f "pushNotifications" > /dev/null; then
    echo "pushNotifications application started successfully."
else
    echo "Failed to start pushNotifications application."
fi
