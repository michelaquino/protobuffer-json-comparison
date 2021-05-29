#!/bin/sh
wrk -t100 -c100 -d10s http://localhost:8888/$endpoint
