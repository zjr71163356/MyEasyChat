#!/bin/bash

# Start etcd
echo "Starting etcd..."
etcd &

# Wait for etcd to start
sleep 5

# Switch to ~/MyEasyChat/user/rpc and run the Go application
echo "Switching to ~/MyEasyChat/user/rpc and running Go application..."
cd ~/MyEasyChat/user/rpc || exit
go run . &

# Switch to ~/MyEasyChat/user/api and run the Go application
echo "Switching to ~/MyEasyChat/user/api and running Go application..."
cd ~/MyEasyChat/user/api || exit
go run .