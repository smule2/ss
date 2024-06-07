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
fi

# Start the application in the background with nohup
nohup ./pushNotifications &
