"use strict";
const path = require("path");
const assert = require("yeoman-assert");
const helpers = require("yeoman-test");
const os = require("os");

describe("generator-tokend-module:app", () => {
  beforeAll(() => {
    return helpers
      .run(path.join(__dirname, "../generators/app"))
      .inDir(path.join(os.tmpdir(), "./temp-test"))
      .withPrompts({
        serviceName: "foo",
        gitService: "GitHub",
        depTool: "go mod"
      });
  });

  it("creates files", () => {
    assert.file([
      "main.go",
      "internal",
      "Dockerfile",
      "config.yaml",
      "README.md"
    ]);
  });
});
