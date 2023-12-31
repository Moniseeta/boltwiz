<div align="center" style="margin-bottom: 50px">

# BoltWiZ
<img src="./docs/img/icon.png" width="200">

</div>

 Welcome to BoltWiZ, the ultimate UI tool for navigating and managing your BoltDB files with ease and precision. BoltWiZ simplifies your interaction with BoltDB, providing a user-friendly interface to perform CRUD operations more effectively.

## Key Features

- **List:** Effortlessly view all buckets or key-value pairs within the current hierarchy of your BoltDB file.
- **Search:** Quickly find child buckets or keys that match a specific substring, making data retrieval straightforward and fast.
- **Add:** Intuitively add new buckets or key-value pairs under the current bucket. At the root level, you have the ability to add new buckets.
- **Rename:** Conveniently rename the keys of pairs. (Note: Renaming of buckets is not yet supported, but we're working on it!)
- **Update:** Easily modify the value associated with a key in a pair under your current bucket.
- **Delete:** Safely remove a bucket or a key-value pair within your current hierarchy.

## Getting Started

## Installation Steps:

### Download the Executable

1. Open a terminal.

2. Download the `boltwiz` executable from Github:

> **Linux:**
> ```bash
> wget https://github.com/Moniseeta/boltwiz/releases/download/v0.0.1/boltwiz.linux
> ```
> **Mac OS:**
> ```bash
> wget https://github.com/Moniseeta/boltwiz/releases/download/v0.0.1/boltwiz.macOS
> ```

3. Rename the downloaded executable to boltwiz

> **Linux:**
> ```bash
> mv boltwiz.linux boltwiz
> ```
> **Mac OS:**
> ```bash
> mv boltwiz.macOS boltwiz
> ```

4. Ensure the downloaded file has executable permissions:

   ```bash
   chmod +x boltwiz
   ```

### Run BoltWiZ

5. Run BoltWiZ with the specified `--db-path` argument:

   ```bash
   ./boltwiz --db-path /path/to/bolt.db
   ```

   Replace `/path/to/bolt.db` with the actual path of the BoltDB database you want to open.

   **Note:** If you encounter permission issues, you may need to use `sudo` or adjust file permissions accordingly.

## Additional Options

- For more command-line options and usage details, you can refer to the help documentation:

  ```bash
  ./boltwiz --help
  ```

## Demo
<video width="100%" controls autoplay src="https://github.com/Moniseeta/boltwiz/assets/11961813/699805c4-b02a-4602-928c-6a99987c732e"></video>

## Support and Contributions

We are committed to improving BoltWiZ and welcome feedback and contributions from our user community.

- **Support:** For support, please drop a mail @ [TeamRoyRed](teamroyred@gmail.com).
- **Contributions:** If you're interested in contributing to the project, please see our [Contribution Guidelines](./CONTRIBUTION.md).

Thank you for choosing BoltWiZ for your BoltDB management needs!

Thank you for choosing our project!
