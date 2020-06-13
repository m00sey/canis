# CRUX

Crux is an extensible credentialing platform providing a configurable, easy to use environment for issuing and verifying credentials that conform with decentralized 
identity standards including [W3C decentralized identifiers](https://w3c.github.io/did-core/) (DIDs), [W3C DID resolution](https://w3c-ccg.github.io/did-resolution/), and [W3C verifiable credentials](https://w3c.github.io/vc-data-model/).

## Summary

- [**Why Crux?**](#why-crux)
- [**Features**](#features)
- [**Distributions**](#distributions)
- [**Development**](#development)
- [**License**](#license)

## Why Crux?

Issuing digital credentials requires an institution or organization to have an agent to represent its interest in the digital landscape.
This agent must act as a fiduciary on behalf of the organization, must hold cryptographic keys representing its delegated authority, and it must communicate
via [DIDComm Protocols](https://github.com/hyperledger/indy-hipe/pull/69).  

In addition needing an agent, organizations need a way to instruct this agent how to issue a credential and to whom.  That requires information that is currently stored 
in legacy (in ToIP terms) systems.

Crux serves as a platform for creating, launching and empowering agents to participate in a credentialing ecosystem on any organization's behalf.  In addition,
Crux provides an easy to use RESTful API and extensible data model to allow for endorsing agents on behalf of any hierarchy of organizational structure.

## Features
