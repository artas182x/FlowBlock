import axios from 'axios';
import config from "../config"

class AuthService {
  login(user) {
    return axios
      .post(config.BASE_API_URL + 'login', {
        certificate: user.certificate,
        privateKey: user.privateKey,
        mspid: user.mspid
      })
      .then(response => {
        if (response.data.token) {
          localStorage.setItem('user', JSON.stringify(response.data));
        }

        return response.data;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }
}

export default new AuthService();
