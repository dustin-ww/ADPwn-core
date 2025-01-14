<a id="readme-top"></a>
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/ADPwnLogo.png" alt="Logo" width="150" height="190">
  </a>

<h1 align="center">ADPwn</h1>

  <p align="center">
    <br />
    <a href="https://github.com/dustin-ww/ADPwn/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/dustin-ww/ADPwn/issues/new?labels=bug">Report Bug</a>
    ·
    <a href="https://github.com/dustin-ww/ADPwn/issues/new?labels=enhancement">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


A console-driven framework is being developed to automate the detection and testing of known attack vectors in Active Directory environments, including vulnerabilities such as PrintNightmare, Zerologon, and EternalBlue. The framework aims to streamline the enumeration and exploitation of domains by integrating existing tools like nmap and NetExec. Built in GoLang, the solution features a connection to a Dgraph database for persistent data storage and provides a visual representation of the results, enabling users to analyze and understand the discovered vulnerabilities effectively.

This project is currently only supported on Linux systems and is maintained in the dev branch.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With


* [![Go][GoLang]][Go-url]
* [Dgraph][Dgraph-url]
* [tview for golang][Tview-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

First of all you need following dependencies on you system:

- docker or podman
- golang 

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/dustin-ww/ADPwn
   ```

2. Build and run the Dgraph container in docker or podman via docker-compose.yml in the main directory
    ```sh
       docker-compose run -d
    ```
3. Navigate to /src/cmd folder and run the project locally

    ```sh
       go run .
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

ADPwn can be used through the terminal interface. The usage should be very intuitive. Detailed examples can be found later in the wiki.
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [x] Add UI
- [x] Add Graph Database
- [ ] Add Wiki
- [ ] Add support for adding
  - [x] Users
  - [x] Targets
- [ ] Add enumeration
    - [x] Host & service scan with nmap
    - [ ] Check for guest account & anonymous login
    - [ ] Check for shares & files
    - [ ] Printer nightmare enumeration
    - [ ] Zerologon enumeration
    - [ ] Eternal blue scanning
    - [ ] Enumerate SYSVOL & GPP
- [ ] Add support for attack vectors
    - [ ] Printer nightmare
    - [ ] Zerologon
    - [ ] Support for coercing
    - [ ] Automatic privilege escalation via DACL abuse

    
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

This project is maintained on the dev branch. You can find the latest working release on the main branch.
Previous releases can be found in GitHubs-Releases in the near future, as implementing versioning is one of the current main goals.
<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Personal website: https://dw-security.de/

Project Link: [https://github.com/dustin-ww/ADPwn](https://github.com/dustin-ww/ADPwn)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->

[GoLang]: https://img.shields.io/badge/golang-00ADD8?&style=plastic&logo=go&logoColor=white
[Go-url]: https://golang.org/
[Dgraph.io]: https://dgraph.io/
[Dgraph-url]: https://dgraph.io/
[Tview-url]: https://github.com/rivo/tview
