#!/bin/sh
wrk -s ./wrk/post.lua -t100 -c100 -d10s http://localhost:8888/$endpoint
