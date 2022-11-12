<a name="GunTour Apps"></a>



[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn-Khalid][linkedin-shield]][linkedin-url-1]
[![LinkedIn-Dany][linkedin-shield]][linkedin-url-2]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/GunTour/Back-End">
    <img src="readme/logo.png" alt="Logo" width="100" height="80">
  </a>

<h3 align="center">GunTour</h3>

  <p align="center">
    Climbing Mountains Is Now Easier. Explore Mount Pangrango simply by registering and exploring.
    <br />
    <a href="https://github.com/GunTour/Back-End"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/GunTour/Back-End">View Demo</a>
    ·
    <a href="https://github.com/GunTour/Back-End/issues">Report Bug</a>
    ·
    <a href="https://github.com/GunTour/Back-End/issues">Request Feature</a>
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This application can be used by users as a place for climbers to order tickets to climb the mountain and can rent the tools needed for climbing and use the services of a ranger as a climbing guide. This application has been integrated with online payment merchants
such as virtual accounts, bank transfers, debit cards, and e-wallets.

The main features that this application has are:
- Users
    - Users can register
    - Users can log in
    - Users can apply to become a ranger
- Booking
    - Users can book tickets on the desired schedule
    - Users can rent climbing products/equipment and hire ranger services
    - Users can make online payments
    - Users can see the results of bookings
- Admins
    - Admin can see users who have booked tickers
    - Admin can see a list of rangers and a list of users who have volunteered to become rangers
    - Admin can add, edit, and delete products for rent

Entity Relationship Diagram

[![GunTour-ERD][erd-screenshot]](https://github.com/GunTour/Back-End/tree/main/readme/erd.jpg)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Golang][Go]][go-url]
* [![Framework][Echo]][echo-url]
* [![Oauth][Oauth]][oauth-url]
* [![Gmail][Gmail]][mail-url]
* [![Calendar][Calendar]][calendar-url]
* [![Server][AWS]][aws-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/GunTour/Back-End
   ```
2. Run project
   ```bash
   cd Back-End
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

_For more examples, please refer to the [OPEN API](https://app.swaggerhub.com/apis-docs/khalidrianda/GunTour/1.0.0#/)_

[![GunTour-API][product-screenshot]](https://github.com/GunTour/Back-End/tree/main/readme/GunTour.gif)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] User
    - [x] Register
    - [x] Login
- [x] Ranger
    - [x] Register as Ranger
- [x] Booking
    - [x] Reserve to climb
    - [x] Book items
    - [x] Book ranger

See the [open issues](https://github.com/GunTour/Back-End/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Khalid - [Khalid](https://www.linkedin.com/in/khalidrianda) - khalidrianda12@gmail.com

Dany - [Dany](https://www.linkedin.com/in/mochammaddany) - mochammaddany@gmail.com

Project Link: [https://github.com/GunTour/Back-End](https://github.com/GunTour/Back-End)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/GunTour/Back-End.svg?style=for-the-badge
[contributors-url]: https://github.com/GunTour/Back-End/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/GunTour/Back-End.svg?style=for-the-badge
[forks-url]: https://github.com/GunTour/Back-End/network/members
[stars-shield]: https://img.shields.io/github/stars/GunTour/Back-End.svg?style=for-the-badge
[stars-url]: https://github.com/GunTour/Back-End/stargazers
[issues-shield]: https://img.shields.io/github/issues/GunTour/Back-End.svg?style=for-the-badge
[issues-url]: https://github.com/GunTour/Back-End/issues
[license-shield]: https://img.shields.io/github/license/GunTour/Back-End.svg?style=for-the-badge
[license-url]: https://github.com/GunTour/Back-End/blob/readme/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url-1]: https://linkedin.com/in/khalidrianda
[linkedin-url-2]: https://linkedin.com/in/mochammaddany
[product-screenshot]: readme/GunTour.gif
[erd-screenshot]: readme/erd.jpg
[Go]: https://img.shields.io/github/go-mod/go-version/GunTour/Back-End
[go-url]: https://go.dev/
[Echo]: https://img.shields.io/badge/Echo-v4-9cf
[echo-url]: https://echo.labstack.com/
[Oauth]: https://img.shields.io/badge/OAuth-Google-informational
[oauth-url]: https://developers.google.com/identity/protocols/oauth2
[Gmail]: https://img.shields.io/badge/Gmail-Google-informational
[mail-url]: https://github.com/googleapis/google-api-go-client
[Calendar]: https://img.shields.io/badge/Calender-Google-informational
[calendar-url]: https://github.com/googleapis/google-api-go-client
[AWS]: https://img.shields.io/badge/AWS-EC2-orange
[aws-url]: https://aws.amazon.com/
