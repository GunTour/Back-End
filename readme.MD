<a name="readme-top"></a>



[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/GunTour/Back-End">
    <img src="readme/logo.png" alt="Logo" width="200px">
  </a>

<h3 align="center">GunTour</h3>

  <p align="center">
    This project is a website for mountaineers who want to book tickets and rental goods for climbing.
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
      <a href="#features">Features</a>
      <ul>
        <li><a href="#entity-relationship-diagram">Entity Relationship Diagram</a></li>
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
## Features

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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Entity Relationship Diagram

[![GunTour-ERD][erd-screenshot]](https://github.com/GunTour/Back-End/blob/main/readme/erd.jpg)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

<div>
    <a href="https://go.dev/">
    <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" title="Golang" alt="Golang" width="60"/>
    </a>&nbsp;
    <a href="https://code.visualstudio.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/9/9a/Visual_Studio_Code_1.35_icon.svg" title="Visual Studio Code" alt="Visual Studio Code" width="40"/>
    </a>&nbsp;
    <a href="https://developers.google.com/apis-explorer">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/24px-Google_%22G%22_Logo.svg.png" title="Google" alt="Google" width="40"/></a>&nbsp;
    <a href="https://aws.amazon.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/9/93/Amazon_Web_Services_Logo.svg" title="AWS" alt="AWS" width="50"/></a>&nbsp;
    <a href="https://www.mysql.com/">
    <img src="https://1000logos.net/wp-content/uploads/2020/08/MySQL-Logo.png" title="MySQL" alt="MySQL" width="60"/></a>&nbsp;
    <a href="https://hub.docker.com/">
    <img src="https://www.docker.com/wp-content/uploads/2022/03/vertical-logo-monochromatic.png" title="Docker" alt="Docker" width="50"/></a>&nbsp;
    <a href="https://swagger.io/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/a/ab/Swagger-logo.png" title="Swagger" alt="Swagger" width="40"/></a>&nbsp;
    <a href="https://www.postman.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/c/c2/Postman_%28software%29.png" title="Postman" alt="Postman" width="120"/></a>&nbsp;
    <a href="https://www.cloudflare.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/4/4b/Cloudflare_Logo.svg" title="Cloudflare" alt="Cloudflare" width="120"/></a>&nbsp;
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/GunTour/Back-End
   ```
2. Get env at [config.env](https://drive.google.com/file/d/18dROVFu708Lsr2lgpA7PtkRxEybSmwJ7/view?usp=sharing)
3. Enter your config in `config.env`
   ```bash
   AWS_REGION = "ENTER YOUR AWS REGION"
   AWS_ACCESS_KEY_ID = "ENTER YOUR AWS ACCESS KEY ID"
   AWS_SECRET_ACCESS_KEY = "ENTER YOUR AWS SECRET KEY ID"
   MIDTRANS_SERVER = "ENTER YOUR MIDTRANS SERVER"
   MIDTRANS_CLIENT = "ENTER YOUR MIDTRANS CLIENT"
   CLIENT_ID = "ENTER YOUR GOOGLE OAUTH CLIENT ID"
   CLIENT_SECRET = "ENTER YOUR GOOGLE OAUTH CLIENT SECRET"
   ```
4. Run project
   ```bash
   cd Back-End
   paste your config.env
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Documentation

_For more examples, please refer to the [OPEN API](https://app.swaggerhub.com/apis-docs/khalidrianda/GunTour/1.0.0#/)_

[![GunTour-API][product-screenshot]](https://github.com/GunTour/Back-End/blob/main/readme/GunTour.gif)

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

[![Khalid][khalid]][khalid-url]
- [![LinkedIn-Khalid][linkedin-shield]][linkedin-url-1] [![Email-Khalid][email-shield]][email-1]

[![Dany][dany]][dany-url]
- [![LinkedIn-Dany][linkedin-shield]][linkedin-url-2] [![Email-Dany][email-shield]][email-2]

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
[license-url]: https://github.com/GunTour/Back-End/blob/main/LICENSE.txt
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
[khalid]: https://img.shields.io/badge/-Khalid-black.svg?style=for-the-badge&logo=Khalid&colorB=555
[dany]: https://img.shields.io/badge/-Dany-black.svg?style=for-the-badge&logo=Dany&colorB=555
[khalid-url]: https://github.com/khalidrianda
[dany-url]: https://github.com/mdanys
[email-shield]: https://img.shields.io/badge/gmail-DD0031?style=for-the-badge&logo=gmail&logoColor=white
[email-1]: khalidrianda12@gmail.com
[email-2]: mochammaddany@gmail.com
