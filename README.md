# Qryptic Gateway

The **Qryptic Gateway** enforces access controls using WireGuard and dynamically updates its configuration based on instructions from the Controller. It provides secure connectivity to internal networks and resources.

## Key Features
- WireGuard-powered secure connectivity.
- Dynamic peer management (add/remove users).
- REST API for communication with the Controller.
- Self-hosted deployment using Docker.

## Quick Start
1. Deploy the Gateway on a public EC2 instance.
2. Add the Gateway in the Controller UI.
3. Run the Gateway using the Docker command provided by the Controller.

For detailed setup and configuration instructions, visit the [Qryptic Gateway Documentation](https://docs.qryptic.com/).

---

## License
Qryptic is licensed under the [AGPL v3.0 License](./LICENSE).
