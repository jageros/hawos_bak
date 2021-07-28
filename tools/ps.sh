#!/bin/bash
ps aux|grep ./builder | grep -v grep
ps aux|grep ./builder | grep -v grep | wc -l