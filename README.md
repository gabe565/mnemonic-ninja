<p align="center">
    <img src="/public/images/logo.png?raw=true" title="Mnemonic Ninja Logo">
</p>

# Mnemonic Major Converter
<p align="center">
    <a href="https://travis-ci.org/gabe565/mnemonic-major-converter"><img src="https://img.shields.io/travis/gabe565/mnemonic-major-converter.svg" alt="Build Status"></a>
    <a href="https://mnemonic.ninja"><img src="https://img.shields.io/website-up-down-green-red/https/mnemonic.ninja.svg?label=site%20status" alt="Site Status"></a>
    <a href="https://github.com/gabe565/mnemonic-major-converter/blob/master/LICENSE"><img src="https://img.shields.io/github/license/gabe565/mnemonic-major-converter.svg" alt="License"></a>
</p>

Website to convert between a number and its corresponding word to aid in memorization.     
Live at <https://mnemonic.ninja>

## Requirements
  - PHP >= 7.0.0
  - OpenSSL PHP Extension
  - PDO PHP Extension
  - Mbstring PHP Extension
  - MySQL, PostgreSQL, SQLite, or SQL Server

## Installation
This project is built using [Laravel](https://laravel.com) as an API and [Vue.js](https://vuejs.org) as a front-end controller.

  1. To obtain the code, simply run

      ```
      git clone --recursive https://github.com/gabe565/mnemonic-major-converter
      cd mnemonic-major-converter
      ```

  2. Then, copy `.env.example` to `.env` and edit the file to set environment variables.     
     An application key must be generated by running

      ```
      php artisan key:generate
      ```
  3. To setup the database tables, run

      ```
      php artisan migrate
      ```

  4. Finally, to populate the database with conversion data, run

      ```
      php artisan db:seed
      ```
