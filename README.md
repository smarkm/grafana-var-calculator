# Grafana-Var-Calculator

This Grafana data source plugin allows users to dynamically process queries and return results that can be used as variables in Grafana dashboards. It supports basic mathematical calculations, string processing, range generation, and list generation, providing a flexible way to define and manipulate data.

---

## Features

- ​**Mathematical Calculations**: Perform basic arithmetic operations (e.g., `1 + 2 * 3`).
- ​**String Processing**: Manipulate strings with functions like `toUpperCase`, `toLowerCase`, `substring`, and `replace`.
- ​**Range Generation**: Generate a range of numbers (e.g., `range(1, 5)`).
- ​**List Generation**: Define and return lists of values (e.g., `["a", "b", "c"]`).
- ​**Error Handling**: Provides clear error messages for invalid queries.
- ​**Safe Expression Parsing**: Uses the `expr-eval` library to safely parse and evaluate queries.

---

## Installation

1. Clone this repository or download the plugin files.
2. Place the plugin in the Grafana plugins directory (e.g., `/var/lib/grafana/plugins`).
3. Restart the Grafana server.
4. In Grafana, go to ​**Configuration > Data Sources** and add a new data source using this plugin.

---

### Using Variables in Grafana

1. Define a variable in your Grafana dashboard (e.g., `myVariable`).
2. Set the variable type to ​**Query**.
3. Use the plugin as the data source and enter a query (e.g., `range(1, 5)`).
4. Use the variable in your queries or panels (e.g., `SELECT * FROM metrics WHERE id = '$myVariable'`).

---

## Example Queries

| Query                  | Result                                                                 |
|------------------------|------------------------------------------------------------------------|
| `1 + 2 * 3`            | `[{ "text": "7", "value": "7" }]`                                      |
| `"hello".toUpperCase()`| `[{ "text": "HELLO", "value": "HELLO" }]`                              |
| `range(1, 3)`          | `[{ "text": "1", "value": "1" }, { "text": "2", "value": "2" }, ... ]` |
| `["a", "b", "c"]`      | `[{ "text": "a", "value": "a" }, { "text": "b", "value": "b" }, ... ]` |

---

## Error Handling

If a query is invalid, the plugin will return an empty array and log an error message. Examples of invalid queries include:

- Unsupported syntax (e.g., `1 + * 2`).
- Missing or incorrect function names (e.g., `rang(1, 5)`).
- Invalid string operations (e.g., `"hello".toUpperCase`).

---

## Security

The plugin uses the `expr-eval` library to safely parse and evaluate queries, avoiding the risks associated with `eval`. However, ensure that only trusted users have access to modify queries in production environments.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with clear and descriptive messages.
4. Submit a pull request, explaining the changes and their purpose.

For major changes, please open an issue first to discuss the proposed changes.

---

## License

This plugin is licensed under the [MIT License](LICENSE).

---

## Support

For support or questions, please open an issue in this repository or contact the maintainers.

---

Enjoy using the Query Processor plugin to enhance your Grafana dashboards! 🚀