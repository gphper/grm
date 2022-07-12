import axios from 'axios';
import { ElMessage } from 'element-plus';
import store from '@/store'
import baseUrl from '../../baseUrl'

const service = axios.create({
    baseURL: baseUrl,
    timeout: 10000
});

service.interceptors.request.use(
    config => {
        store.commit("setLoadReq",true)
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
