import axios from 'axios';
import config from "../config"
import authHeader from "@/services/auth-header";

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

    refreshToken() {
        return axios.get(config.BASE_API_URL + 'v1/refresh_token', {headers: authHeader()}).then(
            (response) => {
                if (response.data.token) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response.data;
            },
            (error) => {
                if (error.response.status !== 200) {
                    this.logout()
                    this.$router.push('/login');
                }
            }
        );
    }
}

export default new AuthService();
