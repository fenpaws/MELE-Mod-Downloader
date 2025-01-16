
# MELE Mod Downloader (GO:LE)

This project aims to bridge the gap that exists in the Mass Effect Legendary Edition modding scene.
This tool helps to download large modpacks and individual mods via the direct download function of Nexus Mods.

The ME3Tweaks Mod Manager can handle this, but it's not possible on Linux since it runs in WINE.
This app bridges the auto-download feature to quickly download those mods.

## This project is in its earliest stage. Nothing will work yet. When a stable version is ready, I will announce it and also create an automatic build for it.

## Features

- Download large modpacks and individual mods from Nexus Mods.
- Compatible with Linux environments.
- Supports direct download links for quick access to mods.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/fenpaws/MELE-Mod-Downloader.git
    cd MELE-Mod-Downloader
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Build the application:
    ```sh
    go build -o mele-mod-downloader
    ```

4. Run the application:
    ```sh
    ./mele-mod-downloader
    ```

## Usage

1. Ensure you have a valid Nexus Mods API key.
2. Configure the application with your API key and desired download settings.
3. Use the provided interface to select and download mods or modpacks.

### Using GO:LE

GO:LE operates in two modes: `consumer` and `run`.

- The `consumer` mode sends the Nexus Mods callback URI to GO:LE. For it to function correctly, you need a desktop
  entry. When you click the link in Nexus Mods, it should associate with this desktop file.

- The `run` command listens to what the `consumer` sends and processes it accordingly.

### Consumer

Create a `gole-consume.desktop` entry with the following content. Please adjust the `Path` and
`Exec` specifications as needed:

> This is a Development .desktop file

```
[Desktop Entry]
Name=GO:LE - Consumer
Path=~/GolandProjects/MELE-Mod-Downloader
Exec=go run cmd/main.go consume %u
Terminal=true
Type=Application
MimeType=x-scheme-handler/nxm;
```

### Runner

> This is a Development .desktop file

```
[Desktop Entry]
Name=GO:LE - Mass Effect:LE mod downloader
Path=~/GolandProjects/MELE-Mod-Downloader
Exec=go run main.go run %u
Terminal=true
Type=Application
MimeType=x-scheme-handler/nxm;

```

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](./LICENCE) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## Contact

For any questions or issues, please open an issue on GitHub or contact the project maintainers.

---

**Note:** This project is not affiliated with BioWare, Electronic Arts, or Nexus Mods. All trademarks and registered
trademarks are the property of their respective owners.
