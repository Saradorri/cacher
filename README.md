# Cacher
## Overview
Cacher is a distributed caching system designed to maintain consistency across multiple nodes using a hash ring implementation with virtual nodes. It aims to efficiently distribute data among servers, allowing for easy addition or removal of nodes without significant reorganization of the data.

## Features
- Consistent hashing for even data distribution
- Dynamic node addition/removal support
- Efficient key-to-node mapping
- Scalable architecture for handling growing datasets
- Support for virtual nodes to improve load distribution
## How it Works
- Data is hashed using a consistent hashing algorithm
- The hash value is mapped to a position on the hash ring
- Nodes are represented as virtual nodes on the ring
- Keys are assigned to the nearest node based on their hash
## Key Components
- Hash Ring: Maintains the mapping between keys and nodes
- Node Manager: Handles adding/removing nodes from the system
- Cache Store: Stores actual data associated with each key
## Usage
By using Cacher, you can create a scalable and efficient caching system that maintains consistency across multiple nodes while allowing for easy expansion of your infrastructure. The addition of virtual nodes enhances load distribution and fault tolerance, making it suitable for large-scale distributed systems.

#### To use this project:

Clone the repository: git clone `https://github.com/saradorri/cacher.git`
Install dependencies: `go mod tidy`
Run the example: `go run cmd/api/main.go`
