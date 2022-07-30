import axios from 'axios';
import { ElMessage } from 'element-plus';
import store from '@/store'
import baseUrl from '../../baseUrl'
import Vrouter from "@/router"
const router = Vrouter

const service = axios.create({
    baseURL: baseUrl,
    timeout: 100000
});

service.interceptors.request.use(
    config => {
        store.commit("setLoadReq",true)
        let auth = sessionStorage.getItem('auth')
        if (auth != null) { 
            let obj = JSON.parse(auth) 
            config.headers.Authorization = obj.jwt;
        }
        return config;
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

service.interceptors.response.use(
    response => {
        store.commit("setLoadReq",false)
        if (response.status === 200) {
            if(!response.data.status){
                let auth = sessionStorage.getItem('auth')
                if (auth == null ||response.data.msg == "Jwt Token Expired") {
                    router.push('/login')
                } 
                ElMessage({
                    message: response.data.msg,
                    type: 'error',
                })
            }
            return response.data;
        } else {
            Promise.reject();
        }
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

export default service;
