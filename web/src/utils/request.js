import axios from 'axios';
import { ElMessage } from 'element-plus';
import store from '@/store'
import baseUrl from '../../baseUrl'
import Vrouter from "@/router"
const router = Vrouter

const service = axios.create({
    baseURL: baseUrl,
    timeout: 10000
});

service.interceptors.request.use(
    config => {
        store.commit("setLoadReq",true)
        let token = sessionStorage.getItem('auth')
        if (token != null) {  
            config.headers.Authorization = token;
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
                router.push('/login')
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
