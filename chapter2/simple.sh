#!/usr/bin/env bash
go build -o micro -ldflags "-X main.buildName=micro -X main.buildGitRevision=f8c315083e7b739f0f055ee46a747c8e109d7539-dirty -X main.buildStatus=Modified -X main.buildUser=`whoami` -X main.buildHost=`hostname -f` -X main.buildTime=`date +%Y-%m-%d--%T`"
