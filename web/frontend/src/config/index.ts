if (process.env.NODE_ENV === 'production') {
  module.exports = {
    BASE_API_URL: "api/"
  }
} else {
  module.exports = {
    BASE_API_URL: "http://192.168.121.104/api/"
  }
}

