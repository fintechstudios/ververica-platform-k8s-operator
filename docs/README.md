# Ververica Platform K8s Operator Documentation

## Table of Contents

* [Concepts](#Concepts)
    * [Design](#Design)
    * [Custom Resources](#Custom)
* [Guides](#Guides)

## Concepts

A working knowledge of the Ververica Platform concepts is a prerequisite for understanding the operator.

### Design

The [`design`](design.md) document explains how the operator controls the
Ververica Platform using the Kubernetes systems. The goal of the operator is to fully support
the core features of the Ververica Platform, not to add Flink deployment functionality.
The operator makes use of the AppManager and Platform APIs to make this possible.

### Custom Resources

The operator relies on custom Kubernetes resources to maintain state of Ververica Platform
objects. Please refer to the [mappings](mappings) for more information on the resource representations
between the platforms.  

## Guides

Guides for deploying and using the operator in conjunction with the Ververica Platform
live in the [guides directory](guides).
