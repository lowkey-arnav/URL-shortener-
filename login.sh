#!/bin/bash
source dev.env
export TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$TEST_EMAIL\",\"password\":\"$TEST_PASS\"}" \
  | sed -E 's/.*"access_token":"([^"]+)".*/\1/')
