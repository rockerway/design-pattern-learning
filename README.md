# design-pattern-learning

Learning design patterns with golang, node.js, python...

## Menu

- [design-pattern-learning](#design-pattern-learning)
  - [Menu](#menu)
  - [Languages](#languages)
  - [Start up](#start-up)
    - [Go](#go)
    - [Node JS](#node-js)
    - [Python](#python)
    - [TypeScript](#typescript)
  - [Design Patterns](#design-patterns)
    - [Creational patterns](#creational-patterns)
      - [Abstract factory](#abstract-factory)
      - [Builder](#builder)
      - [Dependency injection](#dependency-injection)
      - [Factory method](#factory-method)
        - [User case](#user-case)
      - [Lazy initialization](#lazy-initialization)
      - [Multiton](#multiton)
      - [Object pool](#object-pool)
      - [Prototype](#prototype)
      - [Resource acquisition is initialization (RAII)](#resource-acquisition-is-initialization-raii)
      - [Singleton](#singleton)
    - [Structural patterns](#structural-patterns)
      - [Adapter, Wrapper, or Translator](#adapter-wrapper-or-translator)
      - [Bridge](#bridge)
      - [Composite](#composite)
      - [Decorator](#decorator)
      - [Extension object](#extension-object)
      - [Facade](#facade)
      - [Flyweight](#flyweight)
      - [Front controller](#front-controller)
      - [Marker](#marker)
      - [Module](#module)
      - [Proxy](#proxy)
      - [Twin](#twin)
    - [Behavioral patterns](#behavioral-patterns)
      - [Blackboard](#blackboard)
      - [Chain of responsibility](#chain-of-responsibility)
      - [Command](#command)
      - [Interpreter](#interpreter)
      - [Iterator](#iterator)
      - [Mediator](#mediator)
      - [Memento](#memento)
      - [Null object](#null-object)
      - [Observer or Publish/subscribe](#observer-or-publishsubscribe)
      - [Servant](#servant)
      - [Specification](#specification)
      - [State](#state)
      - [Strategy](#strategy)
      - [Template method](#template-method)
      - [Visitor](#visitor)
    - [Concurrency patterns](#concurrency-patterns)
      - [Active Object](#active-object)
      - [Balking](#balking)
      - [Binding properties](#binding-properties)
      - [Compute kernel](#compute-kernel)
      - [Double-checked locking](#double-checked-locking)
      - [Event-based asynchronous](#event-based-asynchronous)
      - [Guarded suspension](#guarded-suspension)
      - [Join](#join)
      - [Lock](#lock)
      - [Messaging design pattern (MDP)](#messaging-design-pattern-mdp)
      - [Monitor object](#monitor-object)
      - [Reactor](#reactor)
      - [Read-write lock](#read-write-lock)
      - [Scheduler](#scheduler)
      - [Thread pool](#thread-pool)
      - [Thread-specific storage](#thread-specific-storage)

## Languages

- Go
- Node.js
- Python
- TypeScript (need to learn...)

## Start up

develop code base on docker, so you need to install [Docker](https://docker) first.

### Go

1. change workdir to `go`
2. run command `./dev`

### Node JS

1. change workdir to `nodejs`
2. build docker image `docker build -t design-pattern:node .`
3. run command `./dev`

### Python

1. change workdir to `python`
2. build docker image `docker build -t design-pattern:python .`
3. run command `./dev`

### TypeScript

1. change workdir to `typescript`
2. build docker image `docker build -t design-pattern:typescript .`
3. run command `./dev`

## Design Patterns

### Creational patterns

#### Abstract factory

#### Builder

#### Dependency injection

#### Factory method

`Factory` pattern 將實體化 class 這個動作封裝成另一個 factory class，專門 new 這些 class，可以使用傳入參數決定要創建哪個實例，在 `simple factory` 使用 condition 決定要創建哪個實例，但若要添加或刪減創建 class 的種類，需要更動 condition，違反了 OO 的 Open-Closeed Principle，為了解決 `simple factory` 的耦合問題，`factory method` 將每個要創建的實例都設立一個 factory class，創建實例時只需使用相對應得 factory 來執行創建。

##### User case

使用 `factory method` 建立食物烹煮工廠，讓廚師可以製作出飲料、烤物、燉煮類食物。

#### Lazy initialization

#### Multiton

#### Object pool

#### Prototype

#### Resource acquisition is initialization (RAII)

#### Singleton

### Structural patterns

#### Adapter, Wrapper, or Translator

#### Bridge

#### Composite

#### Decorator

#### Extension object

#### Facade

#### Flyweight

#### Front controller

#### Marker

#### Module

#### Proxy

#### Twin

### Behavioral patterns

#### Blackboard

#### Chain of responsibility

#### Command

#### Interpreter

#### Iterator

#### Mediator

#### Memento

#### Null object

#### Observer or Publish/subscribe

#### Servant

#### Specification

#### State

#### Strategy

#### Template method

#### Visitor

### Concurrency patterns

#### Active Object

#### Balking

#### Binding properties

#### Compute kernel

#### Double-checked locking

#### Event-based asynchronous

#### Guarded suspension

#### Join

#### Lock

#### Messaging design pattern (MDP)

#### Monitor object

#### Reactor

#### Read-write lock

#### Scheduler

#### Thread pool

#### Thread-specific storage
