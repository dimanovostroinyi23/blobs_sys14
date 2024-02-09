"use strict";

const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");
const fsConfig = require("./fs_config");
const operations = require("./operations");

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(
        `Welcome to the super ${chalk.red(
          "generator-tokend-module"
        )} generator!`
      )
    );

    var prompts = [
      {
        name: "serviceName",
        message:
          "What is the name of new service? (e.g. my-service-svc) svc stands for service"
      },
      {
        type: "list",
        name: "gitService",
        message: "Which git hosting are you going to use?",
        choices: ["GitLab", "GitHub", "Other"]
      },
      {
        type: "confirm",
        name: "useDB",
        message: "Would you like to use PostgreSQL database for this service?",
        default: false
      },
      {
        type: "confirm",
        name: "handleHTTP",
        message: "Would you like this service to handle HTTP requests?",
        default: false
      },
      {
        type: "confirm",
        name: "hasDocs",
        message: "Would you have documentation pages for this service?",
        default: false
      },
      {
        name: "packageName",
        message:
          "What is the package name for the new service? (e.g. gitlab.com/tokend/my-service-svc)"
      },
      {
        name: "serviceDescription",
        message: "What is the description of the new service? what does it do?"
      },
      {
        name: "contactName",
        message: "What is your full name?"
      },
      {
        name: "telegramContact",
        message:
          'please write in "t.me/username" format link to your telegram account'
      }
    ];

    return this.prompt(prompts).then(props => {
      this.props = props;

      if (!this.props.hasDocs) {
        fsConfig.dirsToCreate.docs.render = operations.nop;
        fsConfig.filesToCopy.docs.render = operations.nop;
        fsConfig.filesToRender.docs.render = operations.nop;
      }

      if (!this.props.useDB) {
        fsConfig.dirsToCreate.db.render = operations.nop;
        fsConfig.filesToCopy.db.render = operations.nop;
        fsConfig.filesToRender.db.render = operations.nop;
      }

      if (!this.props.handleHTTP) {
        fsConfig.dirsToCreate.http.render = operations.nop;
        fsConfig.filesToCopy.http.render = operations.nop;
        fsConfig.filesToRender.http.render = operations.nop;
      }

      switch (this.props.gitService) {
        case "GitLab":
          fsConfig.travisCI.render = operations.nop;
          fsConfig.githubActions.render = operations.nop;
          fsConfig.golangci.render = operations.nop;
          break;
        case "GitHub":
          fsConfig.gitlabCI.render = operations.nop;
          fsConfig.travisCI.render = operations.nop;
          break;
        default:
          fsConfig.gitlabCI.render = operations.nop;
          fsConfig.golangci.render = operations.nop;
          fsConfig.githubActions.render = operations.nop;
      }

      if (this.props.packageName === "") {
        this.props.packageName = "gitlab.com/tokend/" + this.props.serviceName;
      }

      if (this.props.serviceDescription === "") {
        this.props.serviceDescription = "[//]: # (TODO: add description)";
      }
      if (this.props.telegramContact === "") {
        this.props.telegramContact =
          "[//]: # (TODO: place link to your telegram and email)";
      }
    });
  }

  writing() {
    fsConfig.render(this);
  }

  install() {
    this.log("Installing goimports if not already installed");
    this.spawnCommandSync("go", [
      "install",
      "golang.org/x/tools/cmd/goimports@latest"
    ]);
    this.log(
      "Formatting code and updating import lines with goimports in files:"
    );
    this.spawnCommandSync("goimports", ["-w", "-l", "."]);
  }

  end() {
    this.log("");
    this.log(chalk.green("Generated!"));
    if (this.props.gitService === "GitHub") {
      this.log(chalk.yellow("NOTICE:"));
      this.log(
        "I've generated scripts for github actions for you, but you need to switch `Workflow permissions` to `Read and write permissions` in organizations settings or to set `permissions: write-all` in actions.yaml and tag.yaml if it's personal repository."
      );
      if (this.props.hasDocs) {
        this.log(chalk.yellow("NOTICE:"));
        this.log(
          "You've said that you are using GitHub and want to have GitHub Pages site. I've generated docs/ folder, but you'll need manually set up your GitHub repository for using GitHub Pages and you'll also have to build index.html manually each time."
        );
      }
    }

    if (this.props.gitService === "GitLab") {
      this.log("");
      this.log(chalk.yellow("NOTICE:"));
      this.log(
        "I've generated .gitlab-ci.yml for you, but you'll have to set env variables $DOCKERHUB_USER and $DOCKERHUB_PWD in your GitLab repository and only then foxy will be able to push release images to your docker hub. "
      );
      this.log("");
    }

    this.log("");
    this.log(chalk.green("Thank you for coming. I was glad to help you =)"));
  }
};
