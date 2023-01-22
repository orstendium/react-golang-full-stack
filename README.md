# react-golang-full-stack

This is a custom boilerplate to build a full stack web application using React, Golang and Webpack. It is also configured with webpack-dev-server, eslint, prettier and babel. It is pretty much entirely based upon [simple-react-full-stack](https://github.com/crsandeep/simple-react-full-stack) so check it out. I've just converted it from a npmjs/express backend to Golang.

- [react-golang-full-stack](#react-golang-full-stack)
  - [Introduction](#introduction)
    - [Development mode](#development-mode)
    - [Production mode](#production-mode)
  - [Quick Start](#quick-start)
  - [Documentation](#documentation)
    - [Folder Structure](#folder-structure)
    - [Babel](#babel)
    - [ESLint](#eslint)
    - [Webpack](#webpack)
    - [Webpack dev server](#webpack-dev-server)
    - [NoDemon](#nodemon)
    - [Golang](#golang)
    - [Concurrently](#concurrently)
    - [VSCode + ESLint + Prettier](#vscode--eslint--prettier)
      - [Installation guide](#installation-guide)

## Introduction

[Create React App](https://github.com/facebook/create-react-app) is the quickest way to get started with React development and it requires no build configuration. But it overly simplifes and hides the build config making it difficult to extend once you've developed your skills. It also requires some additional work to integrate it with a Golang backend application.

This is a 'simple' full stack [React](https://reactjs.org/) application with a [Golang](https://golang.org/) backend. This application is configured with [Airbnb's ESLint rules](https://github.com/airbnb/javascript) and formatted through [prettier](https://prettier.io/).

### Development mode

In the development mode, we will have 2 servers running. The front end code will be served by the [webpack dev server](https://webpack.js.org/configuration/dev-server/) which helps with hot and live reloading. The server side Golang code will be compiled using [NoDemon](https://nodemon.io/) which helps by automatically rebuilding and restarting the server whenever you save changes to the backend code.

### Production mode

In the production mode, we will have only 1 server running. All the client side code will be bundled into static files using webpack and it will be served by the Golang application.

## Quick Start

```bash
# Clone the repository
git clone https://github.com/orstendium/react-golang-full-stack

# Go inside the directory
cd react-golang-full-stack

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Start production server
npm start
```

## Documentation

### Folder Structure

All the source code will be inside **src** directory. Inside src, there are the webapp and backend directories. All the frontend code (react, css, js and any other assets) will be in webapp directory. Backend Golang code will be in the backend directory.

### Babel

[Babel](https://babeljs.io/) helps us to write code in the latest version of JavaScript. If an environment does not support certain features natively, Babel will help compile those features down to a supported version. It also helps us to convert JSX to Javascript.

[.babelrc file](https://babeljs.io/docs/usage/babelrc/) is used describe the configurations required for Babel. Below is the .babelrc file which I am using.

```javascript
{
    "presets": ["env", "react"]
}
```

Babel requires plugins to do the transformation. Presets are the set of plugins defined by Babel. Preset **env** allows to use babel-preset-es2015, babel-preset-es2016, and babel-preset-es2017 and it will transform them to ES5. Preset **react** allows us to use JSX syntax and it will transform JSX to Javascript.

### ESLint

[ESLint](https://eslint.org/) is a pluggable and configurable linter tool for identifying and reporting on patterns in JavaScript.

[.eslintrc.json file](<(https://eslint.org/docs/user-guide/configuring)>) (alternatively configurations can we written in Javascript or YAML as well) is used describe the configurations required for ESLint. Below is the .eslintrc.json file which I am using.

```javascript
{
  "extends": ["airbnb"],
  "env": {
    "browser": true
  },
  "rules": {
    "no-console": "off",
    "comma-dangle": "off",
    "react/jsx-filename-extension": "off"
  }
}
```

[We'll be using Airbnb's Javascript Style Guide](https://github.com/airbnb/javascript), used by many JavaScript developers worldwide. Since we're writing client (browser) code, I am setting the **env** to browser. Optionally, we can override the Airbnb's configurations to suit our needs. I have turned off [**no-console**](https://eslint.org/docs/rules/no-console), [**comma-dangle**](https://eslint.org/docs/rules/comma-dangle) and [**react/jsx-filename-extension**](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/jsx-filename-extension.md) rules.

### Webpack

[Webpack](https://webpack.js.org/) is a module bundler. Its main purpose is to bundle JavaScript files for usage in a browser.

[webpack.config.js](https://webpack.js.org/configuration/) file is used to describe the configurations required for webpack. Below is the webpack.config.js file which I am using.

```javascript
const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const {CleanWebpackPlugin} = require("clean-webpack-plugin");

const outputDirectory = "./src/backend/dist";

module.exports = {
  entry: ["babel-polyfill", "./src/webapp/index.js"],
  output: {
    path: path.join(__dirname, outputDirectory),
    filename: "bundle.js"
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      },
      {
        test: /\.(png|woff|woff2|eot|ttf|svg)$/,
        loader: "url-loader",
	options: {
	  limit: 100000
	}
      }
    ]
  },
  devServer: {
    port: 3000,
    open: true,
    historyApiFallback: true,
    proxy: {
      "/api": "http://localhost:8080"
    }
  },
  plugins: [
    new CleanWebpackPlugin(),
    new HtmlWebpackPlugin({
      template: "./public/index.html",
      favicon: "./public/favicon.ico"
    })
  ]
};
```

1.  **entry:** entry: ./src/webapp/index.js is where the application starts executing and webpack starts bundling.
    Note: babel-polyfill is added to support async/await. Read more [here](https://babeljs.io/docs/en/babel-polyfill#usage-in-node-browserify-webpack).
2.  **output path and filename:** the target directory and the filename for the bundled output
3.  **module loaders:** Module loaders are transformations that are applied on the source code of a module. We pass all the js file through [babel-loader](https://github.com/babel/babel-loader) to transform JSX to Javascript. CSS files are passed through [css-loaders](https://github.com/webpack-contrib/css-loader) and [style-loaders](https://github.com/webpack-contrib/style-loader) to load and bundle CSS files. Fonts and images are loaded through url-loader.
4.  **Dev Server:** Configurations for the webpack-dev-server which will be described in coming section.
5.  **plugins:** [clean-webpack-plugin](https://github.com/johnagan/clean-webpack-plugin) is a webpack plugin to remove the build folder(s) before building. [html-webpack-plugin](https://github.com/jantimon/html-webpack-plugin) simplifies creation of HTML files to serve your webpack bundles. It loads the template (public/index.html) and injects the output bundle.

### Webpack dev server

[Webpack dev server](https://webpack.js.org/configuration/dev-server/) is used along with webpack. It provides a development server that provides live reloading for the client side code. This should be used for development only.

The devServer section of webpack.config.js contains the configuration required to run webpack-dev-server which is given below.

```javascript
devServer: {
    port: 3000,
    open: true,
    historyApiFallback: true,
    proxy: {
        "/api": "http://localhost:8080"
    }
}
```

[**Port**](https://webpack.js.org/configuration/dev-server/#devserver-port) specifies the Webpack dev server to listen on this particular port (3000 in this case). When [**open**](https://webpack.js.org/configuration/dev-server/#devserver-open) is set to true, it will automatically open the home page on startup. [historyApiFallback](https://webpack.js.org/configuration/dev-server/#devserverhistoryapifallback) is set to ensure all routes in React are served via the index page, rather than trying to get them from the server. [Proxying](https://webpack.js.org/configuration/dev-server/#devserver-proxy) URLs can be useful when we have a separate API backend development server and we want to send API requests on the same domain. In our case, we have a Golang backend where we want to send the API requests to.

### NoDemon

NoDemon is a utility that will monitor for any changes in the server source code and it will automatically restart the server. This is used in development only.

Below is the config which I am using in the package.json.

```javascript
{
  "scripts": {
    "backend": "nodemon --ext go --exec \"cd ./src/backend && go build && backend || exit 1\"",
  }
}
```

Here, we tell NoDemon to execute go build in order to build the files in the directory src/backend where our server side code resides, output the executable into the same directory, and then execute that executable. CompileDaemon will restart the go server whenever a .go file under src/backend directory is modified.

#### Note

You might need to edit the backend script depending on your development OS. Bash needs the './backend' when calling the backend executable while the Windows console uses just 'backend.'

### Golang

Golang is an open source compiled language. It is used to build our backend API's.

src/backend/backend.go has the entry function, main, for the server application. Below is the src/backend/backend.go file

```golang
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const port = ":8080"

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string `json:"username"`
	}
	user := User{os.Getenv("USERNAME")}
	p, _ := json.Marshal(user)
	w.Write(p)
}

func main() {
	log.Println("Starting Backend")

	r := mux.NewRouter()
	// Define API routes
	r.HandleFunc("/api/username", usernameHandler).Methods("GET")

	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	distDir := filepath.Join(currentDir, "dist")
	spa := spaHandler{staticPath: distDir, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	log.Println("Http Listening")
	http.ListenAndServe(
		port, r)
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	log.Println(path)
	indexFile := filepath.Join(h.staticPath, h.indexPath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, indexFile)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
```

This starts a server and listens on port 8080 for connections. The app responds with the JSON object `{username: <username>}` for requests to the URL (/api/username). It is also configured to serve the static files from **dist** directory.

### Concurrently

[Concurrently](https://github.com/kimmobrunfeldt/concurrently) is used to run multiple npm commands concurrently. I am using it to run the webpack dev server and the backend go server concurrently in the development environment. Below are the npm script commands used.

```javascript
"webapp": "webpack-dev-server --mode development --devtool inline-source-map --hot",
"backend": "nodemon --ext go --exec \"cd ./src/backend && go build && backend || exit 1\"",
"dev": "concurrently \"npm run server\" \"npm run client\""
```

### VSCode + ESLint + Prettier

[VSCode](https://code.visualstudio.com/) is a lightweight but powerful source code editor. [ESLint](https://eslint.org/) takes care of the code-quality. [Prettier](https://prettier.io/) takes care of all the formatting.

#### Installation guide

1.  Install [VSCode](https://code.visualstudio.com/)
2.  Install [ESLint extension](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
3.  Install [Prettier extension](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
4.  Modify the VSCode user settings to add below configuration

    ```javascript
    "eslint.alwaysShowStatus": true,
    "eslint.autoFixOnSave": true,
    "editor.formatOnSave": true,
    "prettier.eslintIntegration": true
    ```

Above, we have modified editor configurations. Alternatively, this can be configured at the project level by following [this article](https://medium.com/@netczuk/your-last-eslint-config-9e35bace2f99).
