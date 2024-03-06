const { https } = require("firebase-functions")
const next = require("next")
const { join } = require("path")

const nextjsDistDir = join("src", require("./src/next.config.js").distDir)

const nextjsServer = next({
  dev: true,
  conf: {
    distDir: nextjsDistDir,
  },
})
const nextjsHandle = nextjsServer.getRequestHandler()

exports.nextjsFunc = https.onRequest((req, res) => {
  return nextjsServer.prepare().then(() => nextjsHandle(req, res))
})
