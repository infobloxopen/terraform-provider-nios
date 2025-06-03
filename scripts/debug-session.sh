#!/bin/bash

# Start the provider in debug mode
make debug-run &
PROVIDER_PID=$!

# Wait for debugger to start
sleep 2

# Export the reattach configuration
export TF_REATTACH_PROVIDERS="{\"registry.terraform.io/infoblox-cto/nios\":{\"Protocol\":\"grpc\",\"Pid\":$PROVIDER_PID,\"Test\":true,\"Addr\":{\"Network\":\"unix\",\"String\":\"/tmp/provider.sock\"}}}"

# Print instructions
echo "Provider started in debug mode with PID: $PROVIDER_PID"
echo "Debugger listening on :2345"
echo ""
echo "TF_REATTACH_PROVIDERS has been exported. You can now:"
echo "1. Start VS Code debugger"
echo "2. Run 'terraform apply' in another terminal"
echo ""
echo "To clean up, press Ctrl+C"

# Wait for Ctrl+C
wait $PROVIDER_PID
