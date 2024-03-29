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
  requestFlow(input) {
    return axios.post(config.BASE_API_URL + 'v1/computation/requestflow', input, { headers: authHeader() });
  }
  downloadFile(fileName) {
    return axios.get(config.BASE_API_URL + 'v1/s3/' + fileName, { headers: authHeader(), responseType: 'blob' });
  }
  requestMedicalData(input) {
    return axios.post(config.BASE_API_URL + 'v1/medicaldata/request', input, {
      headers: authHeader(),
      transformResponse: (res) => {
        if (!res) {
          return JSON.parse("[]")
        }
        res = res.replace('/\\0000/g', '')
        return JSON.parse(res)
          },
      responseType: 'json' });
  }
}

export default new UserService();
