#!/bin/sh
ORIGINAL_OWNER=$(stat -c '%U:%G' /generated)
sudo chown -R yeoman:yeoman /generated
yo --no-insight tokend-module
sudo chown -R $ORIGINAL_OWNER /generated
