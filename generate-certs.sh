#!/bin/bash

certstrap --depot-path=valid init --common-name=YoloSwagginsCA --passphrase=""
certstrap --depot-path=valid request-cert --common-name=server --ip=127.0.0.1 --passphrase=""
certstrap --depot-path=valid sign server --CA=YoloSwagginsCA
certstrap --depot-path=valid request-cert --common-name=client --ip=127.0.0.1 --passphrase=""
certstrap --depot-path=valid sign client --CA=YoloSwagginsCA

certstrap --depot-path=invalid init --common-name=RandomCA --passphrase=""
certstrap --depot-path=invalid request-cert --common-name=client-bad --ip=127.0.0.1 --passphrase=""
certstrap --depot-path=invalid sign client-bad --CA=RandomCA
