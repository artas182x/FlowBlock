import axios from 'axios';
import authHeader from './auth-header';
import config from '../config'

class UserService {
  getUserTokens() {
    return axios.get(config.BASE_API_URL + 'v1/computation/usertokens', { headers: authHeader() });
  }
  getUserQueue() {
    return axios.get(config.BASE_API_URL + 'v1/computation/queue', { headers: authHeader() });
  }
  startComputation(tokenId) {
    return axios.post(config.BASE_API_URL + 'v1/computation/token/' + tokenId + '/start',  "", { headers: authHeader() });
  }
  getAvailableMethods(chainCodeName) {
    return axios.get(config.BASE_API_URL + 'v1/computation/availablemethods/' + chainCodeName, { headers: authHeader() });
  }
  requestToken(input) {
    return axios.post(config.BASE_API_URL + 'v1/computation/requesttoken', input, { headers: authHeader() });
  }
}

export default new UserService();
